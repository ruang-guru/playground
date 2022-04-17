package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

//kita ingin hanya ada 5 goroutine yang berjalan bersamaan
func main() {
	ctx := context.TODO()
	sem := semaphore.NewWeighted(int64(5))
	totProcess := 10
	for i := 1; i <= totProcess; i++ {
		if err := sem.Acquire(ctx, 1); err != nil {
			fmt.Printf("Failed to acquire semaphore: %v", err)
			break
		}
		go func(v int) {
			defer sem.Release(1)
			longRunningProcess(v)
		}(i)
	}
	//karena command dibawah ini langsung meminta 5 token
	//dan semua token sudah diambil saat looping
	//maka terjadi blocking disini, hingga semua goroutine
	//memanggil sem.Release
	if err := sem.Acquire(ctx, int64(5)); err != nil {
		fmt.Printf("Failed to acquire semaphore: %v", err)
	}
}
func longRunningProcess(taskID int) {
	fmt.Println(
		time.Now().Format("15:04:05"),
		"Running task with ID",
		taskID)
	time.Sleep(2 * time.Second) // melakukan sesuatu
}
