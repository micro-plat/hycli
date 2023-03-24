package md

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Line 每一行信息
type Line struct {
	Text   string
	LineID int
}

// Lines 表的每一行
type Lines struct {
	Lines [][]*Line
}

// readMarkdown 读取md文件
func readMarkdown(name string) ([]*Line, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readMarkdownByReader(bufio.NewReader(f))
}

// readMarkdown 读取md文件
func readMarkdownByReader(rd *bufio.Reader) ([]*Line, error) {
	lines := make([]*Line, 0, 64)
	num := 0
	for {
		num++
		line, err := rd.ReadString('\n')
		if line == "" && (err != nil || io.EOF == err) {
			break
		}
		line = strings.Trim(strings.Trim(line, "\n"), "\t")
		if strings.TrimSpace(line) == "" {
			continue
		}
		lines = append(lines, &Line{Text: line, LineID: num})
	}
	return lines, nil
}
func NewLine(idx int, name string, tp string, label string) *Line {
	return &Line{
		LineID: idx,
		Text:   fmt.Sprintf("|%s|%s||是||%s|", name, tp, label),
	}
}
