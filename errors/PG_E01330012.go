// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01330012 struct {
}

func (e *PG_E01330012) Error() string {
	return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01330012) Message() string {
	return "加盟店自由項目2の値が最大バイト数を超えています。"
}

func (e *PG_E01330012) CanRetry() bool {
	return false
}
