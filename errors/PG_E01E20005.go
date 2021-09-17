// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01E20005 struct {
}

func (e *PG_E01E20005) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01E20005) Message() string {
	return "継続課金の課金最小間隔日数が最大桁数を超えています。"
}

func (e *PG_E01E20005) CanRetry() bool {
	return false
}
