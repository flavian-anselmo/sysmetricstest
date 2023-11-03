package main

import (
	"fmt"
	"log"
	"time"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/process"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"

)

func main() {

	get_load_avg()
	get_no_of_processes_running()
	get_host_info()
	get_swap_memory_stats()
	get_virtual_memory_stats()
	get_disk_info()

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


func get_swap_memory_stats(){
	// Fetch detailed system information
	swap_info, err := mem.SwapMemory()
	if err != nil {
		log.Fatal(err)
	}


	fmt.Printf(" ---------------------\n")
	fmt.Printf("| SWAP MEMEMORY INFORMATION  |\n")
	fmt.Printf(" ---------------------\n")
	fmt.Println("Free:", (float64(swap_info.Free))/1048576, "MB or ", (float64(swap_info.Free)/1073741824), "GB")
	fmt.Println("Percentage Used:", swap_info.UsedPercent, "%")
	fmt.Println("Used: ",float64(swap_info.Used)/1048576)
	fmt.Println("Total Memory Swapped: ", float64(swap_info.Total)/ 1048576 ,"MB or ", float64(swap_info.Total)/ 1073741824)
}


func get_virtual_memory_stats(){
		// Fetch detailed system information
	vm_info, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" ---------------------\n")
	fmt.Printf("| VIRTUAL  MEMEMORY INFORMATION  |\n")
	fmt.Printf(" ---------------------\n")
	fmt.Println("Total:", (float64(vm_info.Total))/1048576, "MB or ", (float64(vm_info.Total)/1073741824), "GB")
	fmt.Println("Available:", (float64(vm_info.Available))/1048576, "MB or ", (float64(vm_info.Available)/1073741824), "GB")
	fmt.Println("RAM Used:", (float64(vm_info.Used))/1048576, "MB or ", (float64(vm_info.Used)/1073741824), "GB")
	fmt.Println("RAM Used Percentage: ", (float64(vm_info.UsedPercent)),"%")
	fmt.Println("Free:", (float64(vm_info.Free))/1048576, "MB or ", (float64(vm_info.Free)/1073741824), "GB")
}


func get_disk_info(){

	// Fetch disk information
	diskInfo, err := disk.Usage("/")
	if err != nil {
		log.Fatal(err)
	}


	// Print disk information and elapsed time
	fmt.Printf(" ---------------------\n")
	fmt.Printf("| DISK MOUNT INFORMATION  |\n")
	fmt.Printf(" ---------------------\n")
	fmt.Printf("Total: %f GB\n", float64(diskInfo.Total)/1073741824)
	fmt.Printf("Free: %f GB \n", float64(diskInfo.Free)/1073741824)
	fmt.Printf("Used: %f GB \n", float64(diskInfo.Used)/1073741824)
	fmt.Printf("Usage: %.2f%%\n", diskInfo.UsedPercent)
}