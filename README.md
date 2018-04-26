# gostack

Simple Stack Machine

## Installation

To install and deploy the source, you need to install these packages,

- golang

clone this repository and run.

## Test Run

There are 5 test files arranged in order of execution.
- tokenizer_test.go
- ast_element_test.go
- parser_test.go
- stack_test.go
- executor_test.go

```sh
$ cd src
$ go test
```

## Build

```sh
$ cd src/main
$ go build
```

## Deployment

```sh
$ cd src/main
$ ./main -h

Usage of ./main:
  -input string
    	input file (default "./input.txt")
```

```sh
$ ./main -input=input.txt
 4 4 8

$ ./main -input=input2.txt
 1 1 2
```
