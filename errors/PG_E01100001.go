// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01100001 struct {
}

func (e *PG_E01100001) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01100001) Message() string {
	return "取引パスワードが指定されていません。"
}

func (e *PG_E01100001) Code() string {
	return "E01100001"
}

func (e *PG_E01100001) CanRetry() bool {
	return false
}
