// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E41170002 struct{
}

func (e *PG_E41170002) Error() string{
    return "カード会社未対応エラー(ホスト検出)再入力をカード所有者に依頼してください。"
}

func (e *PG_E41170002) Message() string{
    return "入力されたカード会社に対応していません。別のカード番号を入力してください。"
}