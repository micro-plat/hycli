package md

// md2Tbl 读取markdown文件并转换为MarkDownDB对象
func md2conf(fn string) (Tables, error) {
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
