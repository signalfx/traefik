package log

// From https://github.com/containous/traefik/blob/a09dfa3ce10f222ab6e10ac11386d330698ff6f8/log/fields.go

// Log entry name
const (
	EntryPointName      = "entryPointName"
	RouterName          = "routerName"
	MiddlewareName      = "middlewareName"
	MiddlewareType      = "middlewareType"
	ProviderName        = "providerName"
	ServiceName         = "serviceName"
	MetricsProviderName = "metricsProviderName"
	TracingProviderName = "tracingProviderName"
	ServerName          = "serverName"
)
