package storydw

import (
	"dolaxome/igdw/storylink"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Wget(url, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(n)
	return nil
}

func PrintInfo(username, url, filepath string, timestamp int64) {
	fmt.Println("[!]Username: " + username)
	fmt.Println("[!]Story timestamp: " + FormatTimestamp())
	fmt.Println("[!]Download " + url + " to " + filepath)
}

func DownloadAll() {
	users, err := storylink.GetStoriesAll()
	if err != nil {
		return
	}

	for _, user := range users {
		counter := 0
		DownloadInstaUser(user, &counter)
	}
}

func DownloadInstaUser(user storylink.InstUser, counter *int) {
	for _, story := range user.Stories {
		pathy := BuildPath(user.Username, story.Url, *counter)
		PrintInfo(user.Username, story.Url, pathy, story.Timestamp)
		err := Wget(story.Url, pathy)
		if err != nil {
			fmt.Println(err)
		}
		if _, err := os.Stat(pathy); os.IsNotExist(err) {
			PrintInfo(user.Username, story.Url, pathy, story.Timestamp)
			err = Wget(story.Url, pathy)
			if err != nil {
				fmt.Println(err)
			}
		}
		*counter++
	}
}
