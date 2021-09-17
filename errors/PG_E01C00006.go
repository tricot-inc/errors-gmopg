// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01C00006 struct {
}

func (e *PG_E01C00006) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01C00006) Message() string {
	return "自宅電話番号に数字以外の文字が含まれています。"
}

func (e *PG_E01C00006) Code() string {
	return "E01C00006"
}

func (e *PG_E01C00006) CanRetry() bool {
	return false
}
