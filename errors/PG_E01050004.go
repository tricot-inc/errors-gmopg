// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01050004 struct {
}

func (e *PG_E01050004) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01050004) Message() string {
	return "指定した処理区分の処理は実行できません。"
}

func (e *PG_E01050004) Code() string {
	return "E01050004"
}

func (e *PG_E01050004) CanRetry() bool {
	return false
}
