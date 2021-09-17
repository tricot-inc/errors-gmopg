// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01520002 struct {
}

func (e *PG_E01520002) Error() string {
	return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01520002) Message() string {
	return "ユーザー利用端末情報に無効な値が設定されています。"
}

func (e *PG_E01520002) Code() string {
	return "E01520002"
}

func (e *PG_E01520002) CanRetry() bool {
	return false
}
