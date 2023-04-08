package handler

import json2 "encoding/json"

func jsonError(msg string) []byte {
	error := struct {
		Message string `json:"message"`
	}{
		msg,
	}

	json, err := json2.Marshal(error)

	if err != nil {
		return []byte(err.Error())
	}
	return json
}
