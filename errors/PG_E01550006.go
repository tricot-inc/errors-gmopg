// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01550006 struct {
}

func (e *PG_E01550006) Error() string {
	return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01550006) Message() string {
	return "日時情報文字列に無効な文字が含まれます。"
}

func (e *PG_E01550006) CanRetry() bool {
	return false
}
