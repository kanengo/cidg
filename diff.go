package main

import (
	"os/exec"
	"path/filepath"
	"strings"
)

//git diff --name-only HEAD~1

type diffFile struct {
	Path     string
	Package  string
	FileName string
}

func getHeadDiffFiles() ([]diffFile, error) {
	cmd := exec.Command("git", "diff", "--name-only", "HEAD~1")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	fileList := strings.Split(string(output), "\n")

	ret := make([]diffFile, 0, len(fileList))
	for _, file := range fileList {
		if file == "" {
			continue
		}
		pkg, _ := packageForFile(file)
		ret = append(ret, diffFile{
			Path:     file,
			Package:  pkg,
			FileName: filepath.Base(file),
		})
	}
	return ret, nil
}
