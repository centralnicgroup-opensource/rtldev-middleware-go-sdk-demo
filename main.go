package main

import (
	"fmt"

	CL "github.com/hexonet/go-sdk/v3/apiclient"
)

func main() {
	//sessionless API communication
	fmt.Println("--- SESSIONLESS API COMMUNICATION ---")
	cl := CL.NewAPIClient()
	cl.SetCredentials("test.user", "test.passw0rd") //username, password
	cl.SetRemoteIPAddress("1.2.3.4")                //for active ip filter setting
	//cl.SetOTP("12345678") to provide your 2FA otp code
	cl.UseOTESystem()
	cl.EnableDebugMode()
	r := cl.Request(map[string]interface{}{
		"COMMAND": "StatusAccount",
	})
	if r.IsSuccess() {
		fmt.Println("Command succeeded.")
	} else {
		fmt.Println("Command failed.")
	}
	fmt.Println()

	//session based API communication
	fmt.Println("--- SESSION BASED API COMMUNICATION ---")
	cl = CL.NewAPIClient()
	cl.SetCredentials("test.user", "test.passw0rd") //username, password
	cl.SetRemoteIPAddress("1.2.3.4")                //for active ip filter setting
	cl.UseOTESystem()
	cl.EnableDebugMode()
	r = cl.Login()
	// or cl.Login("12345678") to provide your 2FA otp code
	if r.IsSuccess() {
		fmt.Println("Login succeeded.")
		r = cl.Request(map[string]interface{}{
			"COMMAND": "StatusAccount",
		})
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
