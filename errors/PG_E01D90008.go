// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01D90008 struct {
}

func (e *PG_E01D90008) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01D90008) Message() string {
	return "商品の注文情報の書式が正しくありません。"
}

func (e *PG_E01D90008) Code() string {
	return "E01D90008"
}

func (e *PG_E01D90008) CanRetry() bool {
	return false
}
