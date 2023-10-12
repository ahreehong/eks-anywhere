package framework

import (
	"os"
)

const (
	// oob related stuff
	tinkerbellBMCConsumerURL            = "T_TINKERBELL_BMC_CONSUMER_URL"
	tinkerbellBMCHMACSecret             = "T_TINKERBELL_BMC_HMAC_SECRETS"
	tinkerbellBMCTimestampHeader        = "T_TINKERBELL_BMC_TIMESTAMP_HEADER"
	tinkerbellBMCIncludedPayloadHeaders = "T_TINKERBELL_BMC_INCLUDED_PAYLOAD_HEADERS"
	// tinkerbellBMCStaticHeaders          = "T_TINKERBELL_BMC_STATIC_HEADERS"
)

var requiredOOBEnvVars = []string{
	tinkerbellBMCConsumerURL,
	tinkerbellBMCHMACSecret,
	tinkerbellBMCTimestampHeader,
	tinkerbellBMCIncludedPayloadHeaders,
	// tinkerbellBMCStaticHeaders,
}

func RequiredOOBEnvVars() []string {
	return requiredOOBEnvVars
}

func WithOOBConfiguration() ClusterE2ETestOpt {
	return func(e *ClusterE2ETest) {
		checkRequiredEnvVars(e.T, requiredOOBEnvVars)
		consumerURL := os.Getenv(tinkerbellBMCConsumerURL)
		HMACSecret := os.Getenv(tinkerbellBMCHMACSecret)
		timestampHeader := os.Getenv(tinkerbellBMCTimestampHeader)
		includedPayloadHeaders := os.Getenv(tinkerbellBMCIncludedPayloadHeaders)
		// staticHeaders := os.Getenv(tinkerbellBMCStaticHeaders)
		err := os.Setenv("TINKERBELL_BMC_CONSUMER_URL", consumerURL)
		if err != nil {
			e.T.Fatalf("unable to set TINKERBELL_BMC_CONSUMER_URL: %v", err)
		}
		err = os.Setenv("TINKERBELL_BMC_HMAC_SECRETS", HMACSecret)
		if err != nil {
			e.T.Fatalf("unable to set TINKERBELL_BMC_HMAC_SECRETS: %v", err)
		}
		err = os.Setenv("TINKERBELL_BMC_TIMESTAMP_HEADER", timestampHeader)
		if err != nil {
			e.T.Fatalf("unable to set TINKERBELL_BMC_TIMESTAMP_HEADER: %v", err)
		}
		err = os.Setenv("TINKERBELL_BMC_INCLUDED_PAYLOAD_HEADERS", includedPayloadHeaders)
		if err != nil {
			e.T.Fatalf("unable to set TINKERBELL_BMC_INCLUDED_PAYLOAD_HEADERS: %v", err)
		}

		e.WithOOBConfiguration = true
	}
}
