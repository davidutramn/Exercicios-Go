package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const GOOGLE_CHROME_BOOKMARKS_PATH = "/home/davi/.config/google-chrome/Default/Bookmarks"

type ChromeBookmarksFile struct {
	Roots struct {
		BookmarkBar ChromeBookmarkBar `json:"bookmark_bar"`
	} `json:"roots"`
}

type ChromeBookmarkBar struct {
	Children []ChromeBookmarkFolder `json:"children"`
}

type ChromeBookmarkFolder struct {
	Children []ChromeBookmark `json:"children"`
	Name     string           `json:"name"`
	Type     string           `json:"type"`
}

type ChromeBookmark struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

func readChromeBookmarksFile() ChromeBookmarksFile {
	file, err := os.Open(GOOGLE_CHROME_BOOKMARKS_PATH)
	if err != nil {
		if i := os.IsNotExist(err); i {
			fmt.Println("Google Chrome bookmark file not found in:", GOOGLE_CHROME_BOOKMARKS_PATH)
		} else {
			fmt.Println("Can't open file:", err)
		}
		os.Exit(0)
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var result ChromeBookmarksFile

	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		fmt.Println("Unmarshal error", err)
	}

	return result
}
