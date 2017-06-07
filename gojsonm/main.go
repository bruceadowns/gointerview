package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type myjson struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (m *myjson) String() string {
	return fmt.Sprintf("%d: %s", m.ID, m.Name)
}

func main() {
	j := &myjson{1, "foo"}
	fmt.Printf("%s\n", j)

	m, err := json.Marshal(j)
	if err == nil {
		fmt.Printf("%s\n", m)
	} else {
		fmt.Printf("%s\n", err)
	}

	buf := bytes.Buffer{}
	buf.WriteString("{\"id\":1,\"name\":\"foo\"}")
	myj := &myjson{}
	err = json.Unmarshal(buf.Bytes(), &myj)
	if err == nil {
		fmt.Printf("%s\n", myj)
	} else {
		fmt.Printf("%s\n", err)
	}
}
