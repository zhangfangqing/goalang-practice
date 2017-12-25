package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type text struct {
	Year       string
	Month      string
	Day        string
	Num        int
	Safe_title string
	Alt        string
	Img        string
}

func main() {
	for i := 1; i <= 571; i++ {
		url := "https://xkcd.com/" + strconv.Itoa(i) + "/info.0.json"
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		var t text
		if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
			log.Fatal(err)
			resp.Body.Close()
		}
		resp.Body.Close()
		fmt.Printf("%s-%s-%s	%d	%s	%s	%s\n", t.Year, t.Month, t.Day, t.Num, t.Safe_title, t.Alt, t.Img)
	}
}
