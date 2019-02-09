package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gocv.io/x/gocv"
)

func defineKeyboardMapping() map[int]string {
	m := map[int]string{
		48:  "0",
		49:  "1",
		50:  "2",
		51:  "3",
		52:  "4",
		53:  "5",
		54:  "6",
		55:  "7",
		56:  "8",
		57:  "9",
		119: "w",
		99:  "c",
		32:  " ",
		117: "u",
	}
	return m
}

func defineKeyToCategory() map[string]string {
	m := map[string]string{
		"0": "whine",
		"1": "whine_one_chuck",
		"2": "whine_two_chucks",
		"3": "whine_three_chucks",
		"4": "whine_four_chucks",
		"5": "whine_five_chucks",
		"6": "whine_six_chucks",
		"7": "whine_seven_chucks",
		"8": "whine_eight_chucks",
		"9": "mew",
		"u": "unknown",
	}
	return m
}

func makeFullPath(dir string, fileName os.FileInfo) string {
	return filepath.Join(dir, fileName.Name())
}

func Map(files []os.FileInfo, f func(string, os.FileInfo) string) []string {
	fullPaths := make([]string, len(files))
	for i, file := range files {
		fullPaths[i] = f(getFolder(), file)
	}
	return fullPaths
}

func getFolder() string {
	homeDirectory := "."
	return homeDirectory
}

func getListOfFilesIn(dir string) []string {

	f, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	fullPaths := make([]string, 0, len(files))
	for _, file := range files {
		fullPaths = append(fullPaths, path.Join(dir, file.Name()))
	}

	return fullPaths

}

func showImageToUser(fileName string) string {
	window := gocv.NewWindow(fileName)
	defer window.Close()
	img := gocv.IMRead(fileName, gocv.IMReadColor)
	keyMap := defineKeyboardMapping()
	categoryMap := defineKeyToCategory()
	for {
		window.IMShow(img)
		key, ok := keyMap[window.WaitKey(1)]
		if ok {
			callCategory, ok := categoryMap[key]
			if ok {
				return callCategory
			} else {
				fmt.Println("That key is not mapped to a call type")
			}
		}
	}

}

func getListOfCategories(catMap map[string]string) []string {
	keys := make([]string, 0, len(catMap))
	for _, v := range catMap {
		keys = append(keys, v)
	}
	return keys
}

func createCategoryFoldersIfNonexistentIn(dir string) {
	possibleCategories := getListOfCategories(defineKeyToCategory())

	for _, cat := range possibleCategories {
		folderName := path.Join(dir, cat)
		if _, err := os.Stat(folderName); os.IsNotExist(err) {
			os.Mkdir(folderName, os.ModePerm)
		}
	}
}

func isANonCategoricalFolder(folder os.FileInfo) bool {
	isNotCat := true
	dirsWeCreated := getListOfCategories(defineKeyToCategory())
	for _, dir := range dirsWeCreated {
		if dir == folder.Name() {
			isNotCat = false
		}
	}
	return isNotCat
}

func findNonEmptyFolderIn(dir string) string {

	dirToReturn := "None"

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	// exclude the files
	dirs := make([]os.FileInfo, 0, len(files))
	for _, file := range files {
		if file.IsDir() && isANonCategoricalFolder(file) {
			dirs = append(dirs, file)
		}
	}

	// return a folder with at least one png
	for _, dir := range dirs {
		// exclude invisible folders
		if dir.Name()[:1] != "." {
			files := getListOfFilesIn(dir.Name())
			for _, f := range files {
				if strings.HasSuffix(f, ".png") {
					dirToReturn = dir.Name()
					break
				}
			}
		}
	}
	if dirToReturn == "None" {
		fmt.Println("No folders left with pngs to examine. Exiting.")
		os.Exit(3)
	}
	return dirToReturn
}

func moveFileToCategory(file string, category string, parentDir string) {
	fileName := filepath.Base(file)
	newLocation := filepath.Join(parentDir, category, fileName)

	fmt.Println("moving", file, "to", newLocation)

	err := os.Rename(file, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	topLevelDirectory := getFolder()

	// create the folders if they don't exist
	// this looks to the next directory up from the `directory` due to how things are organized
	createCategoryFoldersIfNonexistentIn(topLevelDirectory)

	// pick a directory from the available directories
	directory := findNonEmptyFolderIn(topLevelDirectory)
	fmt.Println("using", directory)

	files := getListOfFilesIn(directory)

	for _, file := range files {
		fmt.Println(file)
	}

	// fullPaths := Map(files, makeFullPath)

	for _, file := range files {

		if strings.HasSuffix(file, ".png") {
			category := showImageToUser(file)
			moveFileToCategory(file, category, topLevelDirectory)
			fmt.Println(category)

		}
	}
	// move each photo

	// write to csv
}
