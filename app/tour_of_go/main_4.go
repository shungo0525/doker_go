package main

import (
	"fmt"
	"math"
	"time"
)

type Vertex struct {
	X, Y float64
}

// メソッド
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 上記のメソッドを関数で書いた場合
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// メソッドに引数がある場合
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// インターフェース
type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// error型を返す => Error()が実行される
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}


func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // メソッドの場合
// 	fmt.Println(Abs(v)) // 関数の場合

  v.Scale(10)
  fmt.Println(v.Abs())

  var i I = T{"hello"}
  i.M()

  var ii interface{}
  describe(ii)

  ii = 42
  describe(ii)

  // 型アサーション
  var iii interface{} = "hello"
  s, ok := iii.(string)
  fmt.Println(s, ok)

  f, ok := iii.(float64)
  fmt.Println(f, ok)

  a := Person{"Arthur Dent", 42}
  fmt.Println(a) // fmt.Println(a.String())と同じ

  if err := run(); err != nil {
    fmt.Println(err)
  }
}
