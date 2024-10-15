#!/bin/bash

# Execute Go tests in the parent directory and all its subdirectories
reset && go test ./... $1