package game

import (
	"fmt"
	"strconv"
	"strings"
)

type Cmd int

const (
	Bird Cmd = iota
	Pipe
	Start
	End
)

// From go to lua
// TODO: variable no of params
type Message struct {
	cmd    Cmd
	params []int
}

// B:24:32?P:12:22? ->
// Message{obj: bird, param1 : 24, param2: 32} and Message{obj: pipe, param1 : 12, param2: 22}

var stream string //TODO: SHould not use global var

func MessageFromStr(chunk string) (Message, error) {

	stream += chunk
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
		msg.cmd = Bird
	} else if obj == "P" {
		msg.cmd = Pipe
	} else {
		return msg, fmt.Errorf("Invalid first char:%s", str_splits[0])
	}

    for _, p_str := range str_splits[1:]{
        param, err := strconv.Atoi(strings.TrimSpace(p_str))
        if err != nil {
            return msg, fmt.Errorf("Invalid second or third fields:%+v", err)
        }
        msg.params = append(msg.params, param)
    }

	return msg, nil
}

func MessageToStr(msg *Message) (string, error) {
	str := ""
	switch msg.cmd {
	case Bird:
		str += "B:"
	case Pipe:
		str += "P:"
	case Start:
		str += "S:"
	case End:
		str += "E:"
	default:
		return "", fmt.Errorf("No obj %v\n", msg.cmd)
	}

    for _, param := range msg.params {
        str += strconv.Itoa(param) + ":"
    }
    str = str[:len(str)-1]
    str += "?"
	return str, nil
}
