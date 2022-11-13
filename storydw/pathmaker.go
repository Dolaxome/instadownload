package storydw

import (
	"net/url"
	"os"
	"path"
	"strconv"
	"time"
)

func StripQueryString(inputUrl string) string {
	u, err := url.Parse(inputUrl)
	if err != nil {
		panic(err)
	}
	u.RawQuery = ""
	return u.String()
}

func FileExt(url string) string {
	filename := path.Base(StripQueryString(url))
	return path.Ext(filename)
}

func FormatTimestamp(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format(time.RFC3339)
}

func CreateIfNotExist(dir string) (err error) {
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
	}
	return
}

func BuildPath(username, url string, timestamp int64) string {
	dirname := path.Join("stories", username)
	CreateIfNotExist(dirname)
	ext := FileExt(url)
	ts := FormatTimestamp(timestamp)
	filename := username + "-" + ts + strconv.FormatInt(timestamp, 10) + ext
	p := path.Join(dirname, filename)
	return p
}
