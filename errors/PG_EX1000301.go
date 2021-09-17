// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_EX1000301 struct {
}

func (e *PG_EX1000301) Error() string {
	return "トークンエラー 指定されたトークンが存在しません。"
}

func (e *PG_EX1000301) Message() string {
	return "決済処理に失敗しました。もう一度カード番号を入力してください。"
}

func (e *PG_EX1000301) Code() string {
	return "EX1000301"
}

func (e *PG_EX1000301) CanRetry() bool {
	return false
}
