package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
//	"github.com/lidozeneli/PennApps_Spr_2016/series_now"
)
type Titles struct {
	Title []string `json:"titles"`
}

var title_literal []string = []string{
	"alaska-the-last-frontier-exposed",
	"ax-men-logged-and-loaded",
	"beachfront-bargain-hunt",
	"galavant",
	"guys-grocery-games",
	"kc-undercover",
	"the-circus-inside-the-greatest-political-show-on-earth",
	"beachfront-bargain-hunt",
	"liv-maddie",
	"the-circus-inside-the-greatest-political-show-on-earth",
	"alaska-the-last-frontier",
	"ax-men",
	"caribbean-life",
	"downton-abbey-on-masterpiece",
	"hoarders",
	"keeping-up-with-the-kardashians",
	"shameless",
	"worst-cooks-in-america",
	"caribbean-life",
	"alaska-the-last-frontier",
	"billions",
	"cutthroat-kitchen",
	"hoarders",
	"hollywood-teen-medium",
	"island-life",
	"live-to-tell",
	"married-by-mom-and-dad",
	"mercy-street",
	"the-x-files",
	"island-life",
	"greatest-party-story-ever",
	}

func main() {
	/*stack, err := series_now.GetSeries("http://www.tvguide.com/new-tonight/")
	if err != nil {
		fmt.Printf("GetSeries: %v:", err)
	}
	valid := series_now.Validate(stack)*/
	fmt.Println("Starting JSON server on port :8080")
	http.HandleFunc("/series/series_titles", seriesTitles)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func seriesTitles(w http.ResponseWriter, r *http.Request) {
	t := Titles{}
	t.Title = title_literal
	
	output, _ := json.Marshal(&t)
	fmt.Fprintf(w, string(output))
}	
