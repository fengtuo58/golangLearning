import (
	"time"
)
type Employee struct{
	ID int
	Name string
	Address string
	Position string
	Dob time.Time
	Salary int
	ManagerId int 
}

var dilbert Employee


