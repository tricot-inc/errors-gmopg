// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01800010 struct{
}

func (e *PG_E01800010) Error() string{
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01800010) Message() string{
    return "暗証番号は利用できません。"
}