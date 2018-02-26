/*
Execution instructions:

$ go run main.go
or
$ go run --race main.go
*/

package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

/*
Data structures for time aware map:

* key - string
* time aware value  - byte array, timestamp
* concurrent map - key to set of time aware values

Note that scope visibility is local to this file.
Upper-case (public scope) would be needed to expose.
*/

type keyType string

type valueType []byte

type timeAwareValueType struct {
	data      valueType
	timestamp time.Time
}

type values []timeAwareValueType

type dataMapType map[keyType]values

type timeAwareMap struct {
	sync.RWMutex
	data dataMapType
}

func buildKey(prefix string, i int, k string) keyType {
	return keyType(fmt.Sprintf("%s_%d_%s", prefix, i, k))
}

func buildTimeAwareValue(v valueType) (res timeAwareValueType) {
	res.data = v
	res.timestamp = time.Now()
	return
}

/*
Api:

* put(key, value)
* time-complexity is amortized O(1)
* (based on golang append builtin)

* get(key, timestamp)
* time-complexity is O(v) where v is number of values for key
* tho likely use-case will return latest value in O(1)
*/

func (tam *timeAwareMap) put(k keyType, v valueType) error {
	if len(k) == 0 {
		return fmt.Errorf("empty key")
	}

	tam.Lock()
	defer tam.Unlock()

	tav, ok := tam.data[k]
	if !ok {
		tav = make(values, 0)
	}
	tam.data[k] = append(tav, buildTimeAwareValue(v))

	return nil
}

func (tam *timeAwareMap) get(k keyType, ts time.Time) (valueType, error) {
	tam.RLock()
	defer tam.RUnlock()

	tav, ok := tam.data[k]
	if ok {
		for i := len(tav) - 1; i >= 0; i-- {
			if !tav[i].timestamp.After(ts) {
				return tav[i].data, nil
			}
		}
	}

	return nil, fmt.Errorf("unknown key")
}

/*
Database is defined and initialized as a global variable
*/

var db = timeAwareMap{data: make(dataMapType)}

/*
Main:

* its purpose is solely to exercise the time aware map
* these tests would generally be in a separate _test.go file
* leveraging the testing package using tdd best practice
*/

func main() {
	log.Print("begin")

	testBasic()
	testBasicTable(0)
	testTimeScew()
	testTimeTravel()
	testConcurrent()

	log.Printf("database dump")
	for k, v := range db.data {
		for _, vv := range v {
			log.Printf("%s: %s: %d", k, string(vv.data), vv.timestamp.UnixNano())
		}
	}

	log.Print("done")
}

/*
Test cases:

* testBasic - basic test from instructions
* testBasicTable - table driven test cases
* testTimeScew - forward and backward timestamps
* testTimeTravel - query at successive timestamps
* testConcurrent - concurrent r/w access
*/

func testBasic() {
	log.Print("basic test")

	//put('a', 'v1') @ timestamp=1
	//get('a', 1) -> 'v1'
	//get('a', 2) -> 'v1'
	//get('b', 2) -> None

	tk := buildKey("testBasic", 0, "a")
	tv := valueType{'v', '1'}
	if err := db.put(tk, tv); err != nil {
		log.Fatalf("put error: %s", err)
	}

	ts := time.Now()
	v, err := db.get(tk, ts)
	if err != nil {
		log.Fatalf("get(%s) error: %s", tk, err)
	}
	if !bytes.Equal(tv, v) {
		log.Fatalf("get(%s) error: %s != %s", tk, string(tv), string(v))
	}

	ts = time.Now()
	v, err = db.get(tk, ts)
	if err != nil {
		log.Fatalf("get(%s) error: %s", tk, err)
	}
	if !bytes.Equal(tv, v) {
		log.Fatalf("get(%s) error: %s != %s", tk, string(tv), string(v))
	}

	tk = "b"
	if _, err := db.get(tk, time.Now()); err == nil {
		log.Fatalf("get(%s) expected error", tk)
	}
}

func testBasicTable(idx int) {
	log.Printf("basic table test %d", idx)

	const prefix = "testBasicTable"
	tt := []struct {
		k keyType
		v valueType
	}{
		{k: buildKey(prefix, idx, "a"), v: []byte("v1")},
		{k: buildKey(prefix, idx, "a"), v: []byte("v2")},
		{k: buildKey(prefix, idx, "foo"), v: []byte{'b', 'a', 'r'}},
		{k: buildKey(prefix, idx, "answer"), v: []byte{42}},
		{k: "", v: nil},
		{k: buildKey(prefix, idx, "key_with_zero_array"), v: []byte{}},
		{k: buildKey(prefix, idx, "key_with_nil_array"), v: nil},
	}

	for _, t := range tt {
		if err := db.put(t.k, t.v); len(t.k) > 0 && err != nil {
			log.Fatalf("put error: %s", err)
		}

		v, err := db.get(t.k, time.Now())
		if len(t.k) > 0 && err != nil {
			log.Fatalf("get(%s) error: %s", t.k, err)
		}
		if !bytes.Equal(t.v, v) {
			log.Fatalf("get(%s) error: %s != %s", t.k, string(t.v), string(v))
		}
	}
}

