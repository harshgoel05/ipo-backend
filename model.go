package main

import "time"

type DMIPO struct {
	LogoUrl     string         `json:"logoUrl"`     // URL to the logo image of the company
	Link        string         `json:"link"`        // Link to more information about the IPO
	Symbol      string         `json:"symbol"`      // Stock symbol
	Name        string         `json:"name"`        // Name of the company issuing the IPO
	StartDate   time.Time      `json:"startDate"`   // Date when the IPO starts
	EndDate     time.Time      `json:"endDate"`     // Date when the IPO ends
	ListingDate time.Time      `json:"listingDate"` // Date when the IPO will be listed on the stock exchange
	PriceRange  DMPriceRange   `json:"priceRange"`  // Price range for the IPO shares
	Slug        string         `json:"slug"`        // URL-friendly identifier
	GmpUrl      string         `json:"gmpUrl"`      // URL for grey market premium details
	Details     DMIPODetail    `json:"details"`     // Detailed information about the IPO
	GmpTimeline []DMGmpDetails `json:"gmpTimeline"` // Grey market premium timeline
}

type DMIPODetail struct {
	IssueSize     string    `json:"issueSize"`      // Total size of the IPO issue
	SizePerLot    int       `json:"sizePerLot"`     // Number of shares per lot
	Schedule      []DMEvent `json:"schedule"`       // List of events and their respective dates
	About         string    `json:"about"`          // Description of the company
	MinInvestment int       `json:"min_investment"` // Minimum investment required
	Strengths     []string  `json:"strengths"`      // List of company strengths
	Risks         []string  `json:"risks"`          // List of company risks
}

type DMPriceRange struct {
	Min float64 `json:"min"` // Minimum price in the range
	Max float64 `json:"max"` // Maximum price in the range
}

type DMEvent struct {
	Event      string    `json:"event"`      // Description of the event (e.g., issue_open, listing_date)
	Date       time.Time `json:"date"`       // Date when the event occurs
	EventTitle string    `json:"eventTitle"` // Title of the event
}

type DMGmpDetails struct {
	Date  time.Time `json:"date"`  // Date of the GMP record
	Price float64   `json:"price"` // Grey market premium price
}
