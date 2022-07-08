package main

import (
	"encoding/json"
)

type Struct1 struct {
	Type string `json:"type"`
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func UnmarshallWithoutRawMessage(n int) {
	jsonValue := json.RawMessage(`{
		"type": "test",
		"data": {
			"id": "1"
		}
	}`)
	for i := 0; i < n; i++ {
		var obj Struct1
		json.Unmarshal(jsonValue, &obj)
	}
}

type Struct2 struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}
type Struct2Data struct {
	ID string `json:"id"`
}

func UnmarshallWithRawMessage01(n int) {
	jsonValue := json.RawMessage(`{
		"type": "test",
		"data": {
			"id": "1"
		}
	}`)
	for i := 0; i < n; i++ {
		var obj Struct2
		json.Unmarshal(jsonValue, &obj)
	}
}

func UnmarshallWithRawMessage02(n int) {
	jsonValue := json.RawMessage(`{
		"type": "test",
		"data": {
			"id": "1"
		}
	}`)
	switch string(jsonValue) {
	case "":

	}
	for i := 0; i < n; i++ {
		var obj Struct2
		json.Unmarshal(jsonValue, &obj)

		var data Struct2Data
		json.Unmarshal(obj.Data, &data)
	}
}
