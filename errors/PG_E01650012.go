// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01650012 struct {
}

func (e *PG_E01650012) Error() string {
	return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01650012) Message() string {
	return "指定されたショップは、指定されたサイトに属していません。"
}

func (e *PG_E01650012) CanRetry() bool {
	return false
}
