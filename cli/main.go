package main

import (
	"fmt"
	"github.com/ahmetson/client-lib"
	"github.com/ahmetson/config-lib/service"
	"github.com/ahmetson/datatype-lib/data_type/key_value"
	"github.com/ahmetson/datatype-lib/message"
	"github.com/ahmetson/os-lib/arg"
	"github.com/ahmetson/os-lib/net"
	"github.com/ahmetson/os-lib/process"
	"github.com/pebbe/zmq4"
	"os"
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
		params := key_value.New().Set("title", title).Set("description", description)

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

		params := key_value.New().Set("number", number)

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
			Parameters: key_value.New(),
		}

		reply, err := backend.Request(&req)
		if err != nil {
			panic(err)
		}

		fmt.Printf("replied: %v\n", reply)
	} else if cmd == "close" {
		fmt.Printf("flag --port must point to the manager of the service\n")
		req := message.Request{
			Command:    "heartbeat",
			Parameters: key_value.New(),
		}

		reply, err := backend.Request(&req)
		if err != nil {
			panic(err)
		}
		if !reply.IsOK() {
			fmt.Printf("heartbeat failed with '%s' error, is it closed or run --same-process to check is it running", reply.ErrorMessage())
			return
		}

		req = message.Request{
			Command:    "close",
			Parameters: key_value.New(),
		}

		reply, err = backend.Request(&req)
		if err != nil {
			panic(err)
		}
		fmt.Printf("reply: ok? %v error message: %s\n", reply.IsOK(), reply.ErrorMessage())

		fmt.Printf("close signal send, check in other terminal that app is closed\n")
	} else if cmd == "same-process" {
		if !arg.FlagExist("port-2") {
			panic("flag --port-2 not set")
		}

		port2 := arg.FlagValue("port-2")

		portNum, err := strconv.Atoi(port)
		if err != nil {
			panic(err)
		}
		portNum2, err := strconv.Atoi(port2)
		if err != nil {
			panic(err)
		}

		used := net.IsPortUsed("localhost", portNum)
		used2 := net.IsPortUsed("localhost", portNum2)
		if !used && !used2 {
			fmt.Printf("both ports are not used\n")
			return
		}
		if !used {
			fmt.Printf("port %s not used\n", port)
		}
		if !used2 {
			fmt.Printf("port %s not used\n", port2)
		}

		pid := uint64(0)
		pid2 := uint64(0)
		if used {
			pid, err = process.PortToPid(portNum)
			if err != nil {
				panic(err)
			}
			fmt.Printf("port %s pid %d\n", port, pid)
		}
		if used2 {
			pid2, err = process.PortToPid(portNum2)
			if err != nil {
				panic(err)
			}
			fmt.Printf("port %s pid %d\n", port2, pid2)
		}

		if used && used2 {
			fmt.Printf("same process? %v\n", pid == pid2)
		}
	} else if cmd == "kill" {
		if !arg.FlagExist("pid") {
			panic("flag --pid not set")
		}
		pid := arg.FlagValue("pid")

		pidNum, err := strconv.Atoi(pid)
		if err != nil {
			panic(err)
		}

		proc, err := os.FindProcess(pidNum)
		if err != nil {
			panic(err)
		}

		err = proc.Kill()
		if err != nil {
			panic(err)
		}

		fmt.Printf("process with id %s killed\n", pid)
	} else if cmd == "units" {
		fmt.Printf("fetch the units from service manager")

		rule := service.NewServiceDestination("github.com/ahmetson/today-do")

		req := message.Request{
			Command:    "units",
			Parameters: key_value.New().Set("rule", *rule),
		}
		fmt.Printf("Request: %v\n", req)

		reply, err := backend.Request(&req)
		if err != nil {
			panic(err)
		}

		if !reply.IsOK() {
			fmt.Printf("failed to reply: %s\n", reply.ErrorMessage())
		}
		fmt.Printf("replied parameters: %v\n", reply.ReplyParameters())
	}
}
