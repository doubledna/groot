package notify

// Message : custom message struct
// Notifications are only sent when a task is created, completed, or failed
type Message struct {
}

// LarkMessage lark message struct
type LarkMessage struct {
	MsgType string  `json:"msg_type"`
	Content Content `json:"content"`
}
type InnerContent struct {
	Tag  string `json:"tag"`
	Text string `json:"text"`
	Href string `json:"href,omitempty"`
}
type ZhCN struct {
	Title   string           `json:"title"`
	Content [][]InnerContent `json:"content"`
}
type Post struct {
	ZhCN ZhCN `json:"zh-CN"`
}
type Content struct {
	Post Post `json:"post"`
}
