package message

// Text 文本消息
type Text struct {
	Content string `json:"content,omitempty"`
}

// Image 图片消息
type Image struct {
	MediaID string `json:"media_id,omitempty"`
	PicURL  string `json:"pic_url,omitempty"`
}

// Link 链接消息
type Link struct {
	Title  string `json:"title"`
	PicURL string `json:"picurl,omitempty"`
	Desc   string `json:"desc,omitempty"`
	URL    string `json:"url"`
}

// Miniprogram 小程序消息
type Miniprogram struct {
	Title      string `json:"title"`
	PicMediaID string `json:"pic_media_id"`
	AppID      string `json:"appid"`
	Page       string `json:"page"`
}
