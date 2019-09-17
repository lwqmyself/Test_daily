package utils

import "github.com/astaxie/beego/logs"

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		logs.Error(err.Error())
	}
	count, err := result.RowsAffected()
	if err != nil {
		logs.Error(err.Error())
	}
	return count, err
}
