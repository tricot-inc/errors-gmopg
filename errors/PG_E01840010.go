// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01840010 struct {
}

func (e *PG_E01840010) Error() string {
	return "入力パラメータエラー 入力パラメータの組み合わせが不正です。"
}

func (e *PG_E01840010) Message() string {
	return "必要な入力パラメータが指定されていません。"
}

func (e *PG_E01840010) CanRetry() bool {
	return false
}
