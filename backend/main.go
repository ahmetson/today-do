package main

import (
	"fmt"
	"github.com/ahmetson/datatype-lib/data_type/key_value"
	"github.com/ahmetson/datatype-lib/message"
	"github.com/ahmetson/handler-lib/sync_replier"
	"github.com/ahmetson/service-lib"
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

	wg, err := todayDo.Start()
	if err != nil {
		panic(err)
	}

	println("waiting for the operations...")

	wg.Wait()

	println("close the app")
}
