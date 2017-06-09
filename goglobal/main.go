package main

import "fmt"

var globalMap map[int]string

//var globalMap map[int]string = make(map[int]string)
//var globalMap = make(map[int]string)

func init() {
	globalMap = make(map[int]string)
}

func main() {
	fmt.Printf("%+v\n", globalMap)

	globalMap[1] = "foo"
	globalMap[2] = "bar"
	fmt.Printf("%+v\n", globalMap)

	for k, v := range globalMap {
		fmt.Printf("%d: %s\n", k, v)
	}

	_str := "underscore"
	fmt.Printf("%s\n", _str)
}

/*
func getBackend() (*Backend,err){
  select {
  case be <- backendQueue:
    return be,nil
  case <-time.After(100*time.Millisecond):
    return new be
  }
}

func queueBackend(be) {
  select {
  case backendQueue <- be:
    // good
  case <-time.After(1 * time.Second):
    be.Close()
  }
}
*/
