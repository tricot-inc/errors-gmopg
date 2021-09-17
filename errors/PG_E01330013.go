// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01330013 struct {
}

func (e *PG_E01330013) Error() string {
	return "入力パラメータエラー 再入力をカード所有者に依頼してください。"
}

func (e *PG_E01330013) Message() string {
	return "加盟店自由項目2の値に利用できない文字が含まれています。"
}

func (e *PG_E01330013) Code() string {
	return "E01330013"
}

func (e *PG_E01330013) CanRetry() bool {
	return false
}
