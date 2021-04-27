package main
//
//type Handler interface {
//	ServerHttp(w ResponseWriter, r *Request)
//}
//
//func ListenAndServe(address string, h Handler) error
//

func main() {
	db := database{"shoes": 50, "socks":5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d) 
}

type database map[string]dollars

func (db database) ServerHttp(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s \n", item, price)
	}
}

func (db database) ServerHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
	}
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item %q \n", item)
		return
	}
	fmt.Fprintf(w, "%s\n",price)
default:
	w.WriterHeader(http.StatusNotFound)
	fmt.Fprintf(w, "no such page: %s\n", req.URL)
}
type HandlerFunc func(w ResponseWrietr, r *Request)
func (f HandlerFunc) ServeHTTP func(w ResponseWriter, r *Request) {
	
}

