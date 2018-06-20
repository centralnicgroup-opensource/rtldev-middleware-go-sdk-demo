package main

import (
	"fmt"

	"github.com/hexonet/go-sdk/client"
)

func main() {
	//sessionless API communication
	fmt.Println("--- SESSIONLESS API COMMUNICATION ---")
	cl := client.NewClient()
	cl.SetCredentials("test.user", "test.passw0rd", "")
	cl.UseOTESystem()
	cmd := map[string]string{
		"COMMAND": "StatusAccount",
	}
	r := cl.Request(cmd)
	if r.IsSuccess() {
		fmt.Println("Command succeeded.")
	} else {
		fmt.Println("Command failed.")
	}
	fmt.Println()

	//session based API communication
	fmt.Println("--- SESSION BASED API COMMUNICATION ---")
	cl = client.NewClient()
	cl.SetCredentials("test.user", "test.passw0rd", "") //username, password, otp code (2FA)
	cl.UseOTESystem()
	r = cl.Login()
	if r.IsSuccess() {
		fmt.Println("Login succeeded.")
		cmd := map[string]string{
			"COMMAND": "StatusAccount",
		}
		r = cl.Request(cmd)
		if r.IsSuccess() {
			fmt.Println("Command succeeded.")
			r = cl.Logout()
			if r.IsSuccess() {
				fmt.Println("Logout succeeded.")
			} else {
				fmt.Println("Logout failed.")
			}
		} else {
			fmt.Println("Command failed.")
		}
	} else {
		fmt.Println("Login failed.")
	}
}
