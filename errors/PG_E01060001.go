// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01060001 struct {
}

func (e *PG_E01060001) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01060001) Message() string {
	return "利用金額が指定されていません。"
}

func (e *PG_E01060001) Code() string {
	return "E01060001"
}

func (e *PG_E01060001) CanRetry() bool {
	return false
}
