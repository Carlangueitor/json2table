package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/olekukonko/tablewriter"
)


func main() {
	jsonBytes, _ := ioutil.ReadAll(os.Stdin)
	jsonData := handleInput(jsonBytes)
	table := createTable(jsonData)
	fmt.Print(table)
}

func handleInput(jsonBytes []byte) ([]map[string]interface{}) {
	var out []map[string]interface{}
	json.Unmarshal(jsonBytes, &out)
	return out
}

func createTable(data []map[string]interface{}) (string) {
	header := make([]string, 0)
	for key := range data[0] {
		header = append(header, key)
	}
	buffer := new(bytes.Buffer)

	table := tablewriter.NewWriter(buffer)
	table.SetHeader(header)

	for index := range data {
		var row []string;
		for key := range header {
			cell := fmt.Sprintf("%v", data[index][header[key]])
			row = append(row, cell)
		}
		table.Append(row)
	}
	table.Render()
	return buffer.String()
}
