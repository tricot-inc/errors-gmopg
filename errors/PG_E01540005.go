// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01540005 struct{
}

func (e *PG_E01540005) Error() string{
    return "リンク決済呼び出しエラー 設定を確認してください。"
}

func (e *PG_E01540005) Message() string{
    return "決済キャンセル時URLが最大文字数を超えています。"
}