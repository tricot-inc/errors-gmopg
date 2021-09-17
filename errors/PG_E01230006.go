// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01230006 struct {
}

func (e *PG_E01230006) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01230006) Message() string {
	return "カード登録連番に数字以外の文字が含まれています。"
}

func (e *PG_E01230006) Code() string {
	return "E01230006"
}

func (e *PG_E01230006) CanRetry() bool {
	return false
}
