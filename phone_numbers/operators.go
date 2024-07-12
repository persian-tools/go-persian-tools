package phonenumbers

import "errors"

type Operator string
type SimType string

var ErrInvalidFormat = errors.New("invalid phone number format")

// List of valid prefixes
var Prefixes = []string{"+98", "98", "0098", "0"}

const (
	ShatelMobile Operator = "ShatelMobile"
	MCI          Operator = "MCI"
	Irancell     Operator = "Irancell"
	Taliya       Operator = "Taliya"
	RightTel     Operator = "RightTel"

	Permanent SimType = "Permanent"
	Credit    SimType = "Credit"
)

type OperatorDetails struct {
	base     string
	province []string
	model    string
	operator Operator
	simTypes []SimType
}

func (o Operator) Details() map[string]OperatorDetails {
	switch o {
	case MCI:
		return MCIMap
	case Taliya:
		return TALIYA
	case RightTel:
		return RIGHTTEL
	case Irancell:
		return IRANCELL
	case ShatelMobile:
		return SHATELMOBILE
	default:
		return nil
	}
}

func (od *OperatorDetails) GetProvinceList() []string {
	return od.province
}

func (od *OperatorDetails) GetBase() string {
	return od.base
}

func (od *OperatorDetails) GetModel() string {
	return od.model
}

func (od *OperatorDetails) GetOperator() Operator {
	return od.operator
}

func (od *OperatorDetails) GetSimTypeList() []SimType {
	return od.simTypes
}
