package yask_test

import (
	"os"

	"github.com/fcg-xvii/go-tools/text/config"
)

var (
	yaFolderID string
	yaAPIKey   string
)

func init() {
	if f, err := os.Open("test_data/ya.config"); err == nil {
		config.SplitToVals(f, "::", &yaFolderID, &yaAPIKey)
		f.Close()
	}
}
