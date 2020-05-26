package test

import (
	"fmt"
	"qp_game_platform/game/204_redmahjong/msg"
	"time"

	msg103 "qp_game_platform/game/103_bairenniuniu/msg"
	"reflect"
	"testing"
)

func SwitchToIndex(mjData int32) int32 {
	return (mjData&0xF0>>4)*9 + (mjData & 0x0F) - 1
}
func SwitchToIndex2(mjData int32) int32 {
	return (mjData&0xF0>>4)/9 + (mjData & 0x0F) + 1
}

func TestISExist(t *testing.T) {
	//var src = []int32{6, 10, 14, 2, 3, 11, 8, 13, 13, 9, 15, 6, 2, 4, 9, 11, 12}
	//var des = []int32{3, 4, 6, 6, 8, 9, 9, 10, 11, 11, 12, 13, 13, 14, 2, 2, 14}
	var src = []int32{3, 3, 4, 4, 5, 6, 6, 7, 7, 8, 9, 9, 11, 11, 12, 13, 13}
	var des = []int32{5, 3, 13, 4, 3, 11, 15, 10, 8, 5, 3, 9, 13, 9, 6, 12, 4}

	fmt.Println("差异:", srcIsExist(des, src))
}
func TestCardMax(t *testing.T) {
	var poker1 = []int32{0x19, 0x1C, 0x1A}
	var poker2 = []int32{0x01, 0x08, 0x06}

	p1 := ReorderPokerA(poker1, 0)
	p2 := ReorderPokerA(poker2, 0)
	fmt.Printf("p1:%v ,p2:%v\n", p1, p2)

	//for k, _ := range p1 {
	//	if GetLogicValue(p1[k]) > GetLogicValue(p2[k]) {
	//		fmt.Printf("%v===true,p1[%v],p2[%v]\n", k,p1[k],p2[k])
	//	} else if GetLogicValue(p1[k]) < GetLogicValue(p2[k]) {
	//		fmt.Printf("%v===false,p1[%v],p2[%v]\n", k,p1[k],p2[k])
	//	}
	//}
	for k := len(p1) - 1; k >= 0; k-- {
		if GetLogicValue(p1[k]) > GetLogicValue(p2[k]) {
			fmt.Printf("%v===true,p1[%v],p2[%v]\n", k, p1[k], p2[k])
		} else if GetLogicValue(p1[k]) < GetLogicValue(p2[k]) {
			fmt.Printf("%v===false,p1[%v],p2[%v]\n", k, p1[k], p2[k])
		}
	}
}

func TestPlanes(t *testing.T) {
	var HavePokers = []int32{36, 53, 39, 7, 23, 24, 8, 56}

	netMold := make([]int32, 0)
	outCard := make([]int32, 0)
	netMold = append(netMold, 8, 7)
	moldNum := conversionPokerType(HavePokers)

	moldNum[3] = make([]int32, 0)
	moldNum[3] = append(moldNum[3], 12, 10, 9, 5, 4, 3)
	outCard = PlanesHandle(5, netMold, moldNum)
	fmt.Printf("mold=%v,出牌 outcard=%v\n", moldNum, outCard)
}

func TestUnary(t *testing.T) {
	cron := NewCronTab()
	result := TernaryOperator(false, cron.f1, cron.f2)
	fmt.Printf("result=%v,%v %T\n", reflect.TypeOf(result).Kind(), reflect.Func, reflect.Func)
}

// 测试麻将最优出牌
func TestOutCard(t *testing.T) {
	//	GetOutCard()
	// GetOptimalCard()
	var cards = map[int32]int32{0x07: 1, 0x08: 1, 0x09: 1, 0x15: 2, 0x19: 2, 0x21: 1, 0x24: 1, 0x27: 1, 0x28: 1, 0x29: 1, 0x31: 2} //胡牌

	// fmt.Println(CheckPingHu(cards))
	replaceCard := func(card int32, tempCards map[int32]int32) map[int32]int32 {
		if tempCards[card] != 0 {
			tempCards[card] --
		}
		tempCards[0x35]++
		return tempCards
	}
	var minFlag int32 = 0x08    // 最多缺7
	MagicMahjong := cards[0x35] //记录癞子
	var OutCardsAsc = make(map[int32]int32, 0)
	for k, _ := range cards {
		if k != 0x35 {
			tempData := replaceCard(k, cards)
			queCount := GetMininum(tempData)
			if minFlag > queCount {
				minFlag = queCount
				OutCardsAsc[minFlag] = k
			}
			cards[k]++
			cards[0x35] = MagicMahjong
			if MagicMahjong == 0 {
				delete(cards, 0x35)
			}
		}
	}

}

