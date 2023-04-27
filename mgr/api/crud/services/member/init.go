package member

import (
	"github.com/micro-plat/hydra"
)

func init() {
	hydra.S.Micro("/member/login", loginHandle)
	hydra.S.Micro("/menu/list/get", getMenuHandle)
	hydra.S.Micro("/user/info/get", getUserHandle)
	hydra.S.Micro("/system/info/get", getSystemHandle)
}

// exp = append(exp, "/*/file/upload", "/*/member/login", "/*/system/info/get")
