package main

import (
	"fmt"
	"sync"
)

func main() {
	featureA := make(chan string, 4)
	featureB := make(chan string, 4)
	mainBranch := make(chan string, 4)
	mergeCh := make(chan string, 2)
	done := make(chan struct{})

	var devWg sync.WaitGroup
	var listenerWg sync.WaitGroup
	// i := 1

	devWg.Add(6)

	// for range 2 {
	// 	go developer(fmt.Sprintf("Dev-%d", i), "feature-A", featureA, &devWg)
	// 	i++
	// }

	// for range 2 {
	// 	go developer(fmt.Sprintf("Dev-%d", i), "feature-B", featureB, &devWg)
	// 	i++
	// }

	// for range 2 {
	// 	go developer(fmt.Sprintf("Dev-%d", i), "main", mainBranch, &devWg)
	// 	i++
	// }
	go developer("Dev-1", "feature-A", featureA, &devWg)
	go developer("Dev-2", "feature-A", featureA, &devWg)
	go developer("Dev-3", "feature-B", featureB, &devWg)
	go developer("Dev-4", "feature-B", featureB, &devWg)
	go developer("Dev-5", "main", mainBranch, &devWg)
	go developer("Dev-6", "main", mainBranch, &devWg)

	go func() {
		devWg.Wait()
		close(featureA)
		close(featureB)
		close(mainBranch)

	}()

	listenerWg.Add(3)
	go branchListener("featureA", featureA, mergeCh, &listenerWg)
	go branchListener("featureB", featureB, mergeCh, &listenerWg)
	go branchListener("main", mainBranch, mergeCh, &listenerWg)

	go mergeHandler(mergeCh, done)

	listenerWg.Wait()
	close(mergeCh)

	<-done

}

func developer(name, branch string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	// messagesString := []string{"fix: login bug",
	// 	"feat: add dashboard",
	// 	"feat: dark mode toggle",
	// 	"refactor: auth module",
	// 	"chore: update deps",
	// 	"docs: update README"}

	messages := map[string]string{
		"Dev-1": "fix: login bug",
		"Dev-2": "feat: add dashboard",
		"Dev-3": "feat: dark mode toggle",
		"Dev-4": "refactor: auth module",
		"Dev-5": "chore: update deps",
		"Dev-6": "docs: update README",
	}

	fmt.Printf("[%s] %s pushing: %s\n", branch, name, messages[name])
	ch <- messages[name]

}

func branchListener(branch string, ch <-chan string, merge chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range ch {
		fmt.Printf("[%s] Received commit: %s\n", branch, msg)
	}

	if branch != "main"{
		merge <- fmt.Sprintf("AUTO-MERGE: %s into main", branch)
	} 
}

func mergeHandler(merge <-chan string, done chan<- struct{}) {
	for msg := range merge {
		fmt.Println(msg)
	}

	done <- struct{}{}
}
