package main

import (
	"encoding/json"
	"errors"
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
	var cfg Config
	pwd, _ := os.Getwd()
	configPath := filepath.Join(pwd, "config.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {

		//warn of missing config
		fmt.Println(errors.New("config.json does not exist, running in demo mode: command=echo, port=3000"))

		//run in demo mode
		cfg = Config{3000, "echo"}

	} else {
		configJson, _ := os.Open(configPath)

		err := json.NewDecoder(configJson).Decode(&cfg)
		if err != nil {
			panic(err)
		}
	}

	//create server
	m := martini.Classic()

	//catch all GETs
	m.Get("/**", func() (int, string) {
		return 404, "Nothing to GET here."
	})
	//catch POST on '/'
	m.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var params Params
		err := json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		//run command
		go run(cfg.Command, params)
	})

	//start server
	m.RunOnAddr(fmt.Sprintf(":%d", cfg.Port))
}

func run(cmd string, params Params) (output string) {
	cmd = cmd + " " + params.Params

	//log input
	fmt.Println("input:", cmd)

	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	//log output
	fmt.Println("output:", strings.TrimSpace(string(out)))

	return strings.TrimSpace(string(out))
}
