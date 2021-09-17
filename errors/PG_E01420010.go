// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01420010 struct{
}

func (e *PG_E01420010) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01420010) Message() string {
    return "仮売上有効期間を超えています。"
}

func (e *PG_E01420010) CanRetry() bool {
    return false
}