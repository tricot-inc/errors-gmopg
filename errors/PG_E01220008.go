// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01220008 struct{
}

func (e *PG_E01220008) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01220008) Message() string{
    return "会員IDの書式が正しくありません。"
}