package main

import (
	"fmt"

	"golang.org/x/tools/go/packages"
)

func packageForFile(file string) (string, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedModule,
	}
	pkgs, err := packages.Load(cfg, "file="+file)
	if err != nil {
		return "", err
	}
	if len(pkgs) == 0 {
		return "", fmt.Errorf("no package found for %s", file)
	}
	// fmt.Println("file pkg", file, pkgs)
	return pkgs[0].PkgPath, nil
}
