# cpu_usage
**Lightweight Runtime Conflict Detection for CPU Efficient Transaction Processing**

### Paper Information
- **Author(s):** Naveen Kumar Bandaru
- **Published In:** International Journal of Intelligent Systems and Applications in Engineering (IJISAE)
- **Publication Date:** June, 2023
- **ISSN:**  2147-6799

### Abstract
High concurrency transaction processing systems often experience performance degradation due to conflicts among simultaneous read and write operations. Conventional mechanisms such as Two Phase Locking and Optimistic Concurrency Control introduce blocking, repeated retries, and significant processor overhead. This work examines the impact of these mechanisms on CPU utilization and scalability in distributed environments. A lightweight runtime conflict detection approach is introduced to identify conflicts earlier during execution and reduce unnecessary computation. Experimental evaluation across multiple cluster sizes demonstrates improved processor efficiency and better scalability in transaction processing systems.

### Major Research Contributions
- **Lightweight Runtime Conflict Detection Mechanism:**  
Introduced a runtime method that detects transactional conflicts early during execution using compact metadata instead of relying on heavy locking or late validation.

- **Processor Efficient Transaction Execution:**  
Designed a conflict management approach that reduces blocking synchronization and repeated transaction retries, leading to lower processor utilization during high concurrency workloads.
- **Distributed Experimental Evaluation:** 
Implemented a transaction processing model using Go based concurrent workers to simulate distributed workloads and evaluate processor utilization across cluster sizes.

- **Scalability Analysis Across Cluster Sizes:**  
Conducted experiments on clusters with 3, 5, 7, 9, and 11 nodes to analyze how CPU utilization changes as transaction processing systems scale.

### Practical Significance and Impact
- **Reduced Processor Utilization:**
The lightweight runtime approach significantly lowers CPU usage compared with conventional locking and optimistic concurrency control mechanisms.

- **Improved Transaction Processing Efficiency:**  
Early conflict detection minimizes wasted computation caused by blocking synchronization and repeated transaction retries.

- **Better Scalability for Distributed Systems:**  
Processor consumption decreases steadily as cluster size increases, demonstrating efficient resource utilization and improved scalability.

- **Suitability for High Concurrency Platforms:**  
The framework supports efficient transaction processing in environments such as distributed databases, cloud systems, and microservice based platforms.

### Experimental Results (Summary)

  | Nodes | Lock Based 2PL CPU %| Lightweight Runtime Detection %| Improvement (%) |
  |-------|---------------------| -------------------------------| ---------------|
  | 3     |  88                 | 55                             | 37.50          |
  | 5     |  84                 | 49                             | 41.67          |
  | 7     |  82                 | 46                             | 43.90          |
  | 9     |  80                 | 43                             | 46.25          |
  | 11    |  79                 | 41                             | 48.10          |

### Citation
Lightweight Runtime Conflict Detection for CPU Efficient Transaction Processing
* Naveen Kumar Bandaru
* International Journal of Intelligent Systems and Applications in Engineering (IJISAE) 
* ISSN 2147-6799
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijisae.org/index.php/IJISAE \
**Author Contact** \
**LinkedIn**: linkedin.com/in/naveen-bandaru | **Email**: naveen.bandaru@gmail.com







