package contact

import "github.com/silenceper/wechat/v2/work/context"

// Contact struct
type Contact struct {
	*context.Context
}

// NewContact 实例
func NewContact(context *context.Context) *Contact {
	basic := new(Contact)
	basic.Context = context
	return basic
}
