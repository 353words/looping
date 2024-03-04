package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Log struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func logHandler(r io.Reader) error {
	dec := json.NewDecoder(r)
	for {
		var log Log
		err := dec.Decode(&log)
		switch {
		case errors.Is(err, io.EOF):
			goto done
		case err != nil:
			return err
		default:
			fmt.Printf("log: %+v\n", log)
		}
	}

done:
	return nil
}

var data = `
{"level": "info", "message": "server starting"}
{"level": "info", "message": "server ready on port 8080"}
{"level": "error", "message": "/update - bad login"}
`

func main() {
	logHandler(strings.NewReader(data))
}
