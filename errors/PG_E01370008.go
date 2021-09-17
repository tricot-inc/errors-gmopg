// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01370008 struct{
}

func (e *PG_E01370008) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01370008) Message() string{
    return "3Dセキュア表示店舗名の書式が正しくありません。"
}