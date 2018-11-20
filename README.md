# Web Executioner
> define commands to be executed remotely with parameters passed over the network... deploy at your own risk.

[![CircleCI](https://circleci.com/gh/TeamMacLean/web-executioner.svg?style=svg)](https://circleci.com/gh/TeamMacLean/web-executioner)


## Usage
Sending a `POST` request with a `json` body to the server will cause the command defined in your `config.json` file.

Your request should have a parameter called "params" which will be appended onto the command and ran, for example if your `POST` request is `{"params":"hello world"}` and in `config.json` your command is `echo`, when the request is received `echo hello world` will be executed on the server.

## Install
```
go get -v -t -d ./...
go test -v ./...
go build server.go
```

## Config
Copy `config.example.json` to `config.json` and change `port` to the port you wish to run the server on and `command` to the command you would like to run when a `POST` request is send to the server.

## Run
```
./server
```