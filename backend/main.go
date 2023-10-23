package main

import (
	"fmt"
	serviceConfig "github.com/ahmetson/config-lib/service"
	"github.com/ahmetson/datatype-lib/data_type/key_value"
	"github.com/ahmetson/datatype-lib/message"
	"github.com/ahmetson/handler-lib/sync_replier"
	"github.com/ahmetson/os-lib/path"
	"github.com/ahmetson/service-lib"
	"path/filepath"
)

const (
	Add  = "add"
	Done = "done"
	List = "list"
)

var tasksAmount uint64

type Task struct {
	Number      uint64 `json:"number,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var list key_value.KeyValue

func onList(request message.RequestInterface) message.ReplyInterface {
	todoList := make([]Task, len(list))

	i := 0
	for _, raw := range list {
		task := raw.(Task)
		todoList[i] = task

		i++
	}

	params := key_value.New().Set("list", todoList)

	return request.Ok(params)
}

func onAdd(request message.RequestInterface) message.ReplyInterface {
	title, err := request.RouteParameters().StringValue("title")
	if err != nil {
		return request.Fail(fmt.Sprintf("request.Parameters.GetString('title'): %v", err))
	}

	description, err := request.RouteParameters().StringValue("description")
	if err != nil {
		return request.Fail(fmt.Sprintf("request.Parameters.GetString('description'): %v", err))
	}

	tasksAmount++
	number := tasksAmount

	task := Task{
		Number:      number,
		Title:       title,
		Description: description,
	}

	list.Set(fmt.Sprintf("%d", number), task)

	params := key_value.New().Set("number", number)

	return request.Ok(params)
}

func onDone(request message.RequestInterface) message.ReplyInterface {
	number, err := request.RouteParameters().Uint64Value("number")
	if err != nil {
		return request.Fail(fmt.Sprintf("request.Parameters.GetUint64('number'): %v", err))
	}

	numberStr := fmt.Sprintf("%d", number)
	_, ok := list[numberStr]
	if !ok {
		return request.Fail(fmt.Sprintf("list['%d'] not found", number))
	}

	delete(list, numberStr)

	return request.Ok(key_value.New())
}

func main() {
	list = key_value.New()
	tasksAmount = 0

	syncReplier := sync_replier.New()
	err := syncReplier.Route(Add, onAdd)
	if err != nil {
		panic(err)
	}
	err = syncReplier.Route(Done, onDone)
	if err != nil {
		panic(err)
	}
	err = syncReplier.Route(List, onList)
	if err != nil {
		panic(err)
	}

	todayDo, err := service.New()
	if err != nil {
		panic(err)
	}

	todayDo.SetHandler("main", syncReplier)
	local := serviceConfig.Local{
		LocalBin: filepath.Join(path.BinPath("../proxy/bin", "test")),
	}
	webProxy := &serviceConfig.Proxy{
		Local:    &local,
		Id:       "web-proxy",
		Url:      "github.com/ahmetson/web-proxy",
		Category: "entry",
	}
	listLocal := serviceConfig.Local{
		LocalBin: filepath.Join(path.BinPath("../list-proxy/bin", "test")),
	}
	listProxy := &serviceConfig.Proxy{
		Local:    &listLocal,
		Id:       "list-proxy",
		Url:      "github.com/ahmetson/today-do/list-local",
		Category: "convert",
	}
	serviceRule := serviceConfig.NewServiceDestination(todayDo.Url())
	fmt.Printf("service rule: %v to %s\n", *serviceRule, todayDo.Url())
	proxyChain, err := serviceConfig.NewProxyChain([]*serviceConfig.Proxy{webProxy, listProxy}, serviceRule)
	if err != nil {
		panic(err)
	}
	err = todayDo.SetProxyChain(proxyChain)
	if err != nil {
		panic(err)
	}

	wg, err := todayDo.Start()
	if err != nil {
		panic(err)
	}

	println("waiting for the operations in backend...")

	wg.Wait()

	println("close the backend app")
}
