// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01110010 struct {
}

func (e *PG_E01110010) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01110010) Message() string {
	return "指定された取引は決済が完了していません。"
}

func (e *PG_E01110010) CanRetry() bool {
	return false
}
