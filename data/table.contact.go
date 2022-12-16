package data

type BUITable struct {
	Current interface{}
	Main    interface{}
}

// contactTB 连接多个UITable
func contactTables(c interface{}, e interface{}) *BUITable {
	return &BUITable{
		Current: c,
		Main:    e,
	}
}
