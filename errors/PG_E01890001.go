// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01890001 struct{
}

func (e *PG_E01890001) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01890001) Message() string{
    return "登録済みカード登録連番が指定されていません。"
}