package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
  // ポインタ
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println(Vertex{1,2})

  // struct(構造体)には.(ドット)でアクセスする
	v := Vertex{1, 2}
  v.X = 4
  fmt.Println(v.X)

  foo := 1
  fmt.Println("foo:", foo)
  Add(&foo)
  fmt.Println("foo:", foo)

  fmt.Println(v1, p, v2, v3)
}

func Add(a *int) { //←aはint型のポインタ変数になる。👉を作ったイメージ。
  *a += 1         //aポインタが指している変数に入る数を+1する
  fmt.Println("aの値(関数内)：", *a)
}
