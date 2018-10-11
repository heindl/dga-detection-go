package modpath

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Abs() (string, error) {
	absPath, err := filepath.Abs(".")
	if err != nil {
		return "", errors.Wrapf(err, "could not get absolute path %s", absPath)
	}

	if _, err := os.Stat(filepath.Join(absPath, "go.mod")); err == nil {
		return absPath, nil
	}
	// Mod file doesn't exist, so check to see if we are in a package directory.
	modulePath, err := getModulePath(absPath)
	if err != nil {
		return "", err
	}
	return modulePath, nil
}

func containsFile(directoryPath, fileName string) (bool, error) {

	files, err := ioutil.ReadDir(directoryPath)
	if err != nil {
		return false, errors.Wrapf(err, "could not read directory %s", directoryPath)
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		areEqual := f.Name() == fileName
		if areEqual { // || (!caseSensitive && strings.EqualFold(f.Name(), fileName)) {
			return true, nil
		}
	}
	return false, nil
}

// getModulePath recursively searches upward from given path for a go.mod file, and errors out if not found.
func getModulePath(path string) (string, error) {

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", errors.Wrap(err, path)
	}

	for {
		rel, err := filepath.Rel("/", abs)
		if err != nil {
			return "", errors.Wrap(err, abs)
		}
		if rel == "." {
			log.Printf(`The given command path [%s] appears to not be within a Go module, which is necessary.
				No 'go.mod' file was found in the directory, or in an parent directory.
				Please read https://github.com/golang/go/wiki/Modules for more information about setting one up.
			`, rel)
			return "", errors.New("given path is not within a module")
		}
		rel = filepath.Join("/", rel)

		hasFile, err := containsFile(rel, "go.mod")
		if err != nil {
			return "", err
		}
		if hasFile {
			return rel, nil
		}
		abs += "/.."
	}

}
