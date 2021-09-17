// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E11310005 struct {
}

func (e *PG_E11310005) Error() string {
	return "リンク決済エラー カード登録の制限に抵触します。"
}

func (e *PG_E11310005) Message() string {
	return "すでにカードを登録している会員は、取引後カード登録を実行できません。"
}

func (e *PG_E11310005) Code() string {
	return "E11310005"
}

func (e *PG_E11310005) CanRetry() bool {
	return false
}
