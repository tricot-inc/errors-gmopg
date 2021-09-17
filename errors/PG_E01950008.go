// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01950008 struct {
}

func (e *PG_E01950008) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01950008) Message() string {
	return "3DS2.0非対応時取り扱いの書式が正しくありません。"
}

func (e *PG_E01950008) CanRetry() bool {
	return false
}
