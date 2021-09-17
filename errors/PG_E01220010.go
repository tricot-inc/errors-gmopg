// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01220010 struct {
}

func (e *PG_E01220010) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01220010) Message() string {
	return "会員IDが重複しています。"
}

func (e *PG_E01220010) CanRetry() bool {
	return false
}
