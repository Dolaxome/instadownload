package storylink

var config = map[string]string{
	"ds_user_id": "45141217438",
	"sessionid":  "45141217438%3ArAUp1DA8WsWZeI%3A17%3AAYc2WgKKFsAWdnLp9S6JGBwvQqRzYV51WPm4mzqdWg",
	"csrftoken":  "x9R8Q1CKIGtXMRy96r0pGge8AzQgLnXG",
	"User-Agent": "Instagram 10.3.2 (iPhone7,2; iPhone OS 9_3_3; en_US; en-US; scale=2.00; 750x1334) AppleWebKit/420+",
}

func SetUserId(s string) {
	config["ds_user_id"] = s
}
func SetSessionId(s string) {
	config["sessionid"] = s
}
func SetCsrfToken(s string) {
	config["csrftoken"] = s
}
