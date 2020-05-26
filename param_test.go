package test

import (
	"fmt"
	"qp_game_platform/game/204_redmahjong/global"
	"qp_game_platform/game/204_redmahjong/msg"
	"reflect"
	"testing"
	"time"
)

var pokerNoGhost = []int32{
	//0方块: 2 3 4 5 6 7 8 9 10 J Q K A
	0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x01,
	//1梅花: 2 3 4 5 6 7 8 9 10 J Q K A
	0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x11,
	//2红桃: 2 3 4 5 6 7 8 9 10 J Q K A
	0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x21,
	//3黑桃: 2 3 4 5 6 7 8 9 10 J Q K A
	0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, 0x31,
}

var pokerHaveGhost = []int32{
	//0方块: 2 3 4 5 6 7 8 9 10 J Q K A
	0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x01,
	//1梅花: 2 3 4 5 6 7 8 9 10 J Q K A
	0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x11,
	//2红桃: 2 3 4 5 6 7 8 9 10 J Q K A
	0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x21,
	//3黑桃: 2 3 4 5 6 7 8 9 10 J Q K A
	0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, 0x31,
	//4鬼：大鬼 小鬼 十进制 79 95
	0x4F, 0x5F,
}

var hzMahjong = []int32{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, //万子
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, //万子
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, //万子
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, //万子
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, //索子
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, //索子
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, //索子
	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, //索子
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, //同子
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, //同子
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, //同子
	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, //同子
	0x35, 0x35, 0x35, 0x35, //红中
}

// 获取单次麻将堆
func GetOneMahjongHeap() []int32 {
	mjData := make(map[int32]int32)
	for _, v := range hzMahjong {
		mjData[v] = v
	}
	var mjheap []int32
	for _, v := range mjData {
		mjheap = append(mjheap, v)
	}
	return mjheap
}

func srcIsExist(src, des []int32) []int32 {
	var diff []int32
	for i := 0; i < len(src); i++ {
		for j := 0; j < len(des); j++ {
			if src[i] == des[j] {
				des = append(des[:j], des[j+1:]...)
				break
			}
		}
	}
	diff = des
	return diff
}

// 测试定时器
func TestTicker(t *testing.T) {
	cron := NewCronTab()
	done := make(chan bool)
	go cron.Go(done)
	go func() {
		time.Sleep(time.Second * 5)
		fmt.Printf("sleep\n")
		cron.cronTime.Stop()
		fmt.Printf("关闭定时器 time=%v\n", time.Now().String())
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-cron.cronTime.C:
			fmt.Println("Current time: ", t)
		}
	}
}

type CronTab struct {
	cronTime *time.Ticker
}

func (c *CronTab) Go(done chan bool) {
	time.Sleep(10 * time.Second)
	done <- true
}
func NewCronTab() *CronTab {
	return &CronTab{
		cronTime: time.NewTicker(time.Second * 2),
	}
}
func ExampleNewTicker() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}

const (
	//花色掩码
	maskColor = 0xF0
	//数值掩码
	maskValue = 0x0F
)

//比大小 K最大
func CompareK(poker int32, poker1 int32) bool {

	if GetValue(poker) > GetValue(poker1) {
		return true
	} else if GetValue(poker) == GetValue(poker1) {

		if GetColor(poker) > GetColor(poker1) {
			return false
		} else {
			return true
		}

	} else {
		return false
	}

}

//比大小 A最大
func CompareA(poker int32, poker1 int32) bool {

	if GetValue(poker) == 0x01 {
		poker = 0x0e
	}
	if GetValue(poker1) == 0x01 {
		poker1 = 0x0e
	}
	return CompareK(poker, poker1)
}

func GetColor(poker int32) int32 {
	return (poker & maskColor) / 16
}

func GetValue(poker int32) int32 {
	return poker & maskValue
}

func ReorderPokerA(pokers []int32, number int) []int32 {

	if number == len(pokers) {
		return pokers
	}

	for i := number + 1; i < len(pokers); i++ {

		if CompareA(pokers[number], pokers[i]) {
			pokers[number], pokers[i] = pokers[i], pokers[number]
		}

	}

	return ReorderPokerA(pokers, number+1)

}

//获取逻辑值
func GetLogicValue(poker int32) int32 {
	if GetValue(poker) == 0x01 {
		return 0x0E
	} else {
		return GetValue(poker)
	}
}

