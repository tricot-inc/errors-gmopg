// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01B70010 struct {
}

func (e *PG_E01B70010) Error() string {
	return "入力パラメータエラー 設定を確認してください。"
}

func (e *PG_E01B70010) Message() string {
	return "請求先住所の州または都道府県番号を指定する場合は請求先住所の国番号を省略できません。"
}

func (e *PG_E01B70010) Code() string {
	return "E01B70010"
}

func (e *PG_E01B70010) CanRetry() bool {
	return false
}
