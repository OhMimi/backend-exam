package main

import (
	"fmt"
	"sync"
)

type Employee struct {
	ID int
}

func (e *Employee) GetName() string {
	return fmt.Sprintf("employee[%d]", e.ID)
}

func (e *Employee) ProcessTask(wg *sync.WaitGroup, itemChan <-chan Item, statisticsChan chan<- StatisticsData) {
	defer wg.Done()
	itemCount := 0
	totalElapsedTime := 0

	fmt.Printf("%s is ready.\n", e.GetName())

	for item := range itemChan {
		fmt.Printf("%s start process item: %s.\n", e.GetName(), item.GetName())
		item.Process()
		fmt.Printf("%s end process item: %s elapsed: %d second.\n", e.GetName(), item.GetName(), item.GetProcessTime())
		itemCount++
		totalElapsedTime += item.GetProcessTime()
	}

	statisticsChan <- StatisticsData{
		EmployeeID:         e.ID,
		ProcessedItemCount: itemCount,
		ElapsedTime:        totalElapsedTime,
	}

	fmt.Printf("%s has finished all tasks.\n", e.GetName())
}