// 获取飞机
func PlanesHandle(cardType int, nearestCards []int32, handCards map[int32][]int32) []int32 {
	outCard := make([]int32, 0)
	ContinuityNum := make([]int32, 0)
	if len(handCards[3]) >= len(nearestCards) {
		for _, card := range handCards[3] {
			if GetLogicValue(card) > GetLogicValue(nearestCards[0]) {
				outCard = append(outCard, card)
			}
		}
	}
	for i := len(outCard) - 1; i > 0; i-- {
		if outCard[i-1]-1 == outCard[i] {
			if len(ContinuityNum) > 0 {
				if ContinuityNum[len(ContinuityNum)-1] != outCard[i] {
					ContinuityNum = append(ContinuityNum, outCard[i])
				}
				if ContinuityNum[len(ContinuityNum)-1] != outCard[i-1] {
					ContinuityNum = append(ContinuityNum, outCard[i-1])
				}
			} else {
				ContinuityNum = append(ContinuityNum, outCard[i], outCard[i-1])
			}
		} else {
			if len(ContinuityNum) < 2 {
				ContinuityNum = make([]int32, 0)
			}
		}
	}
	if len(ContinuityNum) >= len(nearestCards) {
		n := len(ContinuityNum) - len(nearestCards)
		outCard = ContinuityNum[:len(ContinuityNum)-n]
	} else {
		outCard = make([]int32, 0)
	}
	if len(outCard) > 0 {
		switch cardType {
		case 5: //三不带
		case 6: //三带一
			if len(handCards[1]) >= len(outCard) {
				outCard = append(outCard, handCards[1][:len(outCard)]...)
			}
		case 7: //三带对
			if len(handCards[2]) >= len(outCard) {
				outCard = append(outCard, handCards[2][:len(outCard)]...)
			}
		}
	}
	return outCard
}

// 三目运算符
func TernaryOperator(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		switch reflect.ValueOf(trueVal).Type().String() {
		case "func()":
			runFunc := trueVal.(func())
			runFunc()
		}
		return trueVal
	}
	switch reflect.ValueOf(falseVal).Type().String() {
	case "func()":
		runFunc := falseVal.(func())
		runFunc()
	}
	return falseVal
}

func (l *CronTab) f1() {
	fmt.Printf("true 1\n")
}

func (l *CronTab) f2(i int32) {
	fmt.Printf("false 1\n")
}
func SwitchToIndex1(mjData int32) int32 {
	return (mjData&0xF0>>4)*9 + (mjData & 0x0F) - 1
}
func CheckPingHu1(userMahjong map[int32]int32) (int32, int32) {
	var data [31]int32
	for k, v := range userMahjong {
		if k == global.MagicMahjong {
			continue
		}

		data[SwitchToIndex1(k)] = v
	}
	//定义变量
	var i, queCount int32
	var jiangFlag bool
	for i < 31 {
		if data[i] == 0 {
			i++
			continue
		}

		//数值
		value := i + 1%9
		if value == 0 {
			value = 9
		}

		if data[i] == 2 {
			if jiangFlag {
				queCount++
			} else {
				jiangFlag = true
			}
			data[i] = 0
			continue
		}
		if data[i] == 3 {
			data[i] = 0
			continue
		}
		for data[i] > 0 {
			//判断是否是顺子
			if value <= 0x09-2 && data[i+1] > 0 && data[i+2] > 0 {
				for j := 0; j < 3; j++ {
					data[i+int32(j)]--
				}
				continue
			}

			switch data[i] {
			case 1:
				if (value == 0x08 && data[i+1] > 0) || data[i+1] > 0 || data[i+2] > 0 {
					if data[i+1] != 0 {
						data[i+1] --
					}
					if value != 0x08 && data[i+2] != 0 {
						data[i+2]--
					}
					queCount++
				} else {
					if jiangFlag {
						queCount += 2
					} else {
						jiangFlag = true
						queCount++
					}
				}
			case 4:
				if jiangFlag {
					queCount += 2
				} else {
					queCount++
				}
			}
			data[i] = 0
		}
	}
	code := int32(0)
	if queCount == 0 && userMahjong[global.MagicMahjong] != 1 || queCount > 0 && queCount == userMahjong[global.MagicMahjong] {
		code |= global.CHR_PING_HU
		code |= global.WIK_HU

		if queCount > 0 {
			code |= global.CHR_MAGIC
		}
	}
	return code, queCount
}

