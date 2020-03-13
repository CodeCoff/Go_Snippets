package main

import (
	"fmt"
	"math"
	"net/http"
)

//--------------------------------
//Structs: Go doesn't have classes (whaaaat!). It uses Structs.
//remote control car example
type car struct {
	gas_pedal      uint16
	brake_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

func main() {
	fmt.Println("Welcome to go!")
	squareroot()
	sum()

	x := 15
	a := &x //memory address

	fmt.Println(*a, a)

	//change value
	*a = 5

	fmt.Println(*a)

	//------------------------------------------
	//simple web server
	//http.HandleFunc("/", indexHandler) //this is a router
	//http.ListenAndServe(":8000", nil)  //this is our server
	//------------------------------------------
	//structs implementation
	a_car := car{gas_pedal: 25000,
		brake_pedal:    0,
		steering_wheel: 12500,
		top_speed_kmh:  160}

	fmt.Println(a_car.gas_pedal)
}

//here w  is a writer, request
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go is cool!")
}

func squareroot() {
	fmt.Println("Square root of 9 is:", math.Sqrt(9))
}

func sum() {
	sum := 0
	nums := []int{1, 2, 3, 4}
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}

func str(a, b string) (string, string) {
	return a, b
}
