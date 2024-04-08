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
    Start
)

// From go to lua
type Message struct {
	Obj    Object
	Param1 int
	Param2 int
}

// B:24:32?P:12:22? ->
// Message{obj: bird, param1 : 24, param2: 32} and Message{obj: pipe, param1 : 12, param2: 22}

var stream string //TODO: SHould not use global var

func MessageFromStr(chunk string) (Message, error) {

	stream += chunk
	//prev_stream := stream
	idx := strings.Index(stream, "?")
	if idx == -1 {
		return Message{}, nil
	}

	cmd := stream[:idx]
	stream = stream[idx+1:]

	str_splits := strings.Split(cmd, ":")
	var msg Message

	obj := strings.TrimSpace(str_splits[0])
    
	if obj == "B" {
		msg.Obj = Bird
	} else if obj == "P" {
		msg.Obj = Pipe
	} else {
		return msg, fmt.Errorf("Invalid first char:%s", str_splits[0])
	}

	param1, err := strconv.Atoi(strings.TrimSpace(str_splits[1]))
	param2, err := strconv.Atoi(strings.TrimSpace(str_splits[2]))

	if err != nil {
		return msg, fmt.Errorf("Invalid second or third fields:%s|%s", str_splits[1], str_splits[2])
	}

	msg.Param1 = param1
	msg.Param2 = param2

	return msg, nil
}

func MessageToStr(msg *Message) (string, error) {
	str := ""
	if msg.Obj == Bird {
		str += "B:"
	} else if msg.Obj == Pipe {
		str += "P:"
	} else if msg.Obj == Start {
		str += "S:" //start cmd
	}else {
        return "", fmt.Errorf("No obj %v\n", msg.Obj)
    }

	str += strconv.Itoa(msg.Param1) + ":" + strconv.Itoa(msg.Param2) + "?"
	return str, nil
}
