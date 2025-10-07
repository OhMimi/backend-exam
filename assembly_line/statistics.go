package main

import "fmt"

type StatisticsData struct {
	EmployeeID         int
	ProcessedItemCount int
	ElapsedTime        int // 單位為秒
}

func (s StatisticsData) PrintContent() {
	fmt.Printf("employee[%d] processed %d items, elapsed time: %d seconds.\n", s.EmployeeID, s.ProcessedItemCount, s.ElapsedTime)
}
