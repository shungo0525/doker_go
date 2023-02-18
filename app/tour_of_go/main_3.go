package main

import (
  "fmt"
  "golang.org/x/tour/pic"
  "golang.org/x/tour/wc"
  "strings"
  "math"
)

type Vertex struct {
	X int
	Y int
}

type Vertex2 struct {
	Lat, Long float64
}

// 連想配列
var m = map[string]Vertex2{
	"Bell Labs": Vertex2{
		40.68433, -74.39967,
	},
	"Google": Vertex2{
		37.42202, -122.08408,
	},
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

//   var a [2]string
//   a[0] = "Hello"
//   a[1] = "World"
//   fmt.Println(a)

  // ref: https://maku77.github.io/p/cjosvz3/ (スライスについて)
  // arr[m:n] => 元の配列の「m ～ n-1」の領域の要素を参照可能なスライスを取得する
  names := [4]string{"a", "b", "c", "d"}
  fmt.Println(names)
  // スライス作成
  a := names[0:2] // => ["a", "b"]
  b := names[1:3] // => ["b", "c"]
  fmt.Println(a, b)

  // スライスを変更すると元の値も変更される
  b[0] = "x"
  fmt.Println(a, b)
  fmt.Println(names)

  s := []int{2, 3, 5, 7, 11, 13}

  s = s[1:4]
  fmt.Println(s)

  s = s[:2]
	fmt.Println(s)

  s = s[1:]
  fmt.Println(s)

  // ref: https://qiita.com/tchnkmr/items/10071a53a8bce87b62a3
  // - 配列は固定長、スライスは可変長
  // - 配列はappend()で拡張できない、スライスはできる
  // - 配列は変数への代入でコピーが作れる、スライスはcopy()を使用しないと作れない

  // make: スライスを作成する
  aa := make([]int, 5)
  printSlice("aa", aa)

  var ss []int
  fmt.Println(ss)

  // append works on nil slices.
  ss = append(ss, 0,1,2,3,4)
  fmt.Println(ss)

  pic.Show(Pic)

  fmt.Println(m)

  mm := make(map[string]int)
 	mm["Answer"] = 42
 	fmt.Println(mm)

 	vv, ok := mm["Answer"]
 	fmt.Println(vv, ok)

 	wc.Test(WordCount)

 	hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }
  fmt.Println(hypot(5, 12)) // math.Sqrt(5*5 + 12*12) => 13
  fmt.Println(compute(hypot)) // math.Sqrt(3*3 + 4*4) => 5
  fmt.Println(compute(math.Pow)) // math.Por(3^4) => 81

  pos := adder()
  fmt.Println(pos(1)) // 1
  fmt.Println(pos(2)) // 3
  fmt.Println(pos(3)) // 6
}

func Add(a *int) { //←aはint型のポインタ変数になる。👉を作ったイメージ。
  *a += 1         //aポインタが指している変数に入る数を+1する
  fmt.Println("aの値(関数内)：", *a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s: len=%d, cap=%d, %v\n",	s, len(x), cap(x), x)
}

// Exercise
func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8((x + y) / 2)
		}
	}
	return pic
}

// Exercise
func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	for _, w := range strings.Fields(s) {
		wc[w]++
	}
	return wc
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// クロージャー
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}