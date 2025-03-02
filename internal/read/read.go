package read

import (
	"fmt"
	"github.com/HanksJCTsai/XReader/internal/logger"
	"github.com/HanksJCTsai/XReader/internal/util"
	"os"
	"time"
)

func ReadFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return util.WrapError("File reading failed", err)
	}
	logger.Log.Info("File read successfully", "path", filePath)
	fmt.Println("File Contents:")
	fmt.Println(string(content))
	return nil
}

func GetFileInfo(filePath string, inclContents bool) error {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return util.WrapError("Get this file info", err)
	}
	logger.Log.Info("Got file info successfully", "path", filePath)
	fmt.Printf("File Name: %s\n", fileInfo.Name())
	fmt.Printf("File Size: %d 位元組\n", fileInfo.Size())
	fmt.Printf("Last modify time: %s\n", fileInfo.ModTime().Format(time.RFC1123Z))
	if inclContents {
		content, err := os.ReadFile(filePath)
		if err != nil {
			return util.WrapError("File reading failed", err)
		}
		logger.Log.Info("File read successfully", "path", filePath)
		fmt.Println("File Contents:")
		fmt.Println(string(content))
	}
	return nil
}
