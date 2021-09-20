// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01040003 struct {
}

func (e *PG_E01040003) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01040003) Message() string {
	return "オーダーIDが最大文字数を超えています。"
}

func (e *PG_E01040003) Code() string {
	return "E01040003"
}

func (e *PG_E01040003) CanRetry() bool {
	return false
}
