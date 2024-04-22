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
	KeyPress
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

func MessageFromStr(chunk string) (*Message, error) {

	stream += chunk
	idx := strings.Index(stream, "?")
	if idx == -1 {
        return nil, nil
	}

	cmd := stream[:idx]
	stream = stream[idx+1:]

	str_splits := strings.Split(cmd, ":")
	var msg Message

	obj := strings.TrimSpace(str_splits[0])

	switch obj {
	case "B":
		msg.cmd = Bird
	case "P":
		msg.cmd = Pipe
	case "S":
		msg.cmd = Start
	case "E":
		msg.cmd = End
	case "K":
		msg.cmd = KeyPress
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
