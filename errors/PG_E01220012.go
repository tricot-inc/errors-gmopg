// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01220012 struct {
}

func (e *PG_E01220012) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01220012) Message() string {
	return "会員IDの長さが正しくありません。"
}

func (e *PG_E01220012) Code() string {
	return "E01220012"
}

func (e *PG_E01220012) CanRetry() bool {
	return false
}
