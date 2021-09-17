// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E00000000 struct{
}

func (e *PG_E00000000) Error() string {
    return "結果通知プログラム疎通確認用 疎通確認用なので、対処する必要性はありません。"
}

func (e *PG_E00000000) Message() string {
    return "特になし"
}

func (e *PG_E00000000) CanRetry() bool {
    return false
}