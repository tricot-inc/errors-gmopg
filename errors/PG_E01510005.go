// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01510005 struct{
}

func (e *PG_E01510005) Error() string {
    return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01510005) Message() string {
    return "購買情報文字列の文字数が間違っています。"
}

func (e *PG_E01510005) CanRetry() bool {
    return false
}