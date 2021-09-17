// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01160007 struct {
}

func (e *PG_E01160007) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01160007) Message() string {
	return "ボーナス分割回数に数字以外の文字が含まれています。"
}

func (e *PG_E01160007) CanRetry() bool {
	return false
}
