package main

import (
	"flag"
	"fmt"
	"gophercises/cyoa"
	"log"
	"net/http"
	"os"
)

func main(){
	port:=flag.Int("port", 3000, "Port to start cyoa webapp")
	filename:=flag.String("file","gopher.json","Json file which contains the complete story")
	flag.Parse()
	fmt.Printf("Using %s for story \n", *filename)
	f, err:= os.Open(*filename)
	if err!=nil{
		log.Println("Error while opening file : ",err)
		return
	}
	story, err:=cyoa.GetStory(f)
	if err!=nil{
		log.Printf("Error parsing the story in file %s.\n", *filename)
		log.Println(err)
		return
	}
	h:=cyoa.NewHandler(story)
	fmt.Printf("Starting server on port :%d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}