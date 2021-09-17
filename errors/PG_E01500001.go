// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01500001 struct {
}

func (e *PG_E01500001) Error() string {
	return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01500001) Message() string {
	return "ショップ情報文字列が設定されていません。"
}

func (e *PG_E01500001) CanRetry() bool {
	return false
}
