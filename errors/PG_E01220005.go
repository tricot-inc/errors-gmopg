// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01220005 struct{
}

func (e *PG_E01220005) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01220005) Message() string{
    return "会員IDが最大桁数を超えています。"
}