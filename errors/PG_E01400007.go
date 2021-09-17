// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01400007 struct {
}

func (e *PG_E01400007) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01400007) Message() string {
	return "加盟店自由項目返却フラグに”0”,”1”以外の値が指定されています。"
}

func (e *PG_E01400007) CanRetry() bool {
	return false
}
