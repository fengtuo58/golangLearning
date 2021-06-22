package main
import(
	"http"
)
func main() {
	
}

func httpGetBody(url string) (interface{}, error){
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type Memo struct {
	f Func
	mu sync.Mutex
	cache map[string]result
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{f:f, cache:make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	res,ok := memo.cache[key]
	
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
		memo.mu.Lock()
		res, ok := memo.cache[key]
		if !ok {
			res.value, res.err = memo.f(key)
			memo.cahe[key] = res
		}
		memo.mu.Unlock() 
	}
	return res.value, res.err
}

m := memo.New(httpGetBody)

for url := range incomingURLs() {
	start := time.Now()
	value, err := m.Get(url)
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
}



