package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var Status bool
var Salary float32
var Time time.Time

func VariablesRest() {
	Name = "John"
	Status = true
	Salary = 1000.50
	Time = time.Now()

	fmt.Println(Name)
	fmt.Println(Status)
	fmt.Println(Salary)
	fmt.Println(Time)

}

func TextToConverter(number int) (bool, string) {
	text := strconv.Itoa(number)
	return true, text
}
