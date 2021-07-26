package main

import (
	"fmt"
	"log"
	"net/http"
)

var TAGS_URL = "http://localhost:1313/tlist.html"
var CATEGORIES_URL = "http://localhost:1313/clist.html"

func main() {
	hugoCompleter := HugoCompleter{
		tagsUrl:       TAGS_URL,
		categoriesUrl: CATEGORIES_URL,
	}

	http.HandleFunc("/tags", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got requests for fetching tags list")
		tags, err := hugoCompleter.GetTagsJson()
		if err != nil {
			log.Println("Error in tagsJson")
		}
		w.Write(tags)
	})

	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got requests for fetching categories list")
		categories, err := hugoCompleter.GetCategoriesJson()
		if err != nil {
			log.Println("Error in getCategoriesJson")
		}

		w.Write(categories)
	})

	// listen to port

	addr := "localhost:5050"
	log.Println("Listing on 5050 port.")
	log.Println(banner(addr))
	http.ListenAndServe(addr, nil)

}

func banner(addr string) string {
	fetchTags := addr + "/tags"
	fetchCategories := addr + "/categories"
	banner := "\n=========Usage=============\n"
	banner += fmt.Sprintf("Fetch Tags: %v \n", fetchTags)
	banner += fmt.Sprintf("Fetch Categories: %v\n", fetchCategories)
	banner += "=========================="

	return banner
}
