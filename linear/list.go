package linear

import (
	"fmt"
	"strings"
)

/**
 * 一个支持在尾端插入的列表，同时支持查询、查新、自动扩容
 */
type List struct {
	caps        int
	used        int
	extendRatio int
	items       []int
}

func NewList() *List {
	return &List{
		caps:        8,
		used:        0,
		extendRatio: 2,
		items:       make([]int, 8),
	}
}

func (list *List) Capacityint() int {
	return list.caps
}

func (list *List) Usedint() int {
	return list.used
}

func (list *List) IsOutOfBorder(index int) bool {
	return index < 0 || index > list.caps
}

func (list *List) Get(index int) int {
	/**
	 * 如果没有超出边界，就直接从切片中取出来
	 */
	if list.IsOutOfBorder(index) {
		panic("Out-Of-Border-Error")
	}
	return list.items[index]
}

func (list *List) Set(item int, index int) {
	/**
	 * 如果没有超出边界，就直接更新切片
	 */
	if list.IsOutOfBorder(index) {
		panic("Out-Of-Border-Error")
	}
	list.items[index] = item
}

func (list *List) extend() {
	if list.used == list.caps {
		temp := make([]int, list.extendRatio*list.caps)
		copy(temp, list.items)
		list.items = temp
		list.caps = list.caps * list.extendRatio
	}
}

func (list *List) String() string {
	temp := []string{}
	for i := range list.items {
		temp = append(temp, fmt.Sprintf("%d", list.items[i]))
	}
	return "[" + strings.Join(temp, ",") + "]"
}

func (list *List) Append(item int) {
	/*
	 * 检查一下是不是需要扩容，如果有需要就给它扩容上去
	 */
	if list.used == list.caps {
		list.extend()
	}
	list.items[list.used] = item
}

func (list *List) Remove(index int) int {
	if list.IsOutOfBorder(index) {
		panic("Out-Of-Border-Error")
	}

	num := list.items[index]

	for i := index; i < list.used; i++ {
		list.items[i] = list.items[i+1]
	}
	list.used = list.used - 1
	return num
}
