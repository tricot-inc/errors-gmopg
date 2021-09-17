// Code generated from gen-error/main.go; DO NOT EDIT
package errors

type PG_E92000002 struct{
}

func (e *PG_E92000002) Error() string{
    return "流量制限オーバーエラー 同時処理数が規定値を超えています。"
}

func (e *PG_E92000002) Message() string{
    return "只今、大変込み合っていますので、しばらく時間をあけて再度決済を行ってください。"
}