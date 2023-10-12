package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

func parseDo() {
	jsonFile, err := os.Open("dialogue.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	dst := &bytes.Buffer{}
	if err := json.Compact(dst, data); err != nil {
		panic(err)
	}

	file, err := os.OpenFile("dialogue.jsonl", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(append(dst.Bytes(), '\n'))
	if err != nil {
		panic(err)
	}
}
