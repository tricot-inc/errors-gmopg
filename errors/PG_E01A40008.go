// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01A40008 struct{
}

func (e *PG_E01A40008) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01A40008) Message() string {
    return "カード会員の不審行為情報の書式が正しくありません。"
}

func (e *PG_E01A40008) CanRetry() bool {
    return false
}