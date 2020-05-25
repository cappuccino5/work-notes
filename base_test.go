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



// 测试听胡
func TestTingHu(t *testing.T) {
	var userMj = map[int32]map[int32]int32{0: {
		0x01: 3, 0x03: 3, 0x06: 1, 0x07: 1, 0x08: 1, 0x15: 1,0x16: 1,
	},
		1: {
			0x01: 3, 0x03: 3, 0x06: 1, 0x07: 1, 0x08: 1, 0x18: 1,0x19: 1,
		},
		2: {0x01: 2, 0x03: 3, 0x06: 1, 0x07: 1, 0x08: 1,0x17:1, 0x18: 1,0x19: 1,}}

	respCode := GetOutMjData(userMj[2])
	fmt.Printf("code :%v,hex:%X\n", respCode, respCode)
}

// 获取最佳出牌数据，出牌后缺的听胡数最少
func GetOutMjData(userMjData map[int32]int32) map[int32][]int32 {
	replaceCard := func(card int32, tempCards map[int32]int32) map[int32]int32 {
		if tempCards[card] != 0 {
			tempCards[card] --
		}
		tempCards[0x35]++
		return tempCards
	}
	MagicMahjong := userMjData[global.MagicMahjong] //记录癞子
	var OutCardsAsc = make(map[int32]int32, 0)
	var array []int32
	var tingHuCards = make(map[int32][]int32, 0)
	for k, num := range userMjData {
		if num <= 0 {
			delete(userMjData, k)
			continue
		}
		if k != 0x35 {
			queCount := GetMininum(replaceCard(k, userMjData))
			if queCount == 1 {
				OutCardsAsc[k] = num
				array = append(array, k)
			}
			userMjData[k]++
			userMjData[global.MagicMahjong] = MagicMahjong
			if MagicMahjong == 0 {
				delete(userMjData, global.MagicMahjong)
			}
		}
	}
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[j], array[i] = array[i], array[j]
			}
		}
	}
	fmt.Println("array:", array)
	for k, num := range OutCardsAsc {
		for i := 1; i < len(array); i++ {
			if array[i] == array[i-1]+1 && num >= 2 {
				if array[i-1]&0x0F <= 7 {
					tingHuCards[k] = append(tingHuCards[k], array[i]+1, array[i-1]-1)
				} else if array[i]&0x0F > 7 {
					tingHuCards[k] = append(tingHuCards[k], array[i-1]-1)
				}
			} else if array[i] != k && num == 1 && OutCardsAsc[array[i]] == 1 {
				tingHuCards[k] = append(tingHuCards[k], array[i])
			}
		}
	}
	return tingHuCards
}

//获取最低需要癞子数
func GetMininum(userMahjong map[int32]int32) int32 {
	var data [31]int32
	for k, v := range userMahjong {
		if k == global.MagicMahjong {
			continue
		}

		data[SwitchToIndex(k)] = v
	}

	var Minimum int32
	for k, v := range data {
		var queCount int32
		//取出一对将
		tempData := data
		switch v {
		case 0:
			queCount += 2
		case 1:
			queCount++
			tempData[k] -= 1
		case 2:
			tempData[k] -= 2
		case 3:
			tempData[k] -= 2
		case 4:
			tempData[k] -= 2
		}

		for i := 0; i < 31; {
			if tempData[i] == 0 {
				i++
				continue
			}

			//数值
			value := i%9 + 1
			for tempData[i] > 0 {
				//判断是否是顺子
				if value <= 0x09-2 && tempData[i+1] > 0 && tempData[i+2] > 0 {
					for j := 0; j < 3; j++ {
						tempData[i+j]--
					}
					continue
				}

				switch tempData[i] {
				case 1:
					if value <= 0x08 && ((value == 0x08 && tempData[i+1] > 0) || tempData[i+1] > 0 || tempData[i+2] > 0) {
						if tempData[i+1] != 0 {
							tempData[i+1] --
						}
						if value != 0x08 && tempData[i+2] != 0 {
							tempData[i+2]--
						}
						queCount++
					} else {
						queCount += 2
					}
				case 2:
					queCount++
				case 3:
				case 4:
					queCount += 2
				}

				tempData[i] = 0
			}
		}

		if Minimum == 0 && queCount != 0 {
			Minimum = queCount
		}
		if queCount < Minimum {
			Minimum = queCount
		}
	}

	return Minimum
}
