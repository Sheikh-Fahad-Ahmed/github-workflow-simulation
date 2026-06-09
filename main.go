package main

import (
	"fmt"
	"sync"
)

func main() {
	featureA := make(chan string, 4)
	// featureB := make(chan string, 4)
	// mainBranch := make(chan string, 4)
	mergeCh := make(chan string, 2)
	// done := make(chan struct{})

	var devWg sync.WaitGroup
	var listenerWg sync.WaitGroup

	for i := 0; i < 6; i++ {
		devWg.Add(1)
		go developer(fmt.Sprintf("Dev-%d", i+1), "featureA", featureA, &devWg)
	}

	listenerWg.Add(6)
	for range 6 {
		go branchListener("featureA", featureA, mergeCh, &listenerWg)
	}

}

func developer(name, branch string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	messages := []string{"fix: login bug",
		"feat: add dashboard",
		"feat: dark mode toggle",
		"refactor: auth module",
		"chore: update deps",
		"docs: update README"}

	for _, msg := range messages {
		ch <- fmt.Sprintf("[%s] %s pushing: %s", branch, name, msg)
	}

}

func branchListener(branch string, ch <-chan string, merge chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
}
