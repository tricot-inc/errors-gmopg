// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01B80005 struct {
}

func (e *PG_E01B80005) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01B80005) Message() string {
	return "カード会員のメールアドレスが最大桁数を超えています。"
}

func (e *PG_E01B80005) Code() string {
	return "E01B80005"
}

func (e *PG_E01B80005) CanRetry() bool {
	return false
}
