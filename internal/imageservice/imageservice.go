package imageservice

import (
	"internshipApplicationTemplate/internal/config"
	"os"
	"strings"
)

var r []string

func New(cfg config.Config) error {
	dir := cfg.FilePath
	filenames, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, name := range filenames {
		id := strings.Split(name.Name(), ".")[0]
		r = append(r, id)
	}

	return nil
}
