// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costexplorer

type AccountScope string

// Enum values for AccountScope
const (
	AccountScopePayer  AccountScope = "PAYER"
	AccountScopeLinked AccountScope = "LINKED"
)

func (enum AccountScope) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AccountScope) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Context string

// Enum values for Context
const (
	ContextCostAndUsage Context = "COST_AND_USAGE"
	ContextReservations Context = "RESERVATIONS"
)

func (enum Context) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Context) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Dimension string

// Enum values for Dimension
const (
	DimensionAz                 Dimension = "AZ"
	DimensionInstanceType       Dimension = "INSTANCE_TYPE"
	DimensionLinkedAccount      Dimension = "LINKED_ACCOUNT"
	DimensionOperation          Dimension = "OPERATION"
	DimensionPurchaseType       Dimension = "PURCHASE_TYPE"
	DimensionRegion             Dimension = "REGION"
	DimensionService            Dimension = "SERVICE"
	DimensionUsageType          Dimension = "USAGE_TYPE"
	DimensionUsageTypeGroup     Dimension = "USAGE_TYPE_GROUP"
	DimensionRecordType         Dimension = "RECORD_TYPE"
	DimensionOperatingSystem    Dimension = "OPERATING_SYSTEM"
	DimensionTenancy            Dimension = "TENANCY"
	DimensionScope              Dimension = "SCOPE"
	DimensionPlatform           Dimension = "PLATFORM"
	DimensionSubscriptionId     Dimension = "SUBSCRIPTION_ID"
	DimensionLegalEntityName    Dimension = "LEGAL_ENTITY_NAME"
	DimensionDeploymentOption   Dimension = "DEPLOYMENT_OPTION"
	DimensionDatabaseEngine     Dimension = "DATABASE_ENGINE"
	DimensionCacheEngine        Dimension = "CACHE_ENGINE"
	DimensionInstanceTypeFamily Dimension = "INSTANCE_TYPE_FAMILY"
	DimensionBillingEntity      Dimension = "BILLING_ENTITY"
	DimensionReservationId      Dimension = "RESERVATION_ID"
)

func (enum Dimension) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Dimension) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Granularity string

// Enum values for Granularity
const (
	GranularityDaily   Granularity = "DAILY"
	GranularityMonthly Granularity = "MONTHLY"
	GranularityHourly  Granularity = "HOURLY"
)

func (enum Granularity) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Granularity) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type GroupDefinitionType string

// Enum values for GroupDefinitionType
const (
	GroupDefinitionTypeDimension GroupDefinitionType = "DIMENSION"
	GroupDefinitionTypeTag       GroupDefinitionType = "TAG"
)

func (enum GroupDefinitionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum GroupDefinitionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LookbackPeriodInDays string

// Enum values for LookbackPeriodInDays
const (
	LookbackPeriodInDaysSevenDays  LookbackPeriodInDays = "SEVEN_DAYS"
	LookbackPeriodInDaysThirtyDays LookbackPeriodInDays = "THIRTY_DAYS"
	LookbackPeriodInDaysSixtyDays  LookbackPeriodInDays = "SIXTY_DAYS"
)

func (enum LookbackPeriodInDays) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LookbackPeriodInDays) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Metric string

// Enum values for Metric
const (
	MetricBlendedCost           Metric = "BLENDED_COST"
	MetricUnblendedCost         Metric = "UNBLENDED_COST"
	MetricAmortizedCost         Metric = "AMORTIZED_COST"
	MetricNetUnblendedCost      Metric = "NET_UNBLENDED_COST"
	MetricNetAmortizedCost      Metric = "NET_AMORTIZED_COST"
	MetricUsageQuantity         Metric = "USAGE_QUANTITY"
	MetricNormalizedUsageAmount Metric = "NORMALIZED_USAGE_AMOUNT"
)

func (enum Metric) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Metric) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OfferingClass string

// Enum values for OfferingClass
const (
	OfferingClassStandard    OfferingClass = "STANDARD"
	OfferingClassConvertible OfferingClass = "CONVERTIBLE"
)

func (enum OfferingClass) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OfferingClass) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PaymentOption string

// Enum values for PaymentOption
const (
	PaymentOptionNoUpfront         PaymentOption = "NO_UPFRONT"
	PaymentOptionPartialUpfront    PaymentOption = "PARTIAL_UPFRONT"
	PaymentOptionAllUpfront        PaymentOption = "ALL_UPFRONT"
	PaymentOptionLightUtilization  PaymentOption = "LIGHT_UTILIZATION"
	PaymentOptionMediumUtilization PaymentOption = "MEDIUM_UTILIZATION"
	PaymentOptionHeavyUtilization  PaymentOption = "HEAVY_UTILIZATION"
)

func (enum PaymentOption) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PaymentOption) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TermInYears string

// Enum values for TermInYears
const (
	TermInYearsOneYear    TermInYears = "ONE_YEAR"
	TermInYearsThreeYears TermInYears = "THREE_YEARS"
)

func (enum TermInYears) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TermInYears) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
