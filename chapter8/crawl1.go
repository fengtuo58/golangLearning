import (
	"string"
)


func crawl(url string) []string{
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	go func() {
		worklist <- os.Args[1:]
	}()

	seen := make(map[string]bool)
	for list := range worklist {
		if !seen[link] {
			seen[link] = true
			go func(link string) {
				worklist <- crawl(link)
			}(link)
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int //number of pending sends to worklist
	n++
	go func() {
		worklist <- os.Args[1:]
	}()
	seen := make(map[string]bool)
	for; n >= 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
			}
			n++
			go func(link string) {
				worklist <- crawl(link)
			}(link)
		}
	}
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func(){ worklist <- os.Args[1:]}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {worklist <- foundLinks}()
			}
		}
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func main() {
	fmt.Println("Commencing countdonw.")
	tick := time.Tick(1*time.Second)
	for xountdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	lauch()
}



