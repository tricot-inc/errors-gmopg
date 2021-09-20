// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01350008 struct {
}

func (e *PG_E01350008) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01350008) Message() string {
	return "MDの書式が正しくありません。"
}

func (e *PG_E01350008) Code() string {
	return "E01350008"
}

func (e *PG_E01350008) CanRetry() bool {
	return false
}
