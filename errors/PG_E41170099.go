// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E41170099 struct{
}

func (e *PG_E41170099) Error() string{
    return "カード番号チェックディジットエラー(ホスト検出) 再入力をカード所有者に依頼してください。"
}

func (e *PG_E41170099) Message() string{
    return "カード番号に誤りがあります。再度確認して入力してください。"
}