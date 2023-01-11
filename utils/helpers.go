package utils

import (
	"encoding/json"
)

func ConvertToJson(toPrint any) any {
	jsonData, err := json.Marshal(toPrint) //  used to convert a Go value into a BYTE SLICE in a specific encoding format. Usually to JSON, can also be XML, time, etc
	if err != nil {
		panic(err)
	}

	jsonString := string(jsonData)                      // Convert the byte slice to stringified JSON
	return json.Unmarshal([]byte(jsonString), &toPrint) // Convert the stringified JSON to JSON
}

func ConvertByteToJson(toPrint []byte) any {

	jsonString := string(toPrint)                       // Convert the byte slice to stringified JSON
	return json.Unmarshal([]byte(jsonString), &toPrint) // Convert the stringified JSON to JSON
}
