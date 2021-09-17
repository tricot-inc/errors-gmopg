// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E01860008 struct{
}

func (e *PG_E01860008) Error() string{
    return "カード番号マスクフラグ カード番号マスクフラグの書式が正しくありません。"
}

func (e *PG_E01860008) Message() string{
    return "カード番号マスクフラグの書式が正しくありません。"
}