package controller

import (
	"math/rand"
	"testing"
	"time"
)


var Method = make([]int32, 0)

func init(){
	Method = append(Method, 1, 2, 4, 7)
}

// 斐波拉契方法
func GetFloorWay(n int) int32 {
	if n < 0 {
		return 0
	}
	if n < len(Method) {
		return Method[n]
	}
	return GetFloorWay(n-1) + GetFloorWay(n-2) + GetFloorWay(n-3) + GetFloorWay(n-4)
}

// 测试斐波拉契 <上楼梯案例>
func TestFeiBoLaqi(t *testing.T) {
	var step = 10
	fmt.Println("method:", Method)
	for i := 4; i < step; i++ {
		Method = append(Method, GetFloorWay(i))
	}
	fmt.Println("step:", step, "Method:", Method)
}


func TestTradeQueue(t *testing.T) {
	queue := NewTradeQueue()
	queue.Push(map[int]string{1: "adas"})
	queue.Push(map[int]string{2: "123"})
	queue.Push(map[int]string{3: "abc"})

	err := queue.Push(map[int]string{4: "bn"})
	if err != nil {
		t.Log(err)
	}
	queue.List()
	val, err := queue.Pop()
	if err != nil {
		t.Log(err)
	}
	t.Log(val)
	queue.List()
}

// 选择排序 200000万的数据量运行47s左右
func TestSelectSortArray(t *testing.T) {
	var arr []int
	var maxSize = 200000
	nowTime1 := time.Now().Unix()
	rand.Seed(time.Now().Unix())
	for i := 0; i < maxSize; i++ {
		n := rand.Int() % 100
		arr = append(arr, n)
	}
	t.Log("size:", len(arr), "select 原数据", arr)

	func(array []int) {
		for i := 0; i < len(array); i++ {
			max := array[i]
			maxIndex := i
			for j := i + 1; j < len(array); j++ {
				if array[j] < max {
					max = array[j]
					maxIndex = j
				}
			}
			if maxIndex != i {
				array[maxIndex], array[i] = arr[i], array[maxIndex]
			}
		}
	}(arr)
	nowTime2 := time.Now().Unix()
	t.Log("修改后 :", arr, "耗时:", nowTime2-nowTime1, "size:", len(arr))
}

// 插入排序  200000万的数据量运行13s左右
func TestInsertArray(t *testing.T) {
	var arr []int
	var maxSize = 200000
	nowTime1 := time.Now().Unix()
	rand.Seed(time.Now().Unix())
	for i := 0; i < maxSize; i++ {
		n := rand.Int() % 100
		arr = append(arr, n)
	}
	t.Log("size:", len(arr), "insert 原数据", arr)
	func(array []int) {
		for i := 1; i < len(array); i++ {
			insertVal := array[i]
			insertIndex := i - 1
			for insertIndex >= 0 && array[insertIndex] > insertVal {
				array[insertIndex+1] = array[insertIndex]
				insertIndex --
			}
			if insertIndex+1 != i {
				array[insertIndex+1] = insertVal
			}
		}

	}(arr)
	nowTime2 := time.Now().Unix()
	t.Log("修改后 :", arr, "耗时:", nowTime2-nowTime1, "size:", len(arr))
}

// 快速排序  200000万的数据量运行1s内
func TestQuickSort(t *testing.T) {
	var arr []int
	var maxSize = 11
	nowTime1 := time.Now().Unix()
	rand.Seed(time.Now().Unix())
	for i := 0; i < maxSize; i++ {
		n := rand.Int() % 100
		arr = append(arr, n)
	}
	t.Log("size:", len(arr), "quick 原数据", arr)
	quickSort(0, len(arr)-1, arr)
	nowTime2 := time.Now().Unix()
	t.Log("修改后 :", arr, "耗时:", nowTime2-nowTime1, "size:", len(arr))
}

// 快排使用递归
func quickSort(front, rear int, array []int) {
	left := front
	right := rear
	pivot := array[(front+rear)/2]
	for left < right {
		for array[left] > pivot {
			left++
		}
		for array[right] < pivot {
			right--
		}
		if left >= right {
			break
		}
		array[left], array[right] = array[right], array[left]
		if array[right] == pivot {
			left++
		}
		if array[left] == pivot {
			right--
		}
	}
	if left == right {
		left++
		right--
	}
	if rear > left {
		quickSort(left, rear, array)
	}
	if front < right {
		quickSort(front, right, array)
	}
}
