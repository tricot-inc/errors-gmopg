// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01890002 struct{
}

func (e *PG_E01890002) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01890002) Message() string{
    return "指定された登録済みカードが存在しません。"
}