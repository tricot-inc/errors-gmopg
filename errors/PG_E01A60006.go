// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01A60006 struct {
}

func (e *PG_E01A60006) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01A60006) Message() string {
	return "前年の取引回数に数字以外の文字が含まれています。"
}

func (e *PG_E01A60006) CanRetry() bool {
	return false
}
