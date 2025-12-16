package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/goccy/go-yaml"
)

//go list -deps -f '{{.ImportPath}} ===  {{.GoFiles}}' |grep "$(go list -m)"

var configPath = flag.String("config_path", ".cidg.yml", "config path")

func main() {
	flag.Parse()

	var cfg Config

	content, err := os.ReadFile(*configPath)
	if err != nil {
		panic(err)
	}
	yaml.Unmarshal(content, &cfg)

	if err := run(&cfg); err != nil {
		fmt.Println(fmt.Errorf("cidg failed: %w", err))
		os.Exit(1)
		return
	}
}

func run(cfg *Config) error {
	// fmt.Println("moduleList:", moduleList)
	execDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// execDir := filepath.Dir(executable)
	affectModules := make(map[string]struct{})
	diffs, err := getHeadDiffFiles(cfg)
	if err != nil {
		return err
	}

	for _, mod := range cfg.ModuleList {
		mod = filepath.Clean(mod)
		for _, diff := range diffs {
			if strings.HasPrefix(diff.Path, mod) {
				affectModules[mod] = struct{}{}
				break
			} else if slices.Contains(cfg.GlobalFiles, diff.Path) {
				affectModules[mod] = struct{}{}
				break
			}
		}
		if _, ok := affectModules[mod]; ok {
			continue
		}
		deps, err := listModuleDeps(execDir, mod)
		if err != nil {
			fmt.Println("listModuleDeps err:", err)
			return err
		}
		for _, diff := range diffs {
			depPkg, ok := deps[diff.Package]
			if ok && slices.Contains(depPkg.Files, diff.FileName) {
				affectModules[mod] = struct{}{}
				break
			}
		}
	}

	resultList := make([]string, 0, len(affectModules))

	for _, mod := range cfg.ModuleList {
		if _, ok := affectModules[mod]; ok {
			resultList = append(resultList, mod)
		}
	}

	jsonStr, _ := json.Marshal(resultList)
	fmt.Println(string(jsonStr))
	return nil
}
