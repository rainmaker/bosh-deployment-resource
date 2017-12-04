// Package partners provides access to the Google Partners API.
//
// See https://developers.google.com/partners/
//
// Usage example:
//
//   import "google.golang.org/api/partners/v2"
//   ...
//   partnersService, err := partners.New(oauthHttpClient)
package partners // import "google.golang.org/api/partners/v2"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "partners:v2"
const apiName = "partners"
const apiVersion = "v2"
const basePath = "https://partners.googleapis.com/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Analytics = NewAnalyticsService(s)
	s.ClientMessages = NewClientMessagesService(s)
	s.Companies = NewCompaniesService(s)
	s.Exams = NewExamsService(s)
	s.Leads = NewLeadsService(s)
	s.Offers = NewOffersService(s)
	s.UserEvents = NewUserEventsService(s)
	s.UserStates = NewUserStatesService(s)
	s.Users = NewUsersService(s)
	s.V2 = NewV2Service(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Analytics *AnalyticsService

	ClientMessages *ClientMessagesService

	Companies *CompaniesService

	Exams *ExamsService

	Leads *LeadsService

	Offers *OffersService

	UserEvents *UserEventsService

	UserStates *UserStatesService

	Users *UsersService

	V2 *V2Service
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAnalyticsService(s *Service) *AnalyticsService {
	rs := &AnalyticsService{s: s}
	return rs
}

type AnalyticsService struct {
	s *Service
}

func NewClientMessagesService(s *Service) *ClientMessagesService {
	rs := &ClientMessagesService{s: s}
	return rs
}

type ClientMessagesService struct {
	s *Service
}

func NewCompaniesService(s *Service) *CompaniesService {
	rs := &CompaniesService{s: s}
	rs.Leads = NewCompaniesLeadsService(s)
	return rs
}

type CompaniesService struct {
	s *Service

	Leads *CompaniesLeadsService
}

func NewCompaniesLeadsService(s *Service) *CompaniesLeadsService {
	rs := &CompaniesLeadsService{s: s}
	return rs
}

type CompaniesLeadsService struct {
	s *Service
}

func NewExamsService(s *Service) *ExamsService {
	rs := &ExamsService{s: s}
	return rs
}

type ExamsService struct {
	s *Service
}

func NewLeadsService(s *Service) *LeadsService {
	rs := &LeadsService{s: s}
	return rs
}

type LeadsService struct {
	s *Service
}

func NewOffersService(s *Service) *OffersService {
	rs := &OffersService{s: s}
	rs.History = NewOffersHistoryService(s)
	return rs
}

type OffersService struct {
	s *Service

	History *OffersHistoryService
}

func NewOffersHistoryService(s *Service) *OffersHistoryService {
	rs := &OffersHistoryService{s: s}
	return rs
}

type OffersHistoryService struct {
	s *Service
}

func NewUserEventsService(s *Service) *UserEventsService {
	rs := &UserEventsService{s: s}
	return rs
}

type UserEventsService struct {
	s *Service
}

func NewUserStatesService(s *Service) *UserStatesService {
	rs := &UserStatesService{s: s}
	return rs
}

type UserStatesService struct {
	s *Service
}

func NewUsersService(s *Service) *UsersService {
	rs := &UsersService{s: s}
	return rs
}

type UsersService struct {
	s *Service
}

func NewV2Service(s *Service) *V2Service {
	rs := &V2Service{s: s}
	return rs
}

type V2Service struct {
	s *Service
}

// AdWordsManagerAccountInfo: Information about a particular AdWords
// Manager Account.
// Read more at https://support.google.com/adwords/answer/6139186
type AdWordsManagerAccountInfo struct {
	// CustomerName: Name of the customer this account represents.
	CustomerName string `json:"customerName,omitempty"`

	// Id: The AdWords Manager Account id.
	Id int64 `json:"id,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "CustomerName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CustomerName") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AdWordsManagerAccountInfo) MarshalJSON() ([]byte, error) {
	type noMethod AdWordsManagerAccountInfo
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Analytics: Analytics data for a `Company` within a single day.
type Analytics struct {
	// Contacts: Instances of users contacting the `Company`
	// on the specified date.
	Contacts *AnalyticsDataPoint `json:"contacts,omitempty"`

	// EventDate: Date on which these events occurred.
	EventDate *Date `json:"eventDate,omitempty"`

	// ProfileViews: Instances of users viewing the `Company` profile
	// on the specified date.
	ProfileViews *AnalyticsDataPoint `json:"profileViews,omitempty"`

	// SearchViews: Instances of users seeing the `Company` in Google
	// Partners Search results
	// on the specified date.
	SearchViews *AnalyticsDataPoint `json:"searchViews,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Contacts") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Contacts") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Analytics) MarshalJSON() ([]byte, error) {
	type noMethod Analytics
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AnalyticsDataPoint: Details of the analytics events for a `Company`
// within a single day.
type AnalyticsDataPoint struct {
	// EventCount: Number of times the type of event occurred.
	// Meaning depends on context (e.g. profile views, contacts, etc.).
	EventCount int64 `json:"eventCount,omitempty"`

	// EventLocations: Location information of where these events occurred.
	EventLocations []*LatLng `json:"eventLocations,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EventCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EventCount") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AnalyticsDataPoint) MarshalJSON() ([]byte, error) {
	type noMethod AnalyticsDataPoint
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AnalyticsSummary: Analytics aggregated data for a `Company` for a
// given date range.
type AnalyticsSummary struct {
	// ContactsCount: Aggregated number of times users contacted the
	// `Company`
	// for given date range.
	ContactsCount int64 `json:"contactsCount,omitempty"`

	// ProfileViewsCount: Aggregated number of profile views for the
	// `Company` for given date range.
	ProfileViewsCount int64 `json:"profileViewsCount,omitempty"`

	// SearchViewsCount: Aggregated number of times users saw the
	// `Company`
	// in Google Partners Search results for given date range.
	SearchViewsCount int64 `json:"searchViewsCount,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ContactsCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ContactsCount") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AnalyticsSummary) MarshalJSON() ([]byte, error) {
	type noMethod AnalyticsSummary
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AvailableOffer: Available Offers to be distributed.
type AvailableOffer struct {
	// Available: The number of codes for this offer that are available for
	// distribution.
	Available int64 `json:"available,omitempty"`

	// CountryOfferInfos: Offer info by country.
	CountryOfferInfos []*CountryOfferInfo `json:"countryOfferInfos,omitempty"`

	// Description: Description of the offer.
	Description string `json:"description,omitempty"`

	// Id: ID of this offer.
	Id int64 `json:"id,omitempty,string"`

	// MaxAccountAge: The maximum age of an account [in days] to be
	// eligible.
	MaxAccountAge int64 `json:"maxAccountAge,omitempty"`

	// Name: Name of the offer.
	Name string `json:"name,omitempty"`

	// OfferLevel: Level of this offer.
	//
	// Possible values:
	//   "OFFER_LEVEL_UNSPECIFIED" - Unset.
	//   "OFFER_LEVEL_DENY_PROBLEM" - Users/Agencies that have no offers
	// because of a problem.
	//   "OFFER_LEVEL_DENY_CONTRACT" - Users/Agencies that have no offers
	// due to contractural agreements.
	//   "OFFER_LEVEL_MANUAL" - Users/Agencies that have a
	// manually-configured limit.
	//   "OFFER_LEVEL_LIMIT_0" - Some Agencies don't get any offers.
	//   "OFFER_LEVEL_LIMIT_5" - Basic level gets 5 per month.
	//   "OFFER_LEVEL_LIMIT_15" - Agencies with adequate AHI and spend get
	// 15/month.
	//   "OFFER_LEVEL_LIMIT_50" - Badged partners (even in grace) get 50 per
	// month.
	OfferLevel string `json:"offerLevel,omitempty"`

	// OfferType: Type of offer.
	//
	// Possible values:
	//   "OFFER_TYPE_UNSPECIFIED" - Unset.
	//   "OFFER_TYPE_SPEND_X_GET_Y" - AdWords spend X get Y.
	//   "OFFER_TYPE_VIDEO" - Youtube video.
	//   "OFFER_TYPE_SPEND_MATCH" - Spend Match up to Y.
	OfferType string `json:"offerType,omitempty"`

	// QualifiedCustomer: Customers who qualify for this offer.
	QualifiedCustomer []*OfferCustomer `json:"qualifiedCustomer,omitempty"`

	// QualifiedCustomersComplete: Whether or not the list of qualified
	// customers is definitely complete.
	QualifiedCustomersComplete bool `json:"qualifiedCustomersComplete,omitempty"`

	// ShowSpecialOfferCopy: Should special text be shown on the offers
	// page.
	ShowSpecialOfferCopy bool `json:"showSpecialOfferCopy,omitempty"`

	// Terms: Terms of the offer.
	Terms string `json:"terms,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Available") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Available") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AvailableOffer) MarshalJSON() ([]byte, error) {
	type noMethod AvailableOffer
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Certification: A user's information on a specific certification.
type Certification struct {
	// Achieved: Whether this certification has been achieved.
	Achieved bool `json:"achieved,omitempty"`

	// CertificationType: The type of certification, the area of expertise.
	//
	// Possible values:
	//   "CERTIFICATION_TYPE_UNSPECIFIED" - Unchosen.
	//   "CT_ADWORDS" - AdWords certified.
	//   "CT_YOUTUBE" - YouTube certified.
	//   "CT_VIDEOADS" - VideoAds certified.
	//   "CT_ANALYTICS" - Analytics certified.
	//   "CT_DOUBLECLICK" - DoubleClick certified.
	//   "CT_SHOPPING" - Shopping certified.
	//   "CT_MOBILE" - Mobile certified.
	//   "CT_DIGITAL_SALES" - Digital sales certified.
	//   "CT_ADWORDS_SEARCH" - AdWords Search certified.
	//   "CT_ADWORDS_DISPLAY" - AdWords Display certified.
	//   "CT_MOBILE_SITES" - Mobile Sites certified.
	CertificationType string `json:"certificationType,omitempty"`

	// Expiration: Date this certification is due to expire.
	Expiration string `json:"expiration,omitempty"`

	// LastAchieved: The date the user last achieved certification.
	LastAchieved string `json:"lastAchieved,omitempty"`

	// Warning: Whether this certification is in the state of warning.
	Warning bool `json:"warning,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Achieved") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Achieved") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Certification) MarshalJSON() ([]byte, error) {
	type noMethod Certification
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CertificationExamStatus: Status for a Google Partners certification
// exam.
type CertificationExamStatus struct {
	// NumberUsersPass: The number of people who have passed the
	// certification exam.
	NumberUsersPass int64 `json:"numberUsersPass,omitempty"`

	// Type: The type of certification exam.
	//
	// Possible values:
	//   "CERTIFICATION_EXAM_TYPE_UNSPECIFIED" - Unchosen.
	//   "CET_ADWORDS_FUNDAMENTALS" - Adwords Fundamentals exam.
	//   "CET_ADWORDS_ADVANCED_SEARCH" - AdWords advanced search exam.
	//   "CET_ADWORDS_ADVANCED_DISPLAY" - AdWords advanced display exam.
	//   "CET_VIDEO_ADS" - VideoAds exam.
	//   "CET_DOUBLECLICK" - DoubleClick exam.
	//   "CET_ANALYTICS" - Analytics exam.
	//   "CET_SHOPPING" - Shopping exam.
	//   "CET_MOBILE" - Mobile exam.
	//   "CET_DIGITAL_SALES" - Digital Sales exam.
	//   "CET_MOBILE_SITES" - Mobile Sites exam.
	Type string `json:"type,omitempty"`

	// ForceSendFields is a list of field names (e.g. "NumberUsersPass") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NumberUsersPass") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *CertificationExamStatus) MarshalJSON() ([]byte, error) {
	type noMethod CertificationExamStatus
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CertificationStatus: Google Partners certification status.
type CertificationStatus struct {
	// ExamStatuses: List of certification exam statuses.
	ExamStatuses []*CertificationExamStatus `json:"examStatuses,omitempty"`

	// IsCertified: Whether certification is passing.
	IsCertified bool `json:"isCertified,omitempty"`

	// Type: The type of the certification.
	//
	// Possible values:
	//   "CERTIFICATION_TYPE_UNSPECIFIED" - Unchosen.
	//   "CT_ADWORDS" - AdWords certified.
	//   "CT_YOUTUBE" - YouTube certified.
	//   "CT_VIDEOADS" - VideoAds certified.
	//   "CT_ANALYTICS" - Analytics certified.
	//   "CT_DOUBLECLICK" - DoubleClick certified.
	//   "CT_SHOPPING" - Shopping certified.
	//   "CT_MOBILE" - Mobile certified.
	//   "CT_DIGITAL_SALES" - Digital sales certified.
	//   "CT_ADWORDS_SEARCH" - AdWords Search certified.
	//   "CT_ADWORDS_DISPLAY" - AdWords Display certified.
	//   "CT_MOBILE_SITES" - Mobile Sites certified.
	Type string `json:"type,omitempty"`

	// UserCount: Number of people who are certified,
	UserCount int64 `json:"userCount,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExamStatuses") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExamStatuses") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CertificationStatus) MarshalJSON() ([]byte, error) {
	type noMethod CertificationStatus
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Company: A company resource in the Google Partners API. Once
// certified, it qualifies
// for being searched by advertisers.
type Company struct {
	// AdditionalWebsites: URL of the company's additional websites used to
	// verify the dynamic badges.
	// These are stored as full URLs as entered by the user, but only the
	// TLD will
	// be used for the actual verification.
	AdditionalWebsites []string `json:"additionalWebsites,omitempty"`

	// AutoApprovalEmailDomains: Email domains that allow users with a
	// matching email address to get
	// auto-approved for associating with this company.
	AutoApprovalEmailDomains []string `json:"autoApprovalEmailDomains,omitempty"`

	// BadgeTier: Partner badge tier
	//
	// Possible values:
	//   "BADGE_TIER_NONE" - Tier badge is not set.
	//   "BADGE_TIER_REGULAR" - Agency has regular partner badge.
	//   "BADGE_TIER_PREMIER" - Agency has premier badge.
	BadgeTier string `json:"badgeTier,omitempty"`

	// CertificationStatuses: The list of Google Partners certification
	// statuses for the company.
	CertificationStatuses []*CertificationStatus `json:"certificationStatuses,omitempty"`

	// CompanyTypes: Company type labels listed on the company's profile.
	//
	// Possible values:
	//   "COMPANY_TYPE_UNSPECIFIED" - Unchosen.
	//   "FULL_SERVICE_AGENCY" - Handles all aspects of the advertising
	// process.
	//   "MEDIA_AGENCY" - Focuses solely on an advertiser's media placement.
	//   "CREATIVE_AGENCY" - Plans/executes advertising campaigns.
	//   "CDIGITAL_AGENCY" - Like a
	// FULL_SERVICE_AGENCY,
	// but specializing in digital.
	//   "SEM_SEO" - Increases visibility in search engine result pages.
	//   "PERFORMANCE_MARKETING" - Drives promotional efforts for immediate
	// impact.
	//   "ADVERTISING_TOOL_DEVELOPMENT" - Focuses on bid management,
	// conversion, reporting.
	//   "PR" - Establishes favorable relationship with public through
	// low/no-cost
	// communications.
	//   "SELF_MANAGED" - Does not manage other company's accounts, manages
	// own marketing programs.
	//   "RESELLER" - Full-service AdWords account management for local
	// businesses.
	CompanyTypes []string `json:"companyTypes,omitempty"`

	// ConvertedMinMonthlyBudget: The minimum monthly budget that the
	// company accepts for partner business,
	// converted to the requested currency code.
	ConvertedMinMonthlyBudget *Money `json:"convertedMinMonthlyBudget,omitempty"`

	// Id: The ID of the company.
	Id string `json:"id,omitempty"`

	// Industries: Industries the company can help with.
	//
	// Possible values:
	//   "INDUSTRY_UNSPECIFIED" - Unchosen.
	//   "I_AUTOMOTIVE" - The automotive industry.
	//   "I_BUSINESS_TO_BUSINESS" - The business-to-business industry.
	//   "I_CONSUMER_PACKAGED_GOODS" - The consumer packaged goods industry.
	//   "I_EDUCATION" - The education industry.
	//   "I_FINANCE" - The finance industry.
	//   "I_HEALTHCARE" - The healthcare industry.
	//   "I_MEDIA_AND_ENTERTAINMENT" - The media and entertainment industry.
	//   "I_RETAIL" - The retail industry.
	//   "I_TECHNOLOGY" - The technology industry.
	//   "I_TRAVEL" - The travel industry.
	Industries []string `json:"industries,omitempty"`

	// LocalizedInfos: The list of localized info for the company.
	LocalizedInfos []*LocalizedCompanyInfo `json:"localizedInfos,omitempty"`

	// Locations: The list of all company locations.
	// If set, must include the
	// primary_location
	// in the list.
	Locations []*Location `json:"locations,omitempty"`

	// Name: The name of the company.
	Name string `json:"name,omitempty"`

	// OriginalMinMonthlyBudget: The unconverted minimum monthly budget that
	// the company accepts for partner
	// business.
	OriginalMinMonthlyBudget *Money `json:"originalMinMonthlyBudget,omitempty"`

	// PrimaryAdwordsManagerAccountId: The Primary AdWords Manager Account
	// id.
	PrimaryAdwordsManagerAccountId int64 `json:"primaryAdwordsManagerAccountId,omitempty,string"`

	// PrimaryLanguageCode: The primary language code of the company, as
	// defined by
	// <a href="https://tools.ietf.org/html/bcp47">BCP 47</a>
	// (IETF BCP 47, "Tags for Identifying Languages").
	PrimaryLanguageCode string `json:"primaryLanguageCode,omitempty"`

	// PrimaryLocation: The primary location of the company.
	PrimaryLocation *Location `json:"primaryLocation,omitempty"`

	// ProfileStatus: The public viewability status of the company's
	// profile.
	//
	// Possible values:
	//   "COMPANY_PROFILE_STATUS_UNSPECIFIED" - Unchosen.
	//   "HIDDEN" - Company profile does not show up publicly.
	//   "PUBLISHED" - Company profile can only be viewed by the profile's
	// URL
	// and not by Google Partner Search.
	//   "SEARCHABLE" - Company profile can be viewed by the profile's
	// URL
	// and by Google Partner Search.
	ProfileStatus string `json:"profileStatus,omitempty"`

	// PublicProfile: Basic information from the company's public profile.
	PublicProfile *PublicProfile `json:"publicProfile,omitempty"`

	// Ranks: Information related to the ranking of the company within the
	// list of
	// companies.
	Ranks []*Rank `json:"ranks,omitempty"`

	// Services: Services the company can help with.
	//
	// Possible values:
	//   "SERVICE_UNSPECIFIED" - Unchosen.
	//   "S_ADVANCED_ADWORDS_SUPPORT" - Help with advanced AdWords support.
	//   "S_ADVERTISING_ON_GOOGLE" - Help with advertising on Google.
	//   "S_AN_ENHANCED_WEBSITE" - Help with an enhanced website.
	//   "S_AN_ONLINE_MARKETING_PLAN" - Help with an online marketing plan.
	//   "S_MOBILE_AND_VIDEO_ADS" - Help with mobile and video ads.
	//   "S_MOBILE_WEBSITE_SERVICES" - Help with mobile websites.
	Services []string `json:"services,omitempty"`

	// SpecializationStatus: The list of Google Partners specialization
	// statuses for the company.
	SpecializationStatus []*SpecializationStatus `json:"specializationStatus,omitempty"`

	// WebsiteUrl: URL of the company's website.
	WebsiteUrl string `json:"websiteUrl,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AdditionalWebsites")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AdditionalWebsites") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Company) MarshalJSON() ([]byte, error) {
	type noMethod Company
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CompanyRelation: A CompanyRelation resource representing information
// about a user's
// affiliation and standing with a company in Partners.
type CompanyRelation struct {
	// Address: The primary address for this company.
	Address string `json:"address,omitempty"`

	// BadgeTier: Whether the company is a Partner.
	//
	// Possible values:
	//   "BADGE_TIER_NONE" - Tier badge is not set.
	//   "BADGE_TIER_REGULAR" - Agency has regular partner badge.
	//   "BADGE_TIER_PREMIER" - Agency has premier badge.
	BadgeTier string `json:"badgeTier,omitempty"`

	// CompanyAdmin: Indicates if the user is an admin for this company.
	CompanyAdmin bool `json:"companyAdmin,omitempty"`

	// CompanyId: The ID of the company. There may be no id if this is
	// a
	// pending company.5
	CompanyId string `json:"companyId,omitempty"`

	// CreationTime: The timestamp of when affiliation was
	// requested.
	// @OutputOnly
	CreationTime string `json:"creationTime,omitempty"`

	// InternalCompanyId: The internal company ID.
	// Only available for a whitelisted set of api clients.
	InternalCompanyId string `json:"internalCompanyId,omitempty"`

	// IsPending: The flag that indicates if the company is pending
	// verification.
	IsPending bool `json:"isPending,omitempty"`

	// LogoUrl: A URL to a profile photo, e.g. a G+ profile photo.
	LogoUrl string `json:"logoUrl,omitempty"`

	// ManagerAccount: The AdWords manager account # associated this
	// company.
	ManagerAccount int64 `json:"managerAccount,omitempty,string"`

	// Name: The name (in the company's primary language) for the company.
	Name string `json:"name,omitempty"`

	// PhoneNumber: The phone number for the company's primary address.
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// PrimaryAddress: The primary location of the company.
	PrimaryAddress *Location `json:"primaryAddress,omitempty"`

	// PrimaryCountryCode: The primary country code of the company.
	PrimaryCountryCode string `json:"primaryCountryCode,omitempty"`

	// PrimaryLanguageCode: The primary language code of the company.
	PrimaryLanguageCode string `json:"primaryLanguageCode,omitempty"`

	// ResolvedTimestamp: The timestamp when the user was
	// approved.
	// @OutputOnly
	ResolvedTimestamp string `json:"resolvedTimestamp,omitempty"`

	// Segment: The segment the company is classified as.
	//
	// Possible values:
	//   "COMPANY_SEGMENT_UNKNOWN" - Default segment indicates an unknown.
	//   "COMPANY_SEGMENT_NAL" - Segment representing a selected group of
	// Partners
	//   "COMPANY_SEGMENT_PSP" - Segment representing Premier SMB Partners,
	// an AdWords partnership program.
	//   "COMPANY_SEGMENT_PPSP" - A segment of Premier SMB Partners that
	// have relationship with Google.
	Segment []string `json:"segment,omitempty"`

	// SpecializationStatus: The list of Google Partners specialization
	// statuses for the company.
	SpecializationStatus []*SpecializationStatus `json:"specializationStatus,omitempty"`

	// State: The state of relationship, in terms of approvals.
	//
	// Possible values:
	//   "USER_COMPANY_REATION_STATE_NONE_SPECIFIED" - Default unspecified
	// value.
	//   "USER_COMPANY_RELATION_STATE_AWAIT_EMAIL" - User has filled in a
	// request to be associated with an company.
	// Now waiting email confirmation.
	//   "USER_COMPANY_RELATION_STATE_AWAIT_ADMIN" - Pending approval from
	// company.
	// Email confirmation will not approve this one.
	//   "USER_COMPANY_RELATION_STATE_APPROVED" - Approved by company.
	State string `json:"state,omitempty"`

	// Website: The website URL for this company.
	Website string `json:"website,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Address") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Address") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CompanyRelation) MarshalJSON() ([]byte, error) {
	type noMethod CompanyRelation
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CountryOfferInfo: Offer info by country.
type CountryOfferInfo struct {
	// GetYAmount: (localized) Get Y amount for that country's offer.
	GetYAmount string `json:"getYAmount,omitempty"`

	// OfferCountryCode: Country code for which offer codes may be
	// requested.
	OfferCountryCode string `json:"offerCountryCode,omitempty"`

	// OfferType: Type of offer country is eligible for.
	//
	// Possible values:
	//   "OFFER_TYPE_UNSPECIFIED" - Unset.
	//   "OFFER_TYPE_SPEND_X_GET_Y" - AdWords spend X get Y.
	//   "OFFER_TYPE_VIDEO" - Youtube video.
	//   "OFFER_TYPE_SPEND_MATCH" - Spend Match up to Y.
	OfferType string `json:"offerType,omitempty"`

	// SpendXAmount: (localized) Spend X amount for that country's offer.
	SpendXAmount string `json:"spendXAmount,omitempty"`

	// ForceSendFields is a list of field names (e.g. "GetYAmount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "GetYAmount") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CountryOfferInfo) MarshalJSON() ([]byte, error) {
	type noMethod CountryOfferInfo
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreateLeadRequest: Request message for CreateLead.
type CreateLeadRequest struct {
	// Lead: The lead resource. The `LeadType` must not be
	// `LEAD_TYPE_UNSPECIFIED`
	// and either `email` or `phone_number` must be provided.
	Lead *Lead `json:"lead,omitempty"`

	// RecaptchaChallenge: <a
	// href="https://www.google.com/recaptcha/">reCaptcha</a> challenge
	// info.
	RecaptchaChallenge *RecaptchaChallenge `json:"recaptchaChallenge,omitempty"`

	// RequestMetadata: Current request metadata.
	RequestMetadata *RequestMetadata `json:"requestMetadata,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Lead") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Lead") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreateLeadRequest) MarshalJSON() ([]byte, error) {
	type noMethod CreateLeadRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// CreateLeadResponse: Response message for CreateLead.
type CreateLeadResponse struct {
	// Lead: Lead that was created depending on the outcome of
	// <a href="https://www.google.com/recaptcha/">reCaptcha</a> validation.
	Lead *Lead `json:"lead,omitempty"`

	// RecaptchaStatus: The outcome of <a
	// href="https://www.google.com/recaptcha/">reCaptcha</a>
	// validation.
	//
	// Possible values:
	//   "RECAPTCHA_STATUS_UNSPECIFIED" - Unchosen.
	//   "RS_NOT_NEEDED" - No reCaptcha validation needed.
	//   "RS_PASSED" - reCaptcha challenge passed.
	//   "RS_FAILED" - reCaptcha challenge failed.
	RecaptchaStatus string `json:"recaptchaStatus,omitempty"`

	// ResponseMetadata: Current response metadata.
	ResponseMetadata *ResponseMetadata `json:"responseMetadata,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Lead") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Lead") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *CreateLeadResponse) MarshalJSON() ([]byte, error) {
	type noMethod CreateLeadResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Date: Represents a whole calendar date, e.g. date of birth. The time
// of day and
// time zone are either specified elsewhere or are not significant. The
// date
// is relative to the Proleptic Gregorian Calendar. The day may be 0
// to
// represent a year and month where the day is not significant, e.g.
// credit card
// expiration date. The year may be 0 to represent a month and day
// independent
// of year, e.g. anniversary date. Related types are
// google.type.TimeOfDay
// and `google.protobuf.Timestamp`.
type Date struct {
	// Day: Day of month. Must be from 1 to 31 and valid for the year and
	// month, or 0
	// if specifying a year/month where the day is not significant.
	Day int64 `json:"day,omitempty"`

	// Month: Month of year. Must be from 1 to 12.
	Month int64 `json:"month,omitempty"`

	// Year: Year of date. Must be from 1 to 9999, or 0 if specifying a date
	// without
	// a year.
	Year int64 `json:"year,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Day") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Day") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Date) MarshalJSON() ([]byte, error) {
	type noMethod Date
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// DebugInfo: Debug information about this request.
type DebugInfo struct {
	// ServerInfo: Info about the server that serviced this request.
	ServerInfo string `json:"serverInfo,omitempty"`

	// ServerTraceInfo: Server-side debug stack trace.
	ServerTraceInfo string `json:"serverTraceInfo,omitempty"`

	// ServiceUrl: URL of the service that handled this request.
	ServiceUrl string `json:"serviceUrl,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ServerInfo") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ServerInfo") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *DebugInfo) MarshalJSON() ([]byte, error) {
	type noMethod DebugInfo
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Empty: A generic empty message that you can re-use to avoid defining
// duplicated
// empty messages in your APIs. A typical example is to use it as the
// request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns
// (google.protobuf.Empty);
//     }
//
// The JSON representation for `Empty` is empty JSON object `{}`.
type Empty struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`
}

// EventData: Key value data pair for an event.
type EventData struct {
	// Key: Data type.
	//
	// Possible values:
	//   "EVENT_DATA_TYPE_UNSPECIFIED" - Unchosen.
	//   "ACTION" - Action data.
	//   "AGENCY_ID" - Agency ID data.
	//   "AGENCY_NAME" - Agency name data.
	//   "AGENCY_PHONE_NUMBER" - Agency phone number data.
	//   "AGENCY_WEBSITE" - Agency website data.
	//   "BUDGET" - Budget data.
	//   "CENTER_POINT" - Center-point data.
	//   "CERTIFICATION" - Certification data.
	//   "COMMENT" - Comment data.
	//   "COUNTRY" - Country data.
	//   "CURRENCY" - Currency data.
	//   "CURRENTLY_VIEWED_AGENCY_ID" - Currently viewed agency ID data.
	//   "DISTANCE" - Distance data.
	//   "DISTANCE_TYPE" - Distance type data.
	//   "EXAM" - Exam data.
	//   "HISTORY_TOKEN" - History token data.
	//   "ID" - Identifier data.
	//   "INDUSTRY" - Industry data.
	//   "INSIGHT_TAG" - Insight tag data.
	//   "LANGUAGE" - Language data.
	//   "LOCATION" - Location  data.
	//   "MARKETING_OPT_IN" - Marketing opt-in data.
	//   "QUERY" - Query data.
	//   "SEARCH_START_INDEX" - Search start index data.
	//   "SERVICE" - Service data.
	//   "SHOW_VOW" - Show vow data.
	//   "SOLUTION" - Solution data.
	//   "TRAFFIC_SOURCE_ID" - Traffic source ID data.
	//   "TRAFFIC_SUB_ID" - Traffic sub ID data.
	//   "VIEW_PORT" - Viewport data.
	//   "WEBSITE" - Website data.
	//   "DETAILS" - Details data.
	//   "EXPERIMENT_ID" - Experiment ID data.
	//   "GPS_MOTIVATION" - Google Partner Search motivation data.
	//   "URL" - URL data.
	//   "ELEMENT_FOCUS" - Element we wanted user to focus on.
	//   "PROGRESS" - Progress when viewing an item \[0-100\].
	Key string `json:"key,omitempty"`

	// Values: Data values.
	Values []string `json:"values,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Key") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Key") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *EventData) MarshalJSON() ([]byte, error) {
	type noMethod EventData
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ExamStatus: A user's information on a specific exam.
type ExamStatus struct {
	// ExamType: The type of the exam.
	//
	// Possible values:
	//   "CERTIFICATION_EXAM_TYPE_UNSPECIFIED" - Unchosen.
	//   "CET_ADWORDS_FUNDAMENTALS" - Adwords Fundamentals exam.
	//   "CET_ADWORDS_ADVANCED_SEARCH" - AdWords advanced search exam.
	//   "CET_ADWORDS_ADVANCED_DISPLAY" - AdWords advanced display exam.
	//   "CET_VIDEO_ADS" - VideoAds exam.
	//   "CET_DOUBLECLICK" - DoubleClick exam.
	//   "CET_ANALYTICS" - Analytics exam.
	//   "CET_SHOPPING" - Shopping exam.
	//   "CET_MOBILE" - Mobile exam.
	//   "CET_DIGITAL_SALES" - Digital Sales exam.
	//   "CET_MOBILE_SITES" - Mobile Sites exam.
	ExamType string `json:"examType,omitempty"`

	// Expiration: Date this exam is due to expire.
	Expiration string `json:"expiration,omitempty"`

	// LastPassed: The date the user last passed this exam.
	LastPassed string `json:"lastPassed,omitempty"`

	// Passed: Whether this exam has been passed and not expired.
	Passed bool `json:"passed,omitempty"`

	// Taken: The date the user last taken this exam.
	Taken string `json:"taken,omitempty"`

	// Warning: Whether this exam is in the state of warning.
	Warning bool `json:"warning,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExamType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExamType") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ExamStatus) MarshalJSON() ([]byte, error) {
	type noMethod ExamStatus
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ExamToken: A token that allows a user to take an exam.
type ExamToken struct {
	// ExamId: The id of the exam the token is for.
	ExamId int64 `json:"examId,omitempty,string"`

	// ExamType: The type of the exam the token belongs to.
	//
	// Possible values:
	//   "CERTIFICATION_EXAM_TYPE_UNSPECIFIED" - Unchosen.
	//   "CET_ADWORDS_FUNDAMENTALS" - Adwords Fundamentals exam.
	//   "CET_ADWORDS_ADVANCED_SEARCH" - AdWords advanced search exam.
	//   "CET_ADWORDS_ADVANCED_DISPLAY" - AdWords advanced display exam.
	//   "CET_VIDEO_ADS" - VideoAds exam.
	//   "CET_DOUBLECLICK" - DoubleClick exam.
	//   "CET_ANALYTICS" - Analytics exam.
	//   "CET_SHOPPING" - Shopping exam.
	//   "CET_MOBILE" - Mobile exam.
	//   "CET_DIGITAL_SALES" - Digital Sales exam.
	//   "CET_MOBILE_SITES" - Mobile Sites exam.
	ExamType string `json:"examType,omitempty"`

	// Token: The token, only present if the user has access to the exam.
	Token string `json:"token,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ExamId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExamId") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ExamToken) MarshalJSON() ([]byte, error) {
	type noMethod ExamToken
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GetCompanyResponse: Response message for GetCompany.
type GetCompanyResponse struct {
	// Company: The company.
	Company *Company `json:"company,omitempty"`

	// ResponseMetadata: Current response metadata.
	ResponseMetadata *ResponseMetadata `json:"responseMetadata,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Company") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Company") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GetCompanyResponse) MarshalJSON() ([]byte, error) {
	type noMethod GetCompanyResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GetPartnersStatusResponse: Response message for
// GetPartnersStatus.
type GetPartnersStatusResponse struct {
	// ResponseMetadata: Current response metadata.
	ResponseMetadata *ResponseMetadata `json:"responseMetadata,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ResponseMetadata") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResponseMetadata") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *GetPartnersStatusResponse) MarshalJSON() ([]byte, error) {
	type noMethod GetPartnersStatusResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// HistoricalOffer: Historical information about a Google Partners
// Offer.
type HistoricalOffer struct {
	// AdwordsUrl: Client's AdWords page URL.
	AdwordsUrl string `json:"adwordsUrl,omitempty"`

	// ClientEmail: Email address for client.
	ClientEmail string `json:"clientEmail,omitempty"`

	// ClientId: ID of client.
	ClientId int64 `json:"clientId,omitempty,string"`

	// ClientName: Name of the client.
	ClientName string `json:"clientName,omitempty"`

	// CreationTime: Time offer was first created.
	CreationTime string `json:"creationTime,omitempty"`

	// ExpirationTime: Time this offer expires.
	ExpirationTime string `json:"expirationTime,omitempty"`

	// LastModifiedTime: Time last action was taken.
	LastModifiedTime string `json:"lastModifiedTime,omitempty"`

	// OfferCode: Offer code.
	OfferCode string `json:"offerCode,omitempty"`

	// OfferCountryCode: Country Code for the offer country.
	OfferCountryCode string `json:"offerCountryCode,omitempty"`

	// OfferType: Type of offer.
	//
	// Possible values:
	//   "OFFER_TYPE_UNSPECIFIED" - Unset.
	//   "OFFER_TYPE_SPEND_X_GET_Y" - AdWords spend X get Y.
	//   "OFFER_TYPE_VIDEO" - Youtube video.
	//   "OFFER_TYPE_SPEND_MATCH" - Spend Match up to Y.
	OfferType string `json:"offerType,omitempty"`

	// SenderName: Name (First + Last) of the partners user to whom the
	// incentive is allocated.
	SenderName string `json:"senderName,omitempty"`

	// Status: Status of the offer.
	//
	// Possible values:
	//   "OFFER_STATUS_UNSPECIFIED" - Unset.
	//   "OFFER_STATUS_DISTRIBUTED" - Offer distributed.
	//   "OFFER_STATUS_REDEEMED" - Offer redeemed.
	//   "OFFER_STATUS_AWARDED" - Offer awarded.
	//   "OFFER_STATUS_EXPIRED" - Offer expired.
	Status string `json:"status,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AdwordsUrl") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AdwordsUrl") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *HistoricalOffer) MarshalJSON() ([]byte, error) {
	type noMethod HistoricalOffer
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LatLng: An object representing a latitude/longitude pair. This is
// expressed as a pair
// of doubles representing degrees latitude and degrees longitude.
// Unless
// specified otherwise, this must conform to the
// <a
// href="http://www.unoosa.org/pdf/icg/2012/template/WGS_84.pdf">WGS84
// st
// andard</a>. Values must be within normalized ranges.
//
// Example of normalization code in Python:
//
//     def NormalizeLongitude(longitude):
//       """Wraps decimal degrees longitude to [-180.0, 180.0]."""
//       q, r = divmod(longitude, 360.0)
//       if r > 180.0 or (r == 180.0 and q <= -1.0):
//         return r - 360.0
//       return r
//
//     def NormalizeLatLng(latitude, longitude):
//       """Wraps decimal degrees latitude and longitude to
//       [-90.0, 90.0] and [-180.0, 180.0], respectively."""
//       r = latitude % 360.0
//       if r <= 90.0:
//         return r, NormalizeLongitude(longitude)
//       elif r >= 270.0:
//         return r - 360, NormalizeLongitude(longitude)
//       else:
//         return 180 - r, NormalizeLongitude(longitude + 180.0)
//
//     assert 180.0 == NormalizeLongitude(180.0)
//     assert -180.0 == NormalizeLongitude(-180.0)
//     assert -179.0 == NormalizeLongitude(181.0)
//     assert (0.0, 0.0) == NormalizeLatLng(360.0, 0.0)
//     assert (0.0, 0.0) == NormalizeLatLng(-360.0, 0.0)
//     assert (85.0, 180.0) == NormalizeLatLng(95.0, 0.0)
//     assert (-85.0, -170.0) == NormalizeLatLng(-95.0, 10.0)
//     assert (90.0, 10.0) == NormalizeLatLng(90.0, 10.0)
//     assert (-90.0, -10.0) == NormalizeLatLng(-90.0, -10.0)
//     assert (0.0, -170.0) == NormalizeLatLng(-180.0, 10.0)
//     assert (0.0, -170.0) == NormalizeLatLng(180.0, 10.0)
//     assert (-90.0, 10.0) == NormalizeLatLng(270.0, 10.0)
//     assert (90.0, 10.0) == NormalizeLatLng(-270.0, 10.0)
type LatLng struct {
	// Latitude: The latitude in degrees. It must be in the range [-90.0,
	// +90.0].
	Latitude float64 `json:"latitude,omitempty"`

	// Longitude: The longitude in degrees. It must be in the range [-180.0,
	// +180.0].
	Longitude float64 `json:"longitude,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Latitude") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Latitude") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LatLng) MarshalJSON() ([]byte, error) {
	type noMethod LatLng
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *LatLng) UnmarshalJSON(data []byte) error {
	type noMethod LatLng
	var s1 struct {
		Latitude  gensupport.JSONFloat64 `json:"latitude"`
		Longitude gensupport.JSONFloat64 `json:"longitude"`
		*noMethod
	}
	s1.noMethod = (*noMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Latitude = float64(s1.Latitude)
	s.Longitude = float64(s1.Longitude)
	return nil
}

// Lead: A lead resource that represents an advertiser contact for a
// `Company`. These
// are usually generated via Google Partner Search (the advertiser
// portal).
type Lead struct {
	// AdwordsCustomerId: The AdWords Customer ID of the lead.
	AdwordsCustomerId int64 `json:"adwordsCustomerId,omitempty,string"`

	// Comments: Comments lead source gave.
	Comments string `json:"comments,omitempty"`

	// CreateTime: Timestamp of when this lead was created.
	CreateTime string `json:"createTime,omitempty"`

	// Email: Email address of lead source.
	Email string `json:"email,omitempty"`

	// FamilyName: Last name of lead source.
	FamilyName string `json:"familyName,omitempty"`

	// GivenName: First name of lead source.
	GivenName string `json:"givenName,omitempty"`

	// GpsMotivations: List of reasons for using Google Partner Search and
	// creating a lead.
	//
	// Possible values:
	//   "GPS_MOTIVATION_UNSPECIFIED" - Unchosen.
	//   "GPSM_HELP_WITH_ADVERTISING" - Advertiser needs help with their
	// advertising.
	//   "GPSM_HELP_WITH_WEBSITE" - Advertiser needs help with their
	// website.
	//   "GPSM_NO_WEBSITE" - Advertiser does not have a website.
	GpsMotivations []string `json:"gpsMotivations,omitempty"`

	// Id: ID of the lead.
	Id string `json:"id,omitempty"`

	// LanguageCode: Language code of the lead's language preference, as
	// defined by
	// <a href="https://tools.ietf.org/html/bcp47">BCP 47</a>
	// (IETF BCP 47, "Tags for Identifying Languages").
	LanguageCode string `json:"languageCode,omitempty"`

	// MarketingOptIn: Whether or not the lead signed up for marketing
	// emails
	MarketingOptIn bool `json:"marketingOptIn,omitempty"`

	// MinMonthlyBudget: The minimum monthly budget lead source is willing
	// to spend.
	MinMonthlyBudget *Money `json:"minMonthlyBudget,omitempty"`

	// PhoneNumber: Phone number of lead source.
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// State: The lead's state in relation to the company.
	//
	// Possible values:
	//   "LEAD_STATE_UNSPECIFIED" - Unchosen.
	//   "LEAD" - Lead not yet contacted.
	//   "CONTACTED" - Lead has been contacted.
	//   "CLIENT" - Lead has become a client.
	//   "OTHER" - Lead in a state not covered by other options.
	State string `json:"state,omitempty"`

	// Type: Type of lead.
	//
	// Possible values:
	//   "LEAD_TYPE_UNSPECIFIED" - Unchosen.
	//   "LT_GPS" - Google Partner Search.
	Type string `json:"type,omitempty"`

	// WebsiteUrl: Website URL of lead source.
	WebsiteUrl string `json:"websiteUrl,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AdwordsCustomerId")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AdwordsCustomerId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Lead) MarshalJSON() ([]byte, error) {
	type noMethod Lead
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListAnalyticsResponse: Response message for
// ListAnalytics.
type ListAnalyticsResponse struct {
	// Analytics: The list of analytics.
	// Sorted in ascending order of
	// Analytics.event_date.
	Analytics []*Analytics `json:"analytics,omitempty"`

	// AnalyticsSummary: Aggregated information across the
	// response's
	// analytics.
	AnalyticsSummary *AnalyticsSummary `json:"analyticsSummary,omitempty"`

	// NextPageToken: A token to retrieve next page of results.
	// Pass this value in the `ListAnalyticsRequest.page_token` field in
	// the
	// subsequent call to
	// ListAnalytics to retrieve the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ResponseMetadata: Current response metadata.
	ResponseMetadata *ResponseMetadata `json:"responseMetadata,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Analytics") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Analytics") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListAnalyticsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListAnalyticsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListCompaniesResponse: Response message for
// ListCompanies.
type ListCompaniesResponse struct {
	// Companies: The list of companies.
	Companies []*Company `json:"companies,omitempty"`

	// NextPageToken: A token to retrieve next page of results.
	// Pass this value in the `ListCompaniesRequest.page_token` field in
	// the
	// subsequent call to
	// ListCompanies to retrieve the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ResponseMetadata: Current response metadata.
	ResponseMetadata *ResponseMetadata `json:"responseMetadata,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Companies") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Companies") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListCompaniesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListCompaniesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListLeadsResponse: Response message for ListLeads.
type ListLeadsResponse struct {
	// Leads: The list of leads.
	Leads []*Lead `json:"leads,omitempty"`

	// NextPageToken: A token to retrieve next page of results.
	// Pass this value in the `ListLeadsRequest.page_token` field in
	// the
	// subsequent call to
	// ListLeads to retrieve the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ResponseMetadata: Current response metadata.
	ResponseMetadata *ResponseMetadata `json:"responseMetadata,omitempty"`

	// TotalSize: The total count of leads for the given company.
	TotalSize int64 `json:"totalSize,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Leads") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Leads") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListLeadsResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListLeadsResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListOffersHistoryResponse: Response for ListOfferHistory.
type ListOffersHistoryResponse struct {
	// CanShowEntireCompany: True if the user has the option to show entire
	// company history.
	CanShowEntireCompany bool `json:"canShowEntireCompany,omitempty"`

	// NextPageToken: Supply this token in a ListOffersHistoryRequest to
	// retrieve the next page.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// Offers: Historical offers meeting request.
	Offers []*HistoricalOffer `json:"offers,omitempty"`

	// ResponseMetadata: Current response metadata.
	ResponseMetadata *ResponseMetadata `json:"responseMetadata,omitempty"`

	// ShowingEntireCompany: True if this response is showing entire company
	// history.
	ShowingEntireCompany bool `json:"showingEntireCompany,omitempty"`

	// TotalResults: Number of results across all pages.
	TotalResults int64 `json:"totalResults,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "CanShowEntireCompany") to unconditionally include in API requests.
	// By default, fields with empty values are omitted from API requests.
	// However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CanShowEntireCompany") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListOffersHistoryResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListOffersHistoryResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListOffersResponse: Response for ListOffer.
type ListOffersResponse struct {
	// AvailableOffers: Available Offers to be distributed.
	AvailableOffers []*AvailableOffer `json:"availableOffers,omitempty"`

	// NoOfferReason: Reason why no Offers are available.
	//
	// Possible values:
	//   "NO_OFFER_REASON_UNSPECIFIED" - Unset.
	//   "NO_OFFER_REASON_NO_MCC" - Not an MCC.
	//   "NO_OFFER_REASON_LIMIT_REACHED" - Offer limit has been reached.
	//   "NO_OFFER_REASON_INELIGIBLE" - Ineligible for offers.
	NoOfferReason string `json:"noOfferReason,omitempty"`

	// ResponseMetadata: Current response metadata.
	ResponseMetadata *ResponseMetadata `json:"responseMetadata,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AvailableOffers") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AvailableOffers") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListOffersResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListOffersResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListUserStatesResponse: Response message for
// ListUserStates.
type ListUserStatesResponse struct {
	// ResponseMetadata: Current response metadata.
	ResponseMetadata *ResponseMetadata `json:"responseMetadata,omitempty"`

	// UserStates: User's states.
	//
	// Possible values:
	//   "USER_STATE_UNSPECIFIED" - Unchosen.
	//   "US_REQUIRES_RECAPTCHA_FOR_GPS_CONTACT" - User must pass <a
	// href="https://www.google.com/recaptcha/">reCaptcha</a> to
	// contact a Partner via Google Partner Search.
	UserStates []string `json:"userStates,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ResponseMetadata") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResponseMetadata") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *ListUserStatesResponse) MarshalJSON() ([]byte, error) {
	type noMethod ListUserStatesResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LocalizedCompanyInfo: The localized company information.
type LocalizedCompanyInfo struct {
	// CountryCodes: List of country codes for the localized company info.
	CountryCodes []string `json:"countryCodes,omitempty"`

	// DisplayName: Localized display name.
	DisplayName string `json:"displayName,omitempty"`

	// LanguageCode: Language code of the localized company info, as defined
	// by
	// <a href="https://tools.ietf.org/html/bcp47">BCP 47</a>
	// (IETF BCP 47, "Tags for Identifying Languages").
	LanguageCode string `json:"languageCode,omitempty"`

	// Overview: Localized brief description that the company uses to
	// advertise themselves.
	Overview string `json:"overview,omitempty"`

	// ForceSendFields is a list of field names (e.g. "CountryCodes") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CountryCodes") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LocalizedCompanyInfo) MarshalJSON() ([]byte, error) {
	type noMethod LocalizedCompanyInfo
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Location: A location with address and geographic coordinates. May
// optionally contain a
// detailed (multi-field) version of the address.
type Location struct {
	// Address: The single string version of the address.
	Address string `json:"address,omitempty"`

	// AddressLine: The following address lines represent the most specific
	// part of any
	// address.
	AddressLine []string `json:"addressLine,omitempty"`

	// AdministrativeArea: Top-level administrative subdivision of this
	// country.
	AdministrativeArea string `json:"administrativeArea,omitempty"`

	// DependentLocality: Dependent locality or sublocality. Used for UK
	// dependent localities, or
	// neighborhoods or boroughs in other locations.
	DependentLocality string `json:"dependentLocality,omitempty"`

	// LanguageCode: Language code of the address. Should be in BCP 47
	// format.
	LanguageCode string `json:"languageCode,omitempty"`

	// LatLng: The latitude and longitude of the location, in degrees.
	LatLng *LatLng `json:"latLng,omitempty"`

	// Locality: Generally refers to the city/town portion of an address.
	Locality string `json:"locality,omitempty"`

	// PostalCode: Values are frequently alphanumeric.
	PostalCode string `json:"postalCode,omitempty"`

	// RegionCode: CLDR (Common Locale Data Repository) region code .
	RegionCode string `json:"regionCode,omitempty"`

	// SortingCode: Use of this code is very country-specific, but will
	// refer to a secondary
	// classification code for sorting mail.
	SortingCode string `json:"sortingCode,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Address") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Address") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Location) MarshalJSON() ([]byte, error) {
	type noMethod Location
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LogMessageRequest: Request message for
// LogClientMessage.
type LogMessageRequest struct {
	// ClientInfo: Map of client info, such as URL, browser navigator,
	// browser platform, etc.
	ClientInfo map[string]string `json:"clientInfo,omitempty"`

	// Details: Details about the client message.
	Details string `json:"details,omitempty"`

	// Level: Message level of client message.
	//
	// Possible values:
	//   "MESSAGE_LEVEL_UNSPECIFIED" - Unchosen.
	//   "ML_FINE" - Message level for tracing information.
	//   "ML_INFO" - Message level for informational messages.
	//   "ML_WARNING" - Message level for potential problems.
	//   "ML_SEVERE" - Message level for serious failures.
	Level string `json:"level,omitempty"`

	// RequestMetadata: Current request metadata.
	RequestMetadata *RequestMetadata `json:"requestMetadata,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ClientInfo") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ClientInfo") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LogMessageRequest) MarshalJSON() ([]byte, error) {
	type noMethod LogMessageRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LogMessageResponse: Response message for
// LogClientMessage.
type LogMessageResponse struct {
	// ResponseMetadata: Current response metadata.
	ResponseMetadata *ResponseMetadata `json:"responseMetadata,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ResponseMetadata") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResponseMetadata") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *LogMessageResponse) MarshalJSON() ([]byte, error) {
	type noMethod LogMessageResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LogUserEventRequest: Request message for
// LogUserEvent.
type LogUserEventRequest struct {
	// EventAction: The action that occurred.
	//
	// Possible values:
	//   "EVENT_ACTION_UNSPECIFIED" - Unchosen.
	//   "SMB_CLICKED_FIND_A_PARTNER_BUTTON_BOTTOM" - Advertiser clicked
	// `Find a partner` bottom button.
	//   "SMB_CLICKED_FIND_A_PARTNER_BUTTON_TOP" - Advertiser clicked `Find
	// a partner` top button.
	//   "AGENCY_CLICKED_JOIN_NOW_BUTTON_BOTTOM" - Agency clicked `Join now`
	// bottom button.
	//   "AGENCY_CLICKED_JOIN_NOW_BUTTON_TOP" - Agency clicked `Join now`
	// top button.
	//   "SMB_CANCELED_PARTNER_CONTACT_FORM" - Advertiser canceled partner
	// contact form.
	//   "SMB_CLICKED_CONTACT_A_PARTNER" - Advertiser started partner
	// contact form.
	//   "SMB_COMPLETED_PARTNER_CONTACT_FORM" - Advertiser completed partner
	// contact form.
	//   "SMB_ENTERED_EMAIL_IN_CONTACT_PARTNER_FORM" - Advertiser entered
	// email in contact form.
	//   "SMB_ENTERED_NAME_IN_CONTACT_PARTNER_FORM" - Advertiser entered
	// name in contact form.
	//   "SMB_ENTERED_PHONE_IN_CONTACT_PARTNER_FORM" - Advertiser entered
	// phone in contact form.
	//   "SMB_FAILED_RECAPTCHA_IN_CONTACT_PARTNER_FORM" - Advertiser failed
	// <a href="https://www.google.com/recaptcha/">reCaptcha</a>
	// in contact form.
	//   "PARTNER_VIEWED_BY_SMB" - Company viewed by advertiser.
	//   "SMB_CANCELED_PARTNER_CONTACT_FORM_ON_GPS" - Advertiser canceled
	// partner contact form on Google Partner Search.
	//   "SMB_CHANGED_A_SEARCH_PARAMETER_TOP" - Advertiser changed a top
	// search parameter.
	//   "SMB_CLICKED_CONTACT_A_PARTNER_ON_GPS" - Advertiser started partner
	// contact form on Google Partner Search.
	//   "SMB_CLICKED_SHOW_MORE_PARTNERS_BUTTON_BOTTOM" - Advertiser clicked
	// `Show more partners` bottom button.
	//   "SMB_COMPLETED_PARTNER_CONTACT_FORM_ON_GPS" - Advertiser completed
	// partner contact form on Google Partner Search.
	//   "SMB_NO_PARTNERS_AVAILABLE_WITH_SEARCH_CRITERIA" - Advertiser saw
	// no partners available with search criteria.
	//   "SMB_PERFORMED_SEARCH_ON_GPS" - Advertiser performed search on
	// Google Partner Search.
	//   "SMB_VIEWED_A_PARTNER_ON_GPS" - Advertiser viewed a partner on
	// Google Partner Search.
	//   "SMB_CANCELED_PARTNER_CONTACT_FORM_ON_PROFILE_PAGE" - Advertiser
	// canceled partner contact form on profile page.
	//   "SMB_CLICKED_CONTACT_A_PARTNER_ON_PROFILE_PAGE" - Advertiser
	// started partner contact form on profile page.
	//   "SMB_CLICKED_PARTNER_WEBSITE" - Advertiser clicked partner website.
	//   "SMB_COMPLETED_PARTNER_CONTACT_FORM_ON_PROFILE_PAGE" - Advertiser
	// completed contact form on profile page.
	//   "SMB_VIEWED_A_PARTNER_PROFILE" - Advertiser viewed a partner
	// profile.
	//   "AGENCY_CLICKED_ACCEPT_TOS_BUTTON" - Agency clicked `accept Terms
	// Of Service` button.
	//   "AGENCY_CHANGED_TOS_COUNTRY" - Agency changed Terms Of Service
	// country.
	//   "AGENCY_ADDED_ADDRESS_IN_MY_PROFILE_PORTAL" - Agency added address
	// in profile portal.
	//   "AGENCY_ADDED_PHONE_NUMBER_IN_MY_PROFILE_PORTAL" - Agency added
	// phone number in profile portal.
	//   "AGENCY_CHANGED_PRIMARY_ACCOUNT_ASSOCIATION" - Agency changed
	// primary account association.
	//   "AGENCY_CHANGED_PRIMARY_COUNTRY_ASSOCIATION" - Agency changed
	// primary country association.
	//   "AGENCY_CLICKED_AFFILIATE_BUTTON_IN_MY_PROFILE_IN_PORTAL" - Agency
	// clicked `affiliate` button in profile portal.
	//   "AGENCY_CLICKED_GIVE_EDIT_ACCESS_IN_MY_PROFILE_PORTAL" - Agency
	// clicked `give edit access` in profile portal.
	//   "AGENCY_CLICKED_LOG_OUT_IN_MY_PROFILE_PORTAL" - Agency clicked `log
	// out` in profile portal.
	//   "AGENCY_CLICKED_MY_PROFILE_LEFT_NAV_IN_PORTAL" - Agency clicked
	// profile portal left nav.
	//   "AGENCY_CLICKED_SAVE_AND_CONTINUE_AT_BOT_OF_COMPLETE_PROFILE" -
	// Agency clicked `save and continue` at bottom of complete profile.
	//   "AGENCY_CLICKED_UNAFFILIATE_IN_MY_PROFILE_PORTAL" - Agency clicked
	// `unaffiliate` in profile portal.
	//   "AGENCY_FILLED_OUT_COMP_AFFILIATION_IN_MY_PROFILE_PORTAL" - Agency
	// filled out company affiliation in profile portal.
	//   "AGENCY_SUCCESSFULLY_CONNECTED_WITH_COMPANY_IN_MY_PROFILE" - Agency
	// successfully connected with company in profile portal.
	//   "AGENCY_CLICKED_CREATE_MCC_IN_MY_PROFILE_PORTAL" - Agency clicked
	// create MCC in profile portal.
	//   "AGENCY_DIDNT_HAVE_AN_MCC_ASSOCIATED_ON_COMPLETE_PROFILE" - Agency
	// did not have an MCC associated on profile portal.
	//   "AGENCY_HAD_AN_MCC_ASSOCIATED_ON_COMPLETE_PROFILE" - Agency had an
	// MCC associated on profile portal.
	//   "AGENCY_ADDED_JOB_FUNCTION_IN_MY_PROFILE_PORTAL" - Agency added job
	// function in profile portal.
	//   "AGENCY_LOOKED_AT_JOB_FUNCTION_DROP_DOWN" - Agency looked at job
	// function drop-down.
	//   "AGENCY_SELECTED_ACCOUNT_MANAGER_AS_JOB_FUNCTION" - Agency selected
	// `account manage` as job function.
	//   "AGENCY_SELECTED_ACCOUNT_PLANNER_AS_JOB_FUNCTION" - Agency selected
	// `account planner` as job function.
	//   "AGENCY_SELECTED_ANALYTICS_AS_JOB_FUNCTION" - Agency selected
	// `Analytics` as job function.
	//   "AGENCY_SELECTED_CREATIVE_AS_JOB_FUNCTION" - Agency selected
	// `creative` as job function.
	//   "AGENCY_SELECTED_MEDIA_BUYER_AS_JOB_FUNCTION" - Agency selected
	// `media buyer` as job function.
	//   "AGENCY_SELECTED_MEDIA_PLANNER_AS_JOB_FUNCTION" - Agency selected
	// `media planner` as job function.
	//   "AGENCY_SELECTED_OTHER_AS_JOB_FUNCTION" - Agency selected `other`
	// as job function.
	//   "AGENCY_SELECTED_PRODUCTION_AS_JOB_FUNCTION" - Agency selected
	// `production` as job function.
	//   "AGENCY_SELECTED_SEO_AS_JOB_FUNCTION" - Agency selected `SEO` as
	// job function.
	//   "AGENCY_SELECTED_SALES_REP_AS_JOB_FUNCTION" - Agency selected
	// `sales rep` as job function.
	//   "AGENCY_SELECTED_SEARCH_SPECIALIST_AS_JOB_FUNCTION" - Agency
	// selected `search specialist` as job function.
	//   "AGENCY_ADDED_CHANNELS_IN_MY_PROFILE_PORTAL" - Agency added
	// channels in profile portal.
	//   "AGENCY_LOOKED_AT_ADD_CHANNEL_DROP_DOWN" - Agency looked at `add
	// channel` drop-down.
	//   "AGENCY_SELECTED_CROSS_CHANNEL_FROM_ADD_CHANNEL" - Agency selected
	// `cross channel` from add channel drop-down.
	//   "AGENCY_SELECTED_DISPLAY_FROM_ADD_CHANNEL" - Agency selected
	// `display` from add channel drop-down.
	//   "AGENCY_SELECTED_MOBILE_FROM_ADD_CHANNEL" - Agency selected
	// `mobile` from add channel drop-down.
	//   "AGENCY_SELECTED_SEARCH_FROM_ADD_CHANNEL" - Agency selected
	// `search` from add channel drop-down.
	//   "AGENCY_SELECTED_SOCIAL_FROM_ADD_CHANNEL" - Agency selected
	// `social` from add channel drop-down.
	//   "AGENCY_SELECTED_TOOLS_FROM_ADD_CHANNEL" - Agency selected `tools`
	// from add channel drop-down.
	//   "AGENCY_SELECTED_YOUTUBE_FROM_ADD_CHANNEL" - Agency selected
	// `YouTube` from add channel drop-down.
	//   "AGENCY_ADDED_INDUSTRIES_IN_MY_PROFILE_PORTAL" - Agency added
	// industries in profile portal.
	//   "AGENCY_CHANGED_ADD_INDUSTRIES_DROP_DOWN" - Agency changed `add
	// industries` drop-down.
	//   "AGENCY_ADDED_MARKETS_IN_MY_PROFILE_PORTAL" - Agency added markets
	// in profile portal.
	//   "AGENCY_CHANGED_ADD_MARKETS_DROP_DOWN" - Agency changed `add
	// markets` drop-down.
	//   "AGENCY_CHECKED_RECIEVE_MAIL_PROMOTIONS_MYPROFILE" - Agency checked
	// `recieve mail promotions` in profile portal.
	//   "AGENCY_CHECKED_RECIEVE_MAIL_PROMOTIONS_SIGNUP" - Agency checked
	// `recieve mail promotions` in sign-up.
	//   "AGENCY_SELECTED_OPT_IN_BETA_TESTS_AND_MKT_RESEARCH" - Agency
	// selected `opt-in beta tests and market research`.
	//   "AGENCY_SELECTED_OPT_IN_BETA_TESTS_IN_MY_PROFILE_PORTAL" - Agency
	// selected `opt-in beta tests` in profile portal.
	//   "AGENCY_SELECTED_OPT_IN_NEWS_IN_MY_PROFILE_PORTAL" - Agency
	// selected `opt-in news` in profile portal.
	//   "AGENCY_SELECTED_OPT_IN_NEWS_INVITATIONS_AND_PROMOS" - Agency
	// selected `opt-in news invitations and promotions`.
	//   "AGENCY_SELECTED_OPT_IN_PERFORMANCE_SUG_IN_MY_PROFILE_PORTAL" -
	// Agency selected `opt-in performance SUG` in profile portal.
	//   "AGENCY_SELECTED_OPT_IN_PERFORMANCE_SUGGESTIONS" - Agency selected
	// `opt-in performance suggestions`.
	//   "AGENCY_SELECTED_OPT_IN_SELECT_ALL_EMAIL_NOTIFICATIONS" - Agency
	// selected `opt-in select all email notifications`.
	//   "AGENCY_SELECTED_SELECT_ALL_OPT_INS_IN_MY_PROFILE_PORTAL" - Agency
	// selected `select all opt-ins` in profile portal.
	//   "AGENCY_CLICKED_BACK_BUTTON_ON_CONNECT_WITH_COMPANY" - Agency
	// clicked back button on `connect with company`.
	//   "AGENCY_CLICKED_CONTINUE_TO_OVERVIEW_ON_CONNECT_WITH_COMPANY" -
	// Agency clicked continue to overview on `connect with company`.
	//   "AGECNY_CLICKED_CREATE_MCC_CONNECT_WITH_COMPANY_NOT_FOUND" - Agency
	// clicked `create MCC connect with company not found`.
	//   "AGECNY_CLICKED_GIVE_EDIT_ACCESS_CONNECT_WITH_COMPANY_NOT_FOUND" -
	// Agency clicked `give edit access connect with company not found`.
	//   "AGECNY_CLICKED_LOG_OUT_CONNECT_WITH_COMPANY_NOT_FOUND" - Agency
	// clicked `log out connect with company not found`.
	//   "AGENCY_CLICKED_SKIP_FOR_NOW_ON_CONNECT_WITH_COMPANY_PAGE" - Agency
	// clicked `skip for now on connect with company page`.
	//   "AGENCY_CLOSED_CONNECTED_TO_COMPANY_X_BUTTON_WRONG_COMPANY" -
	// Agency closed connection to company.
	//   "AGENCY_COMPLETED_FIELD_CONNECT_WITH_COMPANY" - Agency completed
	// field connect with company.
	//   "AGECNY_FOUND_COMPANY_TO_CONNECT_WITH" - Agency found company to
	// connect with.
	//   "AGENCY_SUCCESSFULLY_CREATED_COMPANY" - Agency successfully created
	// company.
	//   "AGENCY_ADDED_NEW_COMPANY_LOCATION" - Agency added new company
	// location.
	//   "AGENCY_CLICKED_COMMUNITY_JOIN_NOW_LINK_IN_PORTAL_NOTIFICATIONS" -
	// Agency clicked community `join now link` in portal notifications.
	//   "AGENCY_CLICKED_CONNECT_TO_COMPANY_LINK_IN_PORTAL_NOTIFICATIONS" -
	// Agency clicked `connect to company` link in portal notifications.
	//   "AGENCY_CLICKED_GET_CERTIFIED_LINK_IN_PORTAL_NOTIFICATIONS" -
	// Agency cliecked `get certified` link in portal notifications.
	//
	// "AGENCY_CLICKED_GET_VIDEO_ADS_CERTIFIED_LINK_IN_PORTAL_NOTIFICATIONS"
	// - Agency clicked `get VideoAds certified` link in portal
	// notifications.
	//   "AGENCY_CLICKED_LINK_TO_MCC_LINK_IN_PORTAL_NOTIFICATIONS" - Agency
	// clicked `link to MCC` link in portal notifications.
	//   "AGENCY_CLICKED_INSIGHT_CONTENT_IN_PORTAL" - Agency clicked
	// `insight content` in portal.
	//   "AGENCY_CLICKED_INSIGHTS_VIEW_NOW_PITCH_DECKS_IN_PORTAL" - Agency
	// clicked `insights view now pitch decks` in portal.
	//   "AGENCY_CLICKED_INSIGHTS_LEFT_NAV_IN_PORTAL" - Agency clicked
	// `insights` left nav in portal.
	//   "AGENCY_CLICKED_INSIGHTS_UPLOAD_CONTENT" - Agency clicked `insights
	// upload content`.
	//   "AGENCY_CLICKED_INSIGHTS_VIEWED_DEPRECATED" - Agency clicked
	// `insights viewed deprecated`.
	//   "AGENCY_CLICKED_COMMUNITY_LEFT_NAV_IN_PORTAL" - Agency clicked
	// `community` left nav in portal.
	//   "AGENCY_CLICKED_JOIN_COMMUNITY_BUTTON_COMMUNITY_PORTAL" - Agency
	// clicked `join community` button in community portal.
	//   "AGENCY_CLICKED_CERTIFICATIONS_LEFT_NAV_IN_PORTAL" - Agency clicked
	// `certifications` left nav in portal.
	//   "AGENCY_CLICKED_CERTIFICATIONS_PRODUCT_LEFT_NAV_IN_PORTAL" - Agency
	// clicked `certifications product` left nav in portal.
	//   "AGENCY_CLICKED_PARTNER_STATUS_LEFT_NAV_IN_PORTAL" - Agency clicked
	// `partner status` left nav in portal.
	//   "AGENCY_CLICKED_PARTNER_STATUS_PRODUCT_LEFT_NAV_IN_PORTAL" - Agency
	// clicked `partner status product` left nav in portal.
	//   "AGENCY_CLICKED_OFFERS_LEFT_NAV_IN_PORTAL" - Agency clicked
	// `offers` left nav in portal.
	//   "AGENCY_CLICKED_SEND_BUTTON_ON_OFFERS_PAGE" - Agency clicked `send`
	// button on offers page.
	//   "AGENCY_CLICKED_EXAM_DETAILS_ON_CERT_ADWORDS_PAGE" - Agency clicked
	// `exam details` on certifications AdWords page.
	//   "AGENCY_CLICKED_SEE_EXAMS_CERTIFICATION_MAIN_PAGE" - Agency clicked
	// `see exams` certifications main page.
	//   "AGENCY_CLICKED_TAKE_EXAM_ON_CERT_EXAM_PAGE" - Agency clicked `take
	// exam` on certifications exam page.
	//   "AGENCY_OPENED_LAST_ADMIN_DIALOG" - Agency opened `last admin`
	// dialog.
	//   "AGENCY_OPENED_DIALOG_WITH_NO_USERS" - Agency opened dialog with no
	// users.
	//   "AGENCY_PROMOTED_USER_TO_ADMIN" - Agency promoted user to admin.
	//   "AGENCY_UNAFFILIATED" - Agency unaffiliated.
	//   "AGENCY_CHANGED_ROLES" - Agency changed roles.
	//   "SMB_CLICKED_COMPANY_NAME_LINK_TO_PROFILE" - Advertiser clicked
	// `company name` link to profile.
	//   "SMB_VIEWED_ADWORDS_CERTIFICATE" - Advertiser viewed AdWords
	// certificate.
	//   "SMB_VIEWED_ADWORDS_SEARCH_CERTIFICATE" - Advertiser viewed AdWords
	// Search certificate.
	//   "SMB_VIEWED_ADWORDS_DISPLAY_CERTIFICATE" - Advertiser viewed
	// AdWords Display certificate.
	//   "SMB_CLICKED_ADWORDS_CERTIFICATE_HELP_ICON" - Advertiser clicked
	// AdWords certificate help icon.
	//   "SMB_VIEWED_ANALYTICS_CERTIFICATE" - Advertiser viewed Analytics
	// certificate.
	//   "SMB_VIEWED_DOUBLECLICK_CERTIFICATE" - Advertiser viewed
	// DoubleClick certificate.
	//   "SMB_VIEWED_MOBILE_SITES_CERTIFICATE" - Advertiser viewed Mobile
	// Sites certificate.
	//   "SMB_VIEWED_VIDEO_ADS_CERTIFICATE" - Advertiser viewed VideoAds
	// certificate.
	//   "SMB_VIEWED_SHOPPING_CERTIFICATE" - Advertiser clicked Shopping
	// certificate help icon.
	//   "SMB_CLICKED_VIDEO_ADS_CERTIFICATE_HELP_ICON" - Advertiser clicked
	// VideoAds certificate help icon.
	//   "SMB_VIEWED_DIGITAL_SALES_CERTIFICATE" - Advertiser viewed Digital
	// Sales certificate.
	//   "CLICKED_HELP_AT_BOTTOM" - Clicked `help` at bottom.
	//   "CLICKED_HELP_AT_TOP" - Clicked `help` at top.
	//   "CLIENT_ERROR" - Client error occurred.
	//   "AGENCY_CLICKED_LEFT_NAV_STORIES" - Agency clicked left nav
	// `stories`.
	//   "CLICKED" - Click occured.
	//   "SMB_VIEWED_MOBILE_CERTIFICATE" - Advertiser clicked Mobile
	// certificate help icon.
	//   "AGENCY_FAILED_COMPANY_VERIFICATION" - Agency failed the company
	// verification.
	//   "VISITED_LANDING" - User visited the landing portion of Google
	// Partners.
	//   "VISITED_GPS" - User visited the Google Partner Search portion of
	// Google Partners.
	//   "VISITED_AGENCY_PORTAL" - User visited the agency portal portion of
	// Google Partners.
	//   "CANCELLED_INDIVIDUAL_SIGN_UP" - User cancelled signing up.
	//   "CANCELLED_COMPANY_SIGN_UP" - User cancelled signing up their
	// company.
	//   "AGENCY_CLICKED_SIGN_IN_BUTTON_TOP" - Agency clicked `Sign in` top
	// button.
	//   "AGENCY_CLICKED_SAVE_AND_CONTINUE_AT_BOT_OF_INCOMPLETE_PROFILE" -
	// Agency clicked `save and continue` at bottom of incomplete profile.
	//   "AGENCY_UNSELECTED_OPT_IN_NEWS_INVITATIONS_AND_PROMOS" - Agency
	// unselected `opt-in news invitations and promotions`.
	//   "AGENCY_UNSELECTED_OPT_IN_BETA_TESTS_AND_MKT_RESEARCH" - Agency
	// unselected `opt-in beta tests and market research`.
	//   "AGENCY_UNSELECTED_OPT_IN_PERFORMANCE_SUGGESTIONS" - Agency
	// unselected `opt-in performance suggestions`.
	//   "AGENCY_SELECTED_OPT_OUT_UNSELECT_ALL_EMAIL_NOTIFICATIONS" - Agency
	// selected `opt-out unselect all email notifications`.
	//   "AGENCY_LINKED_INDIVIDUAL_MCC" - Agency linked their individual
	// MCC.
	//   "AGENCY_SUGGESTED_TO_USER" - Agency was suggested to user for
	// affiliation.
	//   "AGENCY_IGNORED_SUGGESTED_AGENCIES_AND_SEARCHED" - Agency ignored
	// suggested agencies and begin searching.
	//   "AGENCY_PICKED_SUGGESTED_AGENCY" - Agency picked a suggested
	// agency.
	//   "AGENCY_SEARCHED_FOR_AGENCIES" - Agency searched for agencies.
	//   "AGENCY_PICKED_SEARCHED_AGENCY" - Agency picked a searched agency.
	//   "AGENCY_DISMISSED_AFFILIATION_WIDGET" - Agency dismissed
	// affiliation widget.
	//   "AGENCY_CLICKED_INSIGHTS_DOWNLOAD_CONTENT" - Agency clicked on the
	// download link for downloading content.
	//   "AGENCY_PROGRESS_INSIGHTS_VIEW_CONTENT" - Agency user is maklingg
	// progress viewing a content item.
	//   "AGENCY_CLICKED_CANCEL_ACCEPT_TOS_BUTTON" - Agency clicked `cancel
	// Terms Of Service` button.
	//   "SMB_ENTERED_WEBSITE_IN_CONTACT_PARTNER_FORM" - Advertiser entered
	// website in contact form.
	//   "AGENCY_SELECTED_OPT_IN_AFA_MIGRATION" - Agency opted in for
	// migrating their exams to Academy for Ads.
	//   "AGENCY_SELECTED_OPT_OUT_AFA_MIGRATION" - Agency opted out for
	// migrating their exams to Academy for Ads.
	EventAction string `json:"eventAction,omitempty"`

	// EventCategory: The category the action belongs to.
	//
	// Possible values:
	//   "EVENT_CATEGORY_UNSPECIFIED" - Unchosen.
	//   "GOOGLE_PARTNER_SEARCH" - Google Partner Search category.
	//   "GOOGLE_PARTNER_SIGNUP_FLOW" - Google Partner sign-up flow
	// category.
	//   "GOOGLE_PARTNER_PORTAL" - Google Partner portal category.
	//   "GOOGLE_PARTNER_PORTAL_MY_PROFILE" - Google Partner portal
	// my-profile category.
	//   "GOOGLE_PARTNER_PORTAL_CERTIFICATIONS" - Google Partner portal
	// certifications category.
	//   "GOOGLE_PARTNER_PORTAL_COMMUNITY" - Google Partner portal community
	// category.
	//   "GOOGLE_PARTNER_PORTAL_INSIGHTS" - Google Partner portal insights
	// category.
	//   "GOOGLE_PARTNER_PORTAL_CLIENTS" - Google Partner portal clients
	// category.
	//   "GOOGLE_PARTNER_PUBLIC_USER_PROFILE" - Google Partner portal public
	// user profile category.
	//   "GOOGLE_PARTNER_PANEL" - Google Partner panel category.
	//   "GOOGLE_PARTNER_PORTAL_LAST_ADMIN_DIALOG" - Google Partner portal
	// last admin dialog category.
	//   "GOOGLE_PARTNER_CLIENT" - Google Partner client category.
	//   "GOOGLE_PARTNER_PORTAL_COMPANY_PROFILE" - Google Partner portal
	// company profile category.
	//   "EXTERNAL_LINKS" - External links category.
	//   "GOOGLE_PARTNER_LANDING" - Google Partner landing category.
	EventCategory string `json:"eventCategory,omitempty"`

	// EventDatas: List of event data for the event.
	EventDatas []*EventData `json:"eventDatas,omitempty"`

	// EventScope: The scope of the event.
	//
	// Possible values:
	//   "EVENT_SCOPE_UNSPECIFIED" - Unchosen.
	//   "VISITOR" - Based on visitor.
	//   "SESSION" - Based on session.
	//   "PAGE" - Based on page visit.
	EventScope string `json:"eventScope,omitempty"`

	// Lead: Advertiser lead information.
	Lead *Lead `json:"lead,omitempty"`

	// RequestMetadata: Current request metadata.
	RequestMetadata *RequestMetadata `json:"requestMetadata,omitempty"`

	// Url: The URL where the event occurred.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "EventAction") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "EventAction") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *LogUserEventRequest) MarshalJSON() ([]byte, error) {
	type noMethod LogUserEventRequest
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// LogUserEventResponse: Response message for
// LogUserEvent.
type LogUserEventResponse struct {
	// ResponseMetadata: Current response metadata.
	ResponseMetadata *ResponseMetadata `json:"responseMetadata,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ResponseMetadata") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ResponseMetadata") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *LogUserEventResponse) MarshalJSON() ([]byte, error) {
	type noMethod LogUserEventResponse
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Money: Represents an amount of money with its currency type.
type Money struct {
	// CurrencyCode: The 3-letter currency code defined in ISO 4217.
	CurrencyCode string `json:"currencyCode,omitempty"`

	// Nanos: Number of nano (10^-9) units of the amount.
	// The value must be between -999,999,999 and +999,999,999 inclusive.
	// If `units` is positive, `nanos` must be positive or zero.
	// If `units` is zero, `nanos` can be positive, zero, or negative.
	// If `units` is negative, `nanos` must be negative or zero.
	// For example $-1.75 is represented as `units`=-1 and
	// `nanos`=-750,000,000.
	Nanos int64 `json:"nanos,omitempty"`

	// Units: The whole units of the amount.
	// For example if `currencyCode` is "USD", then 1 unit is one US
	// dollar.
	Units int64 `json:"units,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "CurrencyCode") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "CurrencyCode") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Money) MarshalJSON() ([]byte, error) {
	type noMethod Money
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OfferCustomer: Customers qualified for an offer.
type OfferCustomer struct {
	// AdwordsUrl: URL to the customer's AdWords page.
	AdwordsUrl string `json:"adwordsUrl,omitempty"`

	// CountryCode: Country code of the customer.
	CountryCode string `json:"countryCode,omitempty"`

	// CreationTime: Time the customer was created.
	CreationTime string `json:"creationTime,omitempty"`

	// EligibilityDaysLeft: Days the customer is still eligible.
	EligibilityDaysLeft int64 `json:"eligibilityDaysLeft,omitempty"`

	// ExternalCid: External CID for the customer.
	ExternalCid int64 `json:"externalCid,omitempty,string"`

	// GetYAmount: Formatted Get Y amount with currency code.
	GetYAmount string `json:"getYAmount,omitempty"`

	// Name: Name of the customer.
	Name string `json:"name,omitempty"`

	// OfferType: Type of the offer
	//
	// Possible values:
	//   "OFFER_TYPE_UNSPECIFIED" - Unset.
	//   "OFFER_TYPE_SPEND_X_GET_Y" - AdWords spend X get Y.
	//   "OFFER_TYPE_VIDEO" - Youtube video.
	//   "OFFER_TYPE_SPEND_MATCH" - Spend Match up to Y.
	OfferType string `json:"offerType,omitempty"`

	// SpendXAmount: Formatted Spend X amount with currency code.
	SpendXAmount string `json:"spendXAmount,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AdwordsUrl") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AdwordsUrl") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OfferCustomer) MarshalJSON() ([]byte, error) {
	type noMethod OfferCustomer
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// OptIns: A set of opt-ins for a user.
type OptIns struct {
	// MarketComm: An opt-in about receiving email from Partners marketing
	// teams. Includes
	// member-only events and special promotional offers for Google
	// products.
	MarketComm bool `json:"marketComm,omitempty"`

	// PerformanceSuggestions: An opt-in about receiving email with
	// customized AdWords campaign management
	// tips.
	PerformanceSuggestions bool `json:"performanceSuggestions,omitempty"`

	// PhoneContact: An opt-in to allow recieivng phone calls about their
	// Partners account.
	PhoneContact bool `json:"phoneContact,omitempty"`

	// PhysicalMail: An opt-in to receive special promotional gifts and
	// material in the mail.
	PhysicalMail bool `json:"physicalMail,omitempty"`

	// SpecialOffers: An opt-in about receiving email regarding new features
	// and products.
	SpecialOffers bool `json:"specialOffers,omitempty"`

	// ForceSendFields is a list of field names (e.g. "MarketComm") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "MarketComm") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *OptIns) MarshalJSON() ([]byte, error) {
	type noMethod OptIns
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PublicProfile: Basic information from a public profile.
type PublicProfile struct {
	// DisplayImageUrl: The URL to the main display image of the public
	// profile. Being deprecated.
	DisplayImageUrl string `json:"displayImageUrl,omitempty"`

	// DisplayName: The display name of the public profile.
	DisplayName string `json:"displayName,omitempty"`

	// Id: The ID which can be used to retrieve more details about the
	// public profile.
	Id string `json:"id,omitempty"`

	// ProfileImage: The URL to the main profile image of the public
	// profile.
	ProfileImage string `json:"profileImage,omitempty"`

	// Url: The URL of the public profile.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DisplayImageUrl") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DisplayImageUrl") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *PublicProfile) MarshalJSON() ([]byte, error) {
	type noMethod PublicProfile
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Rank: Information related to ranking of results.
type Rank struct {
	// Type: The type of rank.
	//
	// Possible values:
	//   "RANK_TYPE_UNSPECIFIED" - Unchosen.
	//   "RT_FINAL_SCORE" - Total final score.
	Type string `json:"type,omitempty"`

	// Value: The numerical value of the rank.
	Value float64 `json:"value,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Type") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Type") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Rank) MarshalJSON() ([]byte, error) {
	type noMethod Rank
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Rank) UnmarshalJSON(data []byte) error {
	type noMethod Rank
	var s1 struct {
		Value gensupport.JSONFloat64 `json:"value"`
		*noMethod
	}
	s1.noMethod = (*noMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.Value = float64(s1.Value)
	return nil
}

// RecaptchaChallenge: <a
// href="https://www.google.com/recaptcha/">reCaptcha</a> challenge
// info.
type RecaptchaChallenge struct {
	// Id: The ID of the reCaptcha challenge.
	Id string `json:"id,omitempty"`

	// Response: The response to the reCaptcha challenge.
	Response string `json:"response,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Id") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RecaptchaChallenge) MarshalJSON() ([]byte, error) {
	type noMethod RecaptchaChallenge
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// RequestMetadata: Common data that is in each API request.
type RequestMetadata struct {
	// ExperimentIds: Experiment IDs the current request belongs to.
	ExperimentIds []string `json:"experimentIds,omitempty"`

	// Locale: Locale to use for the current request.
	Locale string `json:"locale,omitempty"`

	// PartnersSessionId: Google Partners session ID.
	PartnersSessionId string `json:"partnersSessionId,omitempty"`

	// TrafficSource: Source of traffic for the current request.
	TrafficSource *TrafficSource `json:"trafficSource,omitempty"`

	// UserOverrides: Values to use instead of the user's respective
	// defaults for the current
	// request. These are only honored by whitelisted products.
	UserOverrides *UserOverrides `json:"userOverrides,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ExperimentIds") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ExperimentIds") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RequestMetadata) MarshalJSON() ([]byte, error) {
	type noMethod RequestMetadata
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResponseMetadata: Common data that is in each API response.
type ResponseMetadata struct {
	// DebugInfo: Debug information about this request.
	DebugInfo *DebugInfo `json:"debugInfo,omitempty"`

	// ForceSendFields is a list of field names (e.g. "DebugInfo") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "DebugInfo") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResponseMetadata) MarshalJSON() ([]byte, error) {
	type noMethod ResponseMetadata
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// SpecializationStatus: Agency specialization status
type SpecializationStatus struct {
	// BadgeSpecialization: The specialization this status is for.
	//
	// Possible values:
	//   "BADGE_SPECIALIZATION_UNKNOWN" - Unknown specialization
	//   "BADGE_SPECIALIZATION_ADWORDS_SEARCH" - AdWords Search
	// specialization
	//   "BADGE_SPECIALIZATION_ADWORDS_DISPLAY" - AdWords Display
	// specialization
	//   "BADGE_SPECIALIZATION_ADWORDS_MOBILE" - AdWords Mobile
	// specialization
	//   "BADGE_SPECIALIZATION_ADWORDS_VIDEO" - AdWords Video specialization
	//   "BADGE_SPECIALIZATION_ADWORDS_SHOPPING" - AdWords Shopping
	// specialization
	BadgeSpecialization string `json:"badgeSpecialization,omitempty"`

	// BadgeSpecializationState: State of agency specialization.
	//
	// Possible values:
	//   "BADGE_SPECIALIZATION_STATE_UNKNOWN" - Unknown state
	//   "BADGE_SPECIALIZATION_STATE_PASSED" - Specialization passed
	//   "BADGE_SPECIALIZATION_STATE_NOT_PASSED" - Specialization not passed
	//   "BADGE_SPECIALIZATION_STATE_IN_GRACE" - Specialization in grace
	BadgeSpecializationState string `json:"badgeSpecializationState,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BadgeSpecialization")
	// to unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BadgeSpecialization") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *SpecializationStatus) MarshalJSON() ([]byte, error) {
	type noMethod SpecializationStatus
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// TrafficSource: Source of traffic for the current request.
type TrafficSource struct {
	// TrafficSourceId: Identifier to indicate where the traffic comes
	// from.
	// An identifier has multiple letters created by a team which redirected
	// the
	// traffic to us.
	TrafficSourceId string `json:"trafficSourceId,omitempty"`

	// TrafficSubId: Second level identifier to indicate where the traffic
	// comes from.
	// An identifier has multiple letters created by a team which redirected
	// the
	// traffic to us.
	TrafficSubId string `json:"trafficSubId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "TrafficSourceId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "TrafficSourceId") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *TrafficSource) MarshalJSON() ([]byte, error) {
	type noMethod TrafficSource
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// User: A resource representing a user of the Partners platform.
type User struct {
	// AvailableAdwordsManagerAccounts: This is the list of AdWords Manager
	// Accounts the user has edit access to.
	// If the user has edit access to multiple accounts, the user can choose
	// the
	// preferred account and we use this when a personal account is needed.
	// Can
	// be empty meaning the user has access to no accounts.
	// @OutputOnly
	AvailableAdwordsManagerAccounts []*AdWordsManagerAccountInfo `json:"availableAdwordsManagerAccounts,omitempty"`

	// CertificationStatus: The list of achieved certifications. These are
	// calculated based on exam
	// results and other requirements.
	// @OutputOnly
	CertificationStatus []*Certification `json:"certificationStatus,omitempty"`

	// Company: The company that the user is associated with.
	// If not present, the user is not associated with any company.
	Company *CompanyRelation `json:"company,omitempty"`

	// CompanyVerificationEmail: The email address used by the user used for
	// company verification.
	// @OutputOnly
	CompanyVerificationEmail string `json:"companyVerificationEmail,omitempty"`

	// ExamStatus: The list of exams the user ever taken. For each type of
	// exam, only one
	// entry is listed.
	ExamStatus []*ExamStatus `json:"examStatus,omitempty"`

	// Id: The ID of the user.
	Id string `json:"id,omitempty"`

	// InternalId: The internal user ID.
	// Only available for a whitelisted set of api clients.
	InternalId string `json:"internalId,omitempty"`

	// LastAccessTime: The most recent time the user interacted with the
	// Partners site.
	// @OutputOnly
	LastAccessTime string `json:"lastAccessTime,omitempty"`

	// PrimaryEmails: The list of emails the user has access to/can select
	// as primary.
	// @OutputOnly
	PrimaryEmails []string `json:"primaryEmails,omitempty"`

	// Profile: The profile information of a Partners user, contains all the
	// directly
	// editable user information.
	Profile *UserProfile `json:"profile,omitempty"`

	// PublicProfile: Information about a user's external public profile
	// outside Google Partners.
	PublicProfile *PublicProfile `json:"publicProfile,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g.
	// "AvailableAdwordsManagerAccounts") to unconditionally include in API
	// requests. By default, fields with empty values are omitted from API
	// requests. However, any non-pointer, non-interface field appearing in
	// ForceSendFields will be sent to the server regardless of whether the
	// field is empty or not. This may be used to include empty fields in
	// Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g.
	// "AvailableAdwordsManagerAccounts") to include in API requests with
	// the JSON null value. By default, fields with empty values are omitted
	// from API requests. However, any field with an empty value appearing
	// in NullFields will be sent to the server as null. It is an error if a
	// field in this list has a non-empty value. This may be used to include
	// null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *User) MarshalJSON() ([]byte, error) {
	type noMethod User
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UserOverrides: Values to use instead of the user's respective
// defaults. These are only
// honored by whitelisted products.
type UserOverrides struct {
	// IpAddress: IP address to use instead of the user's geo-located IP
	// address.
	IpAddress string `json:"ipAddress,omitempty"`

	// UserId: Logged-in user ID to impersonate instead of the user's ID.
	UserId string `json:"userId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "IpAddress") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "IpAddress") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UserOverrides) MarshalJSON() ([]byte, error) {
	type noMethod UserOverrides
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UserProfile: The profile information of a Partners user.
type UserProfile struct {
	// Address: The user's mailing address, contains multiple fields.
	Address *Location `json:"address,omitempty"`

	// AdwordsManagerAccount: If the user has edit access to multiple
	// accounts, the user can choose the
	// preferred account and it is used when a personal account is needed.
	// Can
	// be empty.
	AdwordsManagerAccount int64 `json:"adwordsManagerAccount,omitempty,string"`

	// Channels: A list of ids representing which channels the user selected
	// they were in.
	Channels []string `json:"channels,omitempty"`

	// EmailAddress: The email address the user has selected on the Partners
	// site as primary.
	EmailAddress string `json:"emailAddress,omitempty"`

	// EmailOptIns: The list of opt-ins for the user, related to
	// communication preferences.
	EmailOptIns *OptIns `json:"emailOptIns,omitempty"`

	// FamilyName: The user's family name.
	FamilyName string `json:"familyName,omitempty"`

	// GivenName: The user's given name.
	GivenName string `json:"givenName,omitempty"`

	// Industries: A list of ids representing which industries the user
	// selected.
	Industries []string `json:"industries,omitempty"`

	// JobFunctions: A list of ids represnting which job categories the user
	// selected.
	JobFunctions []string `json:"jobFunctions,omitempty"`

	// Languages: The list of languages this user understands.
	Languages []string `json:"languages,omitempty"`

	// Markets: A list of ids representing which markets the user was
	// interested in.
	Markets []string `json:"markets,omitempty"`

	// MigrateToAfa: Whether or not to migrate the user's exam data to
	// Academy for Ads.
	MigrateToAfa bool `json:"migrateToAfa,omitempty"`

	// PhoneNumber: The user's phone number.
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// PrimaryCountryCode: The user's primary country, an ISO 2-character
	// code.
	PrimaryCountryCode string `json:"primaryCountryCode,omitempty"`

	// ProfilePublic: Whether the user's public profile is visible to anyone
	// with the URL.
	ProfilePublic bool `json:"profilePublic,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Address") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Address") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UserProfile) MarshalJSON() ([]byte, error) {
	type noMethod UserProfile
	raw := noMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "partners.analytics.list":

type AnalyticsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists analytics data for a user's associated company.
// Should only be called within the context of an authorized logged in
// user.
func (r *AnalyticsService) List() *AnalyticsListCall {
	c := &AnalyticsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// Server may return fewer analytics than requested.
// If unspecified or set to 0, default value is 30.
// Specifies the number of days in the date range when querying
// analytics.
// The `page_token` represents the end date of the date range
// and the start date is calculated using the `page_size` as the
// number
// of days BEFORE the end date.
// Must be a non-negative integer.
func (c *AnalyticsListCall) PageSize(pageSize int64) *AnalyticsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results that the server returns.
// Typically, this is the value of
// `ListAnalyticsResponse.next_page_token`
// returned from the previous call to
// ListAnalytics.
// Will be a date string in `YYYY-MM-DD` format representing the end
// date
// of the date range of results to return.
// If unspecified or set to "", default value is the current date.
func (c *AnalyticsListCall) PageToken(pageToken string) *AnalyticsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// RequestMetadataExperimentIds sets the optional parameter
// "requestMetadata.experimentIds": Experiment IDs the current request
// belongs to.
func (c *AnalyticsListCall) RequestMetadataExperimentIds(requestMetadataExperimentIds ...string) *AnalyticsListCall {
	c.urlParams_.SetMulti("requestMetadata.experimentIds", append([]string{}, requestMetadataExperimentIds...))
	return c
}

// RequestMetadataLocale sets the optional parameter
// "requestMetadata.locale": Locale to use for the current request.
func (c *AnalyticsListCall) RequestMetadataLocale(requestMetadataLocale string) *AnalyticsListCall {
	c.urlParams_.Set("requestMetadata.locale", requestMetadataLocale)
	return c
}

// RequestMetadataPartnersSessionId sets the optional parameter
// "requestMetadata.partnersSessionId": Google Partners session ID.
func (c *AnalyticsListCall) RequestMetadataPartnersSessionId(requestMetadataPartnersSessionId string) *AnalyticsListCall {
	c.urlParams_.Set("requestMetadata.partnersSessionId", requestMetadataPartnersSessionId)
	return c
}

// RequestMetadataTrafficSourceTrafficSourceId sets the optional
// parameter "requestMetadata.trafficSource.trafficSourceId": Identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *AnalyticsListCall) RequestMetadataTrafficSourceTrafficSourceId(requestMetadataTrafficSourceTrafficSourceId string) *AnalyticsListCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSourceId", requestMetadataTrafficSourceTrafficSourceId)
	return c
}

// RequestMetadataTrafficSourceTrafficSubId sets the optional parameter
// "requestMetadata.trafficSource.trafficSubId": Second level identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *AnalyticsListCall) RequestMetadataTrafficSourceTrafficSubId(requestMetadataTrafficSourceTrafficSubId string) *AnalyticsListCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSubId", requestMetadataTrafficSourceTrafficSubId)
	return c
}

// RequestMetadataUserOverridesIpAddress sets the optional parameter
// "requestMetadata.userOverrides.ipAddress": IP address to use instead
// of the user's geo-located IP address.
func (c *AnalyticsListCall) RequestMetadataUserOverridesIpAddress(requestMetadataUserOverridesIpAddress string) *AnalyticsListCall {
	c.urlParams_.Set("requestMetadata.userOverrides.ipAddress", requestMetadataUserOverridesIpAddress)
	return c
}

// RequestMetadataUserOverridesUserId sets the optional parameter
// "requestMetadata.userOverrides.userId": Logged-in user ID to
// impersonate instead of the user's ID.
func (c *AnalyticsListCall) RequestMetadataUserOverridesUserId(requestMetadataUserOverridesUserId string) *AnalyticsListCall {
	c.urlParams_.Set("requestMetadata.userOverrides.userId", requestMetadataUserOverridesUserId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AnalyticsListCall) Fields(s ...googleapi.Field) *AnalyticsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AnalyticsListCall) IfNoneMatch(entityTag string) *AnalyticsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AnalyticsListCall) Context(ctx context.Context) *AnalyticsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AnalyticsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AnalyticsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/analytics")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.analytics.list" call.
// Exactly one of *ListAnalyticsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListAnalyticsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AnalyticsListCall) Do(opts ...googleapi.CallOption) (*ListAnalyticsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListAnalyticsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists analytics data for a user's associated company.\nShould only be called within the context of an authorized logged in user.",
	//   "flatPath": "v2/analytics",
	//   "httpMethod": "GET",
	//   "id": "partners.analytics.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "pageSize": {
	//       "description": "Requested page size. Server may return fewer analytics than requested.\nIf unspecified or set to 0, default value is 30.\nSpecifies the number of days in the date range when querying analytics.\nThe `page_token` represents the end date of the date range\nand the start date is calculated using the `page_size` as the number\nof days BEFORE the end date.\nMust be a non-negative integer.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results that the server returns.\nTypically, this is the value of `ListAnalyticsResponse.next_page_token`\nreturned from the previous call to\nListAnalytics.\nWill be a date string in `YYYY-MM-DD` format representing the end date\nof the date range of results to return.\nIf unspecified or set to \"\", default value is the current date.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.experimentIds": {
	//       "description": "Experiment IDs the current request belongs to.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.locale": {
	//       "description": "Locale to use for the current request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.partnersSessionId": {
	//       "description": "Google Partners session ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSourceId": {
	//       "description": "Identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSubId": {
	//       "description": "Second level identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.ipAddress": {
	//       "description": "IP address to use instead of the user's geo-located IP address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.userId": {
	//       "description": "Logged-in user ID to impersonate instead of the user's ID.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/analytics",
	//   "response": {
	//     "$ref": "ListAnalyticsResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AnalyticsListCall) Pages(ctx context.Context, f func(*ListAnalyticsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "partners.clientMessages.log":

type ClientMessagesLogCall struct {
	s                 *Service
	logmessagerequest *LogMessageRequest
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Log: Logs a generic message from the client, such as
// `Failed to render component`, `Profile page is running slow`,
// `More than 500 users have accessed this result.`, etc.
func (r *ClientMessagesService) Log(logmessagerequest *LogMessageRequest) *ClientMessagesLogCall {
	c := &ClientMessagesLogCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.logmessagerequest = logmessagerequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ClientMessagesLogCall) Fields(s ...googleapi.Field) *ClientMessagesLogCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ClientMessagesLogCall) Context(ctx context.Context) *ClientMessagesLogCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ClientMessagesLogCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ClientMessagesLogCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logmessagerequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/clientMessages:log")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.clientMessages.log" call.
// Exactly one of *LogMessageResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *LogMessageResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ClientMessagesLogCall) Do(opts ...googleapi.CallOption) (*LogMessageResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &LogMessageResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Logs a generic message from the client, such as\n`Failed to render component`, `Profile page is running slow`,\n`More than 500 users have accessed this result.`, etc.",
	//   "flatPath": "v2/clientMessages:log",
	//   "httpMethod": "POST",
	//   "id": "partners.clientMessages.log",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2/clientMessages:log",
	//   "request": {
	//     "$ref": "LogMessageRequest"
	//   },
	//   "response": {
	//     "$ref": "LogMessageResponse"
	//   }
	// }

}

// method id "partners.companies.get":

type CompaniesGetCall struct {
	s            *Service
	companyId    string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a company.
func (r *CompaniesService) Get(companyId string) *CompaniesGetCall {
	c := &CompaniesGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.companyId = companyId
	return c
}

// Address sets the optional parameter "address": The address to use for
// sorting the company's addresses by proximity.
// If not given, the geo-located address of the request is used.
// Used when order_by is set.
func (c *CompaniesGetCall) Address(address string) *CompaniesGetCall {
	c.urlParams_.Set("address", address)
	return c
}

// CurrencyCode sets the optional parameter "currencyCode": If the
// company's budget is in a different currency code than this one,
// then
// the converted budget is converted to this currency code.
func (c *CompaniesGetCall) CurrencyCode(currencyCode string) *CompaniesGetCall {
	c.urlParams_.Set("currencyCode", currencyCode)
	return c
}

// OrderBy sets the optional parameter "orderBy": How to order addresses
// within the returned company. Currently, only
// `address` and `address desc` is supported which will sorted by
// closest to
// farthest in distance from given address and farthest to closest
// distance
// from given address respectively.
func (c *CompaniesGetCall) OrderBy(orderBy string) *CompaniesGetCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// RequestMetadataExperimentIds sets the optional parameter
// "requestMetadata.experimentIds": Experiment IDs the current request
// belongs to.
func (c *CompaniesGetCall) RequestMetadataExperimentIds(requestMetadataExperimentIds ...string) *CompaniesGetCall {
	c.urlParams_.SetMulti("requestMetadata.experimentIds", append([]string{}, requestMetadataExperimentIds...))
	return c
}

// RequestMetadataLocale sets the optional parameter
// "requestMetadata.locale": Locale to use for the current request.
func (c *CompaniesGetCall) RequestMetadataLocale(requestMetadataLocale string) *CompaniesGetCall {
	c.urlParams_.Set("requestMetadata.locale", requestMetadataLocale)
	return c
}

// RequestMetadataPartnersSessionId sets the optional parameter
// "requestMetadata.partnersSessionId": Google Partners session ID.
func (c *CompaniesGetCall) RequestMetadataPartnersSessionId(requestMetadataPartnersSessionId string) *CompaniesGetCall {
	c.urlParams_.Set("requestMetadata.partnersSessionId", requestMetadataPartnersSessionId)
	return c
}

// RequestMetadataTrafficSourceTrafficSourceId sets the optional
// parameter "requestMetadata.trafficSource.trafficSourceId": Identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *CompaniesGetCall) RequestMetadataTrafficSourceTrafficSourceId(requestMetadataTrafficSourceTrafficSourceId string) *CompaniesGetCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSourceId", requestMetadataTrafficSourceTrafficSourceId)
	return c
}

// RequestMetadataTrafficSourceTrafficSubId sets the optional parameter
// "requestMetadata.trafficSource.trafficSubId": Second level identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *CompaniesGetCall) RequestMetadataTrafficSourceTrafficSubId(requestMetadataTrafficSourceTrafficSubId string) *CompaniesGetCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSubId", requestMetadataTrafficSourceTrafficSubId)
	return c
}

// RequestMetadataUserOverridesIpAddress sets the optional parameter
// "requestMetadata.userOverrides.ipAddress": IP address to use instead
// of the user's geo-located IP address.
func (c *CompaniesGetCall) RequestMetadataUserOverridesIpAddress(requestMetadataUserOverridesIpAddress string) *CompaniesGetCall {
	c.urlParams_.Set("requestMetadata.userOverrides.ipAddress", requestMetadataUserOverridesIpAddress)
	return c
}

// RequestMetadataUserOverridesUserId sets the optional parameter
// "requestMetadata.userOverrides.userId": Logged-in user ID to
// impersonate instead of the user's ID.
func (c *CompaniesGetCall) RequestMetadataUserOverridesUserId(requestMetadataUserOverridesUserId string) *CompaniesGetCall {
	c.urlParams_.Set("requestMetadata.userOverrides.userId", requestMetadataUserOverridesUserId)
	return c
}

// View sets the optional parameter "view": The view of `Company`
// resource to be returned. This must not be
// `COMPANY_VIEW_UNSPECIFIED`.
//
// Possible values:
//   "COMPANY_VIEW_UNSPECIFIED"
//   "CV_GOOGLE_PARTNER_SEARCH"
func (c *CompaniesGetCall) View(view string) *CompaniesGetCall {
	c.urlParams_.Set("view", view)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CompaniesGetCall) Fields(s ...googleapi.Field) *CompaniesGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CompaniesGetCall) IfNoneMatch(entityTag string) *CompaniesGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CompaniesGetCall) Context(ctx context.Context) *CompaniesGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CompaniesGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CompaniesGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/companies/{companyId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"companyId": c.companyId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.companies.get" call.
// Exactly one of *GetCompanyResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *GetCompanyResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CompaniesGetCall) Do(opts ...googleapi.CallOption) (*GetCompanyResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GetCompanyResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a company.",
	//   "flatPath": "v2/companies/{companyId}",
	//   "httpMethod": "GET",
	//   "id": "partners.companies.get",
	//   "parameterOrder": [
	//     "companyId"
	//   ],
	//   "parameters": {
	//     "address": {
	//       "description": "The address to use for sorting the company's addresses by proximity.\nIf not given, the geo-located address of the request is used.\nUsed when order_by is set.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "companyId": {
	//       "description": "The ID of the company to retrieve.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "currencyCode": {
	//       "description": "If the company's budget is in a different currency code than this one, then\nthe converted budget is converted to this currency code.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "How to order addresses within the returned company. Currently, only\n`address` and `address desc` is supported which will sorted by closest to\nfarthest in distance from given address and farthest to closest distance\nfrom given address respectively.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.experimentIds": {
	//       "description": "Experiment IDs the current request belongs to.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.locale": {
	//       "description": "Locale to use for the current request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.partnersSessionId": {
	//       "description": "Google Partners session ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSourceId": {
	//       "description": "Identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSubId": {
	//       "description": "Second level identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.ipAddress": {
	//       "description": "IP address to use instead of the user's geo-located IP address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.userId": {
	//       "description": "Logged-in user ID to impersonate instead of the user's ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "The view of `Company` resource to be returned. This must not be\n`COMPANY_VIEW_UNSPECIFIED`.",
	//       "enum": [
	//         "COMPANY_VIEW_UNSPECIFIED",
	//         "CV_GOOGLE_PARTNER_SEARCH"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/companies/{companyId}",
	//   "response": {
	//     "$ref": "GetCompanyResponse"
	//   }
	// }

}

// method id "partners.companies.list":

type CompaniesListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists companies.
func (r *CompaniesService) List() *CompaniesListCall {
	c := &CompaniesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Address sets the optional parameter "address": The address to use
// when searching for companies.
// If not given, the geo-located address of the request is used.
func (c *CompaniesListCall) Address(address string) *CompaniesListCall {
	c.urlParams_.Set("address", address)
	return c
}

// CompanyName sets the optional parameter "companyName": Company name
// to search for.
func (c *CompaniesListCall) CompanyName(companyName string) *CompaniesListCall {
	c.urlParams_.Set("companyName", companyName)
	return c
}

// GpsMotivations sets the optional parameter "gpsMotivations": List of
// reasons for using Google Partner Search to get companies.
//
// Possible values:
//   "GPS_MOTIVATION_UNSPECIFIED"
//   "GPSM_HELP_WITH_ADVERTISING"
//   "GPSM_HELP_WITH_WEBSITE"
//   "GPSM_NO_WEBSITE"
func (c *CompaniesListCall) GpsMotivations(gpsMotivations ...string) *CompaniesListCall {
	c.urlParams_.SetMulti("gpsMotivations", append([]string{}, gpsMotivations...))
	return c
}

// Industries sets the optional parameter "industries": List of
// industries the company can help with.
//
// Possible values:
//   "INDUSTRY_UNSPECIFIED"
//   "I_AUTOMOTIVE"
//   "I_BUSINESS_TO_BUSINESS"
//   "I_CONSUMER_PACKAGED_GOODS"
//   "I_EDUCATION"
//   "I_FINANCE"
//   "I_HEALTHCARE"
//   "I_MEDIA_AND_ENTERTAINMENT"
//   "I_RETAIL"
//   "I_TECHNOLOGY"
//   "I_TRAVEL"
func (c *CompaniesListCall) Industries(industries ...string) *CompaniesListCall {
	c.urlParams_.SetMulti("industries", append([]string{}, industries...))
	return c
}

// LanguageCodes sets the optional parameter "languageCodes": List of
// language codes that company can support. Only primary
// language
// subtags are accepted as defined by
// <a href="https://tools.ietf.org/html/bcp47">BCP 47</a>
// (IETF BCP 47, "Tags for Identifying Languages").
func (c *CompaniesListCall) LanguageCodes(languageCodes ...string) *CompaniesListCall {
	c.urlParams_.SetMulti("languageCodes", append([]string{}, languageCodes...))
	return c
}

// MaxMonthlyBudgetCurrencyCode sets the optional parameter
// "maxMonthlyBudget.currencyCode": The 3-letter currency code defined
// in ISO 4217.
func (c *CompaniesListCall) MaxMonthlyBudgetCurrencyCode(maxMonthlyBudgetCurrencyCode string) *CompaniesListCall {
	c.urlParams_.Set("maxMonthlyBudget.currencyCode", maxMonthlyBudgetCurrencyCode)
	return c
}

// MaxMonthlyBudgetNanos sets the optional parameter
// "maxMonthlyBudget.nanos": Number of nano (10^-9) units of the
// amount.
// The value must be between -999,999,999 and +999,999,999 inclusive.
// If `units` is positive, `nanos` must be positive or zero.
// If `units` is zero, `nanos` can be positive, zero, or negative.
// If `units` is negative, `nanos` must be negative or zero.
// For example $-1.75 is represented as `units`=-1 and
// `nanos`=-750,000,000.
func (c *CompaniesListCall) MaxMonthlyBudgetNanos(maxMonthlyBudgetNanos int64) *CompaniesListCall {
	c.urlParams_.Set("maxMonthlyBudget.nanos", fmt.Sprint(maxMonthlyBudgetNanos))
	return c
}

// MaxMonthlyBudgetUnits sets the optional parameter
// "maxMonthlyBudget.units": The whole units of the amount.
// For example if `currencyCode` is "USD", then 1 unit is one US
// dollar.
func (c *CompaniesListCall) MaxMonthlyBudgetUnits(maxMonthlyBudgetUnits int64) *CompaniesListCall {
	c.urlParams_.Set("maxMonthlyBudget.units", fmt.Sprint(maxMonthlyBudgetUnits))
	return c
}

// MinMonthlyBudgetCurrencyCode sets the optional parameter
// "minMonthlyBudget.currencyCode": The 3-letter currency code defined
// in ISO 4217.
func (c *CompaniesListCall) MinMonthlyBudgetCurrencyCode(minMonthlyBudgetCurrencyCode string) *CompaniesListCall {
	c.urlParams_.Set("minMonthlyBudget.currencyCode", minMonthlyBudgetCurrencyCode)
	return c
}

// MinMonthlyBudgetNanos sets the optional parameter
// "minMonthlyBudget.nanos": Number of nano (10^-9) units of the
// amount.
// The value must be between -999,999,999 and +999,999,999 inclusive.
// If `units` is positive, `nanos` must be positive or zero.
// If `units` is zero, `nanos` can be positive, zero, or negative.
// If `units` is negative, `nanos` must be negative or zero.
// For example $-1.75 is represented as `units`=-1 and
// `nanos`=-750,000,000.
func (c *CompaniesListCall) MinMonthlyBudgetNanos(minMonthlyBudgetNanos int64) *CompaniesListCall {
	c.urlParams_.Set("minMonthlyBudget.nanos", fmt.Sprint(minMonthlyBudgetNanos))
	return c
}

// MinMonthlyBudgetUnits sets the optional parameter
// "minMonthlyBudget.units": The whole units of the amount.
// For example if `currencyCode` is "USD", then 1 unit is one US
// dollar.
func (c *CompaniesListCall) MinMonthlyBudgetUnits(minMonthlyBudgetUnits int64) *CompaniesListCall {
	c.urlParams_.Set("minMonthlyBudget.units", fmt.Sprint(minMonthlyBudgetUnits))
	return c
}

// OrderBy sets the optional parameter "orderBy": How to order addresses
// within the returned companies. Currently, only
// `address` and `address desc` is supported which will sorted by
// closest to
// farthest in distance from given address and farthest to closest
// distance
// from given address respectively.
func (c *CompaniesListCall) OrderBy(orderBy string) *CompaniesListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// Server may return fewer companies than requested.
// If unspecified, server picks an appropriate default.
func (c *CompaniesListCall) PageSize(pageSize int64) *CompaniesListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results that the server returns.
// Typically, this is the value of
// `ListCompaniesResponse.next_page_token`
// returned from the previous call to
// ListCompanies.
func (c *CompaniesListCall) PageToken(pageToken string) *CompaniesListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// RequestMetadataExperimentIds sets the optional parameter
// "requestMetadata.experimentIds": Experiment IDs the current request
// belongs to.
func (c *CompaniesListCall) RequestMetadataExperimentIds(requestMetadataExperimentIds ...string) *CompaniesListCall {
	c.urlParams_.SetMulti("requestMetadata.experimentIds", append([]string{}, requestMetadataExperimentIds...))
	return c
}

// RequestMetadataLocale sets the optional parameter
// "requestMetadata.locale": Locale to use for the current request.
func (c *CompaniesListCall) RequestMetadataLocale(requestMetadataLocale string) *CompaniesListCall {
	c.urlParams_.Set("requestMetadata.locale", requestMetadataLocale)
	return c
}

// RequestMetadataPartnersSessionId sets the optional parameter
// "requestMetadata.partnersSessionId": Google Partners session ID.
func (c *CompaniesListCall) RequestMetadataPartnersSessionId(requestMetadataPartnersSessionId string) *CompaniesListCall {
	c.urlParams_.Set("requestMetadata.partnersSessionId", requestMetadataPartnersSessionId)
	return c
}

// RequestMetadataTrafficSourceTrafficSourceId sets the optional
// parameter "requestMetadata.trafficSource.trafficSourceId": Identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *CompaniesListCall) RequestMetadataTrafficSourceTrafficSourceId(requestMetadataTrafficSourceTrafficSourceId string) *CompaniesListCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSourceId", requestMetadataTrafficSourceTrafficSourceId)
	return c
}

// RequestMetadataTrafficSourceTrafficSubId sets the optional parameter
// "requestMetadata.trafficSource.trafficSubId": Second level identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *CompaniesListCall) RequestMetadataTrafficSourceTrafficSubId(requestMetadataTrafficSourceTrafficSubId string) *CompaniesListCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSubId", requestMetadataTrafficSourceTrafficSubId)
	return c
}

// RequestMetadataUserOverridesIpAddress sets the optional parameter
// "requestMetadata.userOverrides.ipAddress": IP address to use instead
// of the user's geo-located IP address.
func (c *CompaniesListCall) RequestMetadataUserOverridesIpAddress(requestMetadataUserOverridesIpAddress string) *CompaniesListCall {
	c.urlParams_.Set("requestMetadata.userOverrides.ipAddress", requestMetadataUserOverridesIpAddress)
	return c
}

// RequestMetadataUserOverridesUserId sets the optional parameter
// "requestMetadata.userOverrides.userId": Logged-in user ID to
// impersonate instead of the user's ID.
func (c *CompaniesListCall) RequestMetadataUserOverridesUserId(requestMetadataUserOverridesUserId string) *CompaniesListCall {
	c.urlParams_.Set("requestMetadata.userOverrides.userId", requestMetadataUserOverridesUserId)
	return c
}

// Services sets the optional parameter "services": List of services
// that the returned agencies should provide. If this is
// not empty, any returned agency must have at least one of these
// services,
// or one of the specializations in the "specializations" field.
//
// Possible values:
//   "SERVICE_UNSPECIFIED"
//   "S_ADVANCED_ADWORDS_SUPPORT"
//   "S_ADVERTISING_ON_GOOGLE"
//   "S_AN_ENHANCED_WEBSITE"
//   "S_AN_ONLINE_MARKETING_PLAN"
//   "S_MOBILE_AND_VIDEO_ADS"
//   "S_MOBILE_WEBSITE_SERVICES"
func (c *CompaniesListCall) Services(services ...string) *CompaniesListCall {
	c.urlParams_.SetMulti("services", append([]string{}, services...))
	return c
}

// Specializations sets the optional parameter "specializations": List
// of specializations that the returned agencies should provide. If
// this
// is not empty, any returned agency must have at least one of
// these
// specializations, or one of the services in the "services" field.
//
// Possible values:
//   "BADGE_SPECIALIZATION_UNKNOWN"
//   "BADGE_SPECIALIZATION_ADWORDS_SEARCH"
//   "BADGE_SPECIALIZATION_ADWORDS_DISPLAY"
//   "BADGE_SPECIALIZATION_ADWORDS_MOBILE"
//   "BADGE_SPECIALIZATION_ADWORDS_VIDEO"
//   "BADGE_SPECIALIZATION_ADWORDS_SHOPPING"
func (c *CompaniesListCall) Specializations(specializations ...string) *CompaniesListCall {
	c.urlParams_.SetMulti("specializations", append([]string{}, specializations...))
	return c
}

// View sets the optional parameter "view": The view of the `Company`
// resource to be returned. This must not be
// `COMPANY_VIEW_UNSPECIFIED`.
//
// Possible values:
//   "COMPANY_VIEW_UNSPECIFIED"
//   "CV_GOOGLE_PARTNER_SEARCH"
func (c *CompaniesListCall) View(view string) *CompaniesListCall {
	c.urlParams_.Set("view", view)
	return c
}

// WebsiteUrl sets the optional parameter "websiteUrl": Website URL that
// will help to find a better matched company.
// .
func (c *CompaniesListCall) WebsiteUrl(websiteUrl string) *CompaniesListCall {
	c.urlParams_.Set("websiteUrl", websiteUrl)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CompaniesListCall) Fields(s ...googleapi.Field) *CompaniesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *CompaniesListCall) IfNoneMatch(entityTag string) *CompaniesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CompaniesListCall) Context(ctx context.Context) *CompaniesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CompaniesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CompaniesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/companies")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.companies.list" call.
// Exactly one of *ListCompaniesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListCompaniesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CompaniesListCall) Do(opts ...googleapi.CallOption) (*ListCompaniesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListCompaniesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists companies.",
	//   "flatPath": "v2/companies",
	//   "httpMethod": "GET",
	//   "id": "partners.companies.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "address": {
	//       "description": "The address to use when searching for companies.\nIf not given, the geo-located address of the request is used.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "companyName": {
	//       "description": "Company name to search for.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "gpsMotivations": {
	//       "description": "List of reasons for using Google Partner Search to get companies.",
	//       "enum": [
	//         "GPS_MOTIVATION_UNSPECIFIED",
	//         "GPSM_HELP_WITH_ADVERTISING",
	//         "GPSM_HELP_WITH_WEBSITE",
	//         "GPSM_NO_WEBSITE"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "industries": {
	//       "description": "List of industries the company can help with.",
	//       "enum": [
	//         "INDUSTRY_UNSPECIFIED",
	//         "I_AUTOMOTIVE",
	//         "I_BUSINESS_TO_BUSINESS",
	//         "I_CONSUMER_PACKAGED_GOODS",
	//         "I_EDUCATION",
	//         "I_FINANCE",
	//         "I_HEALTHCARE",
	//         "I_MEDIA_AND_ENTERTAINMENT",
	//         "I_RETAIL",
	//         "I_TECHNOLOGY",
	//         "I_TRAVEL"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "languageCodes": {
	//       "description": "List of language codes that company can support. Only primary language\nsubtags are accepted as defined by\n\u003ca href=\"https://tools.ietf.org/html/bcp47\"\u003eBCP 47\u003c/a\u003e\n(IETF BCP 47, \"Tags for Identifying Languages\").",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "maxMonthlyBudget.currencyCode": {
	//       "description": "The 3-letter currency code defined in ISO 4217.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxMonthlyBudget.nanos": {
	//       "description": "Number of nano (10^-9) units of the amount.\nThe value must be between -999,999,999 and +999,999,999 inclusive.\nIf `units` is positive, `nanos` must be positive or zero.\nIf `units` is zero, `nanos` can be positive, zero, or negative.\nIf `units` is negative, `nanos` must be negative or zero.\nFor example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "maxMonthlyBudget.units": {
	//       "description": "The whole units of the amount.\nFor example if `currencyCode` is `\"USD\"`, then 1 unit is one US dollar.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "minMonthlyBudget.currencyCode": {
	//       "description": "The 3-letter currency code defined in ISO 4217.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "minMonthlyBudget.nanos": {
	//       "description": "Number of nano (10^-9) units of the amount.\nThe value must be between -999,999,999 and +999,999,999 inclusive.\nIf `units` is positive, `nanos` must be positive or zero.\nIf `units` is zero, `nanos` can be positive, zero, or negative.\nIf `units` is negative, `nanos` must be negative or zero.\nFor example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "minMonthlyBudget.units": {
	//       "description": "The whole units of the amount.\nFor example if `currencyCode` is `\"USD\"`, then 1 unit is one US dollar.",
	//       "format": "int64",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "How to order addresses within the returned companies. Currently, only\n`address` and `address desc` is supported which will sorted by closest to\nfarthest in distance from given address and farthest to closest distance\nfrom given address respectively.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. Server may return fewer companies than requested.\nIf unspecified, server picks an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results that the server returns.\nTypically, this is the value of `ListCompaniesResponse.next_page_token`\nreturned from the previous call to\nListCompanies.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.experimentIds": {
	//       "description": "Experiment IDs the current request belongs to.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.locale": {
	//       "description": "Locale to use for the current request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.partnersSessionId": {
	//       "description": "Google Partners session ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSourceId": {
	//       "description": "Identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSubId": {
	//       "description": "Second level identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.ipAddress": {
	//       "description": "IP address to use instead of the user's geo-located IP address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.userId": {
	//       "description": "Logged-in user ID to impersonate instead of the user's ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "services": {
	//       "description": "List of services that the returned agencies should provide. If this is\nnot empty, any returned agency must have at least one of these services,\nor one of the specializations in the \"specializations\" field.",
	//       "enum": [
	//         "SERVICE_UNSPECIFIED",
	//         "S_ADVANCED_ADWORDS_SUPPORT",
	//         "S_ADVERTISING_ON_GOOGLE",
	//         "S_AN_ENHANCED_WEBSITE",
	//         "S_AN_ONLINE_MARKETING_PLAN",
	//         "S_MOBILE_AND_VIDEO_ADS",
	//         "S_MOBILE_WEBSITE_SERVICES"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "specializations": {
	//       "description": "List of specializations that the returned agencies should provide. If this\nis not empty, any returned agency must have at least one of these\nspecializations, or one of the services in the \"services\" field.",
	//       "enum": [
	//         "BADGE_SPECIALIZATION_UNKNOWN",
	//         "BADGE_SPECIALIZATION_ADWORDS_SEARCH",
	//         "BADGE_SPECIALIZATION_ADWORDS_DISPLAY",
	//         "BADGE_SPECIALIZATION_ADWORDS_MOBILE",
	//         "BADGE_SPECIALIZATION_ADWORDS_VIDEO",
	//         "BADGE_SPECIALIZATION_ADWORDS_SHOPPING"
	//       ],
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "view": {
	//       "description": "The view of the `Company` resource to be returned. This must not be\n`COMPANY_VIEW_UNSPECIFIED`.",
	//       "enum": [
	//         "COMPANY_VIEW_UNSPECIFIED",
	//         "CV_GOOGLE_PARTNER_SEARCH"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "websiteUrl": {
	//       "description": "Website URL that will help to find a better matched company.\n.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/companies",
	//   "response": {
	//     "$ref": "ListCompaniesResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *CompaniesListCall) Pages(ctx context.Context, f func(*ListCompaniesResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "partners.companies.leads.create":

type CompaniesLeadsCreateCall struct {
	s                 *Service
	companyId         string
	createleadrequest *CreateLeadRequest
	urlParams_        gensupport.URLParams
	ctx_              context.Context
	header_           http.Header
}

// Create: Creates an advertiser lead for the given company ID.
func (r *CompaniesLeadsService) Create(companyId string, createleadrequest *CreateLeadRequest) *CompaniesLeadsCreateCall {
	c := &CompaniesLeadsCreateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.companyId = companyId
	c.createleadrequest = createleadrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *CompaniesLeadsCreateCall) Fields(s ...googleapi.Field) *CompaniesLeadsCreateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *CompaniesLeadsCreateCall) Context(ctx context.Context) *CompaniesLeadsCreateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *CompaniesLeadsCreateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *CompaniesLeadsCreateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.createleadrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/companies/{companyId}/leads")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"companyId": c.companyId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.companies.leads.create" call.
// Exactly one of *CreateLeadResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *CreateLeadResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *CompaniesLeadsCreateCall) Do(opts ...googleapi.CallOption) (*CreateLeadResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CreateLeadResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates an advertiser lead for the given company ID.",
	//   "flatPath": "v2/companies/{companyId}/leads",
	//   "httpMethod": "POST",
	//   "id": "partners.companies.leads.create",
	//   "parameterOrder": [
	//     "companyId"
	//   ],
	//   "parameters": {
	//     "companyId": {
	//       "description": "The ID of the company to contact.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/companies/{companyId}/leads",
	//   "request": {
	//     "$ref": "CreateLeadRequest"
	//   },
	//   "response": {
	//     "$ref": "CreateLeadResponse"
	//   }
	// }

}

// method id "partners.exams.getToken":

type ExamsGetTokenCall struct {
	s            *Service
	examType     string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetToken: Gets an Exam Token for a Partner's user to take an exam in
// the Exams System
func (r *ExamsService) GetToken(examType string) *ExamsGetTokenCall {
	c := &ExamsGetTokenCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.examType = examType
	return c
}

// RequestMetadataExperimentIds sets the optional parameter
// "requestMetadata.experimentIds": Experiment IDs the current request
// belongs to.
func (c *ExamsGetTokenCall) RequestMetadataExperimentIds(requestMetadataExperimentIds ...string) *ExamsGetTokenCall {
	c.urlParams_.SetMulti("requestMetadata.experimentIds", append([]string{}, requestMetadataExperimentIds...))
	return c
}

// RequestMetadataLocale sets the optional parameter
// "requestMetadata.locale": Locale to use for the current request.
func (c *ExamsGetTokenCall) RequestMetadataLocale(requestMetadataLocale string) *ExamsGetTokenCall {
	c.urlParams_.Set("requestMetadata.locale", requestMetadataLocale)
	return c
}

// RequestMetadataPartnersSessionId sets the optional parameter
// "requestMetadata.partnersSessionId": Google Partners session ID.
func (c *ExamsGetTokenCall) RequestMetadataPartnersSessionId(requestMetadataPartnersSessionId string) *ExamsGetTokenCall {
	c.urlParams_.Set("requestMetadata.partnersSessionId", requestMetadataPartnersSessionId)
	return c
}

// RequestMetadataTrafficSourceTrafficSourceId sets the optional
// parameter "requestMetadata.trafficSource.trafficSourceId": Identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *ExamsGetTokenCall) RequestMetadataTrafficSourceTrafficSourceId(requestMetadataTrafficSourceTrafficSourceId string) *ExamsGetTokenCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSourceId", requestMetadataTrafficSourceTrafficSourceId)
	return c
}

// RequestMetadataTrafficSourceTrafficSubId sets the optional parameter
// "requestMetadata.trafficSource.trafficSubId": Second level identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *ExamsGetTokenCall) RequestMetadataTrafficSourceTrafficSubId(requestMetadataTrafficSourceTrafficSubId string) *ExamsGetTokenCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSubId", requestMetadataTrafficSourceTrafficSubId)
	return c
}

// RequestMetadataUserOverridesIpAddress sets the optional parameter
// "requestMetadata.userOverrides.ipAddress": IP address to use instead
// of the user's geo-located IP address.
func (c *ExamsGetTokenCall) RequestMetadataUserOverridesIpAddress(requestMetadataUserOverridesIpAddress string) *ExamsGetTokenCall {
	c.urlParams_.Set("requestMetadata.userOverrides.ipAddress", requestMetadataUserOverridesIpAddress)
	return c
}

// RequestMetadataUserOverridesUserId sets the optional parameter
// "requestMetadata.userOverrides.userId": Logged-in user ID to
// impersonate instead of the user's ID.
func (c *ExamsGetTokenCall) RequestMetadataUserOverridesUserId(requestMetadataUserOverridesUserId string) *ExamsGetTokenCall {
	c.urlParams_.Set("requestMetadata.userOverrides.userId", requestMetadataUserOverridesUserId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ExamsGetTokenCall) Fields(s ...googleapi.Field) *ExamsGetTokenCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ExamsGetTokenCall) IfNoneMatch(entityTag string) *ExamsGetTokenCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ExamsGetTokenCall) Context(ctx context.Context) *ExamsGetTokenCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ExamsGetTokenCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ExamsGetTokenCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/exams/{examType}/token")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"examType": c.examType,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.exams.getToken" call.
// Exactly one of *ExamToken or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *ExamToken.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *ExamsGetTokenCall) Do(opts ...googleapi.CallOption) (*ExamToken, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ExamToken{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets an Exam Token for a Partner's user to take an exam in the Exams System",
	//   "flatPath": "v2/exams/{examType}/token",
	//   "httpMethod": "GET",
	//   "id": "partners.exams.getToken",
	//   "parameterOrder": [
	//     "examType"
	//   ],
	//   "parameters": {
	//     "examType": {
	//       "description": "The exam type we are requesting a token for.",
	//       "enum": [
	//         "CERTIFICATION_EXAM_TYPE_UNSPECIFIED",
	//         "CET_ADWORDS_FUNDAMENTALS",
	//         "CET_ADWORDS_ADVANCED_SEARCH",
	//         "CET_ADWORDS_ADVANCED_DISPLAY",
	//         "CET_VIDEO_ADS",
	//         "CET_DOUBLECLICK",
	//         "CET_ANALYTICS",
	//         "CET_SHOPPING",
	//         "CET_MOBILE",
	//         "CET_DIGITAL_SALES",
	//         "CET_MOBILE_SITES"
	//       ],
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.experimentIds": {
	//       "description": "Experiment IDs the current request belongs to.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.locale": {
	//       "description": "Locale to use for the current request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.partnersSessionId": {
	//       "description": "Google Partners session ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSourceId": {
	//       "description": "Identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSubId": {
	//       "description": "Second level identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.ipAddress": {
	//       "description": "IP address to use instead of the user's geo-located IP address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.userId": {
	//       "description": "Logged-in user ID to impersonate instead of the user's ID.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/exams/{examType}/token",
	//   "response": {
	//     "$ref": "ExamToken"
	//   }
	// }

}

// method id "partners.leads.list":

type LeadsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists advertiser leads for a user's associated company.
// Should only be called within the context of an authorized logged in
// user.
func (r *LeadsService) List() *LeadsListCall {
	c := &LeadsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// OrderBy sets the optional parameter "orderBy": How to order Leads.
// Currently, only `create_time`
// and `create_time desc` are supported
func (c *LeadsListCall) OrderBy(orderBy string) *LeadsListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": Requested page size.
// Server may return fewer leads than requested.
// If unspecified, server picks an appropriate default.
func (c *LeadsListCall) PageSize(pageSize int64) *LeadsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": A token
// identifying a page of results that the server returns.
// Typically, this is the value of
// `ListLeadsResponse.next_page_token`
// returned from the previous call to
// ListLeads.
func (c *LeadsListCall) PageToken(pageToken string) *LeadsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// RequestMetadataExperimentIds sets the optional parameter
// "requestMetadata.experimentIds": Experiment IDs the current request
// belongs to.
func (c *LeadsListCall) RequestMetadataExperimentIds(requestMetadataExperimentIds ...string) *LeadsListCall {
	c.urlParams_.SetMulti("requestMetadata.experimentIds", append([]string{}, requestMetadataExperimentIds...))
	return c
}

// RequestMetadataLocale sets the optional parameter
// "requestMetadata.locale": Locale to use for the current request.
func (c *LeadsListCall) RequestMetadataLocale(requestMetadataLocale string) *LeadsListCall {
	c.urlParams_.Set("requestMetadata.locale", requestMetadataLocale)
	return c
}

// RequestMetadataPartnersSessionId sets the optional parameter
// "requestMetadata.partnersSessionId": Google Partners session ID.
func (c *LeadsListCall) RequestMetadataPartnersSessionId(requestMetadataPartnersSessionId string) *LeadsListCall {
	c.urlParams_.Set("requestMetadata.partnersSessionId", requestMetadataPartnersSessionId)
	return c
}

// RequestMetadataTrafficSourceTrafficSourceId sets the optional
// parameter "requestMetadata.trafficSource.trafficSourceId": Identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *LeadsListCall) RequestMetadataTrafficSourceTrafficSourceId(requestMetadataTrafficSourceTrafficSourceId string) *LeadsListCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSourceId", requestMetadataTrafficSourceTrafficSourceId)
	return c
}

// RequestMetadataTrafficSourceTrafficSubId sets the optional parameter
// "requestMetadata.trafficSource.trafficSubId": Second level identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *LeadsListCall) RequestMetadataTrafficSourceTrafficSubId(requestMetadataTrafficSourceTrafficSubId string) *LeadsListCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSubId", requestMetadataTrafficSourceTrafficSubId)
	return c
}

// RequestMetadataUserOverridesIpAddress sets the optional parameter
// "requestMetadata.userOverrides.ipAddress": IP address to use instead
// of the user's geo-located IP address.
func (c *LeadsListCall) RequestMetadataUserOverridesIpAddress(requestMetadataUserOverridesIpAddress string) *LeadsListCall {
	c.urlParams_.Set("requestMetadata.userOverrides.ipAddress", requestMetadataUserOverridesIpAddress)
	return c
}

// RequestMetadataUserOverridesUserId sets the optional parameter
// "requestMetadata.userOverrides.userId": Logged-in user ID to
// impersonate instead of the user's ID.
func (c *LeadsListCall) RequestMetadataUserOverridesUserId(requestMetadataUserOverridesUserId string) *LeadsListCall {
	c.urlParams_.Set("requestMetadata.userOverrides.userId", requestMetadataUserOverridesUserId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *LeadsListCall) Fields(s ...googleapi.Field) *LeadsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *LeadsListCall) IfNoneMatch(entityTag string) *LeadsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *LeadsListCall) Context(ctx context.Context) *LeadsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *LeadsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *LeadsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/leads")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.leads.list" call.
// Exactly one of *ListLeadsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListLeadsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *LeadsListCall) Do(opts ...googleapi.CallOption) (*ListLeadsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListLeadsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists advertiser leads for a user's associated company.\nShould only be called within the context of an authorized logged in user.",
	//   "flatPath": "v2/leads",
	//   "httpMethod": "GET",
	//   "id": "partners.leads.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "orderBy": {
	//       "description": "How to order Leads. Currently, only `create_time`\nand `create_time desc` are supported",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Requested page size. Server may return fewer leads than requested.\nIf unspecified, server picks an appropriate default.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A token identifying a page of results that the server returns.\nTypically, this is the value of `ListLeadsResponse.next_page_token`\nreturned from the previous call to\nListLeads.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.experimentIds": {
	//       "description": "Experiment IDs the current request belongs to.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.locale": {
	//       "description": "Locale to use for the current request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.partnersSessionId": {
	//       "description": "Google Partners session ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSourceId": {
	//       "description": "Identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSubId": {
	//       "description": "Second level identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.ipAddress": {
	//       "description": "IP address to use instead of the user's geo-located IP address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.userId": {
	//       "description": "Logged-in user ID to impersonate instead of the user's ID.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/leads",
	//   "response": {
	//     "$ref": "ListLeadsResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *LeadsListCall) Pages(ctx context.Context, f func(*ListLeadsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "partners.offers.list":

type OffersListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the Offers available for the current user
func (r *OffersService) List() *OffersListCall {
	c := &OffersListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// RequestMetadataExperimentIds sets the optional parameter
// "requestMetadata.experimentIds": Experiment IDs the current request
// belongs to.
func (c *OffersListCall) RequestMetadataExperimentIds(requestMetadataExperimentIds ...string) *OffersListCall {
	c.urlParams_.SetMulti("requestMetadata.experimentIds", append([]string{}, requestMetadataExperimentIds...))
	return c
}

// RequestMetadataLocale sets the optional parameter
// "requestMetadata.locale": Locale to use for the current request.
func (c *OffersListCall) RequestMetadataLocale(requestMetadataLocale string) *OffersListCall {
	c.urlParams_.Set("requestMetadata.locale", requestMetadataLocale)
	return c
}

// RequestMetadataPartnersSessionId sets the optional parameter
// "requestMetadata.partnersSessionId": Google Partners session ID.
func (c *OffersListCall) RequestMetadataPartnersSessionId(requestMetadataPartnersSessionId string) *OffersListCall {
	c.urlParams_.Set("requestMetadata.partnersSessionId", requestMetadataPartnersSessionId)
	return c
}

// RequestMetadataTrafficSourceTrafficSourceId sets the optional
// parameter "requestMetadata.trafficSource.trafficSourceId": Identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *OffersListCall) RequestMetadataTrafficSourceTrafficSourceId(requestMetadataTrafficSourceTrafficSourceId string) *OffersListCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSourceId", requestMetadataTrafficSourceTrafficSourceId)
	return c
}

// RequestMetadataTrafficSourceTrafficSubId sets the optional parameter
// "requestMetadata.trafficSource.trafficSubId": Second level identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *OffersListCall) RequestMetadataTrafficSourceTrafficSubId(requestMetadataTrafficSourceTrafficSubId string) *OffersListCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSubId", requestMetadataTrafficSourceTrafficSubId)
	return c
}

// RequestMetadataUserOverridesIpAddress sets the optional parameter
// "requestMetadata.userOverrides.ipAddress": IP address to use instead
// of the user's geo-located IP address.
func (c *OffersListCall) RequestMetadataUserOverridesIpAddress(requestMetadataUserOverridesIpAddress string) *OffersListCall {
	c.urlParams_.Set("requestMetadata.userOverrides.ipAddress", requestMetadataUserOverridesIpAddress)
	return c
}

// RequestMetadataUserOverridesUserId sets the optional parameter
// "requestMetadata.userOverrides.userId": Logged-in user ID to
// impersonate instead of the user's ID.
func (c *OffersListCall) RequestMetadataUserOverridesUserId(requestMetadataUserOverridesUserId string) *OffersListCall {
	c.urlParams_.Set("requestMetadata.userOverrides.userId", requestMetadataUserOverridesUserId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OffersListCall) Fields(s ...googleapi.Field) *OffersListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OffersListCall) IfNoneMatch(entityTag string) *OffersListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OffersListCall) Context(ctx context.Context) *OffersListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OffersListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OffersListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/offers")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.offers.list" call.
// Exactly one of *ListOffersResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListOffersResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OffersListCall) Do(opts ...googleapi.CallOption) (*ListOffersResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListOffersResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the Offers available for the current user",
	//   "flatPath": "v2/offers",
	//   "httpMethod": "GET",
	//   "id": "partners.offers.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "requestMetadata.experimentIds": {
	//       "description": "Experiment IDs the current request belongs to.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.locale": {
	//       "description": "Locale to use for the current request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.partnersSessionId": {
	//       "description": "Google Partners session ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSourceId": {
	//       "description": "Identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSubId": {
	//       "description": "Second level identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.ipAddress": {
	//       "description": "IP address to use instead of the user's geo-located IP address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.userId": {
	//       "description": "Logged-in user ID to impersonate instead of the user's ID.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/offers",
	//   "response": {
	//     "$ref": "ListOffersResponse"
	//   }
	// }

}

