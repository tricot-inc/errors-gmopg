// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01E00008 struct {
}

func (e *PG_E01E00008) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01E00008) Message() string {
	return "取引の出荷方法の書式が正しくありません。"
}

func (e *PG_E01E00008) Code() string {
	return "E01E00008"
}

func (e *PG_E01E00008) CanRetry() bool {
	return false
}
