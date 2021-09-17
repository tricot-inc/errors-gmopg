// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01310001 struct {
}

func (e *PG_E01310001) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01310001) Message() string {
	return "使用端末に”0”,”1”以外の値が指定されています。"
}

func (e *PG_E01310001) CanRetry() bool {
	return false
}