// method id "partners.offers.history.list":

type OffersHistoryListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists the Historical Offers for the current user (or user's
// entire company)
func (r *OffersHistoryService) List() *OffersHistoryListCall {
	c := &OffersHistoryListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// EntireCompany sets the optional parameter "entireCompany": if true,
// show history for the entire company.  Requires user to be admin.
func (c *OffersHistoryListCall) EntireCompany(entireCompany bool) *OffersHistoryListCall {
	c.urlParams_.Set("entireCompany", fmt.Sprint(entireCompany))
	return c
}

// OrderBy sets the optional parameter "orderBy": Comma-separated list
// of fields to order by, e.g.: "foo,bar,baz".
// Use "foo desc" to sort descending.
// List of valid field names is: name, offer_code, expiration_time,
// status,
//     last_modified_time, sender_name, creation_time, country_code,
//     offer_type.
func (c *OffersHistoryListCall) OrderBy(orderBy string) *OffersHistoryListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": Maximum number of
// rows to return per page.
func (c *OffersHistoryListCall) PageSize(pageSize int64) *OffersHistoryListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Token to retrieve
// a specific page.
func (c *OffersHistoryListCall) PageToken(pageToken string) *OffersHistoryListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// RequestMetadataExperimentIds sets the optional parameter
// "requestMetadata.experimentIds": Experiment IDs the current request
// belongs to.
func (c *OffersHistoryListCall) RequestMetadataExperimentIds(requestMetadataExperimentIds ...string) *OffersHistoryListCall {
	c.urlParams_.SetMulti("requestMetadata.experimentIds", append([]string{}, requestMetadataExperimentIds...))
	return c
}

// RequestMetadataLocale sets the optional parameter
// "requestMetadata.locale": Locale to use for the current request.
func (c *OffersHistoryListCall) RequestMetadataLocale(requestMetadataLocale string) *OffersHistoryListCall {
	c.urlParams_.Set("requestMetadata.locale", requestMetadataLocale)
	return c
}

// RequestMetadataPartnersSessionId sets the optional parameter
// "requestMetadata.partnersSessionId": Google Partners session ID.
func (c *OffersHistoryListCall) RequestMetadataPartnersSessionId(requestMetadataPartnersSessionId string) *OffersHistoryListCall {
	c.urlParams_.Set("requestMetadata.partnersSessionId", requestMetadataPartnersSessionId)
	return c
}

// RequestMetadataTrafficSourceTrafficSourceId sets the optional
// parameter "requestMetadata.trafficSource.trafficSourceId": Identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *OffersHistoryListCall) RequestMetadataTrafficSourceTrafficSourceId(requestMetadataTrafficSourceTrafficSourceId string) *OffersHistoryListCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSourceId", requestMetadataTrafficSourceTrafficSourceId)
	return c
}

// RequestMetadataTrafficSourceTrafficSubId sets the optional parameter
// "requestMetadata.trafficSource.trafficSubId": Second level identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *OffersHistoryListCall) RequestMetadataTrafficSourceTrafficSubId(requestMetadataTrafficSourceTrafficSubId string) *OffersHistoryListCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSubId", requestMetadataTrafficSourceTrafficSubId)
	return c
}

