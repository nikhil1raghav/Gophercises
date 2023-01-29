package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gophercises/task/db"
	"os"
)

var ListCmd = &cobra.Command{
	Use: "list",
	Short: "list all incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err:=db.AllTasks()
		if err!=nil{
			fmt.Println("Couln't get tasks ", err)
			os.Exit(1)
		}
		if len(tasks)==0{
			fmt.Println("No tasks to complete. Yay!")
			return
		}
		fmt.Println("Following tasks incomplete")
		for i, task:=range tasks{
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},

}
func init(){
	RootCmd.AddCommand(ListCmd)
}
