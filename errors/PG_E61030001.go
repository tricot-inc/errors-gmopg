// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E61030001 struct{
}

func (e *PG_E61030001) Error() string {
    return "加盟店設定エラー 決済を中止して、問い合わせにて設定状況を確認してください。"
}

func (e *PG_E61030001) Message() string {
    return "ご契約内容エラー/現在のご契約では、ご利用になれません。"
}

func (e *PG_E61030001) CanRetry() bool {
    return false
}