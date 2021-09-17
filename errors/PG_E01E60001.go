// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01E60001 struct{
}

func (e *PG_E01E60001) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01E60001) Message() string{
    return "カード会員のログイン情報の指定が正しくありません。ログイン証跡/ログイン方法/ログイン日時のいずれかの省略はできません。"
}