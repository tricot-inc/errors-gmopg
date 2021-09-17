// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01999998 struct{
}

func (e *PG_E01999998) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01999998) Message() string{
    return "項目1「フォーマットバージョン」に“001“が指定されていません。"
}