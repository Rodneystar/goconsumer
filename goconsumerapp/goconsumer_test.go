package goconsumerapp

import (
	"encoding/json"
	"os"
	"testing"
	"text/template"
)

type kval map[string]interface{}

func Test_jsonGotemplateMap(t *testing.T) {

	jsonstr := "{ \"key1\": \"val1\",\"key2\": \"val2\" }"

	kvalMap := make(kval)
	json.Unmarshal([]byte(jsonstr), &kvalMap)

	for k, v := range kvalMap {
		t.Logf("key: %s, value %s,\n", k, v)
	}

	if kvalMap["key1"] != "val1" {
		t.Errorf("key1 should be val1, but is %s", kvalMap["key1"])
	}
}

func Test_jsonGotemplateMapNested(t *testing.T) {

	jsonstr := "{\n    \"key1\": \"val1\",\n    \"key2\": {\n        \"subkey1\": \"subval1\",\n        \"subkey2\": \"subval2\"\n    }\n}"

	kvalMap := make(kval)
	json.Unmarshal([]byte(jsonstr), &kvalMap)

	for k, v := range kvalMap {
		t.Logf("key: %s, value %s,\n", k, v)
	}

	if kvalMap["key1"] != "val1" {
		t.Errorf("key1 should be val1, but is %s", kvalMap["key1"])
	}

	template.JSEscape(os.Stdout, []byte("{ \"hey\" : \"there\" }"))
	tmp, _ := template.New("json").Parse("{{.key2.subkey2}}")
	tmp.Execute(os.Stdout, kvalMap)
}
