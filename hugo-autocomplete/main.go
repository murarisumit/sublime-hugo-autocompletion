package main

import (
	"log"
	"net/http"
)

var TAGS_URL = "http://localhost:1313/tags-list.html"
var CATEGORIES_URL = "http://localhost:1313/categories-list.html"

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
	http.ListenAndServe(":5050", nil)

}
