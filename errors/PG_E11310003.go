// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E11310003 struct {
}

func (e *PG_E11310003) Error() string {
	return "リンク決済エラー リトライ回数オーバー"
}

func (e *PG_E11310003) Message() string {
	return "この取引はリンク決済を実行できません。"
}

func (e *PG_E11310003) Code() string {
	return "E11310003"
}

func (e *PG_E11310003) CanRetry() bool {
	return false
}
