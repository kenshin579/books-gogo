// Package min is an example for shared memory based concurrent min
// algorithm.
package min

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

// Min returns the minimum number from slice a.
func Min(a []int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, e := range a[1:] {
		if min > e {
			min = e
		}
	}
	return min
}

// ParallelMin returns the minimum number from slice a with
// parallelism of n.
func ParallelMin(a []int, n int) int {
	if len(a) < n {
		return Min(a)
	}
	mins := make([]int, n)
	bucketSize := (len(a) + n - 1) / n
	log.Info("bucketSize: ", bucketSize)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			begin, end := i*bucketSize, (i+1)*bucketSize
			log.Info("[", i, "] begin:", begin, " end:", end)
			if end > len(a) {
				end = len(a)
			}
			mins[i] = Min(a[begin:end])
		}(i)
	}
	wg.Wait()
	return Min(mins)
}