// RequestMetadataUserOverridesIpAddress sets the optional parameter
// "requestMetadata.userOverrides.ipAddress": IP address to use instead
// of the user's geo-located IP address.
func (c *OffersHistoryListCall) RequestMetadataUserOverridesIpAddress(requestMetadataUserOverridesIpAddress string) *OffersHistoryListCall {
	c.urlParams_.Set("requestMetadata.userOverrides.ipAddress", requestMetadataUserOverridesIpAddress)
	return c
}

// RequestMetadataUserOverridesUserId sets the optional parameter
// "requestMetadata.userOverrides.userId": Logged-in user ID to
// impersonate instead of the user's ID.
func (c *OffersHistoryListCall) RequestMetadataUserOverridesUserId(requestMetadataUserOverridesUserId string) *OffersHistoryListCall {
	c.urlParams_.Set("requestMetadata.userOverrides.userId", requestMetadataUserOverridesUserId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OffersHistoryListCall) Fields(s ...googleapi.Field) *OffersHistoryListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *OffersHistoryListCall) IfNoneMatch(entityTag string) *OffersHistoryListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *OffersHistoryListCall) Context(ctx context.Context) *OffersHistoryListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *OffersHistoryListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *OffersHistoryListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/offers/history")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.offers.history.list" call.
