package main

import (
	"fmt"
	"sync"
)

type Post struct {
	views int
	mu    sync.Mutex
}

func (p *Post) inc(wg *sync.WaitGroup) {
	//defer p.mu.Unlock()
	//defer wg.Done()
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()
	p.mu.Lock()
	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myView := Post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		myView.inc(&wg)
	}
	fmt.Println("view:-->", myView.views)

	wg.Wait()

}
