package md

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	logs "github.com/lib4dev/cli/logger"
)

func Mds2Tables(filePath string) (Tables, error) {
	tbs, err := mds2Tables(filePath)
	if err != nil {
		return nil, err
	}
	for _, tb := range tbs {
		tb.Tbs = tbs
	}
	return tbs, nil
}

// Mds2Tables 读取markdown文件并转换为MarkDownDB对象
func mds2Tables(filePath string) (Tables, error) {
	if !strings.Contains(filePath, "*") {
		return md2Tbl(filePath)
	}

	fns := getFiles(filePath)
	tbs := Tables{}
	cnfs := TableConfs{}

	//读取完整表结构
	tbpaths, confpaths := groupPath(fns)
	logs.Log.Info("字典文件:", tbpaths)
	logs.Log.Info("配置文件:", confpaths)
	for _, fn := range tbpaths {
		ntbs, err := md2Tbl(fn)
		if err != nil {
			return nil, err
		}
		tbs = append(tbs, ntbs...)
	}

	//去除重复表
	if err := tbs.Duplicate(); err != nil {
		return nil, err
	}

	//读取配置表结构

	for _, fn := range confpaths {
		confs, err := md2Conf(fn)
		if err != nil {
			return nil, err
		}
		cnfs = append(cnfs, confs...)
	}

	// 按表名合并配置信息
	tbs = tbs.resetConf(cnfs)

	return tbs, nil
}

// md2Tbl 读取markdown文件并转换为MarkDownDB对象
func md2Tbl(fn string) (Tables, error) {
	lines, err := readMarkdown(fn)
	if err != nil {
		return nil, err
	}

	tbs, err := Lines2Table(line2TbLines(lines))
	if err != nil {
		return nil, err
	}
	if err = tbs.Duplicate(); err != nil {
		return nil, err
	}

	return tbs, nil
}

// md2Tbl 读取markdown文件并转换为MarkDownDB对象
func md2Conf(fn string) (TableConfs, error) {
	lines, err := readMarkdown(fn)
	if err != nil {
		return nil, err
	}

	tbs, err := Lines2Conf(line2ConfLines(lines))
	if err != nil {
		return nil, err
	}
	if err = tbs.Duplicate(); err != nil {
		return nil, err
	}

	return tbs, nil
}

// Md2TbByContent 将markown文本转成table
func Md2TbByContent(content string) (Tables, error) {
	reader, err := readMarkdownByReader(bufio.NewReader(strings.NewReader(content)))
	if err != nil {
		return nil, err
	}
	tbs, err := Lines2Table(line2TbLines(reader))
	if err != nil {
		return nil, err
	}
	if err = tbs.Duplicate(); err != nil {
		return nil, err
	}

	return tbs, nil
}

func getFiles(path string) (paths []string) {

	//路径是的具体文件
	_, err := os.Stat(path)
	if err == nil {
		return []string{path}
	}
	//查找匹配的文件
	dir, f := filepath.Split(path)

	regexName := fmt.Sprintf("^%s$", strings.Replace(strings.Replace(f, ".md", "\\.md", -1), "*", "(.+)?", -1))
	reg := regexp.MustCompile(regexName)

	if dir == "" {
		dir = "./"
	}
	files, _ := ioutil.ReadDir(dir)
	for _, f := range files {
		fname := f.Name()
		if strings.HasPrefix(fname, ".") || f.IsDir() {
			continue
		}
		if reg.Match([]byte(fname)) {
			paths = append(paths, filepath.Join(dir, fname))
		}
	}
	return paths
}
func groupPath(p []string) ([]string, []string) {
	file, confs := []string{}, []string{}
	for _, v := range p {
		if strings.Contains(v, ".hi.conf.") {
			confs = append(confs, v)
			continue
		}
		file = append(file, v)
	}
	return file, confs
}