// Exactly one of *ListOffersHistoryResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *ListOffersHistoryResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *OffersHistoryListCall) Do(opts ...googleapi.CallOption) (*ListOffersHistoryResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListOffersHistoryResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists the Historical Offers for the current user (or user's entire company)",
	//   "flatPath": "v2/offers/history",
	//   "httpMethod": "GET",
	//   "id": "partners.offers.history.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "entireCompany": {
	//       "description": "if true, show history for the entire company.  Requires user to be admin.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "orderBy": {
	//       "description": "Comma-separated list of fields to order by, e.g.: \"foo,bar,baz\".\nUse \"foo desc\" to sort descending.\nList of valid field names is: name, offer_code, expiration_time, status,\n    last_modified_time, sender_name, creation_time, country_code,\n    offer_type.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "Maximum number of rows to return per page.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Token to retrieve a specific page.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.experimentIds": {
	//       "description": "Experiment IDs the current request belongs to.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.locale": {
	//       "description": "Locale to use for the current request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.partnersSessionId": {
	//       "description": "Google Partners session ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSourceId": {
	//       "description": "Identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSubId": {
	//       "description": "Second level identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.ipAddress": {
	//       "description": "IP address to use instead of the user's geo-located IP address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.userId": {
	//       "description": "Logged-in user ID to impersonate instead of the user's ID.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/offers/history",
	//   "response": {
	//     "$ref": "ListOffersHistoryResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *OffersHistoryListCall) Pages(ctx context.Context, f func(*ListOffersHistoryResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "partners.userEvents.log":

