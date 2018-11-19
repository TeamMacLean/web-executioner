package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Params struct {
	Params string `json:"params"`
}

type Config struct {
	Port    uint16 `json:"port"`
	Command string `json:"command"`
}

func main() {
	//get config
	pwd, _ := os.Getwd()
	configJson, _ := os.Open(filepath.Join(pwd, "config.json"))

	var cfg Config
	err := json.NewDecoder(configJson).Decode(&cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.Port)

	//start server
	m := martini.Classic()

	m.Get("/**", func() (int, string) {
		return 404, "Nothing to GET here."
	})
	m.Post("/", func(w http.ResponseWriter, r *http.Request) {
		// Unmarshal
		var params Params
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		go run(cfg.Command, params)
	})


	m.RunOnAddr(fmt.Sprintf(":%d", cfg.Port))
	//m.Run() //run on default port (3000)
}

//taken from https://stackoverflow.com/questions/20437336/how-to-execute-system-command-in-golang-with-unknown-arguments
func run(cmd string, msg Params) {
	cmd = cmd + " " + msg.Params
	//fmt.Println("command is:", cmd)
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
}
