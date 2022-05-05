package main

import (
	"fmt"

	"github.com/seaung/ZoomEye-go/zoomeye"
)

func main() {
	z := zoomeye.NewEnvZoomEyeClient()

	domaininfo, err := z.DomainSearch("baidu.com", 0, 1)
	if err != nil {
		fmt.Printf("Can't get domain info error %v\n", err)
		return
	}

	fmt.Printf("Status Code : %d\n", domaininfo.Status)
	fmt.Printf("Total count : %d\n", domaininfo.Total)
	fmt.Printf("Message     : %s\n", domaininfo.Msg)

	for _, info := range domaininfo.List {
		fmt.Printf("The Name  : %s\n", info.Name)
		fmt.Printf("Timestamp : %s\n", info.Timestamp)

		if len(info.Ip) == 0 {
			continue
		}

		for _, addr := range info.Ip {
			fmt.Printf("The IP Address : %s\n", addr)
		}
	}
}
