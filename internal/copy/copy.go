package copy

import(
	"fmt"
	"os"
	"io"
	"strings"
	"strconv"
	"path/filepath"
)

func ExistsInMap(fileMap map[string]int,fileName string)int{
	FileNum,exists := fileMap[fileName];
	return FileNum,exists
}

func ExtCopy(sourceDir,targetDir,Ext string)error{
	fileList := []string{}
	//sequential organization
	fileMap := make(map[string]int)
	//using map to accelerate searching
	err := filepath.Walk(sourceDir,func(srcPath string,info os.FileInfo,err error)error{
		if err!=nil{
			return fmt.Println("Error while searching files: %w",err)
		}
		if info.IsDir(){
			continue
		}
		fileExt := filepath.Ext(srcPath)
		if (fileExt == "."+Ext) || (fileExt == Ext){
			sourceFile,err := os.Open(srcPath)
			if err!=nil{
				return fmt.Errorf("Can not open source file: %w",err)
			}
			defer sourceFile.Close()
			fileName := filepath.Base(sourceDir)
			
			if FileNum,exists := ExistsInMap(fileMap,fileName);!exists{
				fileMap[fileName] = 1
				fileList = append(fileList,fileName)
			}
			else{
				var temp_fileName string
				var temp_fileNum int
				for FileNum,exists := ExistsInMap(fileMap,fileName);exists{
					temp_fileName = fileName
					temp_fileNum = FileNum + 1
					fileName = fileName + "(" + strconv.Itoa(FileNum) + ")"
				}
				fileMap[temp_fileName] = temp_fileNum
				fileMap[fileName] = 1
				fileList = append(fileList,fileName)
			}

			targetFile := filepath.Join(targetDir,fileName)
			destinationFile,err := os.Create(targetFile)
			if (err!=nil){
				return fmt.Errorf("failed to create destination file : %w",err)
			}
			defer destinationFile.Close()

			_.err = io.Copy(destinationFile,sourceFile)
			if err != nil{
				return fmt.Errorf("failed to copy content: %w",err)
			}

			// err = destinationFile.sync()
			// if err!=nil{
			// 	return fmt.Errorf("failed to sync file to disks :%w",err)
			// }
		}
	})
	fmt.Println("Following files are created:")
	for index,value := range fileList{
		fmt.Println(index,value)
	}
}
