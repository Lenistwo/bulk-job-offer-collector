package main

type JobOffer struct {
	JobTitle           string         `json:"jobTitle"`
	Employer           string         `json:"employer"`
	CompanyProfileUrl  string         `json:"companyProfileUrl"`
	JobiconCompanyInfo JobiconCompany `json:"jobiconCompanyInfo"`
	CompanyId          int            `json:"companyId"`
	Logo               string         `json:"logo"`
	LastPublicated     string         `json:"lastPublicated"`
	ExpirationDate     string         `json:"expirationDate"`
	Salary             string         `json:"salary"`
	EmploymentLevel    string         `json:"employmentLevel"`
	JobDescription     string         `json:"jobDescription"`
	OfferType          []string       `json:"offerType"`
	OptionalCv         bool           `json:"optionalCv"`
	CountryName        string         `json:"countryName"`
	MainCategoriesIds  []int          `json:"mainCategoriesIds"`
	Offers             []Offer        `json:"offers"`
	TypesOfContract    []string       `json:"typesOfContract"`
	WorkSchedules      []string       `json:"workSchedules"`
	RemoteWork         bool           `json:"remoteWork"`
	OneClickApply      bool           `json:"oneClickApply"`
	UniqueOfferId      string         `json:"uniqueOfferId"`
}

type JobiconCompany struct {
	IsJobiconCompany bool `json:"isJobiconCompany"`
}

type Offer struct {
	OfferId    int      `json:"offerId"`
	OfferUrl   string   `json:"offerUrl"`
	RegionName string   `json:"regionName"`
	Cities     []string `json:"cities"`
	Label      string   `json:"label"`
}

type Pagination struct {
	PreviousPageLinkVisible bool   `json:"previousPageLinkVisible"`
	PreviousPageUrl         string `json:"previousPageUrl"`
	Pages                   []Page `json:"pages"`
	CurrentPageNumber       int    `json:"currentPageNumber"`
	MaxPages                int    `json:"maxPages"`
	NextPageLinkVisible     bool   `json:"nextPageLinkVisible"`
	NextPageUrl             string `json:"nextPageUrl"`
	ShowPagination          bool   `json:"showPagination"`
}
type Page struct {
	PageNumber     int    `json:"pageNumber"`
	IsCurrent      bool   `json:"isCurrent"`
	IsDotSeparator bool   `json:"isDotSeparator"`
	PageUrl        string `json:"pageUrl,omitempty"`
}

type Output struct {
	Title           string   `json:"title"`
	EmploymentLevel string   `json:"employment_level"`
	Employer        string   `json:"employer"`
	Description     string   `json:"description"`
	Salary          string   `json:"salary"`
	TypesOfContract []string `json:"types_of_contract"`
	ApplicationURL  string   `json:"application_url"`
	ExpiresAt       string   `json:"expires_at"`
}
