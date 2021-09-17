// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E91029998 struct {
}

func (e *PG_E91029998) Error() string {
	return "システムエラー(通信) 購入者に取引失敗を表示し、問い合わせにて状況を確認してください。"
}

func (e *PG_E91029998) Message() string {
	return "決済処理に失敗しました。該当のお取引について、店舗までお問い合わせください。"
}

func (e *PG_E91029998) CanRetry() bool {
	return false
}
