package main

import (
	"fmt"
	"path/filepath"

	"golang.org/x/tools/go/packages"
)

func packageForFile(file string) (string, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedModule,
		Dir:  filepath.Dir(file),
	}
	pkgs, err := packages.Load(cfg, "file="+filepath.Base(file))
	if err != nil {
		return "", err
	}
	if len(pkgs) == 0 {
		return "", fmt.Errorf("no package found for %s", file)
	}

	return pkgs[0].PkgPath, nil
}
