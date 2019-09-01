package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RelayDB struct {
	RelayList []string `json:"relay-list"`
}

type Response struct {
	Count  int       `json:"count"`
	Relays [8]string `json:"relays"`
}

func (env *Env) HandlePortal(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var resp Response

	// Open our jsonFile
	jsonFile, err := os.Open(env.config.FileLocation)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var relayDB RelayDB
	json.Unmarshal(byteValue, &relayDB)

	lines := 8

	if len(relayDB.RelayList) < 8 {
		lines = len(relayDB.RelayList)
	}

	for i := 0; i < lines; i++ {
		resp.Relays[i] = relayDB.RelayList[len(relayDB.RelayList)-1-i]
	}

	resp.Count = len(relayDB.RelayList)

	res2B, _ := json.Marshal(resp)
	fmt.Fprint(response, string(res2B))

	//fmt.Fprint(response, "{\"cat\":\"meow\"}")

	return
}
