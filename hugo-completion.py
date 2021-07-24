import sublime
import sublime_plugin
import urllib.request
import json



AUTOCOMPLETE_DAEMON_EP = "http://localhost:5050"
TAGS_URL = f"{AUTOCOMPLETE_DAEMON_EP}/tags"
CATEGORIES_URL = f"{AUTOCOMPLETE_DAEMON_EP}/categories"


class MyCompletionsListener(sublime_plugin.EventListener):
    def on_query_completions(self, view, prefix, locations):
        loc = locations[0]
        if view.match_selector(loc, "meta.frontmatter.markdown"):   
            current_line = view.substr(view.line(view.sel()[0]))
            frontmatter_type = current_line.split(":")[0]
            completions = []
            if frontmatter_type == "tags":
                completions = fetch_completions(TAGS_URL)
            elif frontmatter_type == "category":
                completions = fetch_completions(CATEGORIES_URL)
            return completions
        return ""
        #======================================================
        #         resp = urllib.request.urlopen(tags_url)
        #         # resp = requests.get("http://localhost:5050/tags")
        #         if resp.getcode() == 200 :
        #             jsonData = json.loads(resp.read())
        #             for tag in jsonData:
        #                 completions.append(tag["name"])
        #         else:
        #             completions.append("Error in fetching data")
        #         #tags = resp.json()
        #         #for tag in tags:
        #         #    pass
        #             #completions.append(tag["name"])
        #         return completions
        # return ""
        #======================================================
        # limit you completions
        
        #======================================================
        #completions = [(v + "\tYour Description", v) for v in arr]
        # completions.append(view.substr(view.line(view.sel()[0])))
        # completions.append(frontmatter_type)
        # if view.match_selector(loc, "meta.frontmatter.markdown"):
        #     return completions
        # return ""
        #======================================================



def fetch_completions(url):
    completions = []
    resp = urllib.request.urlopen(url)
    if resp.getcode() == 200 :
        jsonData = json.loads(resp.read())
        for tag in jsonData:
            completions.append(tag["name"])
    else:
        completions.append("Error in fetching data")
    return completions
      