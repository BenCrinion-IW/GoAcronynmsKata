package main

import (
    "testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"io"
)

func TestCanCallTranslateWithLol(t *testing.T) {
	res, _ := Translate("lol", nil)
	assert.Equal(t, "laugh out loud", res, "they should be equal")
}

func TestCanCallTranslateWithFml(t *testing.T) {
	m := make(map[string]string)
	res, _ := Translate("fml",m)
	assert.Equal(t, "f*** my life", res, "they should be equal")
}

func TestCanCallTranslateWithNotKnown(t *testing.T) {
	m := make(map[string]string)
	_, error := Translate("hrh",m)
	assert.NotNil(t, error)
	assert.Equal(t, "hrh not found", error.Error(), "Error message isn't correct")
}

func TestCanCallTranslateWithWordList(t *testing.T) {
	m := make(map[string]string)
	m["ram"] = "random access memory"
	res, _ := Translate("ram", m)
	assert.Equal(t, "random access memory", res, "they should be equal")
}

func TestCanCallTranslateWithWordListWithMissingValue(t *testing.T) {
	m := make(map[string]string)
	m["ram"] = "random access memory"
	_, error := Translate("bob", m)
	assert.NotNil(t, error)
	assert.Equal(t, "bob not found", error.Error(), "Error message isn't correct")
}

func TestCanCallTranslateUpperCaseAccronym(t *testing.T) {
	m := make(map[string]string)
	res, _ := Translate("FML", m)
	assert.Equal(t, "f*** my life", res, "they should be equal")
}

func TestGetFromServer(t *testing.T){
	resp, _ := http.Get("http://localhost:8090/?a=lol")
	assert.NotNil(t, resp)

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, "laugh out loud", string(body))
}