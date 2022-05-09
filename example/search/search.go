package main

import (
	"fmt"

	"github.com/seaung/ZoomEye-go/zoomeye"
)

func main() {
	z := zoomeye.NewEnvZoomEyeClient()

	matchers, err := z.HostSearch("port:21%20city:beijing", "app,os", 1)
	if err != nil {
		fmt.Printf("Can't get matchers : %v\n", err)
		return
	}

	for _, host := range matchers.Matches {
		fmt.Printf("Jarm : %s\n", host.Jarm)
		fmt.Printf("md5 : %s\n", host.Ico.Md5)
	}
}
