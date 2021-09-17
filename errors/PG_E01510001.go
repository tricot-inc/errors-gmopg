// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01510001 struct {
}

func (e *PG_E01510001) Error() string {
	return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01510001) Message() string {
	return "購買情報文字列が設定されていません。"
}

func (e *PG_E01510001) Code() string {
	return "E01510001"
}

func (e *PG_E01510001) CanRetry() bool {
	return false
}
