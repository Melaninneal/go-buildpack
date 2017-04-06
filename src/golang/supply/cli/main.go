package main

import (
	"golang/supply"
	"os"
	"path/filepath"

	"github.com/cloudfoundry/libbuildpack"
)

func main() {
	buildDir := os.Args[1]
	cacheDir := os.Args[2]
	depsDir := os.Args[3]
	depsIdx := os.Args[4]

	logger := libbuildpack.NewLogger()

	compiler, err := libbuildpack.NewCompiler([]string{buildDir, cacheDir, "", depsDir}, logger)
	err = compiler.CheckBuildpackValid()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(10)
	}

	// err = libbuildpack.RunBeforeCompile(compiler)
	// if err != nil {
	// 	compiler.Log.Error("Before Compile: %s", err.Error())
	// 	os.Exit(12)
	// }

	err = libbuildpack.SetEnvironmentFromSupply(depsDir)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(11)
	}

	gs := supply.Supplier{
		Compiler: compiler,
		DepDir:   filepath.Join(depsDir, depsIdx),
	}

	err = supply.Run(&gs)
	if err != nil {
		os.Exit(12)
	}
}