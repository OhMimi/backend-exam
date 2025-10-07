package main

import (
	"fmt"
	"sync"
	"time"
)

/*
請模擬流水線, 五個員工處理三種物品

☑三種物品數量各十件
☑三種物品處理時間需不一樣
☑物品的處理順序請隨機打亂
☑物品處理需透過 interface 來傳遞
☑每個員工一次只能處理一種物品
☑開始以及結束處理都需要打印紀錄
☑統計總處理時間, 及每個員工處理了多少物品
*/

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Program panic:", r)
		}
	}()

	const numEmployees = 5
	const numItemsPerType = 10

	// produce items
	items := make([]Item, 0, numItemsPerType*3)
	for i := 0; i < numItemsPerType; i++ {
		items = append(items, NewItem(1, i+1))
		items = append(items, NewItem(2, i+1))
		items = append(items, NewItem(3, i+1))
	}

	// shuffle items
	ShuffleItems(items)

	fmt.Println("All items are ready and random sorted.")

	// create channels
	itemChan := make(chan Item, len(items))
	statisticsChan := make(chan StatisticsData, numEmployees)

	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(numEmployees)

	// start employees
	for i := 0; i < numEmployees; i++ {
		employee := &Employee{ID: i + 1}
		go employee.ProcessTask(&wg, itemChan, statisticsChan)
	}
	// send items to itemChan
	for _, item := range items {
		itemChan <- item
	}
	close(itemChan)

	// wait for all employees to finish
	wg.Wait()
	close(statisticsChan)

	fmt.Println("================= Statistics ================")

	// print statistics
	totalProcessedItems := 0
	totalElapsedTime := 0
	for stats := range statisticsChan {
		stats.PrintContent()
		totalProcessedItems += stats.ProcessedItemCount
		totalElapsedTime += stats.ElapsedTime
	}

	realElapsed := time.Since(start)

	fmt.Printf("Total processed items: %d\n", totalProcessedItems)
	fmt.Printf("Total elapsed time (sum of all employees): %d seconds\n", totalElapsedTime)
	fmt.Printf("Real elapsed time: %v seconds\n", realElapsed.Seconds())
}
