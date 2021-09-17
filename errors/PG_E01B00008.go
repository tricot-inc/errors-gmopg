// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01B00008 struct {
}

func (e *PG_E01B00008) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01B00008) Message() string {
	return "請求先住所と配送先住所の一致/不一致の書式が正しくありません。"
}

func (e *PG_E01B00008) CanRetry() bool {
	return false
}
