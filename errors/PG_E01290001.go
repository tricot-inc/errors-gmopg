// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01290001 struct {
}

func (e *PG_E01290001) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01290001) Message() string {
	return "HTTP_ACCEPTが指定されていません。"
}

func (e *PG_E01290001) Code() string {
	return "E01290001"
}

func (e *PG_E01290001) CanRetry() bool {
	return false
}
