package calc

import "testing"

func TestA(t *testing.T) {
  t.Run("正の数", func(t *testing.T) {
    var nums []int = []int{1, 2, 3, 4, 5}
    if !(Summarize(nums) == 15) {
      t.Error(`miss`)
    }
  })
}

