package main

import (
	"fmt"
	"gophercises/secrets-cli"
)

func main(){
	v:=secrets.File("mykey", "./fileData")
	err:=v.Set("secret", "dirty-secret")
	if err!=nil{
		fmt.Println(err)
		return
	}

	val, err:=v.Get("secret")
	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println("Value of secret -> ", val)

}
