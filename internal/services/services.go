package services

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/unim26/seed/internal/colors"
	"github.com/unim26/seed/internal/models"
)

const SeedSignature = "SEED_SNAPSHOT_V1"

// travrse the current directory and return list of folders and files path or error
func travrseFolder() (dirs []string, files []string, e error) {
	currentDir := "."

	tdirs := []string{}
	tfiles := []string{}
	err := filepath.WalkDir(currentDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		if d.Name() != "." && strings.HasPrefix(d.Name(), ".") {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		cleanPath := filepath.ToSlash(path)

		if d.IsDir() {
			if d.Name() != "." {
				tdirs = append(tdirs, cleanPath)
			}
		} else {
			tfiles = append(tfiles, cleanPath)
		}

		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return tdirs, tfiles, nil
}

// create config [ .seed ] file
func CreateSnapshot() (string, error) {
	dirs, files, err := travrseFolder()
	if err != nil {
		return "", err
	}

	snapshot := models.Seed{
		SeedSignature: SeedSignature,
		Directories:   dirs,
		Files:         files,
	}

	jsonData, err := json.Marshal(snapshot)
	if err != nil {
		return "", err
	}

	encodedData := base64.StdEncoding.EncodeToString(jsonData)

	currDir, _ := filepath.Abs(".")
	project := filepath.Base(currDir)
	outputFile := project + "_snapshot.seed"
	err = os.WriteFile(outputFile, []byte(encodedData), 0644)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%sSnapshot saved securely to %s%s", colors.Green, outputFile, colors.Reset), nil
}

// build from seed
func BuildFromSnapshot() (string, error) {
	configFile, err := getConfigFile()
	if err != nil {
		return "", err
	}

	fileContent, err := os.ReadFile(configFile)
	if err != nil {
		return "", err
	}

	bytesData, err := base64.StdEncoding.DecodeString(string(fileContent))
	if err != nil {
		return "", nil
	}

	var seedContent models.Seed
	err = json.Unmarshal(bytesData, &seedContent)
	if err != nil {
		return "", nil
	}

	if seedContent.SeedSignature != SeedSignature {
		return "", errors.New("Invalid or corrupted configuration file.")
	}

	err = createStructure(seedContent)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("\n%s%s build complete.%s", colors.Green, configFile, colors.Reset), nil

}

// create the structure
func createStructure(seedContent models.Seed) error {
	//create all dir
	for _, dir := range seedContent.Directories {
		localPath := filepath.FromSlash(dir)
		err := os.MkdirAll(localPath, 0755)

		if err != nil {
			fmt.Printf("%sFailed folder:  %s (%v)%s\n", colors.Red, localPath, err, colors.Reset)
			continue
		}

	}

	//create all files
	for _, file := range seedContent.Files {
		localPath := filepath.FromSlash(file)
		if _, err := os.Stat(localPath); os.IsNotExist(err) {
			f, err := os.Create(localPath)
			if err != nil {
				fmt.Printf("%sFailed file:  %s (%v)%s\n", colors.Red, localPath, err, colors.Reset)
				continue
			}
			f.Close()
			fmt.Printf("%sCreated file: %s%s\n", colors.Yellow, localPath, colors.Reset)
		} else {
			fmt.Printf("%sSkipped: %s (already exists)%s\n", colors.Gray, localPath, colors.Reset)
		}
	}

	return nil
}

// get config file in current directory
func getConfigFile() (string, error) {
	var requiredpath string
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && strings.HasSuffix(d.Name(), ".seed") {
			requiredpath = path
			return nil
		}

		return nil
	})

	if err != nil {
		return "", nil
	}

	if requiredpath == "" {
		return "", errors.New("No configuration file provided in current directory")
	}

	return requiredpath, nil
}
