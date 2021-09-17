// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01370012 struct{
}

func (e *PG_E01370012) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01370012) Message() string {
    return "3Dセキュア表示店舗名の値が最大バイト数を超えています。"
}

func (e *PG_E01370012) CanRetry() bool {
    return false
}