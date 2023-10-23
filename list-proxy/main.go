package main

import (
	"fmt"
	"github.com/ahmetson/datatype-lib/data_type/key_value"
	"github.com/ahmetson/datatype-lib/message"
	"github.com/ahmetson/service-lib"
	"slices"
)

func main() {
	fmt.Printf("list proxy starting...\n")
	proxy, err := service.NewProxy()
	if err != nil {
		panic(err)
	}
	fmt.Printf("new list proxy created...\n")

	onReply := func(handlerId string, req message.RequestInterface, reply message.ReplyInterface) (message.ReplyInterface, error) {
		commands := []string{"add", "done"}
		if !slices.Contains(commands, req.CommandName()) {
			return reply, nil
		}

		dest := proxy.Dest(handlerId)

		if !reply.IsOK() {
			return reply, nil
		}

		req.Next("list", key_value.New())
		listReply, err := dest.Client.Request(req)
		if err != nil {
			return nil, err
		}
		if !listReply.IsOK() {
			return nil, fmt.Errorf("list reply error: %s", listReply.ErrorMessage())
		}

		for key, param := range reply.ReplyParameters() {
			if listReply.ReplyParameters().Exist(key) {
				continue
			}
			listReply.ReplyParameters().Set(key, param)
		}

		return listReply, nil
	}

	err = proxy.SetReplyHandler(onReply)
	if err != nil {
		panic(err)
	}

	wg, err := proxy.Start()
	if err != nil {
		panic(err)
	}

	wg.Wait()
}
