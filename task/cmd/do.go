package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gophercises/task/db"
	"os"
	"strconv"
)

var DoCmd = &cobra.Command{
	Use: "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg:=range args{
			id, err:=strconv.Atoi(arg)
			if err!=nil{
				fmt.Printf("Failed to parse the argument %s.\n", arg)
			}else{
				ids=append(ids, id)
			}
		}
		tasks, err:=db.AllTasks()
		if err!=nil{
			fmt.Println("Something went wrong")
			os.Exit(1)
		}
		for _,id:=range ids{
			if id<=0 || id>len(tasks){
				fmt.Printf("Invalid id %d.", id)
				continue
			}
			task:=tasks[id-1]
			err:=db.DeleteTask(task.Key)
			if err!=nil{
				fmt.Printf("Couldn't mark task %d as complete : %s\n ", id, err)
			}else{
				fmt.Printf("Marked %d as completed\n",id)
			}

		}

		fmt.Println("Task marked as done")
	},

}
func init(){
	RootCmd.AddCommand(DoCmd)
}
