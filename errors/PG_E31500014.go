// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E31500014 struct{
}

func (e *PG_E31500014) Error() string {
    return "加盟店実行エラー リクエストメソッドがGETになっています。POSTに変更してください。"
}

func (e *PG_E31500014) Message() string {
    return "-"
}

func (e *PG_E31500014) CanRetry() bool {
    return false
}