// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01A30008 struct {
}

func (e *PG_E01A30008) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01A30008) Message() string {
	return "カード会員名と配送先名の一致/不一致の書式が正しくありません。"
}

func (e *PG_E01A30008) Code() string {
	return "E01A30008"
}

func (e *PG_E01A30008) CanRetry() bool {
	return false
}
