// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01800008 struct {
}

func (e *PG_E01800008) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01800008) Message() string {
	return "暗証番号の書式が正しくありません。"
}

func (e *PG_E01800008) Code() string {
	return "E01800008"
}

func (e *PG_E01800008) CanRetry() bool {
	return false
}
