// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01A90008 struct {
}

func (e *PG_E01A90008) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01A90008) Message() string {
	return "ログイン日時の書式が正しくありません。"
}

func (e *PG_E01A90008) Code() string {
	return "E01A90008"
}

func (e *PG_E01A90008) CanRetry() bool {
	return false
}
