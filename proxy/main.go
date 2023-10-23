package main

import (
	"fmt"
	webProxy "github.com/ahmetson/web-proxy"
)

func main() {
	fmt.Printf("proxy starting\n")
	todayDo, err := webProxy.New()
	if err != nil {
		fmt.Printf("proxy failed with: %v", err)
		panic(err)
	}
	fmt.Printf("proxy started")

	wg, err := todayDo.Start()
	if err != nil {
		fmt.Printf("failed to start the proxy: %v", err)
		panic(err)
	}

	println("proxy is waiting for the operations...")

	wg.Wait()

	println("close the proxy app")
}
