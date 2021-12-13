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
name                          time/op
Random/pdqsort_64             2.44µs ± 1%
Random/stdsort_64             2.38µs ± 0%
Random/pdqsort_256            13.1µs ± 7%
Random/stdsort_256            13.8µs ± 0%
Random/pdqsort_1024           62.3µs ± 1%
Random/stdsort_1024           62.2µs ± 0%
Random/pdqsort_4096            289µs ± 0%
Random/stdsort_4096            290µs ± 0%
Random/pdqsort_65536          6.08ms ± 0%
Random/stdsort_65536          6.11ms ± 0%
Sorted/pdqsort_64              249ns ± 0%
Sorted/stdsort_64              709ns ± 1%
Sorted/pdqsort_256             577ns ± 0%
Sorted/stdsort_256            3.36µs ± 0%
Sorted/pdqsort_1024           1.97µs ± 0%
Sorted/stdsort_1024           16.6µs ± 0%
Sorted/pdqsort_4096           7.31µs ± 0%
Sorted/stdsort_4096           78.9µs ± 0%
Sorted/pdqsort_65536           115µs ± 0%
Sorted/stdsort_65536          1.73ms ± 0%
NearlySorted/pdqsort_64        906ns ± 2%
NearlySorted/stdsort_64       1.00µs ± 2%
NearlySorted/pdqsort_256      4.52µs ± 1%
NearlySorted/stdsort_256      5.01µs ± 1%
NearlySorted/pdqsort_1024     21.2µs ± 0%
NearlySorted/stdsort_1024     25.6µs ± 4%
NearlySorted/pdqsort_4096      101µs ± 0%
NearlySorted/stdsort_4096      117µs ± 1%
NearlySorted/pdqsort_65536    2.14ms ± 0%
NearlySorted/stdsort_65536    2.46ms ± 0%
Reversed/pdqsort_64            320ns ± 1%
Reversed/stdsort_64            845ns ± 0%
Reversed/pdqsort_256           831ns ± 1%
Reversed/stdsort_256          3.71µs ± 0%
Reversed/pdqsort_1024         2.88µs ± 1%
Reversed/stdsort_1024         17.9µs ± 1%
Reversed/pdqsort_4096         10.9µs ± 1%
Reversed/stdsort_4096         83.9µs ± 0%
Reversed/pdqsort_65536         172µs ± 0%
Reversed/stdsort_65536        1.80ms ± 0%
NearlyReversed/pdqsort_64     1.17µs ± 1%
NearlyReversed/stdsort_64     1.41µs ± 1%
NearlyReversed/pdqsort_256    5.63µs ± 1%
NearlyReversed/stdsort_256    7.24µs ± 1%
NearlyReversed/pdqsort_1024   28.9µs ± 4%
NearlyReversed/stdsort_1024   38.3µs ± 0%
NearlyReversed/pdqsort_4096    136µs ± 1%
NearlyReversed/stdsort_4096    176µs ± 0%
NearlyReversed/pdqsort_65536  2.73ms ± 1%
NearlyReversed/stdsort_65536  3.59ms ± 0%
Mod8/pdqsort_64                878ns ± 1%
Mod8/stdsort_64                941ns ± 0%
Mod8/pdqsort_256              3.01µs ± 1%
Mod8/stdsort_256              3.78µs ± 0%
Mod8/pdqsort_1024             10.6µs ± 1%
Mod8/stdsort_1024             14.3µs ± 0%
Mod8/pdqsort_4096             45.7µs ± 0%
Mod8/stdsort_4096             59.4µs ± 1%
Mod8/pdqsort_65536             754µs ± 1%
Mod8/stdsort_65536            1.01ms ± 0%
AllEqual/pdqsort_64            250ns ± 0%
AllEqual/stdsort_64            376ns ± 2%
AllEqual/pdqsort_256           578ns ± 0%
AllEqual/stdsort_256          1.02µs ± 1%
AllEqual/pdqsort_1024         1.97µs ± 0%
AllEqual/stdsort_1024         3.69µs ± 0%
AllEqual/pdqsort_4096         7.20µs ± 0%
AllEqual/stdsort_4096         14.0µs ± 0%
AllEqual/pdqsort_65536         115µs ± 0%
AllEqual/stdsort_65536         223µs ± 0%
```

