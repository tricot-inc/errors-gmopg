// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01E90001 struct {
}

func (e *PG_E01E90001) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01E90001) Message() string {
	return "職場電話番号の指定が正しくありません。職場電話の国コード/職場電話番号のいずれかの省略はできません。"
}

func (e *PG_E01E90001) CanRetry() bool {
	return false
}
