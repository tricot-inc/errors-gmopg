// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01890006 struct{
}

func (e *PG_E01890006) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01890006) Message() string{
    return "登録済みカード登録連番に数字以外の文字が含まれています。"
}