package callback

// BaseMsg 基础回调信息
type BaseMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Event        string `xml:"Event"`
	ChangeType   string `xml:"ChangeType,omitempty"`
}

// Attr 扩展属性
type Attr struct {
	Type int       `xml:"Type"`
	Name string    `xml:"Name"`
	Text *TextAttr `xml:"Text,omitempty"`
	Web  *WebAttr  `xml:"Web,omitempty"`
}

// TextAttr 文本类型属性
type TextAttr struct {
	Value string `xml:"Value"` // 文本属性的内容
}

// WebAttr 网页类型属性
type WebAttr struct {
	URL   string `xml:"Url"`   // 网页的url
	Title string `xml:"Title"` // 网页的展示标题
}

// UserCallbackMsg 成员变更回调消息
type UserCallbackMsg struct {
	BaseMsg
	UserID         string `xml:"UserID"`
	NewUserID      string `xml:"NewUserID,omitempty"`
	Name           string `xml:"Name,omitempty"`
	Department     string `xml:"Department,omitempty"`
	IsLeaderInDept string `xml:"IsLeaderInDept"`
	Mobile         string `xml:"Mobile,omitempty"`
	Position       string `xml:"Position,omitempty"`
	Gender         int    `xml:"Gender,omitempty"`
	Email          string `xml:"Email,omitempty"`
	Status         int    `xml:"Status,omitempty"`
	Avatar         string `xml:"Avatar,omitempty"`
	Alias          string `xml:"Alias,omitempty"`
	Telephone      string `xml:"Telephone,omitempty"`
	Address        string `xml:"Address,omitempty"`
	ExtAttr        []Attr `xml:"ExtAttr>Item,omitempty"`
}

// DepartmentCallbackMsg 部门回调信息
type DepartmentCallbackMsg struct {
	BaseMsg
	ID       int    `xml:"Id"`
	Name     string `xml:"Name,omitempty"`
	ParentID int    `xml:"ParentId,omitempty"`
	Order    int    `xml:"Order,omitempty"`
}

// TagCallbackMsg 标签变更通知回调信息
type TagCallbackMsg struct {
	BaseMsg
	TagID         int    `xml:"TagId"`
	AddUserItems  string `xml:"AddUserItems"`
	DelUserItems  string `xml:"DelUserItems"`
	AddPartyItems string `xml:"AddPartyItems"`
	DelPartyItems string `xml:"DelPartyItems"`
}

// BatchJob 异步任务
type BatchJob struct {
	JobID   string `xml:"JobId"`
	JobType string `xml:"JobType"`
	ErrCode int    `xml:"ErrCode"`
	ErrMsg  string `xml:"ErrMsg"`
}

// BatchJobCallbackMsg 异步任务完成通知
type BatchJobCallbackMsg struct {
	BaseMsg
	BatchJob BatchJob `xml:"BatchJob"`
}
