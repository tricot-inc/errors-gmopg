// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01010010 struct {
}

func (e *PG_E01010010) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01010010) Message() string {
	return "ショップIDが一致しません。"
}

func (e *PG_E01010010) Code() string {
	return "E01010010"
}

func (e *PG_E01010010) CanRetry() bool {
	return false
}
