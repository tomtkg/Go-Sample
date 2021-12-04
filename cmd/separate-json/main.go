package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sqweek/dialog"
)

func main() {
	inputPath, _ := dialog.File().Filter("json file", "json").Load()
	byteArray, _ := ioutil.ReadFile(inputPath)

	var jsonObj interface{}
	_ = json.Unmarshal(byteArray, &jsonObj)
	for k, v := range jsonObj.(map[string]interface{}) {
		byteArray, _ = json.Marshal(map[string]interface{}{k: v})
		_ = ioutil.WriteFile(k+".json", byteArray, 0755)
		fmt.Println("Create " + k + ".json")
	}

	fmt.Print("Finished!\nPress the Enter key")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
