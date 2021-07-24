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
	tags := a.fetchTags()
	tags_json, err := json.Marshal(tags)

	if err != nil {
		return nil, err
	}
	return tags_json, nil
}

func (a HugoCompleter) GetCategoriesJson() ([]byte, error) {
	tags := a.fetchCategories()
	tags_json, err := json.Marshal(tags)

	if err != nil {
		return nil, err
	}
	return tags_json, nil
}

func (a HugoCompleter) fetchTags() []Tag {
	log.Printf("Fetching tags from %s \n", a.tagsUrl)
	completionData := a.fetchCompletionData(a.tagsUrl)
	var tags []Tag
	for _, d := range completionData {
		tag := Tag{Name: strings.TrimSpace(d)}
		tags = append(tags, tag)
	}

	return tags
}

func (a HugoCompleter) fetchCategories() []Category {
	log.Printf("Fetching categories from %s \n", a.categoriesUrl)
	completionData := a.fetchCompletionData(a.categoriesUrl)
	var categories []Category
	for _, d := range completionData {
		category := Category{Name: strings.TrimSpace(d)}
		categories = append(categories, category)
	}
	return categories
}

func (a HugoCompleter) fetchCompletionData(url string) []string {
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	v := strings.Split(string(html), ",")
	return v
}
