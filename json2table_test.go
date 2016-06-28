package main

import (
	"testing"
)


func TestHandleInput(t *testing.T) {
	json := "[ { \"foo\": \"bar\"}, { \"foo\": \"candy\"} ]"
	out := handleInput([]byte(json))
	if(out[0]["foo"] != "bar") {
		t.Error("Unable to Parse JSON")
	}
}

func TestCreateTable(t *testing.T) {
	expTable := "+-------+\n|  FOO  |\n+-------+\n| candy |\n+-------+\n"
	json := make([]map[string]interface{}, 0)
	var val interface{} = "candy"
	object := map[string]interface{}{
		"foo": val,
	}
	json = append(json, object)

	t.Logf("json: %v", json)

	table := createTable(json)

	t.Logf("Expecting generated table \n%s\n is equal to expected table \n%s\n",
		table, expTable)
	if(table != expTable) {
		t.Error("Table incorrect")
	}
}
