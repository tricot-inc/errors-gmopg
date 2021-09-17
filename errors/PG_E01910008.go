// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01910008 struct {
}

func (e *PG_E01910008) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01910008) Message() string {
	return "サイト設定のマスクレベル利用有無に”0”,”1”以外の値が指定されています。"
}

func (e *PG_E01910008) Code() string {
	return "E01910008"
}

func (e *PG_E01910008) CanRetry() bool {
	return false
}
