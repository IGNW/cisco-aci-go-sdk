package main

import "fmt"

func main() {
	host, name, pass, insecure, _ := LookupClientEnvars()
	c := InitializeClient(host, name, pass, insecure)

	if c.AuthToken.IsValid() == false {
		fmt.Printf("TOKEN EXPIRES: %v\n", c.AuthToken.Expiry)
	} else {
		fmt.Printf("TOKEN is Good: %s", c.AuthToken.Token)
	}
}
