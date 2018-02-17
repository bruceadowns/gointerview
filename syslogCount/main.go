package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	kv := make(map[time.Time]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)
		if len(fields) < 3 {
			log.Fatal("invalid input")
		}

		m := fields[0]
		d := fields[1]
		t := fields[2]
		//log.Printf("%s %s %s", m, d, t)
		ti, err := time.Parse("Jan 2 15:04:05", fmt.Sprintf("%s %s %s", m, d, t))
		if err != nil {
			log.Fatal(err)
		}

		kv[ti.Round(time.Minute)]++
	}

	for k, v := range kv {
		fmt.Printf("%s %d\n", k, v)
	}
}
