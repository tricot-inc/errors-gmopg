// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01500012 struct {
}

func (e *PG_E01500012) Error() string {
	return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01500012) Message() string {
	return "ショップ情報文字列が他の項目と矛盾しています。"
}

func (e *PG_E01500012) Code() string {
	return "E01500012"
}

func (e *PG_E01500012) CanRetry() bool {
	return false
}
