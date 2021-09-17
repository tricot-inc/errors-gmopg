// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01C60002 struct {
}

func (e *PG_E01C60002) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01C60002) Message() string {
	return "配送先住所の国番号が存在しません。"
}

func (e *PG_E01C60002) CanRetry() bool {
	return false
}
