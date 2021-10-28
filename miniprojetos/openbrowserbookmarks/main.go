package main

import (
	"os/exec"
)

func main() {
	bookmarkFolder := getBookmarkFolder("mangas")
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
