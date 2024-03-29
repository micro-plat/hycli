package output

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/micro-plat/lib4go/errs"
)

// pathExists 文件是否存在
func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Create 创建文件，文件夹 存在时写入则覆盖
func Create(path string, cover bool) (file *os.File, err error) {
	dir := filepath.Dir(path)
	if !pathExists(dir) {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return nil, fmt.Errorf("创建文件夹%s失败:%v", path, err)
		}
	}

	if !pathExists(path) {
		file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return nil, fmt.Errorf("无法打开文件: %s(err:%v)", path, err)
		}
		return
	}

	if !cover {
		return nil, fmt.Errorf("文件: %s已经存在", path)
	}

	// 文件存在且文件头部设置不覆盖标识
	if checkCover(path) {
		return nil, errs.NewError(204, fmt.Errorf("跳过文件: %s不允许覆盖", path))

	}
	file, err = os.OpenFile(path, os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("无法打开文件: %s(err:%v)", path, err)
	}
	return
}

// 返回只可读的文件 外部调用注意Close()
func checkCover(path string) bool {
	file, _ := os.Open(path)
	//读取第一行的
	defer file.Close()
	br := bufio.NewReader(file)
	a, _, _ := br.ReadLine()
	return bytes.HasPrefix(a, []byte("//exclude")) ||
		bytes.HasPrefix(a, []byte("<!--exclude-->"))

}
