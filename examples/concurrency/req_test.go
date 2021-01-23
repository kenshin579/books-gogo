package concurrency

import (
	"fmt"
	"sync"
)

//요청과 응답 pairing 시키기
func ExamplePlusOneService() {
	reqs := make(chan Request)
	defer close(reqs)
	for i := 0; i < 3; i++ {
		go PlusOneService(reqs, i)
	}
	var wg sync.WaitGroup
	for i := 3; i < 53; i += 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resps := make(chan Response)
			reqs <- Request{i, resps}
			fmt.Println(i, "=>", <-resps)
		}(i)
	}
	wg.Wait()
	//Output:
}
