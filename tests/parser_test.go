package tests_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/PaulSonOfLars/gotgbot/v2"
)

func testParseToMd01(t *testing.T, filename string) {
	m := &gotgbot.Message{}
	var b []byte
	var err error
	b, err = ioutil.ReadFile("tests" + string(os.PathSeparator) + filename)
	if err != nil {
		b, err = ioutil.ReadFile(filename)
		if err != nil {
			log.Println("Couldn't read " + filename + ": " + err.Error())
			return
		}
	}

	err = json.Unmarshal(b, m)
	if err != nil {
		log.Println("Couldn't read " + filename + ": " + err.Error())
		return
	}

	md := mdparser.ParseFromMessage(m.ReplyToMessage)
	if md == nil {
		t.Error("md is nil")
	}

	strValue := md.ToString()
	if strValue == "" {
		t.Error("strValue is empty")
		return
	}
}

func TestParseToMd01(t *testing.T) {
	dataFiles := []string{
		"data01.json",
		"data02.json",
	}
	for _, current := range dataFiles {
		testParseToMd01(t, current)
	}
}
