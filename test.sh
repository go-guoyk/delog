#!/bin/bash

go build && ./delog -d testdata/rules -dry
