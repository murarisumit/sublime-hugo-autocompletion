package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Tag struct {
	Name string `json:"name,omitempty"`
}

type Category struct {
	Name string `json:"name,omitempty"`
}

// I know interface should end with -er, but this made sense to me for now, till I learn better
type HugoCompleter struct {
	tagsUrl       string
	categoriesUrl string
}

func (a HugoCompleter) GetTagsJson() ([]byte, error) {
	tags, err := a.fetchTags()
	if err != nil {
		log.Println("Error in fetching the tags")
		return nil, err
	}
	tags_json, err := json.Marshal(tags)

	if err != nil {
		log.Fatalf("Error in getting tags json format")
		return nil, err
	}
	return tags_json, nil
}

func (a HugoCompleter) GetCategoriesJson() ([]byte, error) {
	categories, err := a.fetchCategories()
	if err != nil {
		log.Println("Error in fetching the categories")
		return nil, err
	}
	categories_json, err := json.Marshal(categories)

	if err != nil {
		log.Fatalf("Error in getting categories json format")
		return nil, err
	}
	return categories_json, nil
}

func (a HugoCompleter) fetchTags() ([]Tag, error) {
	log.Printf("Fetching tags from %s \n", a.tagsUrl)
	completionData, err := a.fetchCompletionData(a.tagsUrl)
	if err != nil {
		return nil, err
	}
	var tags []Tag
	for _, d := range completionData {
		tag := Tag{Name: strings.TrimSpace(d)}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (a HugoCompleter) fetchCategories() ([]Category, error) {
	log.Printf("Fetching categories from %s \n", a.categoriesUrl)
	completionData, err := a.fetchCompletionData(a.categoriesUrl)
	if err != nil {
		return nil, err
	}
	var categories []Category
	for _, d := range completionData {
		category := Category{Name: strings.TrimSpace(d)}
		categories = append(categories, category)
	}
	return categories, nil
}

func (a HugoCompleter) fetchCompletionData(url string) ([]string, error) {
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		log.Printf("Error in fetching url data from : %v", url)
		return nil, err
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error in reading data from url : %v", url)
		return nil, err
	}
	v := strings.Split(string(html), ",")
	return v, nil
}
