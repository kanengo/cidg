package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/kanengo/cidg/example/service1"
)

//go list -deps -f '{{.ImportPath}} ===  {{.GoFiles}}' |grep "$(go list -m)"

var modulePath = flag.String("module_path", "./", "module path")

func main() {
	flag.Parse()
	service1.Service1()

	moduleList := strings.Fields(*modulePath)

	if err := run(moduleList); err != nil {
		fmt.Println(fmt.Errorf("cidg failed: %w", err))
		os.Exit(1)
		return
	}
}

func run(moduleList []string) error {
	// fmt.Println("moduleList:", moduleList)
	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}
	execDir := filepath.Dir(executable)
	affectModules := make(map[string]struct{})
	diffs, err := getHeadDiffFiles()
	if err != nil {
		return err
	}

	for _, mod := range moduleList {
		mod = filepath.Clean(mod)
		for _, diff := range diffs {
			if strings.HasPrefix(diff.Path, mod) {
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

	for mod := range affectModules {
		fmt.Println("affect module:", mod)
	}

	return nil
}
