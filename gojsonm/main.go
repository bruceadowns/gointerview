package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type myjson struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (m *myjson) String() string {
	return fmt.Sprintf("%d: %s", m.ID, m.Name)
}

/*

b = ioutil.ReadAll(os.Stdin)
json.Unmarshal(b, &e)

vs

// streaming decoder
json.NewDecoder(os.Stdin).Decode(&e)

*/

func streamDecoder() {
	//buf := bytes.Buffer{}
	//buf.WriteString("{\"id\":1,\"name\":\"foo\"}")
	buf := bytes.NewBufferString("{\"id\":2,\"name\":\"foobar\"}")

	j := &myjson{}
	if err := json.NewDecoder(buf).Decode(&j); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", j)
}

func strMarshaller() {
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

func main() {
	streamDecoder()
	strMarshaller()
}
