package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os/exec"
	"strings"
)

func ListInterfaceAddresses() {
	addr, _ := net.InterfaceAddrs()
	fmt.Println(addr)
}

func CheckV4orV6(ip string) string {
	ip = strings.Split(ip, "/")[0]
	parsedIP := net.ParseIP(ip)
	isV4 := net.IP.To4(parsedIP)
	if isV4 != nil {
		return "V4"
	}

	isV6 := net.IP.To16(parsedIP)
	if isV6 != nil {
		return "V6"
	}
	return "nil"
}

func InterfaceByName(interface_name string) {
	temp, _ := net.InterfaceByName(interface_name)
	ipAddr, _ := temp.Addrs()

	fmt.Println("Normal addresses...")
	for i := range ipAddr {
		fmt.Printf("ip %s address : %s\n", CheckV4orV6(ipAddr[i].String()), ipAddr[i])
	}
	multicastAddr, _ := temp.MulticastAddrs()
	fmt.Println("Multicast addresses...")
	for i := range multicastAddr {
		fmt.Printf("ip %s address : %s\n", CheckV4orV6(multicastAddr[i].String()), multicastAddr[i])
	}
}

func AllInterface() {
	fmt.Println("Hello...")
	interfaces, _ := net.Interfaces()
	fmt.Println("Index	Name")
	for i := range interfaces {
		fmt.Printf("%d	%s\n", interfaces[i].Index, interfaces[i].Name)
	}
}

func AddIPRule(source_ip string, table_name string) {
	//ip rule add from source_ip lookup table_name
	//Check Rule exist or not
	//If Exist then do nothing
	//Does not exist then add the rule
	var flag bool
	rules, _ := exec.Command("ip", "rule", "list").Output()
	for _, line := range strings.Split(strings.TrimSuffix(string(rules), "\n"), "\n") {
		if strings.Contains(line, source_ip) && strings.Contains(line, table_name) {
			fmt.Println("Rule exist..")
			flag = true
			fmt.Println(line)
			continue
		}
	}
	if !flag {
		cmd := exec.Command("sudo", "ip", "rule", "add", "from", source_ip, "lookup", table_name)
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func AddIPRoute(destination string, intermediate string, net_interface string, table_name string) {
	//sudo ip route add default via 192.168.57.1 dev enp0s3 table custom1
	//Check Route exist or not
	//If Exist then do nothing
	//Does not exist then add the route
	var flag bool
	routes, _ := exec.Command("ip", "route", "list", "table", table_name).Output()
	for _, line := range strings.Split(strings.TrimSuffix(string(routes), "\n"), "\n") {
		if strings.Contains(line, destination) {
			fmt.Println("Route exist..")
			flag = true
			fmt.Println(line)
			continue
		}
	}
	if !flag {
		cmd := exec.Command("sudo", "ip", "route", "add", destination, "via", intermediate, "dev", net_interface, "table", table_name)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error here : ", err)
		}
	}
}

func MakeSudo() {
	c1 := exec.Command("echo", "allowme1")
	c2 := exec.Command("sudo", "-S", "ls")

	var stderr bytes.Buffer
	read_end, write_end := io.Pipe()
	c1.Stdout = write_end
	c2.Stdin = read_end
	c2.Stderr = &stderr
	c1.Start()

	go func() {
		defer write_end.Close()
		c1.Wait()
	}()
	err := c2.Run()
	if err != nil {
		fmt.Println(stderr)
	}
}

func main() {
	//InterfaceByName("wlp2s0")
	//AllInterface()
}
