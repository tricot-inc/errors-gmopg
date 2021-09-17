// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E11010999 struct{
}

func (e *PG_E11010999) Error() string{
    return "取引エラー すでに取引が完了している可能性があります。"
}

func (e *PG_E11010999) Message() string{
    return "特になし"
}