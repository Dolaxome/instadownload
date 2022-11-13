package storydw

import (
	"dolaxome/igdw/storylink"
	"fmt"
	"os/exec"
	// "os/exec"
	// "time"
)

func Wget(url, filepath string) error {
	// cmd := exec.Command("wget", "-H")
	cmd := exec.Command("wget", url, "-O", filepath)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	return cmd.Run()
}

func PrintInfo(username, url, filepath string, timestamp int64) {
	fmt.Println("[!]Username: " + username)
	fmt.Println("[!]Story timestamp: " + FormatTimestamp(timestamp))
	fmt.Println("[!]Download " + url + " to " + filepath)
}

func DownloadInstaUser(user storylink.InstUser) {
	for _, story := range user.Stories {
		pathy := BuildPath(user.Username, story.Url, story.Timestamp)
		PrintInfo(user.Username, story.Url, pathy, story.Timestamp)
		err := Wget(story.Url, "D:/Projects/IT/Go/ig-web-scrapper/instadownload/"+pathy)
		if err != nil {
			fmt.Println(err)
		}
		// if _, err := os.Stat(pathy); os.IsNotExist(err) {
		// 	PrintInfo(user.Username, story.Url, pathy, story.Timestamp)
		// 	err = Wget(story.Url, pathy)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// }
	}
}

func DownloadAll() {
	users, err := storylink.GetStoriesAll()
	if err != nil {
		return
	}

	for _, user := range users {
		DownloadInstaUser(user)
	}
}
