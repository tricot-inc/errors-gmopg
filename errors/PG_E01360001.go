// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01360001 struct {
}

func (e *PG_E01360001) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01360001) Message() string {
	return "PaResが指定されていません。"
}

func (e *PG_E01360001) Code() string {
	return "E01360001"
}

func (e *PG_E01360001) CanRetry() bool {
	return false
}
