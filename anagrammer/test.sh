#!/bin/sh

go run main.go
go run main.go -dict=/usr/share/dict/words -phrase='new relic jive'

go run main.go -phrase='foo bar'
go run main.go -dict=words.us -phrase='foo bar'
