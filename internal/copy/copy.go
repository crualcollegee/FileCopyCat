package copy

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ExistsInMap(fileMap map[string]int, fileName string) (int, bool) {
	FileNum_, exists := fileMap[fileName]
	return FileNum_, exists
}

func ExtCopy(sourceDir, targetDir, Ext string, keep bool) error {
	fileList := []string{}
	//sequential organization
	fileMap := make(map[string]int)
	//using map to accelerate searching
	err := filepath.Walk(sourceDir, func(srcPath string, info os.FileInfo, err error) error {
		// /*debug*/ fmt.Println("sourceDir :" ,sourceDir ,", srcPath: ",srcPath)
		if err != nil {
			return fmt.Errorf("Error while searching files: %w", err)
		}
		if info.IsDir() {
			return nil
		}
		fileExt := filepath.Ext(srcPath)
		if (fileExt == "."+Ext) || (fileExt == Ext) {
			sourceFile, err := os.Open(srcPath)
			if err != nil {
				return fmt.Errorf("Can not open source file: %w", err)
			}
			defer sourceFile.Close()
			fileName := filepath.Base(srcPath)

			if _, exists := ExistsInMap(fileMap, fileName); !exists {
				fileMap[fileName] = 1
				fileList = append(fileList, fileName)
			} else {
				var temp_fileName string
				var temp_fileNum int
				for FileNum, exists := ExistsInMap(fileMap, fileName); exists; {
					temp_fileName = fileName
					temp_fileNum = FileNum + 1
					Name_withoutExt := strings.TrimSuffix(fileName, filepath.Ext(fileName))
					fileName = Name_withoutExt + "(" + strconv.Itoa(FileNum) + ")" + filepath.Ext(fileName)
					FileNum, exists = ExistsInMap(fileMap, fileName)
				}
				fileMap[temp_fileName] = temp_fileNum
				fileMap[fileName] = 1
				fileList = append(fileList, fileName)
			}

			targetFile := filepath.Join(targetDir, fileName)
			destinationFile, err := os.Create(targetFile)
			if err != nil {
				return fmt.Errorf("failed to create destination file : %w", err)
			}
			defer destinationFile.Close()

			_, err = io.Copy(destinationFile, sourceFile)
			if err != nil {
				return fmt.Errorf("failed to copy content: %w", err)
			}

			err = destinationFile.Sync()
			if err != nil {
				return fmt.Errorf("failed to sync file to disks :%w", err)
			}

			if !keep {
				err := os.Remove(srcPath)
				if err != nil {
					fmt.Errorf("failed to delete source file:%w", err)
				}
			}

		}
		return nil
	})
	if err != nil {
		fmt.Errorf("failed to Walk:%w", err)
	}
	fmt.Println("Following files are created:")
	for index, value := range fileList {
		fmt.Println(index, value)
	}
	return nil
}
