// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01500005 struct{
}

func (e *PG_E01500005) Error() string{
    return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01500005) Message() string{
    return "ショップ情報文字列の文字数が間違っています。"
}