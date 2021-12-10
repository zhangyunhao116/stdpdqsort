# std-pdqsort

The algorithm is mainly based on pattern-defeating quicksort by Orson Peters.

- Paper: https://arxiv.org/pdf/2106.05123.pdf
- C++  implementation: https://github.com/orlp/pdqsort
- Rust implementation: https://docs.rs/pdqsort/latest/pdqsort/

```
Best        Average     Worst       Memory      Stable      Deterministic
n           n log n     n log n     log n       No          Yes
```



## Features

- **Unstable sort**, may reorder equal elements.
- Disable the optimization from [BlockQuickSort](https://dl.acm.org/doi/10.1145/3274660), since its poor performance in Go.



## Benchmark

Go version: go1.18-a412b5f0d8 linux/amd64

CPU: Intel 11700k(8C16T)

OS: ubuntu 20.04

MEMORY: 16G x 2 (3200MHz)

```text
name                      time/op
Random/pdqsort-64         2.44µs ± 0%
Random/stdsort-64         2.38µs ± 0%
Random/pdqsort-256        13.8µs ± 0%
Random/stdsort-256        13.6µs ± 0%
Random/pdqsort-1024       61.8µs ± 0%
Random/stdsort-1024       62.1µs ± 0%
Random/pdqsort-4096        288µs ± 0%
Random/stdsort-4096        290µs ± 0%
Random/pdqsort-65536      6.02ms ± 0%
Random/stdsort-65536      6.11ms ± 0%
Sorted/pdqsort-64          258ns ± 0%
Sorted/stdsort-64          712ns ± 0%
Sorted/pdqsort-256         587ns ± 0%
Sorted/stdsort-256        3.37µs ± 0%
Sorted/pdqsort-1024       1.98µs ± 0%
Sorted/stdsort-1024       16.6µs ± 0%
Sorted/pdqsort-4096       7.32µs ± 0%
Sorted/stdsort-4096       78.7µs ± 0%
Sorted/pdqsort-65536       115µs ± 0%
Sorted/stdsort-65536      1.74ms ± 0%
Sorted90/pdqsort-64        381ns ± 1%
Sorted90/stdsort-64        748ns ± 0%
Sorted90/pdqsort-256      3.27µs ± 0%
Sorted90/stdsort-256      3.97µs ± 0%
Sorted90/pdqsort-1024     11.9µs ± 0%
Sorted90/stdsort-1024     19.7µs ± 0%
Sorted90/pdqsort-4096     49.5µs ± 0%
Sorted90/stdsort-4096     95.3µs ± 0%
Sorted90/pdqsort-65536     923µs ± 0%
Sorted90/stdsort-65536    2.12ms ± 0%
Reversed/pdqsort-64        330ns ± 1%
Reversed/stdsort-64        844ns ± 0%
Reversed/pdqsort-256       835ns ± 0%
Reversed/stdsort-256      3.70µs ± 0%
Reversed/pdqsort-1024     2.93µs ± 0%
Reversed/stdsort-1024     17.8µs ± 0%
Reversed/pdqsort-4096     11.0µs ± 0%
Reversed/stdsort-4096     83.8µs ± 0%
Reversed/pdqsort-65536     172µs ± 0%
Reversed/stdsort-65536    1.80ms ± 0%
Reversed90/pdqsort-64     1.57µs ± 0%
Reversed90/stdsort-64      942ns ± 0%
Reversed90/pdqsort-256    5.51µs ± 1%
Reversed90/stdsort-256    4.89µs ± 1%
Reversed90/pdqsort-1024   25.8µs ± 0%
Reversed90/stdsort-1024   24.4µs ± 3%
Reversed90/pdqsort-4096   99.5µs ± 1%
Reversed90/stdsort-4096    120µs ± 1%
Reversed90/pdqsort-65536  1.78ms ± 0%
Reversed90/stdsort-65536  2.52ms ± 0%
Mod8/pdqsort-64            914ns ± 2%
Mod8/stdsort-64            947ns ± 0%
Mod8/pdqsort-256          3.03µs ± 0%
Mod8/stdsort-256          3.77µs ± 0%
Mod8/pdqsort-1024         10.7µs ± 0%
Mod8/stdsort-1024         14.3µs ± 0%
Mod8/pdqsort-4096         46.1µs ± 0%
Mod8/stdsort-4096         59.3µs ± 0%
Mod8/pdqsort-65536         755µs ± 0%
Mod8/stdsort-65536        1.01ms ± 0%
```

