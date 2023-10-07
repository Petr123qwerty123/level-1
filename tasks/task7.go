package tasks

import (
	"fmt"
	"sync"
	"wb-level-1/concurrencyMap"
)

func Task7() {
	var wg sync.WaitGroup

	n := 10

	wg.Add(n)

	cm := concurrencyMap.NewConcurrencyMap[int, string]()

	for i := 0; i < n; i++ {
		i := i

		go func() {
			defer wg.Done()

			cm.Set(i, string(rune(i)))
			fmt.Printf("%d:%s\n", i, string(rune(i)))
		}()
	}

	wg.Wait()

	fmt.Println(len(cm.Map) == n)
}