// 测试摸牌是否有反应
func TestRespond(t *testing.T) {
	// var cards = map[int32]int32{0x07: 2, 0x08: 2, 0x09: 2, 0x15: 2, 0x19: 2, 0x24: 2, 0x27: 2} //胡牌
	// var cards = map[int32]int32{3:1, 4:2 ,19:1 ,20:1 ,35:1, 37:2 ,38:1 ,40:1, 53:1 }
	//var cards = map[int32]int32{5: 1, 7: 1, 21: 1, 22: 1, 23: 1, 24: 2, 33: 3, 34: 1, 35: 1, 36: 1, 53: 1} // 胡牌
	var cards = map[int32]int32{1: 4, 2: 1, 3: 1, 6: 1, 7: 1, 9: 1, 21: 1, 24: 1, 33: 1, 40: 1, 53: 1}
	data := make([]*msg.DiskMahjong, 0)
	userDiskMjList := &msg.DiskMahjongList{
		Data: data,
	}
	respond := SendMjResponse(0x035, cards, userDiskMjList, true)
	if respond&0x100 != 0 {
		fmt.Printf("%X\n", respond)
	}
	fmt.Printf("%v,%X\n", respond, GetMjResponse(cards))
}

// 测试顺子
func TestSHUNZI(t *testing.T) {
	var cards = []int32{5, 7, 11, 9, 10, 2, 1, 3, 10, 10, 12, 1, 6, 2, 13, 4, 1}
	mold := conversionPokerType(cards)
	outCards := getSingleContinuity(mold, false)
	outC := getOriginData(outCards, true, cards)
	fmt.Println(outCards, outC)
}

// 测试麻将最优出牌
func TestOutCard1(t *testing.T) {
	//	GetOutCard()
	// GetOptimalCard()
	//var cards = map[int32]int32{0x07: 1, 0x08: 1, 0x09: 1, 0x15: 2, 0x19: 2, 0x21: 1, 0x24: 1, 0x27: 1, 0x28: 1, 0x29: 1, 0x31: 2} //胡牌
	//var cards = map[int32]int32{2: 1, 3: 1, 4: 1, 5: 1, 22: 1, 23: 1, 36: 1, 39: 1, 40: 1, 41: 1, 53: 3}
	var cards = map[int32]int32{2: 1, 3: 1, 5: 1, 6: 2, 8: 1, 9: 1, 19: 1, 21: 0, 22: 1, 25: 1, 34: 1, 35: 1, 38: 0, 39: 1, 40: 0, 41: 0, 53: 1}
	for k, num := range cards {
		fmt.Printf(" %v,%X;", num, k)
	}
	fmt.Printf("\n")
	repond := GetOutMjData(cards)
	fmt.Printf("respond:%X -10*4=%v\n", repond, -10*4)
}

// 测试比牌
func TestBrnnCommpre(t *testing.T) {
	var lotteryPoker = make(map[int32]*msg103.Game_S_LotteryPoker)
	fmt.Printf("%T\n", lotteryPoker)
	lotteryPoker[0] = &msg103.Game_S_LotteryPoker{
		//LotteryPoker: []int32{0x02, 0x03, 0x04, 0x05, 0x06},
		LotteryPoker: []int32{0x32, 0x33, 0x34, 0x35, 0x36},
	}
	lotteryPoker[1] = &msg103.Game_S_LotteryPoker{
		LotteryPoker: []int32{0x12, 0x13, 0x14, 0x15, 0x16},
	}
	lotteryPoker[0].PokerType = Client.getPokerType(lotteryPoker[0].LotteryPoker)
	lotteryPoker[1].PokerType = Client.getPokerType(lotteryPoker[1].LotteryPoker)
	fmt.Println("wins:", Client.getWins(lotteryPoker))
	areaCompare := Client.compareType(lotteryPoker[0].PokerType, lotteryPoker[1].PokerType)
	fmt.Printf("pk1=%v,pk2=%v, comprae:%v\n", lotteryPoker[0].PokerType, lotteryPoker[1].PokerType, areaCompare)
	fmt.Println("-10/10", -10/10)
	//winArea := Client.GetWinArea(lotteryPoker)
	//fmt.Println(winArea, len(winArea.LotteryRecord), len(lotteryPoker))
}

