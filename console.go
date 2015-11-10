package main

import "fmt"


func SendNotificationConsole(descriptions []string) {
	for _, val := range descriptions {
		fmt.Println()
		fmt.Println(val)
		fmt.Println("----------------------------------------------------------------------------------------------------")
	}
}