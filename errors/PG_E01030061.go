// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01030061 struct {
}

func (e *PG_E01030061) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01030061) Message() string {
	return "強制返品はご利用できません。"
}

func (e *PG_E01030061) Code() string {
	return "E01030061"
}

func (e *PG_E01030061) CanRetry() bool {
	return false
}
