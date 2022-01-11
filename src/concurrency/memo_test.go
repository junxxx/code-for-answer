package main

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/junxxx/codeforgetanswer/concurrency/cache"
	"github.com/junxxx/codeforgetanswer/concurrency/memo"
)

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://www.baidu.com",
			"https://www.zhihu.com",
			"https://www.qq.com",
			"http://www.paisen.site",
			"https://www.baidu.com",
			"https://www.zhihu.com",
			"https://www.qq.com",
			"http://www.paisen.site",
			"https://www.baidu.com",
			"https://www.zhihu.com",
			"https://www.qq.com",
			"http://www.paisen.site",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

func TestMemo(t *testing.T) {
	m := memo.New(cache.HttpGetBody)
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}
