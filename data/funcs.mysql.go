package data

import (
	"fmt"
	"strings"
)

var mysqlTypeMap = map[string]string{
	"date":       "datetime",
	"float":      "float",
	"number(1)":  "tinyint",
	"number(2)":  "tinyint",
	"number(3)":  "int",
	"number(4)":  "int",
	"number(5)":  "int",
	"number(6)":  "int",
	"number(7)":  "int",
	"number(8)":  "int",
	"number(9)":  "int",
	"number(10)": "int",
	"number(11)": "bigint",
	"number(12)": "bigint",
	"number(13)": "bigint",
	"number(14)": "bigint",
	"number(15)": "bigint",
	"number(16)": "bigint",
	"number(17)": "bigint",
	"number(18)": "bigint",
	"number(19)": "bigint",
	"number(20)": "bigint",
	"number(21)": "bigint",
	"number#":    "decimal",
	"number@":    "bigint",
	"varchar2@":  "varchar",
}
var defMap = map[string]string{
	"datetime": "CURRENT_TIMESTAMP",
	"date":     "CURRENT_TIMESTAMP",
	"sysdate":  "CURRENT_TIMESTAMP",
	"now()":    "CURRENT_TIMESTAMP",
	"now":      "CURRENT_TIMESTAMP",
	"*":        "'*'",
}

func f_mysql_get_def_value(t string) string {
	if v, ok := defMap[strings.Trim(t, "'")]; ok {
		return v
	}

	return t
}
func mySQLType(t string, min int, max int) string {
	nName := t
	nLen := ""
	hasReange := "@"
	if min != 0 && max != 0 {
		hasReange = "#"
		nLen = fmt.Sprintf("(%d,%d)", min, max)
	} else if min != 0 {
		nLen = fmt.Sprintf("(%d)", min)
	}
	fullName := fmt.Sprintf("%s%s", nName, nLen)
	if x, ok := mysqlTypeMap[fullName]; ok {
		return x
	}
	if x, ok := mysqlTypeMap[fmt.Sprintf("%s%s", nName, hasReange)]; ok {
		return fmt.Sprintf("%s%s", x, nLen)
	}
	return fullName
}