func GetOptimalCard() int32 {
	var cards1 = map[int32]int32{0x35: 3, 0x07: 1, 0x09: 1, 0x15: 3, 0x19: 1, 0x24: 1, 0x26: 1, 0x27: 1, 0x28: 1, 0x29: 1} //胡牌
	fmt.Println(CheckPingHu(cards1))
	var cards = map[int32]int32{0x35: 2, 0x07: 1, 0x09: 1, 0x15: 3, 0x19: 1, 0x24: 1, 0x26: 1, 0x27: 1, 0x28: 1, 0x29: 1, 0x031: 1} //胡牌
	replaceCard := func(card int32, tempCards map[int32]int32) map[int32]int32 {
		if tempCards[card] != 0 {
			tempCards[card] --
		}
		tempCards[0x35]++
		return tempCards
	}
	MagicMahjong := cards[0x35]
	for k, _ := range cards {
		if k != 0x35 {
			code, queCount := CheckPingHu2(replaceCard(k, cards))
			cards[k]++
			cards[0x35] = MagicMahjong
			if MagicMahjong == 0 {
				delete(cards, 0x35)
			}
			fmt.Printf("%v-----%v 缺的数:%v 出的牌:%X\n", cards, code, queCount, k)
		}
	}
	return 1
}

func GetOutCard() {
	// var cards = map[int32]int32{0x01: 4, 0x03: 2, 0x08: 2,  0x019: 2, 0x35: 1}

	var cards = map[int32]int32{0x03: 3, 0x08: 2, 0x019: 2, 0x35: 1}
	fmt.Println(CheckPingHu(cards))
	return
	//var tempCard, tempNum int32
	var MagicMahjong int32
	MagicMahjong = cards[0x35]
	replaceCard := func(card int32, Tempcards map[int32]int32) map[int32]int32 {
		if Tempcards[card] != 0 {
			Tempcards[card] --
		}
		Tempcards[0x35]++
		return Tempcards
	}
	for k, _ := range cards {
		if k != 0x35 {
			code, queCount := CheckPingHu2(replaceCard(k, cards))
			cards[k]++
			cards[0x35] = MagicMahjong
			if MagicMahjong == 0 {
				delete(cards, 0x35)
			}
			fmt.Print(cards, "-----", code, " 缺的数:", queCount, "出的牌:", k, "\n")
		} else {
			code, queCount := CheckPingHu2(replaceCard(k, cards))
			fmt.Print(cards, "else -----", code, " 缺的数:", queCount, "\n")
		}
	}
}

//平胡
func CheckPingHu2(userMahjong map[int32]int32) (int32, int32) {
	var data [31]int32
	for k, v := range userMahjong {
		if k == global.MagicMahjong {
			continue
		}

		data[SwitchToIndex(k)] = v
	}

	var code, Minimum int32
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

		if queCount == 0 && userMahjong[global.MagicMahjong] != 1 || queCount > 0 && queCount == userMahjong[global.MagicMahjong] {
			code |= global.CHR_PING_HU
			code |= global.WIK_HU

			if queCount > 0 {
				code |= global.CHR_MAGIC
			}
			break
		}
	}

	return code, Minimum
}

// 顺子 参数5 7 11 9 10 2 1 3 10 10 12 1 6 2 13 4 1
// 获取单顺子,true 拆炸弹，false 不拆
func getSingleContinuity(poker map[int32][]int32, isBomb bool) []int32 {
	// 不拆炸弹
	if !isBomb {
		delete(poker, 4)
	}
	card := pokerMap2Sort(poker)
	var pokerSlice []int32
	var dataMap = make(map[int32]bool)
	var count int
	keys := card
	for i := 0; i < len(keys)-1; i++ {
		values := keys[i+1]
		if (keys[i] < values) && (keys[i] == values-1) {
			count++
			dataMap[keys[i]] = true
			dataMap[values] = true
		} else {
			if count < 4 {
				count = 0
				dataMap = make(map[int32]bool, 0)
			} else {
				break
			}
		}
	}
	if len(dataMap) >= 5 {
		for k, _ := range dataMap {
			pokerSlice = append(pokerSlice, k)
		}
	}
	return pokerSlice
}

