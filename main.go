package main

import (
	"flag"
	"fmt"
	"os"

	"./core"
	"./utils"
)

func main() {
	banner := utils.UglyBanner()
	fmt.Println(banner)
	target := flag.String("target", "", "[-] (e.g: https://wordpress.site.com)")
	server := flag.String("server", "", "[-] (e.g: http://159.89.121.20 or http://mydomain.com")
	flag.Parse()

	if *target == "" || *server == "" {
		fmt.Println("[-] Required a target.. and target server.. ")
		fmt.Printf("(e.g main.go -target https://www.target.com -server http://burpcollborator.net)\n")
		os.Exit(1)
	}

	serverUser := *server
	targetUser := *target

	//verify if data is from stdin
	file, _ := os.Stdin.Stat()
	if (file.Mode() & os.ModeCharDevice) == 0 {
		core.FromStdin()
	} else {
		if core.IsAlive(targetUser) {
			core.VerifyMethods(targetUser)
			fmt.Println(serverUser)
		}
	}

}
