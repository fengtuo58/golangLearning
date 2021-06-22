package main
import(
	"sync"
)

var mu sync.Mutex
func main(){

}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false //insufficient
	}
	reutrn true
}

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func deposit(amount int) {
	balance += amount
}

var mu sync.RWMutext

var balance int

func balance() int {
	mu.RLock()
	defer mu.RUnLock()
	return balance
}

func deposit(amount int) {
	mu.Lock()
	defer mu.UnLock()
	balance += amount
}
