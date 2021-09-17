// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01080010 struct{
}

func (e *PG_E01080010) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01080010) Message() string {
    return "3Dセキュア利用しない契約ですが、3Dセキュア使用フラグ(TdFlag)が指定されています。"
}

func (e *PG_E01080010) CanRetry() bool {
    return false
}