func conversionPokerType(pokers []int32) map[int32][]int32 {
	var mold = make(map[int32]int32, 0) //牌 数量

	for _, card := range pokers {
		mold[card&0x0F]++
	}

	var moldNum = make(map[int32][]int32, 0) //数量 牌

	for card, num := range mold {
		var cards = moldNum[num]
		moldNum[num] = append(cards, card)
	}

	//排序扑克
	for k, v := range moldNum {
		for i := 0; i < len(v)-1; i++ {
			for j := i + 1; j < len(v); j++ {
				if GetLogicValue(v[i]) < GetLogicValue(v[j]) {
					moldNum[k][i], moldNum[k][j] = moldNum[k][j], moldNum[k][i]
				}
			}
		}
	}

	return moldNum
}

// 获取逻辑数，以便判断顺子
func pokerMap2Sort(pokers map[int32][]int32) []int32 {
	var card []int32
	for k, v := range pokers {
		for i := range v {
			if pokers[k][i] == 0x01 {
				pokers[k][i] = 0x0E
			}
			if pokers[k][i] != 0x02 && pokers[k][i] != 0x0F {
				card = append(card, pokers[k][i])
			}
		}
	}
	for i := 0; i < len(card)-1; i++ {
		for j := i + 1; j < len(card); j++ {
			if card[i] > card[j] {
				card[i], card[j] = card[j], card[i]
			}
		}
	}
	return card
}

// 逻辑值转换为源数据
func getOriginData(card []int32, isSingle bool, pokers []int32) []int32 {
	pokerVal := make([]int32, 0)
	size := len(pokers)
	for j := 0; j < len(card); j++ {
		for i := 0; i < size; i++ {
			if GetLogicValue(pokers[i]) == GetLogicValue(card[j]) {
				if isSingle {
					pokerVal = append(pokerVal, pokers[i])
					break
				} else {
					pokerVal = append(pokerVal, pokers[i])
				}
			}
		}
	}
	return pokerVal
}

func GetMininum(userMahjong map[int32]int32) int32 {
	var data [31]int32
	for k, v := range userMahjong {
		if k == global.MagicMahjong {
			continue
		}
		data[SwitchToIndex(k)] = v
	}
	fmt.Println("data:", userMahjong, data)
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

//摸牌是否有响应
func SendMjResponse(mjData int32, userMahjong map[int32]int32, userDiskMjList *msg.DiskMahjongList, first bool) int32 {
	//操作码
	response := int32(0)

	//暗杠检测
	for k, v := range userMahjong {
		if v != 4 {
			continue
		}

		if !first && mjData != k {
			continue
		}

		response |= global.WIK_AN_GANG
		break
	}

	if response == 0 && userDiskMjList != nil {
		//补杠检测
		for _, v := range userDiskMjList.Data {
			if v.Code != global.WIK_PENG {
				continue
			}

			if mjData != v.Data {
				continue
			}

			response |= global.WIK_BU_GANG
			break
		}
	}

	//红中胡检测
	if userMahjong[global.MagicMahjong] == 4 {
		response |= global.CHR_SI_HONG_ZHONG
		return response
	}

	//七对胡检测
	code := CheckQiDuiHu(userMahjong)
	if code != 0 {
		response |= code
		return response
	}

	//平胡检测
	response |= CheckPingHu(userMahjong)
	return response
}

//摸牌是否有响应
func GetMjResponse(userMahjong map[int32]int32) int32 {
	//操作码
	response := int32(0)

	//红中胡检测
	if userMahjong[global.MagicMahjong] == 4 {
		response |= global.CHR_SI_HONG_ZHONG
		return response
	}

	//七对胡检测
	code := CheckQiDuiHu(userMahjong)
	if code != 0 {
		response |= code
		return response
	}

	//平胡检测
	response |= CheckPingHu(userMahjong)
	return response
}

//七对胡
func CheckQiDuiHu(userMahjong map[int32]int32) int32 {
	count := int32(0)
	queCount := int32(0)
	for k, v := range userMahjong {
		count += v
		if k == global.MagicMahjong {
			continue
		}

		if v != 2 && v != 4 {
			queCount++
		}
	}

	code := int32(0)
	if count == 14 && (queCount == 0 || queCount == userMahjong[global.MagicMahjong]) {
		code |= global.CHR_QI_DUI
		code |= global.WIK_HU

		if queCount > 0 {
			code |= global.CHR_MAGIC
		}
	}

	return code
}

//平胡
func CheckPingHu(userMahjong map[int32]int32) int32 {
	var data [31]int32
	for k, v := range userMahjong {
		if k == global.MagicMahjong {
			continue
		}

		data[SwitchToIndex(k)] = v
	}

	var code int32
	for k, v := range data {
		//取出一对将
		tempData := data
		var queCount int32
		switch v {
		case 1:
			queCount++
			tempData[k]--
		case 2, 3, 4:
			tempData[k] -= 2
		default:
			continue
		}

		for i := 0; i < 31; {
			if tempData[i] == 0 {
				i++
				continue
			}

			//数值
			value := i%9 + 1
			for tempData[i] > 0 {
				//判断是否是刻字
				if tempData[i] == 3 {
					tempData[i] = 0
					continue
				}

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
				case 4:
					queCount += 2
				}

				tempData[i] = 0
			}
		}

		if queCount == 0 && userMahjong[global.MagicMahjong] != 1 || queCount > 0 && queCount == userMahjong[global.MagicMahjong] {
			code |= global.CHR_PING_HU
			code |= global.WIK_HU

			if queCount > 0 {
				code |= global.CHR_MAGIC
			}
			break
		}
	}

	return code
}

