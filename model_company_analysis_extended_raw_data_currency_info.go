/*
 * simply-wallst
 *
 * simply-wallst API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CompanyAnalysisExtendedRawDataCurrencyInfo struct for CompanyAnalysisExtendedRawDataCurrencyInfo
type CompanyAnalysisExtendedRawDataCurrencyInfo struct {
	ReportingUnit                    int64  `json:"reporting_unit,omitempty"`
	ReportingUnitText                string `json:"reporting_unit_text,omitempty"`
	ReportingUnitAbs                 int64  `json:"reporting_unit_abs,omitempty"`
	ReportingUnitTextAbs             string `json:"reporting_unit_text_abs,omitempty"`
	ReportingCurrencyIso             string `json:"reporting_currency_iso,omitempty"`
	ReportingCurrencySymbol          string `json:"reporting_currency_symbol,omitempty"`
	PrimaryTradingItemCurrencyIso    string `json:"primary_trading_item_currency_iso,omitempty"`
	PrimaryTradingItemCurrencySymbol string `json:"primary_trading_item_currency_symbol,omitempty"`
	TradingItemCurrencyIso           string `json:"trading_item_currency_iso,omitempty"`
	TradingItemCurrencySymbol        string `json:"trading_item_currency_symbol,omitempty"`
}
