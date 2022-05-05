package main

import (
	"fmt"

	"github.com/seaung/ZoomEye-go/zoomeye"
)

func main() {

	z := zoomeye.NewEnvZoomEyeClient()

	info, err := z.GetResourcesInfo()

	if err != nil {
		fmt.Printf("Get resource info error %v\n", err)
		return
	}

	fmt.Printf("The plan : %s\n", info.Plan)
	fmt.Printf("Username : %s\n", info.UserInfo.Name)
	fmt.Printf("Role     : %s\n", info.UserInfo.Role)
}
