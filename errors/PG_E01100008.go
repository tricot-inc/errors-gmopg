// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01100008 struct{
}

func (e *PG_E01100008) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01100008) Message() string{
    return "取引パスワードの書式が正しくありません。"
}