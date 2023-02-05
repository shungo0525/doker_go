package main

import (
  "fmt"
  "math"
  "math/rand"
)

var c, python, java bool

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
    // Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
  fmt.Println(add(42, 13))

//   := は変数定義(再代入可能)
  a, b := swap("hello", "world")
  fmt.Println(a, b)

  fmt.Println(split(17))

  var i int
  fmt.Println(i, c, python, java)
  fmt.Println(Small)
  fmt.Println(needFloat(Big))
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needFloat(x float64) float64 {
	return x * 0.1
}