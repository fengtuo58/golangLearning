package main

type result {
	value interface{}
	err error
}

type Func func(string) (interface{}, err)

type request struct {
	key string
	response chan <-result
}

type Memo struct {
	requests chan request
}
type entry struct {
	ready chan struct{}
	res result
}

func New(f Func) *Memo {
	memo := &Memo{request: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string) {
	response := make(chan result)
	memo.request <- request{key, response}
	res := <-response
	return res.Value, res.err
}

func (memo *Memo)Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready:make(chan struct)}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e* entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}


func main() {
	
}
