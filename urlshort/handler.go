package urlshort

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type URLMap struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}

func buildMap(mappedYamlPaths []URLMap) map[string]string {
	pathMap := make(map[string]string)

	for _, mapper := range mappedYamlPaths {
		pathMap[mapper.Path] = mapper.URL
	}
	return pathMap
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		destinationUrl, found := pathsToUrls[r.URL.Path]
		if found {
			http.Redirect(w, r, destinationUrl, 301)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var mappedPaths []URLMap
	err := yaml.Unmarshal(yml, &mappedPaths)
	if err != nil {
		return nil, err
	}
	pathMaps := buildMap(mappedPaths)
	return MapHandler(pathMaps, fallback), nil
}
