package game

import (
	"fmt"
	"strconv"
	"strings"
)

type Cmd int

const (
	BirdCmd Cmd = iota
	PipeCmd
	StartCmd
	EndCmd
	KeyPressCmd
)

// From go to lua
// TODO: variable no of params
type Message struct {
	cmd    Cmd
	params []int
}

func MessageFromStr(cmd string) (*Message, error) {

	str_splits := strings.Split(cmd, ":")
	var msg Message

	obj := strings.TrimSpace(str_splits[0])

	switch obj {
	case "B":
		msg.cmd = BirdCmd
	case "P":
		msg.cmd = PipeCmd
	case "S":
		msg.cmd = StartCmd
	case "E":
		msg.cmd = EndCmd
	case "K":
		msg.cmd = KeyPressCmd
	default:
		return nil, fmt.Errorf("Invalid first char:%s", str_splits[0])
	}

	for _, p_str := range str_splits[1:] {
        p_str = strings.TrimSpace(p_str)
        if len(p_str) == 1 && !(p_str[0] >= '0' && p_str[0] <= '9'){
            // single char
            msg.params = append(msg.params, int(p_str[0]))
        }else{
            param, err := strconv.Atoi(strings.TrimSpace(p_str))
            if err != nil {
                return nil, fmt.Errorf("Invalid %+v\n", p_str)
            }
            msg.params = append(msg.params, param)
        }
	}

	return &msg, nil
}

func MessageToStr(msg *Message) (string, error) {
	str := ""
	switch msg.cmd {
	case BirdCmd:
		str += "B:"
	case PipeCmd:
		str += "P:"
	case StartCmd:
		str += "S:"
	case EndCmd:
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
