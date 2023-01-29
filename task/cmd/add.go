package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gophercises/task/db"
	"os"
	"strings"
)

var AddCmd = &cobra.Command{
	Use: "add",
	Short: "Adds a task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		task:=strings.Join(args, " ")
		_, err:=db.CreateTask(task)
		if err!=nil{
			fmt.Println("Error while adding task ", err)
			os.Exit(1)
		}
		fmt.Printf("Added task \"%s\" to Todo list\n",task)
	},

}
func init(){
	RootCmd.AddCommand(AddCmd)
}
