// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01620005 struct {
}

func (e *PG_E01620005) Error() string {
	return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01620005) Message() string {
	return "セッションタイムアウト値が0~9999の範囲外です。"
}

func (e *PG_E01620005) Code() string {
	return "E01620005"
}

func (e *PG_E01620005) CanRetry() bool {
	return false
}
