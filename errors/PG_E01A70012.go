// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01A70012 struct {
}

func (e *PG_E01A70012) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01A70012) Message() string {
	return "ログイン証跡が最大バイト数を超えています。"
}

func (e *PG_E01A70012) CanRetry() bool {
	return false
}
