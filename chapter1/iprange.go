package main

import(
	"strconv"
	"strings"
	"fmt"
)
func main() {
	
	ip := getCidrIpRange("192.168.23.123")
	fmt.Println(ip)
}

func getCidrIpRange(cidr string) (string) {
	if len(cidr) == 0 {
		//	log.Info(" ip null")
		return ""
	}

	if len(strings.Split(cidr, "/")) == 1 {
		//		log.Info("no / ")
		return  ""
	}

	ip := strings.Split(cidr, "/")[0]
	ipSegs := strings.Split(ip, ".")
	maskLen, err:= strconv.Atoi(strings.Split(cidr, "/")[1])
	if err != nil || maskLen > 32 {
		return ""
	}

	IpValue := ipTransferUnInt(ipSegs)
	maskValue := ComputeMaskValue(maskLen)
	fmt.Println(IpValue)
	fmt.Println(maskValue)
	return IntTransferIp(IpValue & maskValue)
}

func ipTransferUnInt(ips []string ) (int) {
	if len(ips) != 4 {
		return 0
	}
	var intIP int
	for k, v := range ips {
		i, err := strconv.Atoi(v)
			fmt.Println(v)
		if err != nil || i > 255 {
			return 0
		}
		intIP = intIP | i<<uint(8*(3-k))
	}
	return intIP
}

func IntTransferIp(ipValue int ) string {
	var ipSegMax int = 255
	ipRange3 := ipValue & (ipSegMax)
	ipRange2 := (ipValue >> 8) & (ipSegMax)
	fmt.Println(ipRange3)
	ipRange1 := (ipValue >> 16) & (ipSegMax)
	ipRange0 := (ipValue >> 24) & (ipSegMax)
	return strconv.Itoa(ipRange0) + "." + strconv.Itoa(ipRange1) + "." + strconv.Itoa(ipRange2) + "."+ strconv.Itoa(ipRange3)
	
}

func ComputeMaskValue(masklen int) int {
	var maskValue int = 0
	var baseValue int = 1
	for i := 0; i < masklen; i++ {
		maskValue |= baseValue << (31-i)
	}
	return maskValue
}
