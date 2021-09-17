// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_EX1000302 struct {
}

func (e *PG_EX1000302) Error() string {
	return "トークンエラー 指定されたトークンは利用済みです。"
}

func (e *PG_EX1000302) Message() string {
	return "決済処理に失敗しました。もう一度カード番号を入力してください。"
}

func (e *PG_EX1000302) CanRetry() bool {
	return false
}
