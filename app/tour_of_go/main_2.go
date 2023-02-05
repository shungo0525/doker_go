package main

import (
  "fmt"
  "math"
  "runtime"
  "time"
)

func main() {
// 	sum := 0
// 	for i := 0; i < 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)

	sum := 1
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)

  fmt.Println(sqrt(2), sqrt(-4))

  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )

  v := 10
  lim := 11
  fmt.Printf("%v >= %v\n", v, lim)

  fmt.Println(Sqrt(5))

  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux.")
  default:
    // freebsd, openbsd,
    // plan9, windows...
    fmt.Printf("%s.\n", os)
  }

  fmt.Println("When's Saturday?")
  today := time.Now().Weekday()
  switch time.Saturday {
  case today + 0:
    fmt.Println("Today.")
  case today + 1:
    fmt.Println("Tomorrow.")
  case today + 2:
    fmt.Println("In two days.")
  default:
    fmt.Println("Too far away.")
  }

  // 遅延(Last in First out)
  defer fmt.Println("world")
 	fmt.Println("hello")

 	fmt.Println("counting")
  for i := 0; i < 10; i++ {
    defer fmt.Println(i)
  }
  fmt.Println("done")
}

func Sqrt(x float64) float64 {
  z := float64(1)
	for i:=1; i<10; {
	  z -= (z*z - x) / (2*z)
	  i += 1
	}
  return z
}

func pow(x, n, lim float64) float64 {
  // xのn乗
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
