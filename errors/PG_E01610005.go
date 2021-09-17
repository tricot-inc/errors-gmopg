// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01610005 struct {
}

func (e *PG_E01610005) Error() string {
	return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01610005) Message() string {
	return "リトライ回数が0~99の範囲外です。"
}

func (e *PG_E01610005) CanRetry() bool {
	return false
}
