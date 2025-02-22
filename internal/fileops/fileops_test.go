package fileops

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	//Prepare test file
	err := os.WriteFile("test.txt", []byte("Hi i am Aiden Tsai!"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("test.txt")

	//Testing function
	content, err := ReadFile("test.txt")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if content != ("Hi i am Aiden Tsai!") {
		t.Errorf("Expected 'test content', got %s", content)
	}
}
