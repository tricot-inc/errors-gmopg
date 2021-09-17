// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01800050 struct {
}

func (e *PG_E01800050) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01800050) Message() string {
	return "暗証番号が不正です。(0000は使用できません)"
}

func (e *PG_E01800050) CanRetry() bool {
	return false
}
