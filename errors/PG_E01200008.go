// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01200008 struct {
}

func (e *PG_E01200008) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01200008) Message() string {
	return "サイトパスワードの書式が正しくありません。"
}

func (e *PG_E01200008) Code() string {
	return "E01200008"
}

func (e *PG_E01200008) CanRetry() bool {
	return false
}
