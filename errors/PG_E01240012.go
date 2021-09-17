// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01240012 struct {
}

func (e *PG_E01240012) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01240012) Message() string {
	return "指定された会員IDがファイル内で重複しています(※洗替時)"
}

func (e *PG_E01240012) Code() string {
	return "E01240012"
}

func (e *PG_E01240012) CanRetry() bool {
	return false
}
