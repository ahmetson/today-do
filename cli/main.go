package main

import (
	"fmt"
	"github.com/ahmetson/client-lib"
	"github.com/ahmetson/common-lib/data_type/key_value"
	"github.com/ahmetson/common-lib/message"
	"github.com/ahmetson/os-lib/arg"
	"github.com/pebbe/zmq4"
	"strconv"
	"time"
)

func main() {
	exist := arg.FlagExist("cmd")
	if !exist {
		panic("flag --cmd=<> not set")
	}
	if !arg.FlagExist("port") {
		panic("flag --port not set")
	}

	cmd := arg.FlagValue("cmd")
	port := arg.FlagValue("port")

	title := time.Now().String()
	description := fmt.Sprintf("this is the text that was set on %s", title)
	backend, err := client.NewRaw(zmq4.REP, fmt.Sprintf("tcp://localhost:%s", port))
	if err != nil {
		panic(err)
	}
	backend.Timeout(time.Second)
	backend.Attempt(1)

	fmt.Printf("executing\n")

	if cmd == "add" {
		params := key_value.Empty().Set("title", title).Set("description", description)

		req := message.Request{
			Command:    "add",
			Parameters: params,
		}

		fmt.Printf("request to the backend\n")

		reply, err := backend.Request(&req)
		if err != nil {
			panic(err)
		}

		fmt.Printf("replied: %v\n", reply)
	} else if cmd == "done" {
		if !arg.FlagExist("number") {
			panic("flag --number not set")
		}

		numberStr := arg.FlagValue("number")
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			panic(err)
		}

		params := key_value.Empty().Set("number", number)

		req := message.Request{
			Command:    "done",
			Parameters: params,
		}

		fmt.Printf("request to the backend\n")

		reply, err := backend.Request(&req)
		if err != nil {
			panic(err)
		}

		fmt.Printf("replied: %v\n", reply)
	} else if cmd == "list" {
		req := message.Request{
			Command:    "list",
			Parameters: key_value.Empty(),
		}

		reply, err := backend.Request(&req)
		if err != nil {
			panic(err)
		}

		fmt.Printf("replied: %v\n", reply)
	} else if cmd == "close" {
		fmt.Printf("flag --port must point to the manager of the service")
		req := message.Request{
			Command:    "close",
			Parameters: key_value.Empty(),
		}

		err = backend.Submit(&req)
		if err != nil {
			panic(err)
		}

		fmt.Printf("close signal send, check in other terminal that app is closed\n")
	}
}