type UserEventsLogCall struct {
	s                   *Service
	logusereventrequest *LogUserEventRequest
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// Log: Logs a user event.
func (r *UserEventsService) Log(logusereventrequest *LogUserEventRequest) *UserEventsLogCall {
	c := &UserEventsLogCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.logusereventrequest = logusereventrequest
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserEventsLogCall) Fields(s ...googleapi.Field) *UserEventsLogCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UserEventsLogCall) Context(ctx context.Context) *UserEventsLogCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UserEventsLogCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UserEventsLogCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.logusereventrequest)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/userEvents:log")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.userEvents.log" call.
// Exactly one of *LogUserEventResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *LogUserEventResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *UserEventsLogCall) Do(opts ...googleapi.CallOption) (*LogUserEventResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &LogUserEventResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Logs a user event.",
	//   "flatPath": "v2/userEvents:log",
	//   "httpMethod": "POST",
	//   "id": "partners.userEvents.log",
	//   "parameterOrder": [],
	//   "parameters": {},
	//   "path": "v2/userEvents:log",
	//   "request": {
	//     "$ref": "LogUserEventRequest"
	//   },
	//   "response": {
	//     "$ref": "LogUserEventResponse"
	//   }
	// }

}

// method id "partners.userStates.list":

type UserStatesListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists states for current user.
func (r *UserStatesService) List() *UserStatesListCall {
	c := &UserStatesListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// RequestMetadataExperimentIds sets the optional parameter
// "requestMetadata.experimentIds": Experiment IDs the current request
// belongs to.
func (c *UserStatesListCall) RequestMetadataExperimentIds(requestMetadataExperimentIds ...string) *UserStatesListCall {
	c.urlParams_.SetMulti("requestMetadata.experimentIds", append([]string{}, requestMetadataExperimentIds...))
	return c
}

// RequestMetadataLocale sets the optional parameter
// "requestMetadata.locale": Locale to use for the current request.
func (c *UserStatesListCall) RequestMetadataLocale(requestMetadataLocale string) *UserStatesListCall {
	c.urlParams_.Set("requestMetadata.locale", requestMetadataLocale)
	return c
}

// RequestMetadataPartnersSessionId sets the optional parameter
// "requestMetadata.partnersSessionId": Google Partners session ID.
func (c *UserStatesListCall) RequestMetadataPartnersSessionId(requestMetadataPartnersSessionId string) *UserStatesListCall {
	c.urlParams_.Set("requestMetadata.partnersSessionId", requestMetadataPartnersSessionId)
	return c
}

// RequestMetadataTrafficSourceTrafficSourceId sets the optional
// parameter "requestMetadata.trafficSource.trafficSourceId": Identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *UserStatesListCall) RequestMetadataTrafficSourceTrafficSourceId(requestMetadataTrafficSourceTrafficSourceId string) *UserStatesListCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSourceId", requestMetadataTrafficSourceTrafficSourceId)
	return c
}

