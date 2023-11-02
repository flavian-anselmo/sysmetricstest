package main

import (
	"fmt"
	"log"
	"time"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/process"
	"github.com/shirou/gopsutil/host"


)

func main() {

	get_load_avg()
	get_no_of_processes_running()
	get_host_info()

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

}


func get_no_of_processes_running(){
	// Measure the time it takes to fetch running processes
	startTime := time.Now()

	// Get a list of all running processes
	processes, err := process.Processes()
	if err != nil {
		log.Fatal(err)
	}

	// Calculate elapsed time
	elapsedTime := time.Since(startTime)

	// Print the list of running processes and elapsed time
	fmt.Printf(" ---------------------\n")
	fmt.Printf("| PROCESSES RUNNING:  |\n")
	fmt.Printf(" ---------------------\n")
	fmt.Println("Number of processes: ", len(processes))
	fmt.Printf("Elapsed Time: %v\n", elapsedTime)


	// for _, p := range processes {
	// 	name, _ := p.Name()
	// 	pid := p.Pid
	// 	fmt.Printf("Process ID: %d, Name: %s\n", pid, name)
	// }
}


func get_host_info (){


	startTime := time.Now()

	// Fetch detailed system information
	info, err := host.Info()
	if err != nil {
		log.Fatal(err)
	}

	// Calculate elapsed time
	elapsedTime := time.Since(startTime)

	fmt.Printf(" ---------------------\n")
	fmt.Printf("| SYSTEM INFORMATION  |\n")
	fmt.Printf(" ---------------------\n")
	fmt.Printf("Hostname: %s\n", info.Hostname)
	fmt.Printf("OS: %s %s\n", info.OS, info.PlatformVersion)
	fmt.Printf("Architecture: %s\n", info.KernelArch)
	fmt.Println("Uptime:", (float64(info.Uptime)/3600), "hours")
	fmt.Printf("Elapsed Time: %v\n", elapsedTime)
}



