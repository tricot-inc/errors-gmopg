// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01040001 struct{
}

func (e *PG_E01040001) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01040001) Message() string{
    return "オーダーIDが指定されていません。"
}