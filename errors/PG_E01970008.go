// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01970008 struct{
}

func (e *PG_E01970008) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01970008) Message() string{
    return "カード会員作成日の書式が正しくありません。"
}