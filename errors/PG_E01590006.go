// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01590006 struct{
}

func (e *PG_E01590006) Error() string{
    return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01590006) Message() string{
    return "商品コードに無効な文字が含まれます。"
}