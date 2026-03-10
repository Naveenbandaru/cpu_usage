# cpu_usage
**Lightweight Runtime Conflict Detection for CPU Efficient Transaction Processing**

### Paper Information
- **Author(s):** Naveen Kumar Bandaru
- **Published In:** International Journal of IJISAE
- **Publication Date:** June, 2023
- **ISSN:**  2147-6799

### Abstract
This paper addresses the problem of excessive CPU utilization in network congestion monitoring within distributed and cloud infrastructures. Modern systems collect telemetry such as bandwidth usage, packet loss, queue occupancy, and delay statistics, yet conventional frameworks analyze these metrics independently at each node. This repetitive data collection and duplicate analysis create high computational overhead, frequent synchronization, and delayed identification of bottlenecks. As the number of nodes grows, telemetry streams expand proportionally, triggering repeated diagnostics and unnecessary reprocessing that further amplify processor load. Consequently, centralized aggregation and monitoring costs reduce throughput, limit scalability, and compromise service stability across distributed environments.

### Key Contributions
- **Cross‑Node Telemetry Correlation:**  
Introduced a monitoring approach that correlates telemetry across nodes, reducing redundant computation and improving CPU efficiency compared to conventional independent node analysis.

- **Processor‑Efficient Monitoring Design:**  
Developed mechanisms that minimize duplicate diagnostics and synchronization overhead, lowering CPU utilization while maintaining visibility into congestion events across distributed systems.

- **Simulation and Validation:** 
Implemented a Go‑based simulation to model conventional monitoring overhead, validating how repeated local analysis and central aggregation lead to high processor load.

- **Scalability Enhancement:**  
Demonstrated that coordinated telemetry analysis prevents linear growth of CPU consumption with cluster size, enabling scalable monitoring in cloud and distributed infrastructures.

### Relevance & Real-World Impact
- **Reduced CPU Utilization:**
Achieved significant processor efficiency by eliminating redundant monitoring tasks, freeing resources for application workloads and improving responsiveness.

- **Improved Latency and Throughput:**  
Lower monitoring overhead directly reduces latency and increases throughput, enhancing service stability in distributed environments.

- **Scalable Cloud Deployment:**  
Framework supports large clusters without proportional CPU growth, addressing scalability challenges in modern cloud infrastructures.

- **Operational Cost and Energy Savings:**  
Efficient monitoring reduces unnecessary computation, lowering energy consumption and operational costs in data centers.

- **Practical Applicability:**
Provides a reference model for industry and research, offering a processor‑efficient monitoring design suitable for production systems and academic exploration.
 
### Experimental Results (Summary)

  | Nodes | Local Telemetry CPU | Telemetry corelation CPU | Improvment (%)  |
  |-------|---------------------| -------------------------| ----------------|
  | 3     |  72                 | 54                       | 25.0           |
  | 5     |  70                 | 50                       | 28.6           |
  | 7     |  68                 | 47                       | 30.9           |
  | 9     |  67                 | 45                       | 32.8           |
  | 11    |  66                 | 43                       | 34.8           |

### Citation
Cross Node Telemetry for CPU Efficient Congestion Monitoring
* Arunkumar Sambandam
* International Journal of Leading Research Publication 
* *****ISSN E-ISSN: 2582-8010
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
[**************https://www.ijlrp.com/](https://www.ijirmps.org/research-paper.php?id=232943) \
**Author Contact** \
**LinkedIn**: linkedin.com/in/arunkumar-sambandam-9b769b | **Email**: arunkumar.sambandam@yahoo.com







