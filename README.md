# gopp
`gopp` is  very simple go package creation assistant

## Usage

```sh
// $GOPATH/src/github.com/trrrrrys/example
$ gopp | xargs go mod init
go: creating new go.mod: module github.com/trrrrrys/example

$ cat go.mod

module github.com/trrrrrys/example

go 1.19
```

If outside GOPATH
```sh
// ~/ghq/github.com/trrrrrys/example
$ gopp -s ~/ghq | xargs go mod init
go: creating new go.mod: module github.com/trrrrrys/example


$ cat go.mod
module github.com/trrrrrys/example

go 1.19
```
