package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

type Dep struct {
	Package string
	Files   []string
}

func runCommands(cmds ...*exec.Cmd) ([]byte, error) {
	if len(cmds) == 0 {
		return nil, fmt.Errorf("no commands provided")
	}
	for i := 0; i < len(cmds)-1; i++ {
		stdout, err := cmds[i].StdoutPipe()
		if err != nil {
			return nil, err
		}
		cmds[i+1].Stdin = stdout
	}
	for i := 0; i < len(cmds)-1; i++ {
		if err := cmds[i].Start(); err != nil {
			return nil, err
		}
	}
	output, err := cmds[len(cmds)-1].CombinedOutput()
	for i := 0; i < len(cmds)-1; i++ {
		if werr := cmds[i].Wait(); werr != nil && err == nil {
			err = werr
		}
	}
	return output, err
}

func listModuleDeps(execDir, modulePath string) (map[string]Dep, error) {
	modulePath = strings.TrimSpace(modulePath)
	modulePath = filepath.Join(execDir, modulePath)
	modOut, err := exec.Command("go", "-C", modulePath, "list", "-m").Output()
	if err != nil {
		return nil, err
	}
	modName := strings.TrimSpace(string(modOut))
	parts := strings.Split(modName, "/")
	if len(parts) > 3 {
		modName = strings.Join(parts[:3], "/")
	}

	output, err := runCommands(
		exec.Command("go", "-C", modulePath, "list", "-deps", "-f", "{{.ImportPath}}:::{{.GoFiles}}"),
		exec.Command("grep", modName),
	)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	deps := make(map[string]Dep, 0)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":::", 2)
		if len(parts) != 2 {
			continue
		}
		pkg := strings.TrimSpace(parts[0])

		filesStr := strings.TrimSpace(parts[1])
		if strings.HasPrefix(filesStr, "[") && strings.HasSuffix(filesStr, "]") {
			filesStr = strings.TrimSuffix(strings.TrimPrefix(filesStr, "["), "]")
		}
		var files []string
		if filesStr != "" {
			files = strings.Fields(filesStr)
		}
		deps[pkg] = Dep{Package: pkg, Files: files}
	}
	return deps, nil
}