// c测试胡牌 测试平胡
func TestPingHu(t *testing.T) {
	//var cards = map[int32]int32{0x05: 2, 0x07: 2, 0x15: 2, 0x16: 2, 0x17: 2, 0x18: 2, 0x24: 1, 0x35: 1} // 胡牌
	var userCards = map[int32]map[int32]int32{0: {
		0x03: 1, 0x04: 1, 0x05: 1, 0x14: 1, 0x15: 1, 0x16: 3, 0x17: 1, 0x18: 1, 0x19: 1, 0x25: 1, 0x26: 1, 0x27: 1,
	}, 1: {
		0x12: 3, 0x15: 3, 0x16: 3, 0x18: 2, 0x19: 2, 0x35: 1,
	},
		2: {0x23: 3, 0x24: 3, 0x25: 3, 0x26: 2, 0x27: 2, 0x28: 1},
		3: {0x05: 1, 0x07: 1, 0x15: 1, 0x16: 1, 0x17: 1, 0x18: 2, 0x21: 3, 0x22: 1, 0x23: 1, 0x24: 1, 0x35: 1},
		4: {0x07: 2, 0x08: 2, 0x09: 2, 0x15: 1, 0x16: 1, 0x17: 1, 0x19: 2, 0x24: 3},
		5: {0x23: 3, 0x24: 2, 0x25: 2, 0x26: 2, 0x27: 2, 0x09: 3},
	}

	//respond := CheckPingHu(userCards[4])
	//
	//fmt.Printf("resp:%v=%X\n", respond, respond)

	for i, v := range userCards {
		resp := CheckPingHu(v)
		fmt.Printf("i=%v,resp:%v=%X\n", i, resp, resp)
	}
}

// 测试fallthrough
func TestFallthrough(t *testing.T) {

	Fallthrough(1)

}

// 测试time.Timer

func TestTimer(t *testing.T) {
	ti := time.NewTimer(time.Second * 1)
	var num int
	go func() {

		for {
			select {
			case <-ti.C:
				num++
				fmt.Println("ti.c..", time.Now().Format("2006-01-02 15:04:05"))
				ti.Reset(time.Second * 1)
			}
			if num == 5 || num == 10 {
				ti.Stop()
			}
		}
	}()

	time.Sleep(time.Second * 8)
	ti.Reset(time.Second * 1)
	time.Sleep(time.Second * 5)
}

func TestAppendSlice(t *testing.T) {

	var cards1 = []int32{25, 36, 35, 38, 24}
	for k, v := range cards1 {
		if v == 24 {
			cards1 = append(cards1[:k], cards1[k+1:]...)
			break
		}
	}
	fmt.Println("card do 1->", cards1)
	//var cards = []int32{25, 36, 35, 38, 24}
	//fmt.Println("card do 1->", cards)
	//cards = cards[1:]
	//fmt.Println("card  do 2:", cards)
	//
	//var userMjData = make(map[int32]int32, 0)
	//var tempData = make(map[int32]int32, 0)
	//var tempData1 = make(map[int32]int32, 0)
	//userMjData[8] = 2
	//userMjData[9] = 3
	//fmt.Println("old ", userMjData)
	//tempData = userMjData
	//for k, v := range userMjData {
	//	tempData1[k] = v
	//}
	//tempData[8] = 1
	//fmt.Println("new ", userMjData, tempData)
	//
	//fmt.Println("now ", userMjData, tempData, tempData1)
}

