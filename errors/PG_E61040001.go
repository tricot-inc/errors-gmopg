// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E61040001 struct{
}

func (e *PG_E61040001) Error() string {
    return "加盟店設定エラー カード番号を入力することはできません。"
}

func (e *PG_E61040001) Message() string {
    return "現在のご契約では、カード番号を指定した決済処理は許可されていません。"
}

func (e *PG_E61040001) CanRetry() bool {
    return false
}