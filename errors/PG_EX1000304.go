// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_EX1000304 struct {
}

func (e *PG_EX1000304) Error() string {
	return "トークンエラー 指定されたトークンが正しくありません。"
}

func (e *PG_EX1000304) Message() string {
	return "決済処理に失敗しました。もう一度カード番号を入力してください。"
}

func (e *PG_EX1000304) CanRetry() bool {
	return false
}
