// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E11310001 struct {
}

func (e *PG_E11310001) Error() string {
	return "リンク決済エラー 指定されたオーダーはリンク決済呼び出しされていません。"
}

func (e *PG_E11310001) Message() string {
	return "この取引はリンク決済を実行できません。"
}

func (e *PG_E11310001) CanRetry() bool {
	return false
}
