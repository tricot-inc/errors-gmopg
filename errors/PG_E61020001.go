// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E61020001 struct{
}

func (e *PG_E61020001) Error() string{
    return "加盟店設定エラー 再入力をお客様に依頼してください。"
}

func (e *PG_E61020001) Message() string{
    return "指定の決済方法は利用停止になっています。"
}