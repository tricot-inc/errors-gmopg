// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01410010 struct {
}

func (e *PG_E01410010) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01410010) Message() string {
	return "該当取引は操作禁止状態です。"
}

func (e *PG_E01410010) Code() string {
	return "E01410010"
}

func (e *PG_E01410010) CanRetry() bool {
	return false
}
