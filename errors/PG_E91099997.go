// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E91099997 struct {
}

func (e *PG_E91099997) Error() string {
	return "システムエラー(通信)"
}

func (e *PG_E91099997) Message() string {
	return "リクエストされたAPIは存在しません。URLをお確かめください。"
}

func (e *PG_E91099997) Code() string {
	return "E91099997"
}

func (e *PG_E91099997) CanRetry() bool {
	return false
}
