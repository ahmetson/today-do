package main

import (
	"fmt"
	"github.com/ahmetson/client-lib/config"
	"github.com/ahmetson/config-lib/client"
	handlerConfig "github.com/ahmetson/handler-lib/config"
	"github.com/ahmetson/os-lib/arg"
	"github.com/ahmetson/service-lib"
	"github.com/ahmetson/web-lib"
)

func main() {
	if !arg.FlagExist("destination") {
		panic("flag --destination not found")
	}
	destinationId := arg.FlagValue("destination")

	proxy, err := web.New()
	if err != nil {
		panic(err)
	}

	todayDo, err := service.New()
	if err != nil {
		panic(err)
	}

	todayDo.SetHandler("entry", proxy)

	configClient, err := client.New()
	if err != nil {
		panic(err)
	}
	destinationExist, err := configClient.ServiceExist(destinationId)
	if err != nil {
		panic(err)
	}
	if !destinationExist {
		panic(fmt.Sprintf("destination '%s' not set", destinationId))
	}
	serviceConfig, err := configClient.Service(destinationId)
	if err != nil {
		panic(fmt.Sprintf("configClient.Service('%s'): %v", destinationId, err))
	}
	destinationConfig := &config.Client{
		ServiceUrl: serviceConfig.Url,
		Id:         serviceConfig.Handlers[0].Id,
		Port:       serviceConfig.Handlers[0].Port,
		TargetType: handlerConfig.SocketType(serviceConfig.Handlers[0].Type),
	}
	destinationConfig.UrlFunc(config.Url)
	proxy.SetDestination(destinationConfig)

	wg, err := todayDo.Start()
	if err != nil {
		panic(err)
	}

	println("waiting for the operations...")

	wg.Wait()

	println("close the app")
}
