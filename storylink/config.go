package storylink

var config = map[string]string{
	"ds_user_id": "{your_ds_user_id}",
	"sessionid":  "{your_sessionid}",
	"csrftoken":  "{your_cfrtoken}",
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
