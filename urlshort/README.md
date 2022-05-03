# Url shortener

Not a url shortener but a redirector. Reads url path bindings from a yaml file and then redirects the user to the url mapping read from file.

## Components

### Yamlhandler

Reads yaml file, creates a url map and then passes it on to Maphandler

Need to create a struct to unmarshal yaml correctly, then a method to build a map from parsed yaml.


### Maphandler

Reads the url mappings and decides which url to redirect the path to.
If path is not in the mapping, use a fallback handler

Redirection is done using `http.Redirect`


```mermaid
graph LR;
id1[Yaml file]--->id2[Yaml handler]--->id5[Path Map]--->id3[Map handler]--->id4[url]
```

