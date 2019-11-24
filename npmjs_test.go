package npmjs

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	file := "test.txt"
	Get(file)
	_, err := os.Stat(file)
	if err != nil {
		t.Errorf("File not create")
	}
}
