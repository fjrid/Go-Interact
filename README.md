# Go Interact

This tool helps you to make interactive exec module in Golang.

## Instalation

You can easily install with this command

```sh
go get github.com/freekup/go-interact
```

## Usage

To use this package you need to import the package

```go
import "github.com/freekup/go-interact"
```

Then you can initialize command struct with this code

```go
    /*
        command := interact.Initiate(command, ...arguments)
    */
    command := interact.Initialize("ls", "-a", "-l")
```

To execute the command you can run with this code

```go
    err := command.Run()
```

You will see the interactive result. If you want to hide the output, you can set the configuration `silent`. This is the example

```go
    /*
        command := interact.Initiate(command, ...arguments)
    */
    command := interact.Initialize("ls", "-a", "-l")
    command.Silent = true
```

## Features
### Convert to String
If you want to see complete command to string, you can easyly using `String()` function.
```go
    /*
        command := interact.Initiate(command, ...arguments)
    */
    command := interact.Initialize("ls", "-a", "-l")
    strCommand := command.String()
```