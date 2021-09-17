// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01190008 struct {
}

func (e *PG_E01190008) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01190008) Message() string {
	return "サイトIDの書式が正しくありません。"
}

func (e *PG_E01190008) CanRetry() bool {
	return false
}
