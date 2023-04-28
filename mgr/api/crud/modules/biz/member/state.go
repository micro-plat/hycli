package member

import "github.com/micro-plat/hydra"

type MemberState struct {
	UserID     int64  `json:"user_id" m2s:"user_id"`
	UserName   string `json:"user_name" m2s:"user_name"`
	FullName   string `json:"name" m2s:"name"`
	RoleName   string `json:"role_name" m2s:"role_name"`
	SystemID   int    `json:"sys_id"  m2s:"sys_id"`
	SystemName string `json:"sys_name"  m2s:"sys_name"`
	Theme      string `json:"theme"  m2s:"theme"`
	SysIdent   string `json:"ident"  m2s:"ident"`
	RoleID     int    `json:"role_id"  m2s:"role_id"`
	IndexURL   string `json:"index_url"  m2s:"index_url"`
	LoginURL   string `json:"login_url"  m2s:"login_url"`
}

const auth_key = "user.auth.info"

func GetMemberState(ctx hydra.IContext) (*MemberState, error) {
	v, ok := ctx.Meta().Get(auth_key)
	if ok {
		return v.(*MemberState), nil
	}
	ms := &MemberState{}
	if err := ctx.User().Auth().Bind(ms); err != nil {
		return nil, err
	}
	ctx.Meta().SetValue(auth_key, ms)
	return ms, nil
}