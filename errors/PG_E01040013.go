// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01040013 struct {
}

func (e *PG_E01040013) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01040013) Message() string {
	return "オーダーIDに半角英数字と”-”以外の文字が含まれています。"
}

func (e *PG_E01040013) Code() string {
	return "E01040013"
}

func (e *PG_E01040013) CanRetry() bool {
	return false
}
