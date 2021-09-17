// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01060005 struct {
}

func (e *PG_E01060005) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01060005) Message() string {
	return "利用金額が最大桁数を超えています。"
}

func (e *PG_E01060005) Code() string {
	return "E01060005"
}

func (e *PG_E01060005) CanRetry() bool {
	return false
}
