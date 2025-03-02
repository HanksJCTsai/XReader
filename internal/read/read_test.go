package read

import (
	"github.com/HanksJCTsai/XReader/internal/config"
	"github.com/HanksJCTsai/XReader/internal/logger"
	"path/filepath"
	"testing"
)

func setupTest(t *testing.T) {
	// 設置測試環境
	config.Cfg = &config.Config{LogLevel: "info"}
	logger.InitLogger()
}

func TestReadFile(t *testing.T) {
	setupTest(t)
	testFilePath := filepath.Join("..", "..", "testdata", "sample.txt")
	if err := ReadFile(testFilePath); err != nil {
		t.Errorf("Expected no errors, but got: %s", err)
	}
}

func TestGetFileInfo(t *testing.T) {
	setupTest(t)
	testFilePath := filepath.Join("..", "..", "testdata", "sample.txt")
	if err := GetFileInfo(testFilePath, false); err != nil {
		t.Errorf("Expected no errors, but got: %s", err)
	}
}
