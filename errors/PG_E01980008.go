// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01980008 struct{
}

func (e *PG_E01980008) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01980008) Message() string{
    return "カード会員パスワード変更日の書式が正しくありません。"
}