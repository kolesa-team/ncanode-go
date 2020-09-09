package ncanode_test

import (
	"testing"

	"github.com/danikarik/ncanode-go"
	"github.com/stretchr/testify/require"
)

func TestRawVerify(t *testing.T) {
	client, err := ncanode.NewClient("http://127.0.0.1:14579")
	require.NoError(t, err)

	testCases := []struct {
		Name       string
		CMS        string
		VerifyOCSP bool
		VerifyCRL  bool
	}{
		{
			Name:       "Default",
			CMS:        `MIIIXQYJKoZIhvcNAQcCoIIITjCCCEoCAQExDzANBglghkgBZQMEAgEFADAVBgkqhkiG9w0BBwGgCAQGYXNkYXNkoIIGPDCCBjgwggQgoAMCAQICFBwcSttFCGs1KeAEyWMCrpdNTwmkMA0GCSqGSIb3DQEBCwUAMC0xCzAJBgNVBAYTAktaMR4wHAYDVQQDDBXSsNCa0J4gMy4wIChSU0EgVEVTVCkwHhcNMjAwMTI4MDYyMzA0WhcNMjEwMTI3MDYyMzA0WjCBtTEeMBwGA1UEAwwV0KLQldCh0KLQntCSINCi0JXQodCiMRUwEwYDVQQEDAzQotCV0KHQotCe0JIxGDAWBgNVBAUTD0lJTjEyMzQ1Njc4OTAxMTELMAkGA1UEBhMCS1oxHDAaBgNVBAcME9Cd0KPQoC3QodCj0JvQotCQ0J0xHDAaBgNVBAgME9Cd0KPQoC3QodCj0JvQotCQ0J0xGTAXBgNVBCoMENCi0JXQodCi0J7QktCY0KcwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCmScXte7gzv70DrjnQIKL50i\/7EXqV1ocME6UV9sZWrtAh4xW5xdllC7GhOQxpbfijRE+Qn1LrEixyjTjEpAej6k3YhKY7ByBMaeoGhrzDmmjaEo\/cYGIfHKeig8vb0ai98HAtCQNcRGDcOAThe5d\/KKqxh5msZGkz+j8nD5KNLivmYY2Zin7mX6UFOk6T5qCaaUJNxn\/S2BvZNN9mD6eWNpXZZK8yqaMU4PQYVCnJzlCfwYNmpFwSdnd1c5x8hCeO\/nmK7OJj\/Uj2T3cmL5IIwmWzKabjA631Sm3zhg7CD21upGN72AcZPAod4aQlMklJvdyv406lOoV2cK4sEuMPAgMBAAGjggHFMIIBwTAOBgNVHQ8BAf8EBAMCBsAwHQYDVR0lBBYwFAYIKwYBBQUHAwQGCCqDDgMDBAEBMB8GA1UdIwQYMBaAFKaMFjN8uOg1ZwY+XkFXVaKvNFBoMB0GA1UdDgQWBBQRxwe9VMv8zLOBXjnE61ex3H3+qTBeBgNVHSAEVzBVMFMGByqDDgMDAgMwSDAhBggrBgEFBQcCARYVaHR0cDovL3BraS5nb3Yua3ovY3BzMCMGCCsGAQUFBwICMBcMFWh0dHA6Ly9wa2kuZ292Lmt6L2NwczA8BgNVHR8ENTAzMDGgL6AthitodHRwOi8vdGVzdC5wa2kuZ292Lmt6L2NybC9uY2FfcnNhX3Rlc3QuY3JsMD4GA1UdLgQ3MDUwM6AxoC+GLWh0dHA6Ly90ZXN0LnBraS5nb3Yua3ovY3JsL25jYV9kX3JzYV90ZXN0LmNybDByBggrBgEFBQcBAQRmMGQwOAYIKwYBBQUHMAKGLGh0dHA6Ly90ZXN0LnBraS5nb3Yua3ovY2VydC9uY2FfcnNhX3Rlc3QuY2VyMCgGCCsGAQUFBzABhhxodHRwOi8vdGVzdC5wa2kuZ292Lmt6L29jc3AvMA0GCSqGSIb3DQEBCwUAA4ICAQBC4uYCiWSBxh8N6qIp4cQeX3inFNYPANAwwzOrbzWM1K0c0AJnR\/uIvGaLKcJQRhwr6H1leSnysfNjBriUPQ\/wPEDrPPaPCfrrHhkzSxZTqh7lFJg+KuRQIq9N\/R+CKtcaJFqZo0609kBKeEDt0HFZVICXK5fYfFoimHpXLKPjVzyFnpVTrsYx+SnJVVLPTG\/WXR75ZCzMU\/fOjCywd4oTLN7TX4XR5EEPcAig0TjHBlOmWsVzRIXiSMd73g4jI9WZS7TF3WKsgCipb9pb0PRqN7SImppYA5SkIErDnUSRItVsziywMQytvBdZ7tfhskXbtQoIBH\/KucUVu0R6K0fYy4XXEyprJYYG6EmuxDqNXWmbxPYnts9Ck6Xeyht5anUOsb64jzoNosJJmfulA\/IvkTaaKLbgTd3FlzQKBTO+OZEr+pVy8m\/edggblYeDuK\/nCexSFPgpZbso76LHG6I6Atu68xbNInscArPNLrsFM6oVpj\/7Ywoxkk9NiQWcWzA3be1UYKoouToq5CD26c4mcVaCY7adR6xGojzCTZg2jqcCpDAFeQMzshYJRjwlLMOmrmFibxdBO1c68GfbaVnOR4cmV2OczNkqNrhHF1pBMp3W9cMZ+10qUUMtn4oOvXOTfNOqGrNiplwoV\/MwcUGMsqda3FCiZy+0ms28RDRbcDGCAdswggHXAgEBMEUwLTELMAkGA1UEBhMCS1oxHjAcBgNVBAMMFdKw0JrQniAzLjAgKFJTQSBURVNUKQIUHBxK20UIazUp4ATJYwKul01PCaQwDQYJYIZIAWUDBAIBBQCgaTAYBgkqhkiG9w0BCQMxCwYJKoZIhvcNAQcBMBwGCSqGSIb3DQEJBTEPFw0yMDA5MDgxOTM4NThaMC8GCSqGSIb3DQEJBDEiBCBf2SRiX2qxahnMmAfHxQauGBNJDkumdfhD1aEOC6rNuDANBgkqhkiG9w0BAQsFAASCAQB\/pRVa5wgjTv66e+JHbwEft6+YrSRzKBVUKevv9TgSmK3jYClDQsoy9Z6QNr97AowiJGgkfjGKRcLd0f67TaTJYDUZcO+M+ARLzl8sxUGruhgnut5j\/CmDf3ZhZMdrXWy3Duw\/8eaaweBmTsZ8B0tpIqJHZDEymY9adfX1LRxqA0ujuZQhowckuDD3qw4wcdqdSJAJXq+cSMpL1Sm9gQLexWTcMCsHanbuysIq2q1y9mx\/QKB9QftPX01pch1m9UEwiTWPUZRrmUlApXEk0K0zVWgqBqrP7ufEmD2ssnRrzmQcQFhT0PGLpUJvtKTuV2DJWgEzUGjD1VMJMq2hLngc`,
			VerifyOCSP: false,
			VerifyCRL:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			resp, err := client.RawVerify(tc.CMS, tc.VerifyOCSP, tc.VerifyCRL)
			require.NoError(t, err)
			require.True(t, resp.Result.Valid)
		})
	}
}