package main

import (
	"fmt"
	"log"
	"time"
	"github.com/shirou/gopsutil/load"
)

func main() {
	get_load_avg()
}



func get_load_avg() {
	// Measure the time it takes to fetch system load average
	startTime := time.Now()

	// Fetch system load average
	loadAvg, err := load.Avg()
	if err != nil {
		log.Fatal(err)
	}
	
	// Calculate elapsed time
	elapsedTime := time.Since(startTime)

	// Print the load average and elapsed time
	fmt.Printf(" ---------------------\n")
	fmt.Printf("| SYSTEM LOAD AVERAGE:|\n")
	fmt.Printf(" ---------------------\n")

	fmt.Printf("1-minute: %.2f\n", loadAvg.Load1)
	fmt.Printf("5-minute: %.2f\n", loadAvg.Load5)
	fmt.Printf("15-minute: %.2f\n", loadAvg.Load15)
	fmt.Printf("Time Taken: %v\n", elapsedTime)
	fmt.Printf("------------------------\n")

}




