// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01080007 struct {
}

func (e *PG_E01080007) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01080007) Message() string {
	return "3Dセキュア使用フラグに0,1以外の値が指定されています。"
}

func (e *PG_E01080007) Code() string {
	return "E01080007"
}

func (e *PG_E01080007) CanRetry() bool {
	return false
}
