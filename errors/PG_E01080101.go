// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01080101 struct{
}

func (e *PG_E01080101) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01080101) Message() string {
    return "3D必須店舗にも関わらず3Dセキュア使用フラグがOFFになっています。"
}

func (e *PG_E01080101) CanRetry() bool {
    return false
}