// RequestMetadataTrafficSourceTrafficSubId sets the optional parameter
// "requestMetadata.trafficSource.trafficSubId": Second level identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *UserStatesListCall) RequestMetadataTrafficSourceTrafficSubId(requestMetadataTrafficSourceTrafficSubId string) *UserStatesListCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSubId", requestMetadataTrafficSourceTrafficSubId)
	return c
}

// RequestMetadataUserOverridesIpAddress sets the optional parameter
// "requestMetadata.userOverrides.ipAddress": IP address to use instead
// of the user's geo-located IP address.
func (c *UserStatesListCall) RequestMetadataUserOverridesIpAddress(requestMetadataUserOverridesIpAddress string) *UserStatesListCall {
	c.urlParams_.Set("requestMetadata.userOverrides.ipAddress", requestMetadataUserOverridesIpAddress)
	return c
}

// RequestMetadataUserOverridesUserId sets the optional parameter
// "requestMetadata.userOverrides.userId": Logged-in user ID to
// impersonate instead of the user's ID.
func (c *UserStatesListCall) RequestMetadataUserOverridesUserId(requestMetadataUserOverridesUserId string) *UserStatesListCall {
	c.urlParams_.Set("requestMetadata.userOverrides.userId", requestMetadataUserOverridesUserId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UserStatesListCall) Fields(s ...googleapi.Field) *UserStatesListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *UserStatesListCall) IfNoneMatch(entityTag string) *UserStatesListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UserStatesListCall) Context(ctx context.Context) *UserStatesListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UserStatesListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UserStatesListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/userStates")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.userStates.list" call.
