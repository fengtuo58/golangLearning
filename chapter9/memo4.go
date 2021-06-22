package main
import(
	"fmt"
	"net/http"
	"io/ioutil"
	"sync"
	"time"
)
type Func func(key string) (interface{}, error)

type result struct {
	err error
	value interface{}
}
type entry struct {
	res result
	ready chan struct{}
}

func New(f Func) *Memo {
	return &Memo{f:f, cache: make(map[string]*entry)}
}

type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]*entry
}

func(memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready:make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		<-e.ready
		memo.mu.Unlock()
	}
	return 	e.res.value, e.res.err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	incomingUrls := make(chan string)
	go func() {
		urls := [...]string {"www.baidu.com", "www.cctv.com", "www.github.com"}
		for _, url := range urls {
			incomingUrls<-url
		}
	}()
	m := New(httpGetBody)
	for url := range incomingUrls {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}
