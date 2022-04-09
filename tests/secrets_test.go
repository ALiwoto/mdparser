package tests

import (
	"strings"
	"testing"

	"github.com/ALiwoto/mdparser/mdparser"
)

func TestSecrets(t *testing.T) {
	secret1 := "123456"
	secret2 := "7890"
	password := "9876543"

	mdparser.AddSecret(secret1, "$SECRET1")
	mdparser.AddSecret(secret2, "$SECRET2")
	mdparser.AddSecret(password, "$PASSWORD")

	md := mdparser.GetBold("hello there\n")
	md.Italic("an error happened with secret1: " + secret1)
	md.Bold("an error happened with secret2: " + secret2)
	md.Link("this is my password: "+password, "https://google.com")

	strValue := md.ToString()

	if strings.Contains(strValue, secret1) {
		t.Error("secret1 not censored")
	}

	if strings.Contains(strValue, secret2) {
		t.Error("secret2 not censored")
	}

	if strings.Contains(strValue, password) {
		t.Error("password not censored")
	}
}
