# Web Executioner
> define commands to be executed remotely with paramaters passed over the network... deploy at your own risk.

[![CircleCI](https://circleci.com/gh/TeamMacLean/web-executioner.svg?style=svg)](https://circleci.com/gh/TeamMacLean/web-executioner)

## Install
```
go get -v -t -d ./...
go test -v ./...
go build server.go
```

## Config
Copy `config.example.json` to `config.json` and change `port` to the port you wish to run the server on and `command` to the command you would like to run when a `POST` request is send to the server.

## Running
```
./server
```