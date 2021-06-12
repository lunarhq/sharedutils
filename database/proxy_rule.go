package database

type ProxyRuleCreateParams struct {
	Name                 *string
	AccountID            *string
	Endpoint             *string
	HeaderApiKeyKey      *string
	HeaderApiKeyValue    *string
	HeaderApiSecretKey   *string
	HeaderApiSecretValue *string
	WhitelistedDomains   *string
}

type ProxyRuleUpdateParams struct {
	Name *string
}

type ProxyRuleListParams struct {
	AccountID *string
}
