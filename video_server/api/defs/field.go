package defs

type VideoInfo struct {
	UserId       int32  `json:"user_id"`
	Name         string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type Comment struct {
	Id          string `json:"id"`
	VideoInfoId string `json:"video_info_id"`
	UserId      int32  `json:"user_id"`
	Content     string `json:"content"`
	Time        string `json:"time"`
}

type ListComments []Comment

type SimpleSession struct {
	Username string //login name
	TTL      int64
}

type ReponseMsg struct {
	Code int32                  `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

type User struct {
	Id        int32  `json:"id"`
	LoginName string `json:"login_name"`
	Pwd       string `json:"pwd"`
	Nickname  string `json:"nickname"`
	Photo     string `json:"photo"`
}
