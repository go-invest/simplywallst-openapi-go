/*
 * simply-wallst
 *
 * simply-wallst API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CompanyDataAnalysis struct for CompanyDataAnalysis
type CompanyDataAnalysis struct {
	Id                string                  `json:"id,omitempty"`
	SharePrice        float32                 `json:"share_price,omitempty"`
	MarketCap         float32                 `json:"market_cap,omitempty"`
	IntrinsicDiscount float32                 `json:"intrinsic_discount,omitempty"`
	Pe                float32                 `json:"pe,omitempty"`
	Pb                float32                 `json:"pb,omitempty"`
	Peg               float32                 `json:"peg,omitempty"`
	Roe               float32                 `json:"roe,omitempty"`
	Roa               float32                 `json:"roa,omitempty"`
	Eps               float32                 `json:"eps,omitempty"`
	DebtEquity        float32                 `json:"debt_equity,omitempty"`
	AnalystCount      int32                   `json:"analyst_count,omitempty"`
	Dividend          CompanyAnalysisDividend `json:"dividend,omitempty"`
	Future            CompanyAnalysisFuture   `json:"future,omitempty"`
	Past              CompanyAnalysisPast     `json:"past,omitempty"`
}
