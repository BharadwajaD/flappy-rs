package server

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)


func TestTCPReadString(t *testing.T){
    sreader := strings.NewReader("K:")
    breader := bufio.NewReader(sreader)
    trw := tcprw {rw: bufio.NewReadWriter(breader, nil)}

    str, err := trw.ReadString('?')
    fmt.Printf("Message: %s and err: %+v\n", str, err)


    sreader.Reset("k?K:q?")
    str, err = trw.ReadString('?')
    fmt.Printf("Message1: %s and err: %+v\n", str, err)

    str, err = trw.ReadString('?')
    fmt.Printf("Message2: %s and err: %+v\n", str, err)
}
