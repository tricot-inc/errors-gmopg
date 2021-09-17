// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01600001 struct{
}

func (e *PG_E01600001) Error() string{
    return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01600001) Message() string{
    return "会員情報チェック文字列が設定されていません。"
}