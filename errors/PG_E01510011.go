// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01510011 struct {
}

func (e *PG_E01510011) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01510011) Message() string {
	return "利用日の値が指定可能範囲外です。"
}

func (e *PG_E01510011) CanRetry() bool {
	return false
}
