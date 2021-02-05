//Use Fatal like os.exit(1) -unsuccessful termination-
package main

import (
	"fmt"
	"log"
	"math"
)

func sqrt(x float64) (float64, int) {
	if x < 0 {
		return x, -1
	}
	return math.Sqrt(x), 0
}

func main() {
	var res, n = sqrt(2)
	check(res, n)
	res, n = (sqrt(-2))
	check(res, n)
}

func check(res float64, n int) {
	if n == -1 {
		log.Fatal("number negative without sqrt ", res)
	}
	fmt.Println(res)
}
