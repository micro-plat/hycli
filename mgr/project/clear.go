package project

import (
	"path/filepath"

	"github.com/codeskyblue/go-sh"
	logs "github.com/lib4dev/cli/logger"
)

//Clear 清理缓存
func Clear(dir string) error {
	if err := run(dir, "npm", "install", "--no-optional", "--verbose"); err != nil {
		return err
	}
	return run(dir, "npm", "install")
}
func run(dir string, name string, args ...interface{}) error {
	session := sh.InteractiveSession()
	session.SetDir(filepath.Join("./", dir))
	logs.Log.Info(append([]interface{}{name}, args...)...)
	session.Command(name, args...)
	return session.Run()
}
