package main

include(
	"fmt"
	"os"
	"github.com/spf13/cobra"
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
		short: "find files with specific extension and copy to a directory"
		Args: cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string){
			if err:= CopyFile(sourceDir,targetDir,Ext);err!=nil{
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