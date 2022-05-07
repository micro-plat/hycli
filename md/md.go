package md

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

//Mds2Tables 读取markdown文件并转换为MarkDownDB对象
func mds2Tables(filePath string) (Tables, error) {
	if !strings.Contains(filePath, "*") {
		return md2Tbl(filePath)
	}

	fns := getFiles(filePath)
	tbs := Tables{}
	for _, fn := range fns {
		ntbs, err := md2Tbl(fn)
		if err != nil {
			return nil, err
		}
		tbs = append(tbs, ntbs...)
	}
	if err := tbs.Duplicate(); err != nil {
		return nil, err
	}

	return tbs, nil
}

//md2Tbl 读取markdown文件并转换为MarkDownDB对象
func md2Tbl(fn string) (Tables, error) {
	lines, err := readMarkdown(fn)
	if err != nil {
		return nil, err
	}

	tbs, err := Lines2Table(line2Lines(lines))
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
