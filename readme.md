# IP Seeker

A simple application to apply concepts learned in the course ["Learn Golang from Scratch! Develop a COMPLETE APPLICATION!"](https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa) (pt-Br).

It's a Command Line Tool (**CLI**) to find the IP and Server name based on a given host.

## Dependencies

It was built using GoLang version 1.21.4.

This application uses external dependencies:

- github.com/urfave/cli

## How to use

You can compile it and then run the program:

```sh
go build . &&\
./ip-seeker
```

To run directly from the compiler, you can run:

```sh
go run .
```

### Parameters

```sh
NAME:
   IP Seeker - Seek IP and server information from the Internet

USAGE:
   ip-seeker [global options] command [command options] [arguments...]

COMMANDS:
   ip       Seek IP information from the Internet
   server   Seek server information from the Internet
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
   ``````

## Author

<img src="https://avatars.githubusercontent.com/u/5307216?v=4" width="50" height="50">

[Fabio Almeida](https://github.com/fabioalmeidaweb)