//平胡
func CheckPingHu3(userMahjong map[int32]int32) int32 {
	var data [31]int32
	var IsJang = false
	var card int32
	for k, v := range userMahjong {
		if k == global.MagicMahjong {
			continue
		}
		if v == 2 && !IsJang {
			IsJang = true
			card = k
		}
		data[SwitchToIndex(k)] = v
	}
	var code int32
	for k, v := range data {
		var queCount int32
		switch v {
		case 1:
			queCount++
		case 2:
		default:
			continue
		}
		//fmt.Printf(" do 1  que=%v v=%v  data:%v \n", queCount, v, data)
		//取出一对将
		tempData := data
		if !IsJang {
			tempData[k] = 0
		} else {
			tempData[card] = 0
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
				//fmt.Printf("for 2 que=%v i=%v  data:%v \n", queCount, i, tempData)
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
		if queCount == 0 && userMahjong[global.MagicMahjong] != 1 || queCount > 0 && queCount == userMahjong[global.MagicMahjong] {
			code |= global.CHR_PING_HU
			code |= global.WIK_HU

			if queCount > 0 {
				code |= global.CHR_MAGIC
			}
			fmt.Println("========================")
			break
		}
	}

	return code
}

//
func Fallthrough(num int) {
	switch num {
	case 0:
		fallthrough
	case 1:
		fmt.Println(" case <----1")
		fallthrough
	case 2:
		fmt.Println(" case <----2")
	case 3:

		fmt.Println(" case <----3")
	case 4:

		fmt.Println(" case <----4")

	}
}

//补扛是否有响应
func GetBuGangResponse(chairID int32, mjData int32, userListMahjong map[int32]map[int32]int32) map[int32]int32 {
	//操作码
	response := make(map[int32]int32)

	for k, v := range userListMahjong {
		if k == chairID {
			continue
		}
		v[mjData]++
		//胡牌检测
		code := CheckPingHu(v)
		if code != 0 {
			response[k] |= code
		}
		//红中胡检测
		if v[global.MagicMahjong] == 4 {
			response[k] |= global.CHR_SI_HONG_ZHONG
		}
		//七对胡检测
		code = CheckQiDuiHu(v)
		if code != 0 {
			response[k] |= code
			return response
		}
		v[mjData]--
		if v[mjData] == 0 {
			delete(v, mjData)
		}
	}

	return response
}

//听胡
func CheckTingHu(userMahjong map[int32]int32) map[int32][]int32 {
	var tindCards = make(map[int32][]int32)
	var data [31]int32
	for k, v := range userMahjong {
		if k == global.MagicMahjong {
			continue
		}

		data[SwitchToIndex(k)] = v
	}

	var code int32
	for k, v := range data {
		var response = make([]int32, 0)
		//取出一对将
		tempData := data
		var queCount int32
		switch v {
		case 1:
			queCount++
			response = append(response, int32(k))
			tempData[k]--
		case 2, 3, 4:
			tempData[k] -= 2
		default:
			continue
		}

		for i := 0; i < 31; {
			if tempData[i] == 0 {
				i++
				continue
			}

			//数值
			value := i%9 + 1
			for tempData[i] > 0 {
				//判断是否是刻字
				if tempData[i] == 3 {
					tempData[i] = 0
					response = append(response, int32(i))
					continue
				}

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
							response = append(response, int32(i))
						}
						if value != 0x08 && tempData[i+2] != 0 {
							tempData[i+2]--
							response = append(response, int32(i))
						}
						queCount++
					} else {
						queCount += 2
					}
				case 2:
					queCount++
					response = append(response, int32(i))
				case 4:
					queCount += 2
				}

				tempData[i] = 0
			}
		}

		if queCount == 0 && userMahjong[global.MagicMahjong] != 1 || queCount > 0 && queCount == userMahjong[global.MagicMahjong] {
			code |= global.CHR_PING_HU
			code |= global.WIK_HU
			fmt.Printf("hu:%v\n", response)
			tindCards[int32(k)] = response
			if queCount > 0 {
				code |= global.CHR_MAGIC
			}
			break
		}
	}

	return tindCards
}

