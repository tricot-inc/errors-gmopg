// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E90010001 struct {
}

func (e *PG_E90010001) Error() string {
	return "2重送信エラー(通信) 管理画面にて状態をご確認ください。後発のリクエストが正常に通知されない場合があります。"
}

func (e *PG_E90010001) Message() string {
	return "現在処理を行っているのでもうしばらくお待ちください。"
}

func (e *PG_E90010001) CanRetry() bool {
	return false
}
