// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01660013 struct {
}

func (e *PG_E01660013) Error() string {
	return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01660013) Message() string {
	return "言語パラメータにサポートされない値が設定されています。"
}

func (e *PG_E01660013) Code() string {
	return "E01660013"
}

func (e *PG_E01660013) CanRetry() bool {
	return false
}
