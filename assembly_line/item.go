package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type BaseItem struct {
	ID          int
	ItemType    int
	ProcessTime int // 模擬處理時間，單位為秒
}

func (i BaseItem) Process() {
	time.Sleep(time.Duration(i.ProcessTime) * time.Second)
}

func (i BaseItem) GetName() string {
	return fmt.Sprintf("item%d-%d", i.ItemType, i.ID)
}

func (i BaseItem) GetProcessTime() int {
	return i.ProcessTime
}

type Item1 struct {
	BaseItem
}
type Item2 struct {
	BaseItem
}

type Item3 struct {
	BaseItem
}

type Item interface {
	// Process 這是一個耗時操作
	Process()
	// GetName 返回物品名稱
	GetName() string
	// GetProcessTime 返回處理時間
	GetProcessTime() int
}

func NewItem(itemType int, id int) Item {
	switch itemType {
	case 1:
		return Item1{
			BaseItem: BaseItem{
				ID:          id,
				ItemType:    1,
				ProcessTime: 1,
			},
		}
	case 2:
		return Item2{
			BaseItem: BaseItem{
				ID:          id,
				ItemType:    2,
				ProcessTime: 2,
			},
		}
	case 3:
		return Item3{
			BaseItem: BaseItem{
				ID:          id,
				ItemType:    3,
				ProcessTime: 3,
			},
		}
	default:
		panic(fmt.Sprintf("ItemType %d is not supported", itemType))
	}
}

// ShuffleItems 隨機打亂物品順序
func ShuffleItems(items []Item) {
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
}
