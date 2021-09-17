// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01D10010 struct{
}

func (e *PG_E01D10010) Error() string {
    return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01D10010) Message() string {
    return "配送先住所の州または都道府県番号を指定する場合は配送先住所の国番号を省略できません。"
}

func (e *PG_E01D10010) CanRetry() bool {
    return false
}