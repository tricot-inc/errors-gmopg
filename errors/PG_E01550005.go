// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01550005 struct {
}

func (e *PG_E01550005) Error() string {
	return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01550005) Message() string {
	return "日時情報文字列の文字数が間違っています。"
}

func (e *PG_E01550005) Code() string {
	return "E01550005"
}

func (e *PG_E01550005) CanRetry() bool {
	return false
}
