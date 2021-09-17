// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01510010 struct {
}

func (e *PG_E01510010) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01510010) Message() string {
	return "利用日の書式が正しくありません。"
}

func (e *PG_E01510010) CanRetry() bool {
	return false
}
