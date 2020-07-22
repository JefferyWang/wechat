package externalcontact

import "github.com/silenceper/wechat/v2/work/context"

// ExternalContact struct
type ExternalContact struct {
	*context.Context
}

// NewExternalContact 实例
func NewExternalContact(context *context.Context) *ExternalContact {
	extContact := new(ExternalContact)
	extContact.Context = context
	return extContact
}
