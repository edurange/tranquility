package main

import (
	"net/http"
	"os/exec"
	"os"
	"encoding/json"
	"time"
	"bytes"
	"strings"
)

type RO struct {
	User string `json:"user"`
	Time int64 `json:"time"`
	Command string `json:"command"`
}

func main(){
	var record RO

	url := os.Getenv("url")
	uuid := os.Getenv("UUID")

	cmd := exec.Command("whoami")				//current user
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	record.User = string(out)

	t := time.Now()								//current time
	record.Time = t.Unix()

	listy := os.Args[1:]
	record.Command = strings.Join(listy, " ")

	send, _ := json.Marshal(&record)

	logger := &http.Client{}
	loggerReq, err := http.NewRequest("POST", url + "/logger", bytes.NewBuffer(send))
	if err != nil{
		panic(err)
	}
	loggerReq.Header.Add("Content-Type", "application/json")
	loggerReq.Header.Add("uuid", uuid)

	//send json
	_, loggerErr := logger.Do(loggerReq)
	if loggerErr != nil{
		panic(loggerErr)
	}
}