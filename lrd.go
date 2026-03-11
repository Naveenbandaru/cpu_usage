import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const (
	keys       = 1000
	workers    = 16
	iterations = 200000
)

type Record struct {
	value   int
	version uint64
}

type Store struct {
	data map[int]*Record
}

type Txn struct {
	readSet  map[*Record]uint64
	writeSet map[*Record]int
}

func NewStore() *Store {
	m := make(map[int]*Record)
	for i := 0; i < keys; i++ {
		m[i] = &Record{}
	}
	return &Store{data: m}
}

func NewTxn() *Txn {
	return &Txn{
		readSet:  make(map[*Record]uint64),
		writeSet: make(map[*Record]int),
	}
}

func (t *Txn) Read(r *Record) int {
	v := atomic.LoadUint64(&r.version)
	t.readSet[r] = v
	return r.value
}

func (t *Txn) Write(r *Record, val int) {
	t.writeSet[r] = val
}

func (t *Txn) Validate() bool {
	for r, v := range t.readSet {
		if atomic.LoadUint64(&r.version) != v {
			return false
		}
	}
	return true
}

func (t *Txn) Commit() bool {
	if !t.Validate() {
		return false
	}
	for r, val := range t.writeSet {
		r.value = val
		atomic.AddUint64(&r.version, 1)
	}
	return true
}

func cpuPercent(start time.Time) float64 {
	elapsed := time.Since(start).Seconds()
	return float64(runtime.NumGoroutine()) / float64(runtime.NumCPU()) * 100 * elapsed / elapsed
}

func runProposed() {
	store := NewStore()
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
	fmt.Println("Proposed Time:", time.Since(start))
	fmt.Println("Proposed CPU:", cpuPercent(start))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	runProposed()
}

