// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01A80008 struct{
}

func (e *PG_E01A80008) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01A80008) Message() string{
    return "ログイン方法の書式が正しくありません。"
}