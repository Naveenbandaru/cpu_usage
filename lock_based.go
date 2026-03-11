import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)
const (
	keys        = 1000
	workers     = 16
	iterations  = 200000
)
type LockRecord struct {
	mu    sync.Mutex
	value int
}
type LockStore struct {
	data map[int]*LockRecord
}
func NewLockStore() *LockStore {
	m := make(map[int]*LockRecord)
	for i := 0; i < keys; i++ {
		m[i] = &LockRecord{}
	}
	return &LockStore{data: m}
}
func (s *LockStore) Read(k int) int {
	r := s.data[k]
	r.mu.Lock()
	v := r.value
	r.mu.Unlock()
	return v
}
func (s *LockStore) Write(k int, v int) {
	r := s.data[k]
	r.mu.Lock()
	r.value = v
	r.mu.Unlock()
}
type OCCRecord struct {
	value   int
	version uint64
}
type OCCStore struct {
	data map[int]*OCCRecord
}
type Txn struct {
	reads  map[*OCCRecord]uint64
	writes map[*OCCRecord]int
}
func NewOCCStore() *OCCStore {
	m := make(map[int]*OCCRecord)
	for i := 0; i < keys; i++ {
		m[i] = &OCCRecord{}
	}
	return &OCCStore{data: m}
}

func NewTxn() *Txn {
	return &Txn{
		reads:  make(map[*OCCRecord]uint64),
		writes: make(map[*OCCRecord]int),
	}
}
func (t *Txn) Read(r *OCCRecord) int {
	v := atomic.LoadUint64(&r.version)
	t.reads[r] = v
	return r.value
}
func (t *Txn) Write(r *OCCRecord, v int) {
	t.writes[r] = v
}
func (t *Txn) Commit() bool {
	for r, v := range t.reads {
		if atomic.LoadUint64(&r.version) != v {
			return false
		}
	}
	for r, val := range t.writes {
		r.value = val
		atomic.AddUint64(&r.version, 1)
	}
	return true
}
func cpuPercent(start time.Time) float64 {
	elapsed := time.Since(start).Seconds()
	return float64(runtime.NumGoroutine()) / float64(runtime.NumCPU()) * 100 * elapsed / elapsed
}
func runLock() {
	store := NewLockStore()
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < iterations; j++ {
				k := rand.Intn(keys)
				if rand.Intn(2) == 0 {
					store.Read(k)
				} else {
					store.Write(k, rand.Int())
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Lock Based Time:", time.Since(start))
	fmt.Println("Lock Based CPU:", cpuPercent(start))
}
func runOCC() {
	store := NewOCCStore()
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < iterations; j++ {
				for {
					t := NewTxn()
					k := rand.Intn(keys)
					r := store.data[k]
					t.Read(r)
					t.Write(r, rand.Int())
					if t.Commit() {
						break
					}
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Optimistic Time:", time.Since(start))
	fmt.Println("Optimistic CPU:", cpuPercent(start))
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	runLock()
	runOCC()
}
