// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01D50006 struct {
}

func (e *PG_E01D50006) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01D50006) Message() string {
	return "購入されたプリペイドカードまたはギフトカードの総数に数字以外の文字が含まれています。"
}

func (e *PG_E01D50006) Code() string {
	return "E01D50006"
}

func (e *PG_E01D50006) CanRetry() bool {
	return false
}
