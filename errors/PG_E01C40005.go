// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01C40005 struct {
}

func (e *PG_E01C40005) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01C40005) Message() string {
	return "職場電話番号が最大桁数を超えています。"
}

func (e *PG_E01C40005) Code() string {
	return "E01C40005"
}

func (e *PG_E01C40005) CanRetry() bool {
	return false
}
