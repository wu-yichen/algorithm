package bubble_sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	test := []struct {
		nums []int
		name string
	}{
		{
			name: "test1",
			nums: func(r int) []int {
				rand.Seed(time.Now().UnixNano())
				var nums []int
				for i := 0; i < 10; i++ {
					num := rand.Intn(r)
					nums = append(nums, num)
				}
				return nums
			}(100),
		},
		{
			name: "test2",
			nums: func(r int) []int {
				rand.Seed(time.Now().UnixNano())
				var nums []int
				for i := 0; i < 10; i++ {
					num := rand.Intn(r)
					nums = append(nums, num)
				}
				return nums
			}(20),
		},
		{
			name: "test3",
			nums: func(r int) []int {
				rand.Seed(time.Now().UnixNano())
				var nums []int
				for i := 0; i < 10; i++ {
					num := rand.Intn(r)
					nums = append(nums, num)
				}
				return nums
			}(10),
		},
		{
			name: "test4",
			nums: func(r int) []int {
				rand.Seed(time.Now().UnixNano())
				var nums []int
				for i := 0; i < 10; i++ {
					num := rand.Intn(r)
					nums = append(nums, num)
				}
				return nums
			}(500),
		},
		{
			name: "test5",
			nums: []int{1, 2, 4, 6, 3},
		},
	}

	for _, value := range test {
		t.Run(value.name, func(t *testing.T) {
			arr := value.nums
			arr2 := make([]int, len(arr))
			copy(arr2, arr)
			bubbleSort(arr)
			sort.Ints(arr2)
			if !reflect.DeepEqual(arr2, arr) {
				t.Error("wrong not equal")
			}
		})
	}
}
