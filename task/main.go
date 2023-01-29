package main

import (
	"gophercises/task/cmd"
	"gophercises/task/db"
	homeDir"github.com/mitchellh/go-homedir"
	"log"
	"path/filepath"
)

func main(){
	home,_:=homeDir.Dir()
	dbPath:=filepath.Join(home, "tasks.db")
	err:=db.Init(dbPath)
	if err!=nil{
		log.Println(err)
		return
	}
	cmd.RootCmd.Execute()


}
