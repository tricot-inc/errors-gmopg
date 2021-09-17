// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01200007 struct {
}

func (e *PG_E01200007) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01200007) Message() string {
	return "サイトIDが正しくありません。"
}

func (e *PG_E01200007) CanRetry() bool {
	return false
}
