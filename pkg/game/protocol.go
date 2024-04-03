package game

import (
	"fmt"
	"strconv"
	"strings"
)

type Object int

const (
	Bird Object = iota
	Pipe
)

// From go to lua
type Message struct {
	obj    Object
	param1 int
	param2 int
}

// B:24:32 -> Message{obj: bird, param1 : 24, param2: 32}
func MessageFromStr(str string) (Message, error) {
	str_splits := strings.Split(str, ":")
	var msg Message

	if str_splits[0] == "B" {
		msg.obj = Bird
	} else if str_splits[0] == "P" {
		msg.obj = Pipe
	} else {
		return msg, fmt.Errorf("Invalid first char:%s", str_splits[0])
	}

	param1, err := strconv.Atoi(str_splits[1])
	param2, err := strconv.Atoi(str_splits[2])

	if err != nil {
		return msg, fmt.Errorf("Invalid second or third fields:%s|%s", str_splits[1], str_splits[2])
	}
	msg = Message{obj: Bird, param1: param1, param2: param2}

	return msg, nil
}

func MessageToStr(msg *Message) (string, error) {
	str := ""
    if msg.obj == Bird {
        str += "B:"
    } else if msg.obj == Pipe {
        str += "P:"
    }else {
        return "", fmt.Errorf("Object %v is not defined", msg.obj)
    }

    str += strconv.Itoa(msg.param1) + ":" + strconv.Itoa(msg.param2) + "\n"
	return str, nil
}
