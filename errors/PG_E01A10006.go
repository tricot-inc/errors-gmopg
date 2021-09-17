// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01A10006 struct {
}

func (e *PG_E01A10006) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01A10006) Message() string {
	return "過去24時間のカード追加の試行回数に数字以外の文字が含まれています。"
}

func (e *PG_E01A10006) CanRetry() bool {
	return false
}
