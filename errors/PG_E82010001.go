// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E82010001 struct {
}

func (e *PG_E82010001) Error() string {
	return "実行の排他エラー 同じサイトから同時に実行されました。このエラーは無視して頂いても構いません。先発の実行分で正常に実行されます。"
}

func (e *PG_E82010001) Message() string {
	return "実行中にエラーが発生しました。処理は開始されませんでした。"
}

func (e *PG_E82010001) Code() string {
	return "E82010001"
}

func (e *PG_E82010001) CanRetry() bool {
	return false
}
