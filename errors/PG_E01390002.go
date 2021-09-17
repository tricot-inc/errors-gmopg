// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01390002 struct {
}

func (e *PG_E01390002) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01390002) Message() string {
	return "指定されたサイトIDと会員IDの会員が存在しません。"
}

func (e *PG_E01390002) CanRetry() bool {
	return false
}
