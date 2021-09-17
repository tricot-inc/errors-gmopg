// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01D20008 struct{
}

func (e *PG_E01D20008) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01D20008) Message() string{
    return "納品先電子メールアドレスの書式が正しくありません。"
}