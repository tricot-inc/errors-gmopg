// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01D70008 struct {
}

func (e *PG_E01D70008) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01D70008) Message() string {
	return "商品の発売予定日の書式が正しくありません。"
}

func (e *PG_E01D70008) CanRetry() bool {
	return false
}
