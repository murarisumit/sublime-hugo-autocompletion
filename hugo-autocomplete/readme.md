Prereq:
* It expects data from hugo for autocompletion. It gets data from `tags.html` and `categories.html` at root directory of generated static site.
* Hugo server should be running, else it would not be able to fetch above data.


A page that return all the tags for hugo, you can use below hugo snippet to get all tags and categories. 
```
{{range $name, $taxonomy := .Site.Taxonomies.tags}}
	{{$name}},
{{end}}
end
```
```
{{range $name, $taxonomy := .Site.Taxonomies.category}}
	{{$name}},
{{end}}
end # this is delimiter to know it's end now.
```


## To run the program
make run


### To get tags
curl localhost:5050/tags

### To get categories
curl localhost:5050/categories

