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
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/net"



)

func main() {
	startTime := time.Now()
	get_load_avg()
	get_no_of_processes_running()
	get_host_info()
	get_swap_memory_stats()
	get_virtual_memory_stats()
	get_disk_info()
	get_network_stats()
	// get_cpu_usage()
	elapsedTime := time.Since(startTime)
	fmt.Printf("Elapsed Time: %v\n", elapsedTime)
}



func get_load_avg() {
	startTime := time.Now()
	// Fetch system load average
	loadAvg, err := load.Avg()
	if err != nil {
		log.Fatal(err)
	}
	elapsedTime := time.Since(startTime)
	// Print the load average and elapsed time
	fmt.Printf(" ---------------------\n")
	fmt.Printf("| SYSTEM LOAD AVERAGE:|\n")
	fmt.Printf(" ---------------------\n")

	fmt.Printf("1-minute: %.2f\n", loadAvg.Load1)
	fmt.Printf("5-minute: %.2f\n", loadAvg.Load5)
	fmt.Printf("15-minute: %.2f\n", loadAvg.Load15)
	fmt.Printf("Elapsed Time: %v\n", elapsedTime)

}


func get_no_of_processes_running(){

	startTime := time.Now()

	// Get a list of all running processes
	processes, err := process.Processes()
	if err != nil {
		log.Fatal(err)
	}

	elapsedTime := time.Since(startTime)

	// Print the list of running processes and elapsed time
	fmt.Printf(" ---------------------\n")
	fmt.Printf("| PROCESSES RUNNING:  |\n")
	fmt.Printf(" ---------------------\n")
	fmt.Println("Number of processes: ", len(processes))
	fmt.Printf("Elapsed Time: %v\n", elapsedTime)

}


func get_host_info (){



	// Fetch detailed system information
	info, err := host.Info()
	if err != nil {
		log.Fatal(err)
	}



	fmt.Printf(" ---------------------\n")
	fmt.Printf("| SYSTEM INFORMATION  |\n")
	fmt.Printf(" ---------------------\n")
	fmt.Printf("Hostname: %s\n", info.Hostname)
	fmt.Printf("OS: %s %s\n", info.OS, info.PlatformVersion)
	fmt.Printf("Architecture: %s\n", info.KernelArch)
	fmt.Println("Uptime:", (float64(info.Uptime)/3600), "hours")
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

func get_cpu_usage(){
	cpuPercent, err := cpu.Percent(5*time.Second, false)
	if err != nil {
		log.Fatal(err)
	}
	
	vm, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" ---------------------\n")
	fmt.Printf("| CPU USAGE            |\n")
	fmt.Printf(" ---------------------\n")	
	fmt.Printf("Go System Metrics:\n")
	fmt.Printf("CPU Usage: %.2f%%\n", cpuPercent[0])
	fmt.Printf("Total Memory: %d MB\n", vm.Total/1024/1024)
}

func get_network_stats(){

	netIOCounters, err := net.IOCounters(false)
	if err != nil {
		log.Fatal(err)
	}


	fmt.Printf(" ---------------------\n")
	fmt.Printf("| NETWWORK STATS   |\n")
	fmt.Printf(" ---------------------\n")
	for _, io := range netIOCounters {
		fmt.Printf("Name: %s\n", io.Name)
		fmt.Printf("  Bytes Sent: %d\n", io.BytesSent)
		fmt.Printf("  Bytes Received: %d\n", io.BytesRecv)
		fmt.Printf("  Packets Sent: %d\n", io.PacketsSent)
		fmt.Printf("  Packets Received: %d\n", io.PacketsRecv)
	}
}