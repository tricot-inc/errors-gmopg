// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01220001 struct{
}

func (e *PG_E01220001) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01220001) Message() string{
    return "会員IDが指定されていません。"
}