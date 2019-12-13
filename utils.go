package controller

import (
	"errors"
	"fmt"
	"reflect"
)

type TradeQueue struct {
	maxSize int
	TradMap []map[int]string
	tail    int
	head    int
	next    *TradeQueue
	pre     *TradeQueue
}

const MaxSize = 1024 + 1

// 队列
func (q *TradeQueue) Push(tradMap map[int]string) (err error) {
	if q.Full() {
		return errors.New("queue full")
	}
	q.TradMap[q.tail] = tradMap
	q.tail = (q.tail + 1 + q.maxSize) % q.maxSize
	return
}

func (q *TradeQueue) Pop() (val map[int]string, err error) {
	if q.Empty() {
		err = errors.New("queue is empty")
		return
	}
	val = q.TradMap[q.head]
	q.head = (q.head + 1 + q.maxSize) % q.maxSize
	return
}

func (q *TradeQueue) List() {
	size := q.Size()
	if size == 0 {
		fmt.Println("size is nil")
		return
	}
	head := q.head
	for i := 0; i < size; i++ {
		fmt.Printf("queue[%v]=%v\n", head, q.TradMap[head])
		head = (head + 1) % q.maxSize
	}
}

func (q *TradeQueue) Full() bool {
	return (q.tail+1)%q.maxSize == q.head
}
func (q *TradeQueue) Empty() bool {
	return q.tail == q.head
}
func (q *TradeQueue) Size() int {
	return (q.tail + q.maxSize - q.head) % q.maxSize
}
func NewTradeQueue() *TradeQueue {
	return &TradeQueue{
		maxSize: MaxSize, //MaxSize,
		head:    0,
		tail:    0,
		TradMap: make([]map[int]string, MaxSize),
	}
}

// 链表环形
func InsertQueue(head *TradeQueue, newTrade *TradeQueue) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	newTrade.next = temp.next
	newTrade.pre = temp
	if temp.next != nil {
		temp.next.pre = newTrade
	}
	temp.next = newTrade
}
// 判断参数是否为空
func empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}
