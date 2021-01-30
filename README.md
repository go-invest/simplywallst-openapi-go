# Go API client for openapi

simply-wallst API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.simplywall.st*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CompaniesApi* | [**ListDevelopmentInfo**](docs/CompaniesApi.md#listdevelopmentinfo) | **Post** /api/company/developments/info | 
*CompaniesApi* | [**ListDevelopments**](docs/CompaniesApi.md#listdevelopments) | **Get** /api/company/developments/{id} | 
*CompaniesApi* | [**ListEstimateCoverages**](docs/CompaniesApi.md#listestimatecoverages) | **Get** /api/company/estimates/coverage/{id} | 
*CompaniesApi* | [**ListNews**](docs/CompaniesApi.md#listnews) | **Get** /api/news/{id} | 
*CompaniesApi* | [**ListOwnerships**](docs/CompaniesApi.md#listownerships) | **Get** /api/company/ownership/shareholders/{id} | 
*CompaniesApi* | [**ListPrices**](docs/CompaniesApi.md#listprices) | **Get** /api/company/price/{id} | 
*CompaniesApi* | [**ReadCompanies**](docs/CompaniesApi.md#readcompanies) | **Get** /api/company/stocks/{exchange}/{sector}/{ticker}/{company} | 
*CompetitorsApi* | [**ListCompetitors**](docs/CompetitorsApi.md#listcompetitors) | **Get** /api/competitors/{id} | 
*DashboardApi* | [**ListIndustryPerformance**](docs/DashboardApi.md#listindustryperformance) | **Get** /dashboard/industry-performance/{country} | 
*DashboardApi* | [**ListInternationalMarket**](docs/DashboardApi.md#listinternationalmarket) | **Get** /dashboard/international-markets | 
*DashboardApi* | [**ListMarketPerformance**](docs/DashboardApi.md#listmarketperformance) | **Get** /dashboard/market-performance/{country} | 
*IndustriesApi* | [**ListIndustry**](docs/IndustriesApi.md#listindustry) | **Get** /api/industry/company/{id} | 
*IndustriesApi* | [**ListIndustryCountry**](docs/IndustriesApi.md#listindustrycountry) | **Get** /industry/0/{country} | 
*IndustriesApi* | [**ListIndustryTree**](docs/IndustriesApi.md#listindustrytree) | **Get** /api/industry/tree | 
*ScreenerApi* | [**ListGridViewSearch**](docs/ScreenerApi.md#listgridviewsearch) | **Post** /api/grid/view/search | 
*ScreenerApi* | [**ListGrids**](docs/ScreenerApi.md#listgrids) | **Post** /api/grid/filter | 


## Documentation For Models

 - [Companies](docs/Companies.md)
 - [Company](docs/Company.md)
 - [CompanyAnalysis](docs/CompanyAnalysis.md)
 - [CompanyAnalysisDividend](docs/CompanyAnalysisDividend.md)
 - [CompanyAnalysisExtended](docs/CompanyAnalysisExtended.md)
 - [CompanyAnalysisExtendedAnalysis](docs/CompanyAnalysisExtendedAnalysis.md)
 - [CompanyAnalysisExtendedAnalysisDividend](docs/CompanyAnalysisExtendedAnalysisDividend.md)
 - [CompanyAnalysisExtendedAnalysisFuture](docs/CompanyAnalysisExtendedAnalysisFuture.md)
 - [CompanyAnalysisExtendedAnalysisHealth](docs/CompanyAnalysisExtendedAnalysisHealth.md)
 - [CompanyAnalysisExtendedAnalysisMisc](docs/CompanyAnalysisExtendedAnalysisMisc.md)
 - [CompanyAnalysisExtendedAnalysisPast](docs/CompanyAnalysisExtendedAnalysisPast.md)
 - [CompanyAnalysisExtendedAnalysisValue](docs/CompanyAnalysisExtendedAnalysisValue.md)
 - [CompanyAnalysisExtendedData](docs/CompanyAnalysisExtendedData.md)
 - [CompanyAnalysisExtendedIndustryAverages](docs/CompanyAnalysisExtendedIndustryAverages.md)
 - [CompanyAnalysisExtendedIndustryAveragesAll](docs/CompanyAnalysisExtendedIndustryAveragesAll.md)
 - [CompanyAnalysisExtendedRaw](docs/CompanyAnalysisExtendedRaw.md)
 - [CompanyAnalysisExtendedRawData](docs/CompanyAnalysisExtendedRawData.md)
 - [CompanyAnalysisExtendedRawDataCurrencyInfo](docs/CompanyAnalysisExtendedRawDataCurrencyInfo.md)
 - [CompanyAnalysisExtendedRawDataIndustries](docs/CompanyAnalysisExtendedRawDataIndustries.md)
 - [CompanyAnalysisExtendedRawDataIndustriesData](docs/CompanyAnalysisExtendedRawDataIndustriesData.md)
 - [CompanyAnalysisExtendedRawDataInsiderTransactions](docs/CompanyAnalysisExtendedRawDataInsiderTransactions.md)
 - [CompanyAnalysisExtendedRawDataInsiderTransactionsData](docs/CompanyAnalysisExtendedRawDataInsiderTransactionsData.md)
 - [CompanyAnalysisExtendedRawDataMarketCap](docs/CompanyAnalysisExtendedRawDataMarketCap.md)
 - [CompanyAnalysisExtendedRawDataMembers](docs/CompanyAnalysisExtendedRawDataMembers.md)
 - [CompanyAnalysisExtendedRawDataMembersData](docs/CompanyAnalysisExtendedRawDataMembersData.md)
 - [CompanyAnalysisFuture](docs/CompanyAnalysisFuture.md)
 - [CompanyAnalysisPast](docs/CompanyAnalysisPast.md)
 - [CompanyDataAnalysis](docs/CompanyDataAnalysis.md)
 - [CompanyDataInfo](docs/CompanyDataInfo.md)
 - [CompanyDataInfoCeo](docs/CompanyDataInfoCeo.md)
 - [CompanyDataInfoIndustry](docs/CompanyDataInfoIndustry.md)
 - [CompanyDataScore](docs/CompanyDataScore.md)
 - [CompanyInfo](docs/CompanyInfo.md)
 - [CompanyMeta](docs/CompanyMeta.md)
 - [CompanyScore](docs/CompanyScore.md)
 - [Competitors](docs/Competitors.md)
 - [Country](docs/Country.md)
 - [CountryField](docs/CountryField.md)
 - [CountryFields](docs/CountryFields.md)
 - [CountryLinks](docs/CountryLinks.md)
 - [Development](docs/Development.md)
 - [DevelopmentEvent](docs/DevelopmentEvent.md)
 - [DevelopmentEventType](docs/DevelopmentEventType.md)
 - [DevelopmentEventTypes](docs/DevelopmentEventTypes.md)
 - [DevelopmentEvents](docs/DevelopmentEvents.md)
 - [DevelopmentInfo](docs/DevelopmentInfo.md)
 - [DevelopmentInfoAdd](docs/DevelopmentInfoAdd.md)
 - [Developments](docs/Developments.md)
 - [EstimateCoverage](docs/EstimateCoverage.md)
 - [EstimateCoverageBrokers](docs/EstimateCoverageBrokers.md)
 - [EstimateCoverageBrokersAnalyst](docs/EstimateCoverageBrokersAnalyst.md)
 - [EstimateCoverageBrokersAnalysts](docs/EstimateCoverageBrokersAnalysts.md)
 - [EstimateCoverages](docs/EstimateCoverages.md)
 - [GridViewSearch](docs/GridViewSearch.md)
 - [GridViewSearchAdd](docs/GridViewSearchAdd.md)
 - [Grids](docs/Grids.md)
 - [GridsAdd](docs/GridsAdd.md)
 - [GridsMeta](docs/GridsMeta.md)
 - [Industry](docs/Industry.md)
 - [IndustryCountry](docs/IndustryCountry.md)
 - [IndustryCountryCountries](docs/IndustryCountryCountries.md)
 - [IndustryCountryData](docs/IndustryCountryData.md)
 - [IndustryCountryLinks](docs/IndustryCountryLinks.md)
 - [IndustryPerformance](docs/IndustryPerformance.md)
 - [IndustryPerformanceData](docs/IndustryPerformanceData.md)
 - [IndustryPerformanceDate](docs/IndustryPerformanceDate.md)
 - [IndustryPerformanceDateData](docs/IndustryPerformanceDateData.md)
 - [IndustryTree](docs/IndustryTree.md)
 - [InternationalMarket](docs/InternationalMarket.md)
 - [InternationalMarketData](docs/InternationalMarketData.md)
 - [InternationalMarketReturn](docs/InternationalMarketReturn.md)
 - [InternationalMarketReturnDate](docs/InternationalMarketReturnDate.md)
 - [MarketPerformance](docs/MarketPerformance.md)
 - [MarketPerformanceData](docs/MarketPerformanceData.md)
 - [MarketPerformanceDate](docs/MarketPerformanceDate.md)
 - [MarketPerformanceDateData](docs/MarketPerformanceDateData.md)
 - [New](docs/New.md)
 - [News](docs/News.md)
 - [Ownership](docs/Ownership.md)
 - [OwnershipOwner](docs/OwnershipOwner.md)
 - [Ownerships](docs/Ownerships.md)
 - [Price](docs/Price.md)
 - [Prices](docs/Prices.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



