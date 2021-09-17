// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01D60005 struct {
}

func (e *PG_E01D60005) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01D60005) Message() string {
	return "購入されたプリペイドカードまたはギフトカードの通貨コードが3桁ではありません。"
}

func (e *PG_E01D60005) CanRetry() bool {
	return false
}
