// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01050002 struct{
}

func (e *PG_E01050002) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01050002) Message() string {
    return "指定された処理区分は定義されていません。"
}

func (e *PG_E01050002) CanRetry() bool {
    return false
}