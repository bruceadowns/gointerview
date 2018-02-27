#!/bin/sh

cat test.txt | cut -b-12 | sort | uniq -c
