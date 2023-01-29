package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"

	"gophercises/quietHN/hn"
)

func main() {
	// parse flags
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "the number of top stories to display")
	flag.Parse()

	tpl := template.Must(template.ParseFiles("./index.gohtml"))

	http.HandleFunc("/", handler(numStories, tpl))

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	sc:=storyCache{
		numStories: 40,
		duration: 1*time.Minute,
	}
	//keep rotating cache
	go func(){
		ticker:=time.NewTicker(30*time.Second)
		for {
			<-ticker.C
			temp := storyCache{
				numStories: 40,
				duration:   1 * time.Minute,
			}
			//this will refresh temp cache
			temp.stories()
			sc.mutex.Lock()
			sc.cache = temp.cache
			sc.expiration = time.Now().Add(sc.duration)
			sc.mutex.Unlock()
		}
	}()
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		stories, err:=sc.stories()
		if err!=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		stories=stories[0:30]
		data := templateData{
			Stories: stories,
			Time:    time.Now().Sub(start),
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
	}
}

func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(hnItem hn.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

// item is the same as the hn.Item, but adds the Host field
type item struct {
	hn.Item
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}
type storyCache struct{
	cache []item
	expiration time.Time
	duration time.Duration
	numStories int
	mutex *sync.Mutex
}
func (sc *storyCache) stories() ([]item, error){
	sc.mutex.Lock()
	defer sc.mutex.Unlock()
	if sc.expiration.After(time.Now()){
		return sc.cache, nil
	}else{
		cache, err :=getTopItems(sc.numStories)
		if err!=nil{
			return nil, err
		}
		sc.cache=cache
		sc.expiration=time.Now().Add(sc.duration)
		return sc.cache,nil
	}
}

func getTopItems(numStories int) ([]item, error){
	var client hn.Client
	type result struct{
		idx int
		item item
		err error
	}
	resultCh:=make(chan result)
	ids, err := client.TopItems()
	if err != nil {
		return nil, fmt.Errorf("Failed to load top stories")
	}
	var stories []item
	fetchStory:=func(idx, id int){
		hnItem, err:=client.GetItem(id)
		if err!=nil{
			resultCh<- result{idx:idx, err:err}
		}else{
			item:=parseHNItem(hnItem)
			if isStoryLink(item){
				resultCh<-result{idx:idx, item: item}
			}else{
				resultCh<-result{idx:idx, err: fmt.Errorf("No story here")}
			}
		}
	}
	for i:=0;i<numStories;i++{
		id:=ids[i]
		go fetchStory(i, id)
	}
	var results []result

	for i:=0;i<numStories;i++{
		results=append(results, <-resultCh)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].idx<results[j].idx
	})

	for _, res:=range results{
		if res.err==nil{
			stories=append(stories, res.item)
		}
	}

	return stories, nil
}