// 测试抢杠胡
func TestQiangGangeHu(t *testing.T) {
	//var userCards = map[int32]map[int32]int32{0: {
	//	0x03: 1, 0x04: 1, 0x05: 1, 0x14: 1, 0x15: 1, 0x16: 3, 0x17: 1, 0x18: 1, 0x19: 1, 0x25: 1, 0x26: 1, 0x27: 1,
	//}, 1: {
	//	0x12: 3, 0x15: 3, 0x16: 3, 0x18: 2, 0x19: 2, 0x35: 1,
	//},
	//	2: {0x23: 3, 0x24: 3, 0x25: 3, 0x26: 2, 0x27: 2, 0x28: 1},
	//	3: {0x05: 1, 0x07: 1, 0x15: 1, 0x16: 1, 0x17: 1, 0x18: 2, 0x21: 3, 0x22: 1, 0x23: 1, 0x24: 1, 0x35: 1},
	//	4: {0x07: 2, 0x08: 2, 0x09: 2, 0x15: 1, 0x16: 1, 0x17: 1, 0x19: 2, 0x24: 3},
	//	5: {0x23: 3, 0x24: 2, 0x25: 2, 0x26: 2, 0x27: 2, 0x09: 3},
	//}
	var userCards = map[int32]map[int32]int32{0: {
		0x03: 1, 0x04: 1, 0x05: 1, 0x14: 1, 0x15: 1, 0x16: 3, 0x17: 1, 0x18: 1, 0x19: 1, 0x25: 1, 0x27: 1,
	}, 1: {
		0x12: 3, 0x15: 3, 0x16: 3, 0x18: 2, 0x19: 2,
	},
		2: {0x23: 2, 0x18: 2, 0x21: 2, 0x26: 2, 0x03: 2, 0x28: 1, 0x35: 2},
		3: {0x05: 1, 0x07: 1, 0x15: 1, 0x16: 1, 0x17: 1, 0x18: 2, 0x21: 3, 0x22: 1, 0x23: 1, 0x24: 1},
	}
	resp := GetBuGangResponse(0, 0x06, userCards)

	fmt.Println("code:", resp, userCards)
}

// 测试听胡的牌
func TestTindHuCards(t *testing.T) {
	//var userCards = map[int32]map[int32]int32{0: {
	//	0x03: 1, 0x04: 1, 0x05: 1, 0x14: 1, 0x15: 1, 0x16: 3, 0x17: 1, 0x18: 1, 0x19: 1, 0x25: 1, 0x26: 1, 0x27: 1,
	//}, 1: {
	//	0x12: 3, 0x15: 3, 0x16: 3, 0x18: 2, 0x19: 2, 0x35: 1,
	//},
	//	2: {0x23: 3, 0x24: 3, 0x25: 3, 0x26: 2, 0x27: 2, 0x28: 1},
	//	3: {0x05: 1, 0x07: 1, 0x15: 1, 0x16: 1, 0x17: 1, 0x18: 2, 0x21: 3, 0x22: 1, 0x23: 1, 0x24: 1, 0x35: 1},
	//	4: {0x07: 2, 0x08: 2, 0x09: 2, 0x15: 1, 0x16: 1, 0x17: 1, 0x19: 2, 0x24: 3},
	//	5: {0x23: 3, 0x24: 2, 0x25: 2, 0x26: 2, 0x27: 2, 0x09: 3},
	//}
	var userCards = map[int32]map[int32]int32{0: {
		0x03: 1, 0x04: 1, 0x05: 1, 0x14: 1, 0x15: 1, 0x16: 3, 0x17: 1, 0x18: 1, 0x19: 1, 0x25: 1, 0x35: 2,
	}, 1: {
		0x12: 3, 0x15: 3, 0x16: 3, 0x19: 2,
	},
		2: {0x23: 2, 0x18: 2, 0x21: 2, 0x26: 2, 0x03: 2, 0x28: 1, 0x35: 2},
		3: {0x05: 1, 0x07: 1, 0x15: 1, 0x16: 1, 0x17: 1, 0x18: 2, 0x21: 3, 0x22: 1, 0x23: 1, 0x24: 1},
	}
	resp := GetOutMjData(userCards[1])

	fmt.Println("code:", resp, userCards[1])
}

// 测试听胡
func TestTingHu(t *testing.T) {
	var userMj = map[int32]map[int32]int32{0: {
		0x01: 3, 0x03: 3, 0x06: 1, 0x07: 1, 0x08: 1, 0x15: 1, 0x16: 1,
	},
		1: {
			0x01: 3, 0x03: 3, 0x06: 1, 0x07: 1, 0x08: 1, 0x18: 1, 0x19: 1,
		},
		2: {0x01: 2, 0x03: 3, 0x06: 1, 0x07: 1, 0x08: 1, 0x17: 1, 0x18: 1, 0x19: 1},
		3: {0x01: 1, 0x03: 3, 0x06: 1, 0x07: 1, 0x08: 1, 0x23: 1, 0x28: 1, 0x25: 1, 0x35: 1},
		4: {0x04: 3, 0x12: 2, 0x14: 1, 0x17: 1, 0x18: 1, 0x22: 2, 0x27: 1, 0x28: 1, 0x35: 2},
	}

	respCode := GetTingMjData(userMj[4])

	fmt.Printf("code :%v,hex:%X\n", respCode, respCode)

	fmt.Printf("data:%v,%v\n", userMj[4], GetOneMahjongHeap())
}
