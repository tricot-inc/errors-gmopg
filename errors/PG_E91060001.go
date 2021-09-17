// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E91060001 struct {
}

func (e *PG_E91060001) Error() string {
	return "システムエラー(モジュールタイプ) 購入者に取引失敗を表示し、問い合わせにて状況を確認してください。"
}

func (e *PG_E91060001) Message() string {
	return "システムの内部エラーです。発生時刻や呼び出しパラメータをご確認のうえ、お問い合わせください。"
}

func (e *PG_E91060001) Code() string {
	return "E91060001"
}

func (e *PG_E91060001) CanRetry() bool {
	return false
}
