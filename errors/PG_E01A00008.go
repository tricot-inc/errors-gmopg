// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01A00008 struct {
}

func (e *PG_E01A00008) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01A00008) Message() string {
	return "カード登録日の書式が正しくありません。"
}

func (e *PG_E01A00008) Code() string {
	return "E01A00008"
}

func (e *PG_E01A00008) CanRetry() bool {
	return false
}
