// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01A60005 struct {
}

func (e *PG_E01A60005) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01A60005) Message() string {
	return "前年の取引回数が最大桁数を超えています。"
}

func (e *PG_E01A60005) CanRetry() bool {
	return false
}
