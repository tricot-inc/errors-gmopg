// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01E40001 struct{
}

func (e *PG_E01E40001) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01E40001) Message() string{
    return "認証状態が指定されていません。"
}