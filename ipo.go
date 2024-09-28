package main

import "time"

type IPO struct {
	ID                  string     `json:"id"`                  // Unique ID of the IPO
	Status              string     `json:"status"`              // Current status of the IPO (e.g., open, closed)
	Slug                string     `json:"slug"`                // URL-friendly identifier
	Name                string     `json:"name"`                // Name of the company issuing the IPO
	Type                string     `json:"type"`                // Type of IPO (e.g., SME)
	StartDate           time.Time  `json:"startDate"`           // Date when the IPO starts
	EndDate             time.Time  `json:"endDate"`             // Date when the IPO ends
	IsDateFinal         bool       `json:"isDateFinal"`         // Indicates if the IPO dates are final
	ListingDate         time.Time  `json:"listingDate"`         // Date when the IPO will be listed on the stock exchange
	PriceRange          PriceRange `json:"priceRange"`          // Price range for the IPO shares
	MinQty              int        `json:"minQty"`              // Minimum quantity of shares to buy
	SizePerLot          int        `json:"sizePerLot"`          // Number of shares per lot
	IssueSize           string     `json:"issueSize"`           // Total size of the IPO issue
	GmpDetails          GmpDetails `json:"gmpDetails"`          // Details about the grey market premium
	ApplyRecommendation *bool      `json:"applyRecommendation"` // Recommendation for applying to the IPO (true - recommended, false - not recommended, null - no recommendation)
	LastUpdatedAt       time.Time  `json:"lastUpdatedAt"`       // Date when the IPO data was last updated
	CreatedAt           time.Time  `json:"createdAt"`           // Date when the IPO data was created
}

type IPODetail struct {
	ID                  string     `json:"id"`                  // Unique ID of the IPO
	Status              string     `json:"status"`              // Current status of the IPO (e.g., open, closed)
	Slug                string     `json:"slug"`                // URL-friendly identifier
	InfoURL             string     `json:"infoUrl"`             // Link to more information about the IPO
	Name                string     `json:"name"`                // Name of the company issuing the IPO
	CompanyDescription  string     `json:"about"`  // Description of the company
	Symbol              string     `json:"symbol"`              // Stock symbol
	Type                string     `json:"type"`                // Type of IPO (e.g., SME)
	StartDate           time.Time  `json:"startDate"`           // Date when the IPO starts
	EndDate             time.Time  `json:"endDate"`             // Date when the IPO ends
	IsDateFinal         bool       `json:"isDateFinal"`         // Indicates if the IPO dates are final
	ListingDate         time.Time  `json:"listingDate"`         // Date when the IPO will be listed on the stock exchange
	PriceRange          PriceRange `json:"priceRange"`          // Price range for the IPO shares
	MinQty              int        `json:"minQty"`              // Minimum quantity of shares to buy
	SizePerLot          int        `json:"sizePerLot"`          // Number of shares per lot
	Logo                string     `json:"logo"`                // URL to the logo image of the company
	IssueSize           string     `json:"issueSize"`           // Total size of the IPO issue
	ProspectusURL       string     `json:"prospectusUrl"`       // URL to download the IPO prospectus
	Schedule            []Event    `json:"schedule"`            // List of events and their respective dates
	GmpDetails          GmpDetails `json:"gmpDetails"`          // Details about the grey market premium
	ApplyRecommendation *bool      `json:"applyRecommendation"` // Recommendation for applying to the IPO (true - recommended, false - not recommended, null - no recommendation)
	LastUpdatedAt       time.Time  `json:"lastUpdatedAt"`       // Date when the IPO data was last updated
	CreatedAt           time.Time  `json:"createdAt"`           // Date when the IPO data was created
	Quota               Quota      `json:"quota"`               // Quota details for the IPO
}

type PriceRange struct {
	Min float64 `json:"min"` // Minimum price in the range
	Max float64 `json:"max"` // Maximum price in the range
}

type Event struct {
	Event      string    `json:"event"`      // Description of the event (e.g., issue_open, listing_date)
	Date       time.Time `json:"date"`       // Date when the event occurs
	EVentTitle string    `json:"eventTitle"` // Title of the event
}

type GmpDetails struct {
	GmpTimeline    []GmpTimeline `json:"gmpTimeline"`    // Description of the event (e.g., issue_open, listing_date)
	LatestGmpPrice float64       `json:"latestGmpPrice"` // Date when the event occurs
	IsGmpActivated bool          `json:"isGmpActivated"` // Title of the event
}
type GmpTimeline struct {
	Event string    `json:"event"` // Description of the event (e.g., issue_open, listing_date)
	Date  time.Time `json:"date"`  // Date when the event occurs
}

type Quota struct {
	RetailQuota QuotaApplication `json:"retailQuota"` // Retail quota for the IPO
	QibQuota    QuotaApplication `json:"qibQuota"`    // Qualified institutional buyer (QIB) quota for the IPO
	SHNIQuota   QuotaApplication `json:"shniQuota"`   // Small high net-worth individual (SHNI) quota for the IPO
	BHNIQuota   QuotaApplication `json:"bHNIQuota"`   // Big high net-worth individual (BHNI) quota for the IPO
}

type QuotaApplication struct {
	OfferedShares            int `json:"offeredShares"`            // Number of shares offered in the IPO
	AppliedShares            int `json:"appliedShares"`            // Number of shares applied for in the IPO
	MinLotSizePerApplication int `json:"minLotSizePerApplication"` // Minimum lot size per application
	MaxLotSizePerApplication int `json:"maxLotSizePerApplication"` // Maximum lot size per application
}
