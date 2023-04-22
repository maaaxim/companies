package models

import (
	"fmt"
)

type CompanyType struct {
	typeValue string
}

func (c CompanyType) Value() string {
	return c.typeValue
}

var (
	CompanyTypeUnknown               = CompanyType{""}
	CompanyTypeCorporation           = CompanyType{"Corporation"}
	CompanyTypeNonProfit             = CompanyType{"NonProfit"}
	CompanyTypeTagCooperative        = CompanyType{"Cooperative"}
	CompanyTypeTagSoleProprietorship = CompanyType{"Sole Proprietorship"}
)

func NewCompanyTypeFromString(typeValue string) (CompanyType, error) {
	switch typeValue {
	case CompanyTypeCorporation.typeValue:
		return CompanyTypeCorporation, nil
	case CompanyTypeNonProfit.typeValue:
		return CompanyTypeNonProfit, nil
	case CompanyTypeTagCooperative.typeValue:
		return CompanyTypeTagCooperative, nil
	case CompanyTypeTagSoleProprietorship.typeValue:
		return CompanyTypeTagSoleProprietorship, nil
	}

	return CompanyTypeUnknown, fmt.Errorf("unknown CompanyType value: %s", typeValue)
}
