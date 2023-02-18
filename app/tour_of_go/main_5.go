package main

import (
    "fmt"
    "time"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("sleep....")
	time.Sleep(1 * time.Second)
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}


func main() {
	go say("world")
	say("hello")
  // time.sleepしないとgoルーチンの実行完了よりも先にメイン関数の処理が完了してしまう。
	time.Sleep(500 * time.Millisecond)

	s := []int{7, 2, 8, -9, 4, 0}
  c := make(chan int)
  fmt.Println(len(s))
  go sum(s[:len(s)/2], c)
  go sum(s[len(s)/2:], c)
  x, y := <-c, <-c
  // x, yの計算が終わったら下記が実行される
  fmt.Println(x, y, x+y)

  // 第二引数はバッファの長さ => 使用できるチャネル数
  ch := make(chan int, 2)
  ch <- 1
  ch <- 2
  // ch <- 3 // ここを実行するとエラーになる
  fmt.Println(<-ch, <-ch)

  cc := make(chan int, 10)
  go fibonacci(cap(cc), cc)
  // チャネルが閉じられるまで、チャネルから値を繰り返し受信し続けます。
  for i := range cc {
    fmt.Println(i)
  }

  // select
  tick := time.Tick(100 * time.Millisecond)
  boom := time.After(500 * time.Millisecond)
  for {
    select {
    case <-tick:
      fmt.Println("tick.")
    case <-boom:
      fmt.Println("BOOM!")
      return
    default:
      fmt.Println("    .")
      time.Sleep(50 * time.Millisecond)
    }
  }
}
