// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01730006 struct {
}

func (e *PG_E01730006) Error() string {
	return "ファイル内容エラー 設定を確認してください。"
}

func (e *PG_E01730006) Message() string {
	return "商品コードが”0000990”ではありません。"
}

func (e *PG_E01730006) Code() string {
	return "E01730006"
}

func (e *PG_E01730006) CanRetry() bool {
	return false
}