// Exactly one of *ListUserStatesResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListUserStatesResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *UserStatesListCall) Do(opts ...googleapi.CallOption) (*ListUserStatesResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListUserStatesResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists states for current user.",
	//   "flatPath": "v2/userStates",
	//   "httpMethod": "GET",
	//   "id": "partners.userStates.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "requestMetadata.experimentIds": {
	//       "description": "Experiment IDs the current request belongs to.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.locale": {
	//       "description": "Locale to use for the current request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.partnersSessionId": {
	//       "description": "Google Partners session ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSourceId": {
	//       "description": "Identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSubId": {
	//       "description": "Second level identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.ipAddress": {
	//       "description": "IP address to use instead of the user's geo-located IP address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.userId": {
	//       "description": "Logged-in user ID to impersonate instead of the user's ID.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/userStates",
	//   "response": {
	//     "$ref": "ListUserStatesResponse"
	//   }
	// }

}

// method id "partners.users.createCompanyRelation":

type UsersCreateCompanyRelationCall struct {
	s               *Service
	userId          string
	companyrelation *CompanyRelation
	urlParams_      gensupport.URLParams
	ctx_            context.Context
	header_         http.Header
}

// CreateCompanyRelation: Creates a user's company relation. Affiliates
// the user to a company.
func (r *UsersService) CreateCompanyRelation(userId string, companyrelation *CompanyRelation) *UsersCreateCompanyRelationCall {
	c := &UsersCreateCompanyRelationCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.userId = userId
	c.companyrelation = companyrelation
	return c
}

// RequestMetadataExperimentIds sets the optional parameter
// "requestMetadata.experimentIds": Experiment IDs the current request
// belongs to.
func (c *UsersCreateCompanyRelationCall) RequestMetadataExperimentIds(requestMetadataExperimentIds ...string) *UsersCreateCompanyRelationCall {
	c.urlParams_.SetMulti("requestMetadata.experimentIds", append([]string{}, requestMetadataExperimentIds...))
	return c
}

// RequestMetadataLocale sets the optional parameter
// "requestMetadata.locale": Locale to use for the current request.
func (c *UsersCreateCompanyRelationCall) RequestMetadataLocale(requestMetadataLocale string) *UsersCreateCompanyRelationCall {
	c.urlParams_.Set("requestMetadata.locale", requestMetadataLocale)
	return c
}

// RequestMetadataPartnersSessionId sets the optional parameter
// "requestMetadata.partnersSessionId": Google Partners session ID.
func (c *UsersCreateCompanyRelationCall) RequestMetadataPartnersSessionId(requestMetadataPartnersSessionId string) *UsersCreateCompanyRelationCall {
	c.urlParams_.Set("requestMetadata.partnersSessionId", requestMetadataPartnersSessionId)
	return c
}

// RequestMetadataTrafficSourceTrafficSourceId sets the optional
// parameter "requestMetadata.trafficSource.trafficSourceId": Identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *UsersCreateCompanyRelationCall) RequestMetadataTrafficSourceTrafficSourceId(requestMetadataTrafficSourceTrafficSourceId string) *UsersCreateCompanyRelationCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSourceId", requestMetadataTrafficSourceTrafficSourceId)
	return c
}

// RequestMetadataTrafficSourceTrafficSubId sets the optional parameter
// "requestMetadata.trafficSource.trafficSubId": Second level identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *UsersCreateCompanyRelationCall) RequestMetadataTrafficSourceTrafficSubId(requestMetadataTrafficSourceTrafficSubId string) *UsersCreateCompanyRelationCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSubId", requestMetadataTrafficSourceTrafficSubId)
	return c
}

// RequestMetadataUserOverridesIpAddress sets the optional parameter
// "requestMetadata.userOverrides.ipAddress": IP address to use instead
// of the user's geo-located IP address.
func (c *UsersCreateCompanyRelationCall) RequestMetadataUserOverridesIpAddress(requestMetadataUserOverridesIpAddress string) *UsersCreateCompanyRelationCall {
	c.urlParams_.Set("requestMetadata.userOverrides.ipAddress", requestMetadataUserOverridesIpAddress)
	return c
}

// RequestMetadataUserOverridesUserId sets the optional parameter
// "requestMetadata.userOverrides.userId": Logged-in user ID to
// impersonate instead of the user's ID.
func (c *UsersCreateCompanyRelationCall) RequestMetadataUserOverridesUserId(requestMetadataUserOverridesUserId string) *UsersCreateCompanyRelationCall {
	c.urlParams_.Set("requestMetadata.userOverrides.userId", requestMetadataUserOverridesUserId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersCreateCompanyRelationCall) Fields(s ...googleapi.Field) *UsersCreateCompanyRelationCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UsersCreateCompanyRelationCall) Context(ctx context.Context) *UsersCreateCompanyRelationCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UsersCreateCompanyRelationCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersCreateCompanyRelationCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.companyrelation)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/users/{userId}/companyRelation")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"userId": c.userId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.users.createCompanyRelation" call.
// Exactly one of *CompanyRelation or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *CompanyRelation.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *UsersCreateCompanyRelationCall) Do(opts ...googleapi.CallOption) (*CompanyRelation, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &CompanyRelation{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a user's company relation. Affiliates the user to a company.",
	//   "flatPath": "v2/users/{userId}/companyRelation",
	//   "httpMethod": "PUT",
	//   "id": "partners.users.createCompanyRelation",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "requestMetadata.experimentIds": {
	//       "description": "Experiment IDs the current request belongs to.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.locale": {
	//       "description": "Locale to use for the current request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.partnersSessionId": {
	//       "description": "Google Partners session ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSourceId": {
	//       "description": "Identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSubId": {
	//       "description": "Second level identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.ipAddress": {
	//       "description": "IP address to use instead of the user's geo-located IP address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.userId": {
	//       "description": "Logged-in user ID to impersonate instead of the user's ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user. Can be set to \u003ccode\u003eme\u003c/code\u003e to mean\nthe currently authenticated user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/users/{userId}/companyRelation",
	//   "request": {
	//     "$ref": "CompanyRelation"
	//   },
	//   "response": {
	//     "$ref": "CompanyRelation"
	//   }
	// }

}

// method id "partners.users.deleteCompanyRelation":

type UsersDeleteCompanyRelationCall struct {
	s          *Service
	userId     string
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// DeleteCompanyRelation: Deletes a user's company relation.
// Unaffiliaites the user from a company.
func (r *UsersService) DeleteCompanyRelation(userId string) *UsersDeleteCompanyRelationCall {
	c := &UsersDeleteCompanyRelationCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.userId = userId
	return c
}

// RequestMetadataExperimentIds sets the optional parameter
// "requestMetadata.experimentIds": Experiment IDs the current request
// belongs to.
func (c *UsersDeleteCompanyRelationCall) RequestMetadataExperimentIds(requestMetadataExperimentIds ...string) *UsersDeleteCompanyRelationCall {
	c.urlParams_.SetMulti("requestMetadata.experimentIds", append([]string{}, requestMetadataExperimentIds...))
	return c
}

// RequestMetadataLocale sets the optional parameter
// "requestMetadata.locale": Locale to use for the current request.
func (c *UsersDeleteCompanyRelationCall) RequestMetadataLocale(requestMetadataLocale string) *UsersDeleteCompanyRelationCall {
	c.urlParams_.Set("requestMetadata.locale", requestMetadataLocale)
	return c
}

// RequestMetadataPartnersSessionId sets the optional parameter
// "requestMetadata.partnersSessionId": Google Partners session ID.
func (c *UsersDeleteCompanyRelationCall) RequestMetadataPartnersSessionId(requestMetadataPartnersSessionId string) *UsersDeleteCompanyRelationCall {
	c.urlParams_.Set("requestMetadata.partnersSessionId", requestMetadataPartnersSessionId)
	return c
}

// RequestMetadataTrafficSourceTrafficSourceId sets the optional
// parameter "requestMetadata.trafficSource.trafficSourceId": Identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *UsersDeleteCompanyRelationCall) RequestMetadataTrafficSourceTrafficSourceId(requestMetadataTrafficSourceTrafficSourceId string) *UsersDeleteCompanyRelationCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSourceId", requestMetadataTrafficSourceTrafficSourceId)
	return c
}

// RequestMetadataTrafficSourceTrafficSubId sets the optional parameter
// "requestMetadata.trafficSource.trafficSubId": Second level identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *UsersDeleteCompanyRelationCall) RequestMetadataTrafficSourceTrafficSubId(requestMetadataTrafficSourceTrafficSubId string) *UsersDeleteCompanyRelationCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSubId", requestMetadataTrafficSourceTrafficSubId)
	return c
}

// RequestMetadataUserOverridesIpAddress sets the optional parameter
// "requestMetadata.userOverrides.ipAddress": IP address to use instead
// of the user's geo-located IP address.
func (c *UsersDeleteCompanyRelationCall) RequestMetadataUserOverridesIpAddress(requestMetadataUserOverridesIpAddress string) *UsersDeleteCompanyRelationCall {
	c.urlParams_.Set("requestMetadata.userOverrides.ipAddress", requestMetadataUserOverridesIpAddress)
	return c
}

// RequestMetadataUserOverridesUserId sets the optional parameter
// "requestMetadata.userOverrides.userId": Logged-in user ID to
// impersonate instead of the user's ID.
func (c *UsersDeleteCompanyRelationCall) RequestMetadataUserOverridesUserId(requestMetadataUserOverridesUserId string) *UsersDeleteCompanyRelationCall {
	c.urlParams_.Set("requestMetadata.userOverrides.userId", requestMetadataUserOverridesUserId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDeleteCompanyRelationCall) Fields(s ...googleapi.Field) *UsersDeleteCompanyRelationCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UsersDeleteCompanyRelationCall) Context(ctx context.Context) *UsersDeleteCompanyRelationCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UsersDeleteCompanyRelationCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersDeleteCompanyRelationCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/users/{userId}/companyRelation")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"userId": c.userId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.users.deleteCompanyRelation" call.
// Exactly one of *Empty or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Empty.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *UsersDeleteCompanyRelationCall) Do(opts ...googleapi.CallOption) (*Empty, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Empty{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a user's company relation. Unaffiliaites the user from a company.",
	//   "flatPath": "v2/users/{userId}/companyRelation",
	//   "httpMethod": "DELETE",
	//   "id": "partners.users.deleteCompanyRelation",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "requestMetadata.experimentIds": {
	//       "description": "Experiment IDs the current request belongs to.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.locale": {
	//       "description": "Locale to use for the current request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.partnersSessionId": {
	//       "description": "Google Partners session ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSourceId": {
	//       "description": "Identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSubId": {
	//       "description": "Second level identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.ipAddress": {
	//       "description": "IP address to use instead of the user's geo-located IP address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.userId": {
	//       "description": "Logged-in user ID to impersonate instead of the user's ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "The ID of the user. Can be set to \u003ccode\u003eme\u003c/code\u003e to mean\nthe currently authenticated user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/users/{userId}/companyRelation",
	//   "response": {
	//     "$ref": "Empty"
	//   }
	// }

}

// method id "partners.users.get":

type UsersGetCall struct {
	s            *Service
	userId       string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Gets a user.
func (r *UsersService) Get(userId string) *UsersGetCall {
	c := &UsersGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.userId = userId
	return c
}

// RequestMetadataExperimentIds sets the optional parameter
// "requestMetadata.experimentIds": Experiment IDs the current request
// belongs to.
func (c *UsersGetCall) RequestMetadataExperimentIds(requestMetadataExperimentIds ...string) *UsersGetCall {
	c.urlParams_.SetMulti("requestMetadata.experimentIds", append([]string{}, requestMetadataExperimentIds...))
	return c
}

// RequestMetadataLocale sets the optional parameter
// "requestMetadata.locale": Locale to use for the current request.
func (c *UsersGetCall) RequestMetadataLocale(requestMetadataLocale string) *UsersGetCall {
	c.urlParams_.Set("requestMetadata.locale", requestMetadataLocale)
	return c
}

// RequestMetadataPartnersSessionId sets the optional parameter
// "requestMetadata.partnersSessionId": Google Partners session ID.
func (c *UsersGetCall) RequestMetadataPartnersSessionId(requestMetadataPartnersSessionId string) *UsersGetCall {
	c.urlParams_.Set("requestMetadata.partnersSessionId", requestMetadataPartnersSessionId)
	return c
}

// RequestMetadataTrafficSourceTrafficSourceId sets the optional
// parameter "requestMetadata.trafficSource.trafficSourceId": Identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *UsersGetCall) RequestMetadataTrafficSourceTrafficSourceId(requestMetadataTrafficSourceTrafficSourceId string) *UsersGetCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSourceId", requestMetadataTrafficSourceTrafficSourceId)
	return c
}

