package main

import(
	"fmt"
	"os"
	"log"
	"github.com/spf13/cobra"
	"github.com/crualcollegee/FileCopyCat/blob/main/internal/copy/copy.go"
)

var sourceDir string
var targetDir string
var Ext string

func main()  {
	var rootCmd = &cobra.Command{
		Use: "file-organization",
		Short: "manage files with CLI",
	}

	var ExtCopy = &cobra.Command{
		Use:"ExtCopy [source directory] [target directory] [extension]",
		Short: "find files with specific extension and copy to a directory"
		Args: cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string){
			if err:= copy.ExtCopy(sourceDir,targetDir,Ext);err!=nil{
				log.Fatalf("Error during copying files : %v",err)
			}
		}
	}

	rootCmd.AddCommand(ExtCopy)

	if err := rootCmd.Execute();err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}