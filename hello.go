package main

import (
	"errors" 
	"strings"
	"net/http"
	"fmt"
)

func acrynmMap() map[string]string {
	m := make(map[string]string)
	m["fml"] = "f*** my life"
	m["lol"] = "laugh out loud"
	return m
}

func Translate(s string,a map[string]string) (string, error) {
	s = strings.ToLower(s)
	m := acrynmMap()

	for k, v := range a {
		m[k] = v
	}

	if val, ok := m[s]; ok {
		return val,nil
	}

	return "", errors.New(s + " not found")
}

func respond(w http.ResponseWriter, req *http.Request){ 
	param1 := req.URL.Query()["a"][0]
	res, _ := Translate(param1,nil)
	fmt.Fprintf(w,res )
}

func main() {
	http.HandleFunc("/", respond)
	http.ListenAndServe(":8090", nil)
}