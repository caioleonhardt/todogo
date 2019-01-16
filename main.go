package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/caioleonhardt/gorpc/todo"
	"github.com/golang/protobuf/proto"
)

func main() {

	todo := &todo.Todo{
		Title:       "Todo",
		Description: "My simple description using protobuf",
	}

	fmt.Println("GRPC:", todo)
	out, err := proto.Marshal(todo)
	if err != nil {
		log.Fatalln("Failed to encode TODO:", err)
	}
	if err := ioutil.WriteFile("/tmp/file", out, 0644); err != nil {
		log.Fatalln("Failed to write TODO:", err)
	}
}
