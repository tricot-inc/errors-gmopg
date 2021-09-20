// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01B90005 struct {
}

func (e *PG_E01B90005) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01B90005) Message() string {
	return "自宅電話の国コードが最大桁数を超えています。"
}

func (e *PG_E01B90005) Code() string {
	return "E01B90005"
}

func (e *PG_E01B90005) CanRetry() bool {
	return false
}
