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
Random/pdqsort-64         2.43µs ± 0%
Random/stdsort-64         2.38µs ± 0%
Random/pdqsort-256        13.7µs ± 0%
Random/stdsort-256        13.7µs ± 1%
Random/pdqsort-1024       61.7µs ± 0%
Random/stdsort-1024       62.1µs ± 0%
Random/pdqsort-4096        287µs ± 0%
Random/stdsort-4096        290µs ± 0%
Random/pdqsort-65536      6.02ms ± 0%
Random/stdsort-65536      6.11ms ± 0%
Sorted/pdqsort-64          254ns ± 0%
Sorted/stdsort-64          711ns ± 0%
Sorted/pdqsort-256         581ns ± 0%
Sorted/stdsort-256        3.38µs ± 0%
Sorted/pdqsort-1024       1.98µs ± 0%
Sorted/stdsort-1024       16.6µs ± 0%
Sorted/pdqsort-4096       7.34µs ± 0%
Sorted/stdsort-4096       78.8µs ± 0%
Sorted/pdqsort-65536       115µs ± 0%
Sorted/stdsort-65536      1.73ms ± 0%
Sorted90/pdqsort-64        372ns ± 2%
Sorted90/stdsort-64        750ns ± 0%
Sorted90/pdqsort-256      3.25µs ± 0%
Sorted90/stdsort-256      3.95µs ± 0%
Sorted90/pdqsort-1024     11.9µs ± 0%
Sorted90/stdsort-1024     19.6µs ± 0%
Sorted90/pdqsort-4096     49.6µs ± 0%
Sorted90/stdsort-4096     95.2µs ± 0%
Sorted90/pdqsort-65536     922µs ± 0%
Sorted90/stdsort-65536    2.12ms ± 0%
Reversed/pdqsort-64        327ns ± 1%
Reversed/stdsort-64        849ns ± 1%
Reversed/pdqsort-256       831ns ± 1%
Reversed/stdsort-256      3.70µs ± 1%
Reversed/pdqsort-1024     2.91µs ± 1%
Reversed/stdsort-1024     18.4µs ± 5%
Reversed/pdqsort-4096     11.0µs ± 0%
Reversed/stdsort-4096     83.5µs ± 0%
Reversed/pdqsort-65536     172µs ± 0%
Reversed/stdsort-65536    1.80ms ± 0%
Reversed90/pdqsort-64     1.55µs ± 1%
Reversed90/stdsort-64      943ns ± 0%
Reversed90/pdqsort-256    5.41µs ± 1%
Reversed90/stdsort-256    4.87µs ± 0%
Reversed90/pdqsort-1024   24.4µs ± 0%
Reversed90/stdsort-1024   24.5µs ± 3%
Reversed90/pdqsort-4096   99.4µs ± 0%
Reversed90/stdsort-4096    120µs ± 1%
Reversed90/pdqsort-65536  1.78ms ± 0%
Reversed90/stdsort-65536  2.52ms ± 0%
Mod8/pdqsort-64            885ns ± 1%
Mod8/stdsort-64            945ns ± 0%
Mod8/pdqsort-256          3.05µs ± 0%
Mod8/stdsort-256          3.79µs ± 1%
Mod8/pdqsort-1024         10.7µs ± 1%
Mod8/stdsort-1024         14.2µs ± 0%
Mod8/pdqsort-4096         46.1µs ± 0%
Mod8/stdsort-4096         59.4µs ± 0%
Mod8/pdqsort-65536         755µs ± 0%
Mod8/stdsort-65536        1.01ms ± 0%
```