func testTimeScew() {
	log.Print("time scew test")

	k := buildKey("testTimeScew", 0, "a")
	v := valueType("v1")
	if err := db.put(k, v); err != nil {
		log.Fatalf("put error: %s", err)
	}

	ts := time.Now()
	if _, err := db.get(k, ts); err != nil {
		log.Fatalf("get(%s) error: %s", k, err)
	}

	ts = ts.Add(-time.Second)
	if _, err := db.get(k, ts); err == nil {
		log.Fatalf("get(%s) expected error", k)
	}

	ts = ts.Add(time.Second * 2)
	if _, err := db.get(k, ts); err != nil {
		log.Fatalf("get(%s) error: %s", k, err)
	}
}

func testTimeTravel() {
	log.Print("time travel test")

	type valueTimeType struct {
		v  valueType
		ts time.Time
	}
	tt := make([]valueTimeType, 0)

	k := buildKey("testTimeTravel", 0, "a")
	for i := 0; i < 10; i++ {
		v := valueType(fmt.Sprintf("value_%d", i))

		if err := db.put(k, v); err != nil {
			log.Fatalf("put error: %s", err)
		}

		tt = append(tt, valueTimeType{v, time.Now()})

		time.Sleep(1)
	}

	for _, vt := range tt {
		v, err := db.get(k, vt.ts)
		if err != nil {
			log.Fatalf("get(%s) error: %s", k, err)
		}
		if !bytes.Equal(v, vt.v) {
			log.Fatalf("get(%s, %s) error: %s != %s", k, vt.ts, v, vt.v)
		}
	}
}

func testConcurrent() {
	log.Print("concurrent test")

	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		idx := i
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			testBasicTable(idx)
		}()
	}

	wg.Wait()
}

/*
Random feedback thoughts:

The problem took roughly 45m to think about and 45m to implement.
I spent more time on testing, prep, and polish than necessary.

The use of coderpad was confusing.
The problem was clear, but the delivery expectations were not.
i.e. 'deliver as a library' vs 'make sure it runs in coderpad'
My solution was unnecessarily shaped by the single-text-file deliverable.

Possible solutions may include:
* provide instructions as a separate document
* candidate deliver as a github repo link/tar
* provide a cockroach private git repo

Imho, you should expect a more targeted golang/git/oss understanding:
* golang is now mainstream
* gitflow/oss is key to your business model
* candidates should already have a dev env
*/

/*
https://coderpad.io/ZKR74KXG

## Submitting Your Solution

Use any language you're comfortable with and feel free to consult any
documentation, StackOverflow, etc. as you normally would in your day-to-day
work.

Pick your desired language from the drop-down above, or, if it isn't supported by
CoderPad, just let us know and include instructions on how to build and run your
solution on recent-ish Linux or MacOS systems.

You can develop and run your solution right in CoderPad if you want, or work in
your preferred environment and paste your code in when you're ready -- just make
sure that it does run in CoderPad before you're done. Once you're finished, just
let us know!

## Implementing a Time-Aware Map

For this exercise, create a map data structure that supports key-value
storage and retrieval. The structure should be time-aware, meaning that
it can maintain multiple values at different times for each key.

The structure's API should support the following operations:

- Put(key, value)
- Get(key, timestamp)

The time-aware map should be exposed as a self-contained library. Its
interface should be idiomatic in whatever language you chose to use.

## Retrieval Behavior

When retrieving a value for a key at a given timestamp, the map will
return the "most recent" value for that key. Here, "most recent" means
the value with the largest timestamp less than or equal to that passed
to Get.

Sane behavior should be defined for the cases where a key does not exist
in the map or where no value existed yet for the key at the specified
timestamp.

### Analysis

Above each operation, please leave a comment discussing its computational time
complexity.

## Testing

Here are a few test cases to get you started (though you will certainly want to
add more):
```
put('a', 'v1') @ timestamp=1

get('a', 1) -> 'v1'
get('a', 2) -> 'v1'
get('b', 2) -> None
```
*/
