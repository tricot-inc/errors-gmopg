// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01670013 struct {
}

func (e *PG_E01670013) Error() string {
	return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01670013) Message() string {
	return "出力エンコーディングにサポートされない値が設定されています。"
}

func (e *PG_E01670013) Code() string {
	return "E01670013"
}

func (e *PG_E01670013) CanRetry() bool {
	return false
}
