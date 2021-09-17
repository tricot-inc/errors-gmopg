// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01390010 struct {
}

func (e *PG_E01390010) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01390010) Message() string {
	return "指定されたサイトIDと会員IDの会員がすでに存在しています。"
}

func (e *PG_E01390010) Code() string {
	return "E01390010"
}

func (e *PG_E01390010) CanRetry() bool {
	return false
}
