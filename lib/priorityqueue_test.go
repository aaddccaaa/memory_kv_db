package lib

import (
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	// 创建一个大根堆
	pq := NewPriorityQueue([]DataItem{})

	// 添加一些测试数据
	data := []DataItem{
		{Key: "key1", Value: 10},
		{Key: "key2", Value: 5},
		{Key: "key3", Value: 20},
		{Key: "key4", Value: 15},
		{Key: "key5", Value: 25},
		{Key: "key6", Value: 30},
		{Key: "key7", Value: 8},
	}

	// 将数据添加到大根堆中
	for _, item := range data {
		pq.Push(item)
	}

	// 验证堆顶元素是否为最大值
	expectedMaxValue := 30
	if pq.q[0].Value != expectedMaxValue {
		t.Errorf("Expected max value %d, but got %d", expectedMaxValue, pq.q[0].Value)
	}

	// 弹出堆顶元素后，验证堆顶元素是否更新为次大值
	pq.Pop()
	expectedSecondMaxValue := 25
	if pq.q[0].Value != expectedSecondMaxValue {
		t.Errorf("Expected second max value %d, but got %d", expectedSecondMaxValue, pq.q[0].Value)
	}
}

func printSlice(slice []DataItem) {
	fmt.Printf("[")
	for i, v := range slice {
		fmt.Printf("(%v, %v)", v.Key, v.Value)
		if i < len(slice)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("]\n")
}

func TestPriorityQueueSort(t *testing.T) {
	pq := NewPriorityQueue([]DataItem{})
	data := []DataItem{
		{Key: "key1", Value: 10},
		{Key: "key2", Value: 5},
		{Key: "key3", Value: 20},
		{Key: "key4", Value: 15},
		{Key: "key5", Value: 25},
		{Key: "key6", Value: 30},
		{Key: "key7", Value: 8},
	}
	for _, item := range data {
		pq.Push(item)
	}
	printSlice(data)
	printSlice(pq.q)
	pq.Sort()
	printSlice(pq.q)
}
