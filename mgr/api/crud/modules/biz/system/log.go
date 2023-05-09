package system

import (
	"net/http"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/lib4go/errs"
	"{-{.PkgName}-}/modules/biz/member"
)

// SaveLog  保存日志信息
func SaveLog(ctx hydra.IContext, title string, content string) interface{} {
	member, err := member.GetMemberState(ctx)
	if err != nil {
		return err
	}
	return SaveLogByUID(member.UserID, title, content)
}

func SaveLogByUID(uid int64, title string, content string) (r interface{}) {
	rx, err := hydra.C.DB().GetRegularDB().Execute(insertUserLog, map[string]interface{}{
		"user_id": uid,
		"title":   title,
		"content": content,
	})
	if err == nil && rx > 0 {
		return nil
	}
	return errs.NewErrorf(http.StatusNotExtended, "数据出错:%+v,row:%d", err, rx)
}
