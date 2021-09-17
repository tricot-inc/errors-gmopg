// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01E30001 struct {
}

func (e *PG_E01E30001) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01E30001) Message() string {
	return "戻りURLが指定されていません。"
}

func (e *PG_E01E30001) Code() string {
	return "E01E30001"
}

func (e *PG_E01E30001) CanRetry() bool {
	return false
}
