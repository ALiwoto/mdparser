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

func TestParseToMd01(t *testing.T) {
	m := &gotgbot.Message{}
	var b []byte
	var err error
	b, err = ioutil.ReadFile("tests" + string(os.PathSeparator) + "data01.json")
	if err != nil {
		b, err = ioutil.ReadFile("data01.json")
		if err != nil {
			log.Println("Couldn't read data01.json")
			return
		}
	}

	err = json.Unmarshal(b, m)
	if err != nil {
		log.Println("Couldn't unmarshal data01.json")
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
