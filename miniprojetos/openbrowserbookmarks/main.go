package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No bookmark folder provided")
		return
	}
	bookmarkFolderName := os.Args[1]
	bookmarkFolder := getBookmarkFolder(bookmarkFolderName)
	links := getBookmarkFolderLinks(bookmarkFolder)

	args := []string{"--new-window"}
	cmd := exec.Command("google-chrome", append(args, links...)...)
	cmd.Start()
}

func openBrowser(links []string) {
	args := []string{"--new-window"}
	cmd := exec.Command("google-chrome", append(args, links...)...)
	cmd.Start()
}

func getBookmarkFolder(folder string) ChromeBookmarkFolder {
	bookmarksFile := readChromeBookmarksFile()
	var bookmark ChromeBookmarkFolder

	for _, collection := range bookmarksFile.Roots.BookmarkBar.Children {
		if collection.Type == "folder" && collection.Name == folder {
			bookmark = collection
		}
	}

	return bookmark
}

func getBookmarkFolderLinks(folder ChromeBookmarkFolder) []string {
	var links []string

	for _, bookmark := range folder.Children {
		if bookmark.Type == "url" {
			links = append(links, bookmark.Url)
		}
	}

	return links
}