// 获取最佳出牌数据，出牌后缺的听胡数最少
func GetOutMjData(MjData map[int32]int32) int32 {
	replaceCard := func(card int32, tempCards map[int32]int32) map[int32]int32 {
		if tempCards[card] != 0 {
			tempCards[card] --
		}
		tempCards[0x35]++
		return tempCards
	}
	var minFlag int32 = 0x0F // 再差的牌最多缺14张牌
	var userMjData = make(map[int32]int32)
	for k, v := range MjData {
		userMjData[k] = v
	}
	MagicMahjong := userMjData[global.MagicMahjong] //记录癞子
	var OutCardsAsc = make(map[int32]int32, 0)
	for k, num := range userMjData {
		if num <= 0 {
			delete(userMjData, k)
			continue
		}
		if k != 0x35 {
			queCount := GetMininum(replaceCard(k, userMjData))
			if minFlag > queCount {
				minFlag = queCount
				OutCardsAsc[minFlag] = k
			}
			userMjData[k]++
			userMjData[global.MagicMahjong] = MagicMahjong
			if MagicMahjong == 0 {
				delete(userMjData, global.MagicMahjong)
			}
		}
	}
	return OutCardsAsc[minFlag]
}

// 获取最佳出牌数据，出牌后缺的听胡数最少
func GetTingMjData(mjData map[int32]int32) map[int32][]int32 {
	replaceCard := func(card int32, tempCards map[int32]int32) {
		if tempCards[card] != 0 {
			tempCards[card] --
		}
		tempCards[0x35]++
	}
	userMjData := make(map[int32]int32)
	for k, v := range mjData {
		userMjData[k] = v
	}
	MagicMahjong := userMjData[global.MagicMahjong] //记录癞子
	var OutCardsAsc = make(map[int32]int32, 0)
	var tingHuCards = make(map[int32][]int32, 0)
	for k, num := range userMjData {
		if num <= 0 {
			delete(userMjData, k)
			continue
		}
		if k != 0x35 {
			replaceCard(k, userMjData)
			queCount := GetMininum(userMjData)
			if queCount-MagicMahjong <= 1 {
				OutCardsAsc[k] = num
			}
			userMjData[k]++
			userMjData[global.MagicMahjong] = MagicMahjong
			if MagicMahjong == 0 {
				delete(userMjData, global.MagicMahjong)
			}
		}
	}
	heapMj := GetOneMahjongHeap()
	for k, _ := range OutCardsAsc {
		userMjData[k]--
		if userMjData[k] == 0 {
			delete(userMjData, k)
		}
		for _, v := range heapMj {
			if CheckTing(v, userMjData) {
				tingHuCards[k] = append(tingHuCards[k], v)
			}
		}
		userMjData[k]++
	}
	return tingHuCards
}

//获取最低需要癞子数
func GetMininum2(userMahjong map[int32]int32) int32 {
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

//判断是否听牌
func CheckTing(mjData int32, userMahjong map[int32]int32) bool {
	//红中胡检测
	if userMahjong[global.MagicMahjong] >= 3 {
		return true
	}
	var temp = make(map[int32]int32)
	for k, v := range userMahjong {
		temp[k] = v
	}
	temp[mjData]++

	//七对胡检测
	if CheckQiDuiHu(temp)&global.WIK_HU != 0 {
		return true
	}

	//平胡检测
	if CheckPingHu(temp)&global.WIK_HU != 0 {
		return true
	}

	return false
}