// RequestMetadataTrafficSourceTrafficSubId sets the optional parameter
// "requestMetadata.trafficSource.trafficSubId": Second level identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *UsersGetCall) RequestMetadataTrafficSourceTrafficSubId(requestMetadataTrafficSourceTrafficSubId string) *UsersGetCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSubId", requestMetadataTrafficSourceTrafficSubId)
	return c
}

// RequestMetadataUserOverridesIpAddress sets the optional parameter
// "requestMetadata.userOverrides.ipAddress": IP address to use instead
// of the user's geo-located IP address.
func (c *UsersGetCall) RequestMetadataUserOverridesIpAddress(requestMetadataUserOverridesIpAddress string) *UsersGetCall {
	c.urlParams_.Set("requestMetadata.userOverrides.ipAddress", requestMetadataUserOverridesIpAddress)
	return c
}

// RequestMetadataUserOverridesUserId sets the optional parameter
// "requestMetadata.userOverrides.userId": Logged-in user ID to
// impersonate instead of the user's ID.
func (c *UsersGetCall) RequestMetadataUserOverridesUserId(requestMetadataUserOverridesUserId string) *UsersGetCall {
	c.urlParams_.Set("requestMetadata.userOverrides.userId", requestMetadataUserOverridesUserId)
	return c
}

// UserView sets the optional parameter "userView": Specifies what parts
// of the user information to return.
//
// Possible values:
//   "BASIC"
//   "PROFILE"
//   "PUBLIC_PROFILE"
func (c *UsersGetCall) UserView(userView string) *UsersGetCall {
	c.urlParams_.Set("userView", userView)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersGetCall) Fields(s ...googleapi.Field) *UsersGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *UsersGetCall) IfNoneMatch(entityTag string) *UsersGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UsersGetCall) Context(ctx context.Context) *UsersGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UsersGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/users/{userId}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"userId": c.userId,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.users.get" call.
// Exactly one of *User or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *User.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *UsersGetCall) Do(opts ...googleapi.CallOption) (*User, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &User{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets a user.",
	//   "flatPath": "v2/users/{userId}",
	//   "httpMethod": "GET",
	//   "id": "partners.users.get",
	//   "parameterOrder": [
	//     "userId"
	//   ],
	//   "parameters": {
	//     "requestMetadata.experimentIds": {
	//       "description": "Experiment IDs the current request belongs to.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.locale": {
	//       "description": "Locale to use for the current request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.partnersSessionId": {
	//       "description": "Google Partners session ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSourceId": {
	//       "description": "Identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSubId": {
	//       "description": "Second level identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.ipAddress": {
	//       "description": "IP address to use instead of the user's geo-located IP address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.userId": {
	//       "description": "Logged-in user ID to impersonate instead of the user's ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "userId": {
	//       "description": "Identifier of the user. Can be set to \u003ccode\u003eme\u003c/code\u003e to mean the currently\nauthenticated user.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "userView": {
	//       "description": "Specifies what parts of the user information to return.",
	//       "enum": [
	//         "BASIC",
	//         "PROFILE",
	//         "PUBLIC_PROFILE"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/users/{userId}",
	//   "response": {
	//     "$ref": "User"
	//   }
	// }

}

// method id "partners.users.updateProfile":

type UsersUpdateProfileCall struct {
	s           *Service
	userprofile *UserProfile
	urlParams_  gensupport.URLParams
	ctx_        context.Context
	header_     http.Header
}

// UpdateProfile: Updates a user's profile. A user can only update their
// own profile and
// should only be called within the context of a logged in user.
func (r *UsersService) UpdateProfile(userprofile *UserProfile) *UsersUpdateProfileCall {
	c := &UsersUpdateProfileCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.userprofile = userprofile
	return c
}

// RequestMetadataExperimentIds sets the optional parameter
// "requestMetadata.experimentIds": Experiment IDs the current request
// belongs to.
func (c *UsersUpdateProfileCall) RequestMetadataExperimentIds(requestMetadataExperimentIds ...string) *UsersUpdateProfileCall {
	c.urlParams_.SetMulti("requestMetadata.experimentIds", append([]string{}, requestMetadataExperimentIds...))
	return c
}

// RequestMetadataLocale sets the optional parameter
// "requestMetadata.locale": Locale to use for the current request.
func (c *UsersUpdateProfileCall) RequestMetadataLocale(requestMetadataLocale string) *UsersUpdateProfileCall {
	c.urlParams_.Set("requestMetadata.locale", requestMetadataLocale)
	return c
}

// RequestMetadataPartnersSessionId sets the optional parameter
// "requestMetadata.partnersSessionId": Google Partners session ID.
func (c *UsersUpdateProfileCall) RequestMetadataPartnersSessionId(requestMetadataPartnersSessionId string) *UsersUpdateProfileCall {
	c.urlParams_.Set("requestMetadata.partnersSessionId", requestMetadataPartnersSessionId)
	return c
}

// RequestMetadataTrafficSourceTrafficSourceId sets the optional
// parameter "requestMetadata.trafficSource.trafficSourceId": Identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *UsersUpdateProfileCall) RequestMetadataTrafficSourceTrafficSourceId(requestMetadataTrafficSourceTrafficSourceId string) *UsersUpdateProfileCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSourceId", requestMetadataTrafficSourceTrafficSourceId)
	return c
}

// RequestMetadataTrafficSourceTrafficSubId sets the optional parameter
// "requestMetadata.trafficSource.trafficSubId": Second level identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *UsersUpdateProfileCall) RequestMetadataTrafficSourceTrafficSubId(requestMetadataTrafficSourceTrafficSubId string) *UsersUpdateProfileCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSubId", requestMetadataTrafficSourceTrafficSubId)
	return c
}

// RequestMetadataUserOverridesIpAddress sets the optional parameter
// "requestMetadata.userOverrides.ipAddress": IP address to use instead
// of the user's geo-located IP address.
func (c *UsersUpdateProfileCall) RequestMetadataUserOverridesIpAddress(requestMetadataUserOverridesIpAddress string) *UsersUpdateProfileCall {
	c.urlParams_.Set("requestMetadata.userOverrides.ipAddress", requestMetadataUserOverridesIpAddress)
	return c
}

// RequestMetadataUserOverridesUserId sets the optional parameter
// "requestMetadata.userOverrides.userId": Logged-in user ID to
// impersonate instead of the user's ID.
func (c *UsersUpdateProfileCall) RequestMetadataUserOverridesUserId(requestMetadataUserOverridesUserId string) *UsersUpdateProfileCall {
	c.urlParams_.Set("requestMetadata.userOverrides.userId", requestMetadataUserOverridesUserId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersUpdateProfileCall) Fields(s ...googleapi.Field) *UsersUpdateProfileCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UsersUpdateProfileCall) Context(ctx context.Context) *UsersUpdateProfileCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UsersUpdateProfileCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersUpdateProfileCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.userprofile)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/users/profile")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.users.updateProfile" call.
// Exactly one of *UserProfile or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *UserProfile.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *UsersUpdateProfileCall) Do(opts ...googleapi.CallOption) (*UserProfile, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &UserProfile{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a user's profile. A user can only update their own profile and\nshould only be called within the context of a logged in user.",
	//   "flatPath": "v2/users/profile",
	//   "httpMethod": "PATCH",
	//   "id": "partners.users.updateProfile",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "requestMetadata.experimentIds": {
	//       "description": "Experiment IDs the current request belongs to.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.locale": {
	//       "description": "Locale to use for the current request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.partnersSessionId": {
	//       "description": "Google Partners session ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSourceId": {
	//       "description": "Identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSubId": {
	//       "description": "Second level identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.ipAddress": {
	//       "description": "IP address to use instead of the user's geo-located IP address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.userId": {
	//       "description": "Logged-in user ID to impersonate instead of the user's ID.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/users/profile",
	//   "request": {
	//     "$ref": "UserProfile"
	//   },
	//   "response": {
	//     "$ref": "UserProfile"
	//   }
	// }

}

// method id "partners.getPartnersstatus":

type V2GetPartnersstatusCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetPartnersstatus: Gets Partners Status of the logged in user's
// agency.
// Should only be called if the logged in user is the admin of the
// agency.
func (r *V2Service) GetPartnersstatus() *V2GetPartnersstatusCall {
	c := &V2GetPartnersstatusCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// RequestMetadataExperimentIds sets the optional parameter
// "requestMetadata.experimentIds": Experiment IDs the current request
// belongs to.
func (c *V2GetPartnersstatusCall) RequestMetadataExperimentIds(requestMetadataExperimentIds ...string) *V2GetPartnersstatusCall {
	c.urlParams_.SetMulti("requestMetadata.experimentIds", append([]string{}, requestMetadataExperimentIds...))
	return c
}

// RequestMetadataLocale sets the optional parameter
// "requestMetadata.locale": Locale to use for the current request.
func (c *V2GetPartnersstatusCall) RequestMetadataLocale(requestMetadataLocale string) *V2GetPartnersstatusCall {
	c.urlParams_.Set("requestMetadata.locale", requestMetadataLocale)
	return c
}

// RequestMetadataPartnersSessionId sets the optional parameter
// "requestMetadata.partnersSessionId": Google Partners session ID.
func (c *V2GetPartnersstatusCall) RequestMetadataPartnersSessionId(requestMetadataPartnersSessionId string) *V2GetPartnersstatusCall {
	c.urlParams_.Set("requestMetadata.partnersSessionId", requestMetadataPartnersSessionId)
	return c
}

// RequestMetadataTrafficSourceTrafficSourceId sets the optional
// parameter "requestMetadata.trafficSource.trafficSourceId": Identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *V2GetPartnersstatusCall) RequestMetadataTrafficSourceTrafficSourceId(requestMetadataTrafficSourceTrafficSourceId string) *V2GetPartnersstatusCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSourceId", requestMetadataTrafficSourceTrafficSourceId)
	return c
}

// RequestMetadataTrafficSourceTrafficSubId sets the optional parameter
// "requestMetadata.trafficSource.trafficSubId": Second level identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *V2GetPartnersstatusCall) RequestMetadataTrafficSourceTrafficSubId(requestMetadataTrafficSourceTrafficSubId string) *V2GetPartnersstatusCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSubId", requestMetadataTrafficSourceTrafficSubId)
	return c
}

// RequestMetadataUserOverridesIpAddress sets the optional parameter
// "requestMetadata.userOverrides.ipAddress": IP address to use instead
// of the user's geo-located IP address.
func (c *V2GetPartnersstatusCall) RequestMetadataUserOverridesIpAddress(requestMetadataUserOverridesIpAddress string) *V2GetPartnersstatusCall {
	c.urlParams_.Set("requestMetadata.userOverrides.ipAddress", requestMetadataUserOverridesIpAddress)
	return c
}

// RequestMetadataUserOverridesUserId sets the optional parameter
// "requestMetadata.userOverrides.userId": Logged-in user ID to
// impersonate instead of the user's ID.
func (c *V2GetPartnersstatusCall) RequestMetadataUserOverridesUserId(requestMetadataUserOverridesUserId string) *V2GetPartnersstatusCall {
	c.urlParams_.Set("requestMetadata.userOverrides.userId", requestMetadataUserOverridesUserId)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V2GetPartnersstatusCall) Fields(s ...googleapi.Field) *V2GetPartnersstatusCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *V2GetPartnersstatusCall) IfNoneMatch(entityTag string) *V2GetPartnersstatusCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V2GetPartnersstatusCall) Context(ctx context.Context) *V2GetPartnersstatusCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *V2GetPartnersstatusCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *V2GetPartnersstatusCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/partnersstatus")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.getPartnersstatus" call.
// Exactly one of *GetPartnersStatusResponse or error will be non-nil.
// Any non-2xx status code is an error. Response headers are in either
// *GetPartnersStatusResponse.ServerResponse.Header or (if a response
// was returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *V2GetPartnersstatusCall) Do(opts ...googleapi.CallOption) (*GetPartnersStatusResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GetPartnersStatusResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Gets Partners Status of the logged in user's agency.\nShould only be called if the logged in user is the admin of the agency.",
	//   "flatPath": "v2/partnersstatus",
	//   "httpMethod": "GET",
	//   "id": "partners.getPartnersstatus",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "requestMetadata.experimentIds": {
	//       "description": "Experiment IDs the current request belongs to.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.locale": {
	//       "description": "Locale to use for the current request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.partnersSessionId": {
	//       "description": "Google Partners session ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSourceId": {
	//       "description": "Identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSubId": {
	//       "description": "Second level identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.ipAddress": {
	//       "description": "IP address to use instead of the user's geo-located IP address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.userId": {
	//       "description": "Logged-in user ID to impersonate instead of the user's ID.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/partnersstatus",
	//   "response": {
	//     "$ref": "GetPartnersStatusResponse"
	//   }
	// }

}

// method id "partners.updateCompanies":

type V2UpdateCompaniesCall struct {
	s          *Service
	company    *Company
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// UpdateCompanies: Update company.
// Should only be called within the context of an authorized logged in
// user.
func (r *V2Service) UpdateCompanies(company *Company) *V2UpdateCompaniesCall {
	c := &V2UpdateCompaniesCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.company = company
	return c
}

// RequestMetadataExperimentIds sets the optional parameter
// "requestMetadata.experimentIds": Experiment IDs the current request
// belongs to.
func (c *V2UpdateCompaniesCall) RequestMetadataExperimentIds(requestMetadataExperimentIds ...string) *V2UpdateCompaniesCall {
	c.urlParams_.SetMulti("requestMetadata.experimentIds", append([]string{}, requestMetadataExperimentIds...))
	return c
}

// RequestMetadataLocale sets the optional parameter
// "requestMetadata.locale": Locale to use for the current request.
func (c *V2UpdateCompaniesCall) RequestMetadataLocale(requestMetadataLocale string) *V2UpdateCompaniesCall {
	c.urlParams_.Set("requestMetadata.locale", requestMetadataLocale)
	return c
}

// RequestMetadataPartnersSessionId sets the optional parameter
// "requestMetadata.partnersSessionId": Google Partners session ID.
func (c *V2UpdateCompaniesCall) RequestMetadataPartnersSessionId(requestMetadataPartnersSessionId string) *V2UpdateCompaniesCall {
	c.urlParams_.Set("requestMetadata.partnersSessionId", requestMetadataPartnersSessionId)
	return c
}

// RequestMetadataTrafficSourceTrafficSourceId sets the optional
// parameter "requestMetadata.trafficSource.trafficSourceId": Identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *V2UpdateCompaniesCall) RequestMetadataTrafficSourceTrafficSourceId(requestMetadataTrafficSourceTrafficSourceId string) *V2UpdateCompaniesCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSourceId", requestMetadataTrafficSourceTrafficSourceId)
	return c
}

// RequestMetadataTrafficSourceTrafficSubId sets the optional parameter
// "requestMetadata.trafficSource.trafficSubId": Second level identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *V2UpdateCompaniesCall) RequestMetadataTrafficSourceTrafficSubId(requestMetadataTrafficSourceTrafficSubId string) *V2UpdateCompaniesCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSubId", requestMetadataTrafficSourceTrafficSubId)
	return c
}

// RequestMetadataUserOverridesIpAddress sets the optional parameter
// "requestMetadata.userOverrides.ipAddress": IP address to use instead
// of the user's geo-located IP address.
func (c *V2UpdateCompaniesCall) RequestMetadataUserOverridesIpAddress(requestMetadataUserOverridesIpAddress string) *V2UpdateCompaniesCall {
	c.urlParams_.Set("requestMetadata.userOverrides.ipAddress", requestMetadataUserOverridesIpAddress)
	return c
}

// RequestMetadataUserOverridesUserId sets the optional parameter
// "requestMetadata.userOverrides.userId": Logged-in user ID to
// impersonate instead of the user's ID.
func (c *V2UpdateCompaniesCall) RequestMetadataUserOverridesUserId(requestMetadataUserOverridesUserId string) *V2UpdateCompaniesCall {
	c.urlParams_.Set("requestMetadata.userOverrides.userId", requestMetadataUserOverridesUserId)
	return c
}

// UpdateMask sets the optional parameter "updateMask": Standard field
// mask for the set of fields to be updated.
// Required with at least 1 value in FieldMask's paths.
func (c *V2UpdateCompaniesCall) UpdateMask(updateMask string) *V2UpdateCompaniesCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V2UpdateCompaniesCall) Fields(s ...googleapi.Field) *V2UpdateCompaniesCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V2UpdateCompaniesCall) Context(ctx context.Context) *V2UpdateCompaniesCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *V2UpdateCompaniesCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *V2UpdateCompaniesCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.company)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/companies")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.updateCompanies" call.
// Exactly one of *Company or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Company.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *V2UpdateCompaniesCall) Do(opts ...googleapi.CallOption) (*Company, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Company{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Update company.\nShould only be called within the context of an authorized logged in user.",
	//   "flatPath": "v2/companies",
	//   "httpMethod": "PATCH",
	//   "id": "partners.updateCompanies",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "requestMetadata.experimentIds": {
	//       "description": "Experiment IDs the current request belongs to.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.locale": {
	//       "description": "Locale to use for the current request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.partnersSessionId": {
	//       "description": "Google Partners session ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSourceId": {
	//       "description": "Identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSubId": {
	//       "description": "Second level identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.ipAddress": {
	//       "description": "IP address to use instead of the user's geo-located IP address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.userId": {
	//       "description": "Logged-in user ID to impersonate instead of the user's ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Standard field mask for the set of fields to be updated.\nRequired with at least 1 value in FieldMask's paths.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/companies",
	//   "request": {
	//     "$ref": "Company"
	//   },
	//   "response": {
	//     "$ref": "Company"
	//   }
	// }

}

// method id "partners.updateLeads":

type V2UpdateLeadsCall struct {
	s          *Service
	lead       *Lead
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// UpdateLeads: Updates the specified lead.
func (r *V2Service) UpdateLeads(lead *Lead) *V2UpdateLeadsCall {
	c := &V2UpdateLeadsCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.lead = lead
	return c
}

// RequestMetadataExperimentIds sets the optional parameter
// "requestMetadata.experimentIds": Experiment IDs the current request
// belongs to.
func (c *V2UpdateLeadsCall) RequestMetadataExperimentIds(requestMetadataExperimentIds ...string) *V2UpdateLeadsCall {
	c.urlParams_.SetMulti("requestMetadata.experimentIds", append([]string{}, requestMetadataExperimentIds...))
	return c
}

// RequestMetadataLocale sets the optional parameter
// "requestMetadata.locale": Locale to use for the current request.
func (c *V2UpdateLeadsCall) RequestMetadataLocale(requestMetadataLocale string) *V2UpdateLeadsCall {
	c.urlParams_.Set("requestMetadata.locale", requestMetadataLocale)
	return c
}

// RequestMetadataPartnersSessionId sets the optional parameter
// "requestMetadata.partnersSessionId": Google Partners session ID.
func (c *V2UpdateLeadsCall) RequestMetadataPartnersSessionId(requestMetadataPartnersSessionId string) *V2UpdateLeadsCall {
	c.urlParams_.Set("requestMetadata.partnersSessionId", requestMetadataPartnersSessionId)
	return c
}

// RequestMetadataTrafficSourceTrafficSourceId sets the optional
// parameter "requestMetadata.trafficSource.trafficSourceId": Identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *V2UpdateLeadsCall) RequestMetadataTrafficSourceTrafficSourceId(requestMetadataTrafficSourceTrafficSourceId string) *V2UpdateLeadsCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSourceId", requestMetadataTrafficSourceTrafficSourceId)
	return c
}

// RequestMetadataTrafficSourceTrafficSubId sets the optional parameter
// "requestMetadata.trafficSource.trafficSubId": Second level identifier
// to indicate where the traffic comes from.
// An identifier has multiple letters created by a team which redirected
// the
// traffic to us.
func (c *V2UpdateLeadsCall) RequestMetadataTrafficSourceTrafficSubId(requestMetadataTrafficSourceTrafficSubId string) *V2UpdateLeadsCall {
	c.urlParams_.Set("requestMetadata.trafficSource.trafficSubId", requestMetadataTrafficSourceTrafficSubId)
	return c
}

// RequestMetadataUserOverridesIpAddress sets the optional parameter
// "requestMetadata.userOverrides.ipAddress": IP address to use instead
// of the user's geo-located IP address.
func (c *V2UpdateLeadsCall) RequestMetadataUserOverridesIpAddress(requestMetadataUserOverridesIpAddress string) *V2UpdateLeadsCall {
	c.urlParams_.Set("requestMetadata.userOverrides.ipAddress", requestMetadataUserOverridesIpAddress)
	return c
}

// RequestMetadataUserOverridesUserId sets the optional parameter
// "requestMetadata.userOverrides.userId": Logged-in user ID to
// impersonate instead of the user's ID.
func (c *V2UpdateLeadsCall) RequestMetadataUserOverridesUserId(requestMetadataUserOverridesUserId string) *V2UpdateLeadsCall {
	c.urlParams_.Set("requestMetadata.userOverrides.userId", requestMetadataUserOverridesUserId)
	return c
}

// UpdateMask sets the optional parameter "updateMask": Standard field
// mask for the set of fields to be updated.
// Required with at least 1 value in FieldMask's paths.
// Only `state` and `adwords_customer_id` are currently supported.
func (c *V2UpdateLeadsCall) UpdateMask(updateMask string) *V2UpdateLeadsCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *V2UpdateLeadsCall) Fields(s ...googleapi.Field) *V2UpdateLeadsCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *V2UpdateLeadsCall) Context(ctx context.Context) *V2UpdateLeadsCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *V2UpdateLeadsCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *V2UpdateLeadsCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.lead)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/leads")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "partners.updateLeads" call.
// Exactly one of *Lead or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Lead.ServerResponse.Header or (if a response was returned at all) in
// error.(*googleapi.Error).Header. Use googleapi.IsNotModified to check
// whether the returned error was because http.StatusNotModified was
// returned.
func (c *V2UpdateLeadsCall) Do(opts ...googleapi.CallOption) (*Lead, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Lead{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := json.NewDecoder(res.Body).Decode(target); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates the specified lead.",
	//   "flatPath": "v2/leads",
	//   "httpMethod": "PATCH",
	//   "id": "partners.updateLeads",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "requestMetadata.experimentIds": {
	//       "description": "Experiment IDs the current request belongs to.",
	//       "location": "query",
	//       "repeated": true,
	//       "type": "string"
	//     },
	//     "requestMetadata.locale": {
	//       "description": "Locale to use for the current request.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.partnersSessionId": {
	//       "description": "Google Partners session ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSourceId": {
	//       "description": "Identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.trafficSource.trafficSubId": {
	//       "description": "Second level identifier to indicate where the traffic comes from.\nAn identifier has multiple letters created by a team which redirected the\ntraffic to us.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.ipAddress": {
	//       "description": "IP address to use instead of the user's geo-located IP address.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "requestMetadata.userOverrides.userId": {
	//       "description": "Logged-in user ID to impersonate instead of the user's ID.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Standard field mask for the set of fields to be updated.\nRequired with at least 1 value in FieldMask's paths.\nOnly `state` and `adwords_customer_id` are currently supported.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/leads",
	//   "request": {
	//     "$ref": "Lead"
	//   },
	//   "response": {
	//     "$ref": "Lead"
	//   }
	// }

}
