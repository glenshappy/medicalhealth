package defs

type VideoInfo struct {
	UserId int32 `json:"user_id"`
	Name string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type Comment struct {
	Id string `json:"id"`
	VideoInfoId string `json:"video_info_id"`
	UserId int32 `json:"user_id"`
	Content string `json:"content"`
	Time string `json:"time"`
}

type ListComments []Comment

