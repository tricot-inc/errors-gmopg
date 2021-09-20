// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01590005 struct {
}

func (e *PG_E01590005) Error() string {
	return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01590005) Message() string {
	return "商品コードが最大桁数を超えています。"
}

func (e *PG_E01590005) Code() string {
	return "E01590005"
}

func (e *PG_E01590005) CanRetry() bool {
	return false
}
