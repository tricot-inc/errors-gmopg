// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01010008 struct {
}

func (e *PG_E01010008) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01010008) Message() string {
	return "ショップIDに半角英数字以外の文字が含まれているか、13文字を超えています。"
}

func (e *PG_E01010008) Code() string {
	return "E01010008"
}

func (e *PG_E01010008) CanRetry() bool {
	return false
}
