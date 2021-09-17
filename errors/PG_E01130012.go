// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01130012 struct{
}

func (e *PG_E01130012) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01130012) Message() string{
    return "カード会社略称が最大バイト数を超えています。"
}