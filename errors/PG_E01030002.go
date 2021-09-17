// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01030002 struct {
}

func (e *PG_E01030002) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01030002) Message() string {
	return "指定されたIDとパスワードのショップが存在しません。"
}

func (e *PG_E01030002) Code() string {
	return "E01030002"
}

func (e *PG_E01030002) CanRetry() bool {
	return false
}
