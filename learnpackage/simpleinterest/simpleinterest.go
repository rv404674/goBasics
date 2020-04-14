package simpleinterest

import "fmt"

//init wont take any arguemtn and also wont return anything
func init() {
	fmt.Println("Simple INterst Package Initialized")
}

// Note the Capital Letter that Calculate is starting with. This is essential
func Calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
