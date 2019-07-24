package main

import (
	"errors"
	"fmt"
	"os"

	cnab_riff "github.com/projectriff/cnab-riff/pkg"
)

func main() {
	bundlePath, bundleDestinationPath, manifestPath, manifestDestinationPath, err := verifyCommandLineArgs(os.Args)
	if err != nil {
		fmt.Printf("error validating arguments %v\n", err)
		os.Exit(1)
	}

	err = cnab_riff.FinalizeBundle(bundlePath, bundleDestinationPath, manifestPath, manifestDestinationPath)
	if err != nil {
		fmt.Printf("error updating bundle: %v\n", err)
		os.Exit(1)
	}
}

func verifyCommandLineArgs(args []string) (bundleTemplatePath, bundleDestinationPath, manifestPath, manifestDestinationPath string, err error) {
	if len(args) == 1 {
		return "base-duffle.json", "duffle.json", "kab-manifest.yaml", "./cnab/app/kab/manifest.yaml", nil
	}
	if len(args) != 5 {
		return "", "", "", "", errors.New("usage: ./list-images <path/to/duffle.json> </path/to/kab-manifest.yaml> </path/to/cnab-manifest-destination.yaml>")
	}
	return args[1], args[2], args[3], args[4], nil
}
