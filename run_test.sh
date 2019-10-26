#!/bin/bash

go build && ./logdel -d testdata/rules -no-delete -base-date 2019-10-13
