package pkg

import "github.com/rsmrtk/fd-cfg/config/env"

type Config struct {
	ENV                           env.ENV
	PostgresURL                   string
	StripeKey                     string
	GoogleMapsKey                 string
	JWTSecret                     string
	JWTDuration                   string
	TwilioAccountSID              string
	TwilioAuthToken               string
	TwilioVerifyServiceSID        string
	FcmServiceAccount             string
	ApnsAuthKey                   string
	MapStaticApiKey               string
	MailGunDomain                 string
	MailApiKey                    string
	TollGuruApiKey                string
	PagerDutyIntegrationKey       string
	JurnyVoxImplantCredentials    string
	JurnyVoxImplantAddress        string
	StagingJurnyVoxImplantAddress string
	TLSCertFile                   string
	TLSKeyFile                    string
	PhpAPIKey                     string
	PhpRiderAPIStagingURL         string
	PhpRiderAPIProdURL            string
	PhpCustomerFeedbackAPIProdURL string
}
