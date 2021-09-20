// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01270010 struct {
}

func (e *PG_E01270010) Error() string {
	return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01270010) Message() string {
	return "指定された支払回数はご利用できません。"
}

func (e *PG_E01270010) Code() string {
	return "E01270010"
}

func (e *PG_E01270010) CanRetry() bool {
	return false
}
