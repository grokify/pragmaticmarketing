package pragmaticmarketing

import "strings"

/*
https://medium.com/i-want-to-be-a-product-manager-when-i-grow-up/the-pragmatic-marketing-framework-ef29eb7270b1

https://www.pragmaticinstitute.com/resources/articles/product/the-strategic-role-of-product-management-when-development-goes-agile/

*/

const (
	CategoryMarket    = "Market"
	CategoryFocus     = "Focus"
	CategoryBusiness  = "Business"
	CategoryPlanning  = "Planning"
	CategoryPrograms  = "Programs"
	CategoryReadiness = "Readiness"
	CategorySupport   = "Support"
	FocusMarket       = "Market/Business"
	FocusProduct      = "Product/Technical"
	HorizonStrategy   = "Strategic"
	HorizonTactical   = "Tactical"
)

type Activity struct {
	Name     string
	Category string
	Focus    string
	Rank     int
}

func CategoryNames() []string {
	return []string{
		CategoryMarket, CategoryFocus, CategoryBusiness,
		CategoryPlanning, CategoryPrograms, CategoryReadiness, CategorySupport}
}

func FocusNames() []string {
	return []string{FocusMarket, FocusProduct}
}

func Activities() []Activity {
	return []Activity{
		{
			Name:     "Market Problems",
			Category: CategoryMarket,
			Rank:     0,
		},
	}
}

const (
	rawMarketMarket   = "MarketProblems;Win/Loss Analysis;Distinct Competencies"
	rawMarketFocus    = "Market Definition;Distribution Strategy;Product Portfolio"
	rawMarketBusiness = "Business Plan;Pricing;Build, Buy or Partner;Product Profitability"
	rawMarketPlanning = "Positioning;Buying Process;Buyer Personas;User Personas"
	rawMarketPrograms = "Market Plan;Custommer Acquisition;Customer Retention;Program Effectivenss"

	rawProductMarket    = "Competitive Landscape;Asset Assessment"
	rawProductFocus     = "Product Roadmap"
	rawProductBusiness  = "Innovation"
	rawProductPlanning  = "Requirements;Use Scenarios;Stakeholder Communications"
	rawProductPrograms  = "Launch Program;Thought Leadership;Lead Generation;Referrals & References"
	rawProductReadiness = "Sales Process;Collateral;Sales Tools;Channel Training"
	rawProductSupport   = "Presentations & Demos;Special Calls;Event Support;Channel Support"
)

func AllActivities() []Activity {
	activities := []Activity{}
	for _, focusName := range FocusNames() {
		for _, categoryName := range CategoryNames() {
			activities = buildActivities(activities, focusName, categoryName)
		}
	}
	return activities
}

func rawActivities(focus, category string) string {
	if focus == FocusMarket {
		switch category {
		case CategoryMarket:
			return rawMarketMarket
		case CategoryFocus:
			return rawMarketFocus
		case CategoryBusiness:
			return rawMarketBusiness
		case CategoryPlanning:
			return rawMarketPlanning
		case CategoryPrograms:
			return rawMarketPrograms
		default:
			return ""
		}
	}
	switch category {
	case CategoryMarket:
		return rawProductMarket
	case CategoryFocus:
		return rawProductFocus
	case CategoryBusiness:
		return rawProductBusiness
	case CategoryPlanning:
		return rawProductPlanning
	case CategoryPrograms:
		return rawProductPrograms
	case CategoryReadiness:
		return rawProductReadiness
	case CategorySupport:
		return rawProductSupport
	default:
		return ""
	}
}

func buildActivities(activities []Activity, focus, category string) []Activity {
	activityNames := strings.Split(rawActivities(focus, category), ";")
	for i, activityName := range activityNames {
		activities = append(activities, Activity{
			Name:     activityName,
			Category: category,
			Focus:    focus,
			Rank:     i})
	}
	return activities
}
