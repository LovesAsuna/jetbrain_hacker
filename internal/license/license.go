package license

import (
	"github.com/dromara/carbon/v2"
)

type License struct {
	LicenseID          string    `json:"licenseId"`
	LicenseeName       string    `json:"licenseeName"`
	AssigneeName       string    `json:"assigneeName"`
	AssigneeEmail      string    `json:"assigneeEmail"`
	LicenseRestriction string    `json:"licenseRestriction"`
	CheckConcurrentUse bool      `json:"checkConcurrentUse"`
	Products           []Product `json:"products"`
	Metadata           string    `json:"metadata"`
	Hash               string    `json:"hash"`
	GracePeriodDays    int       `json:"gracePeriodDays"`
	AutoProlongated    bool      `json:"autoProlongated"`
	IsAutoProlongated  bool      `json:"isAutoProlongated"`
}

type Product struct {
	Code         string `json:"code"`
	FallbackDate string `json:"fallbackDate"`
	PaidUpTo     string `json:"paidUpTo"`
}

type ProductDto struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func Generate(licenseId, licenseeName, assigneeName, assigneeEmail, time string) (*License, error) {
	fallBackDate := carbon.Now().SetLayout(carbon.DateLayout).String()
	paidUpTo := carbon.ParseByLayout(time, carbon.DateLayout).String()

	codes := []string{"YTD", "QDGO", "MF", "DG", "PS", "QA", "IIE", "YTWE", "FLS", "DLE", "RFU", "PPS", "PCWMP", "II", "TCC", "RSU", "PCC", "RC", "PCE", "FLIJ", "TBA", "DL", "SPP", "QDCLD", "SPA", "DMCLP", "PSW", "GW", "PSI", "IIU", "DMU", "PWS", "HB", "WS", "PCP", "KT", "DCCLT", "RSCLT", "WRS", "RSC", "RRD", "TC", "IIC", "QDPY", "DPK", "DC", "PDB", "DPPS", "QDPHP", "GO", "HCC", "RDCPPP", "QDJVMC", "CL", "DM", "CWML", "FLL", "RR", "QDJS", "RS", "RM", "DS", "MPS", "DPN", "US", "CLN", "DPCLT", "RSV", "MPSIIP", "DB", "QDANDC", "AC", "QDJVM", "PRB", "RD", "CWMR", "SP", "RS0", "DP", "RSF", "PGO", "QDPYC", "PPC", "PC", "EHS", "RSCHB", "FL", "QDNET", "JCD"}
	products := make([]Product, 0, len(codes))
	for _, code := range codes {
		products = append(products, Product{
			Code:         code,
			FallbackDate: fallBackDate,
			PaidUpTo:     paidUpTo,
		})
	}

	license := &License{
		LicenseID:         licenseId,
		LicenseeName:      licenseeName,
		AssigneeName:      assigneeName,
		AssigneeEmail:     assigneeEmail,
		Products:          products,
		Metadata:          "0120230102PPAA013009",
		Hash:              "41472961/0:1563609451",
		GracePeriodDays:   7,
		AutoProlongated:   true,
		IsAutoProlongated: true,
	}
	return license, nil
}
