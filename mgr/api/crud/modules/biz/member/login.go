package member

import (
	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"{-{.PkgName}-}/modules/const/enum"
	"{-{.PkgName}-}/modules/const/field"
)

// 用户登录
func Login(userName string, pwd string, ident string) (*MemberState, error) {
	db := hydra.C.DB().GetRegularDB()
	data, err := db.Query(sqlUserLogin, map[string]interface{}{
		"user_name": userName,
		"ident":     ident,
	})
	if err != nil {
		return nil, errs.NewErrorf(enum.ErrSystemBusy, "暂时无法登录%w", err)
	}
	if data.IsEmpty() {
		return nil, errs.NewError(enum.ErrUserPwdNotExists, "用户名不存在或密码错误")
	}

	//检查用户状态
	userInfo := data.Get(0)
	if userInfo.GetInt(field.UserStatus, -1) != enum.Normal {
		return nil, errs.NewError(enum.ErrUserDisabled, "用户已禁用暂时无法登录")
	}
	if userInfo.GetInt(field.IsLock, -1) == enum.BoolTrue {
		return nil, errs.NewError(enum.ErrUserLocked, "用户已锁定暂时无法登录")
	}

	if userInfo.GetString(field.Password) != pwd {
		//更新用户登录失败次数
		r, err := db.Execute(sqlUpdateLoginByFailed, map[string]interface{}{"user_name": userName})
		if err != nil || r == 0 {
			return nil, errs.NewErrorf(enum.ErrSystemBusy, "系统繁忙登录失败%w", errs.GetDBError(r, err))
		}
		return nil, errs.NewError(enum.ErrUserPwdNotExists, "用户名密码错误")
	}

	//检查角色状态
	if userInfo.GetInt(field.RoleStatus, -1) != enum.Normal ||
		userInfo.GetInt(field.RoleInfoStatus, -1) != enum.Normal {
		return nil, errs.NewError(enum.ErrRoleLocked, "没有权限或角色已锁定暂时无法登录")
	}

	//检查系统状态
	if userInfo.GetInt(field.SystemStatus, -1) != enum.Normal {
		return nil, errs.NewError(enum.ErrSystemLocked, "系统已锁定暂时无法登录")
	}
	m := &MemberState{}
	err = userInfo.ToAnyStruct(m)
	if err != nil {
		return nil, errs.NewErrorf(enum.ErrSystemBusy, "系统繁忙无法登录%w", err)
	}

	//更新用户登录登录时间
	r, err := db.Execute(sqlUpdateLoginBySuccess, userInfo.ToMap())
	if err != nil || r == 0 {
		return nil, errs.NewErrorf(enum.ErrSystemBusy, "系统繁忙无法登录%w", errs.GetDBError(r, err))
	}

	return m, nil
}
