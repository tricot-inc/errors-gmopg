// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E61010002 struct {
}

func (e *PG_E61010002) Error() string {
	return "加盟店設定エラー 問い合わせにてカード会社契約状況を確認してください。"
}

func (e *PG_E61010002) Message() string {
	return "ご利用出来ないカードをご利用になったもしくはカード番号が誤っております。"
}

func (e *PG_E61010002) CanRetry() bool {
	return false
}
