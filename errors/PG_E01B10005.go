// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01B10005 struct {
}

func (e *PG_E01B10005) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01B10005) Message() string {
	return "請求先住所の都市が最大桁数を超えています。"
}

func (e *PG_E01B10005) Code() string {
	return "E01B10005"
}

func (e *PG_E01B10005) CanRetry() bool {
	return false
}
