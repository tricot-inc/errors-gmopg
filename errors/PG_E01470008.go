// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01470008 struct {
}

func (e *PG_E01470008) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01470008) Message() string {
	return "カード登録連番モードの書式が正しくありません。"
}

func (e *PG_E01470008) Code() string {
	return "E01470008"
}

func (e *PG_E01470008) CanRetry() bool {
	return false
}
