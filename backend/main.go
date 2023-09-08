package main

import (
	"fmt"
	"github.com/ahmetson/common-lib/data_type/key_value"
	"github.com/ahmetson/common-lib/message"
	"github.com/ahmetson/handler-lib/sync_replier"
	"github.com/ahmetson/service-lib"
	"sync"
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

func onList(request message.Request) message.Reply {
	todoList := make([]Task, len(list))

	i := 0
	for _, raw := range list {
		task := raw.(Task)
		todoList[i] = task

		i++
	}

	params := key_value.Empty().Set("list", todoList)

	return request.Ok(params)
}

func onAdd(request message.Request) message.Reply {
	title, err := request.Parameters.GetString("title")
	if err != nil {
		return request.Fail(fmt.Sprintf("request.Parameters.GetString('title'): %v", err))
	}

	description, err := request.Parameters.GetString("description")
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

	params := key_value.Empty().Set("number", number)

	return request.Ok(params)
}

func onDone(request message.Request) message.Reply {
	number, err := request.Parameters.GetUint64("number")
	if err != nil {
		return request.Fail(fmt.Sprintf("request.Parameters.GetUint64('number'): %v", err))
	}

	numberStr := fmt.Sprintf("%d", number)
	_, ok := list[numberStr]
	if !ok {
		return request.Fail(fmt.Sprintf("list['%d'] not found", number))
	}

	delete(list, numberStr)

	return request.Ok(key_value.Empty())
}

func main() {
	list = key_value.Empty()
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

	err = todayDo.Start()
	if err != nil {
		panic(err)
	}

	println("started")

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
