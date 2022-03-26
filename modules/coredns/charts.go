package coredns

import "github.com/netdata/go.d.plugin/agent/module"

type (
	// Charts is an alias for module.Charts
	Charts = module.Charts
	// Chart is an alias for module.Chart
	Chart = module.Chart
	// Dims is an alias for module.Dims
	Dims = module.Dims
	// Dim is an alias for module.Dim
	Dim = module.Dim
)

var summaryCharts = Charts{
	{
		ID:    "dns_request_count_total",
		Title: "Number Of DNS Requests",
		Units: "requests/s",
		Fam:   "summary",
		Ctx:   "coredns.dns_request_count_total",
		Dims: Dims{
			{ID: "request_total", Name: "requests", Algo: module.Incremental},
		},
	},
	{
		ID:    "dns_responses_count_total",
		Title: "Number Of DNS Responses",
		Units: "responses/s",
		Fam:   "summary",
		Ctx:   "coredns.dns_responses_count_total",
		Dims: Dims{
			{ID: "response_total", Name: "responses", Algo: module.Incremental},
		},
	},
	{
		ID:    "dns_request_count_total_per_status",
		Title: "Number Of Processed And Dropped DNS Requests",
		Units: "requests/s",
		Fam:   "summary",
		Ctx:   "coredns.dns_request_count_total_per_status",
		Type:  module.Stacked,
		Dims: Dims{
			{ID: "request_per_status_processed", Name: "processed", Algo: module.Incremental},
			{ID: "request_per_status_dropped", Name: "dropped", Algo: module.Incremental},
		},
	},
	{
		ID:    "dns_no_matching_zone_dropped_total",
		Title: "Number Of Dropped DNS Requests Because Of No Matching Zone",
		Units: "requests/s",
		Fam:   "summary",
		Ctx:   "coredns.dns_no_matching_zone_dropped_total",
		Dims: Dims{
			{ID: "no_matching_zone_dropped_total", Name: "dropped", Algo: module.Incremental},
		},
	},
	{
		ID:    "dns_panic_count_total",
		Title: "Number Of Panics",
		Units: "panics/s",
		Fam:   "summary",
		Ctx:   "coredns.dns_panic_count_total",
		Dims: Dims{
			{ID: "panic_total", Name: "panics", Algo: module.Incremental},
		},
	},
	{
		ID:    "dns_requests_count_total_per_proto",
		Title: "Number Of DNS Requests Per Transport Protocol",
		Units: "requests/s",
		Fam:   "summary",
		Ctx:   "coredns.dns_requests_count_total_per_proto",
		Type:  module.Stacked,
		Dims: Dims{
			{ID: "request_per_proto_udp", Name: "udp", Algo: module.Incremental},
			{ID: "request_per_proto_tcp", Name: "tcp", Algo: module.Incremental},
		},
	},
	{
		ID:    "dns_requests_count_total_per_ip_family",
		Title: "Number Of DNS Requests Per IP Family",
		Units: "requests/s",
		Fam:   "summary",
		Ctx:   "coredns.dns_requests_count_total_per_ip_family",
		Type:  module.Stacked,
		Dims: Dims{
			{ID: "request_per_ip_family_v4", Name: "v4", Algo: module.Incremental},
			{ID: "request_per_ip_family_v6", Name: "v6", Algo: module.Incremental},
		},
	},
	//{
	//	ID:    "dns_requests_duration_seconds",
	//	Title: "Number Of DNS Requests Per Bucket",
	//	Units: "requests/s",
	//	Fam:   "summary",
	//	Ctx:   "coredns.dns_requests_duration_seconds",
	//	Type:  module.Stacked,
	//	Dims: Dims{
	//		{ID: "request_duration_seconds_bucket_0.00025", Name: "0.00025s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_0.0005", Name: "0.0005s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_0.001", Name: "0.001s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_0.002", Name: "0.002s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_0.004", Name: "0.004s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_0.008", Name: "0.008s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_0.016", Name: "0.016s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_0.032", Name: "0.032s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_0.064", Name: "0.064s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_0.128", Name: "0.128s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_0.256", Name: "0.256s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_0.512", Name: "0.512s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_1.024", Name: "1.024s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_2.048", Name: "2.048s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_4.096", Name: "4.096s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_8.192", Name: "8.192s", Algo: module.Incremental},
	//		{ID: "request_duration_seconds_bucket_+Inf", Name: "+Inf", Algo: module.Incremental},
	//	},
	//},
	{
		ID:    "dns_requests_count_total_per_type",
		Title: "Number Of DNS Requests Per Type",
		Units: "requests/s",
		Fam:   "summary",
		Ctx:   "coredns.dns_requests_count_total_per_per_type",
		Type:  module.Stacked,
		Dims: Dims{
			{ID: "request_per_type_A", Name: "A", Algo: module.Incremental},
			{ID: "request_per_type_AAAA", Name: "AAAA", Algo: module.Incremental},
			{ID: "request_per_type_MX", Name: "MX", Algo: module.Incremental},
			{ID: "request_per_type_SOA", Name: "SOA", Algo: module.Incremental},
			{ID: "request_per_type_CNAME", Name: "CNAME", Algo: module.Incremental},
			{ID: "request_per_type_PTR", Name: "PTR", Algo: module.Incremental},
			{ID: "request_per_type_TXT", Name: "TXT", Algo: module.Incremental},
			{ID: "request_per_type_NS", Name: "NS", Algo: module.Incremental},
			{ID: "request_per_type_DS", Name: "DS", Algo: module.Incremental},
			{ID: "request_per_type_DNSKEY", Name: "DNSKEY", Algo: module.Incremental},
			{ID: "request_per_type_RRSIG", Name: "RRSIG", Algo: module.Incremental},
			{ID: "request_per_type_NSEC", Name: "NSEC", Algo: module.Incremental},
			{ID: "request_per_type_NSEC3", Name: "NSEC3", Algo: module.Incremental},
			{ID: "request_per_type_IXFR", Name: "IXFR", Algo: module.Incremental},
			{ID: "request_per_type_ANY", Name: "ANY", Algo: module.Incremental},
			{ID: "request_per_type_other", Name: "other", Algo: module.Incremental},
		},
	},
	{
		ID:    "dns_responses_count_total_per_rcode",
		Title: "Number Of DNS Responses Per Rcode",
		Units: "responses/s",
		Fam:   "summary",
		Ctx:   "coredns.dns_responses_count_total_per_rcode",
		Type:  module.Stacked,
		Dims: Dims{
			{ID: "response_per_rcode_NOERROR", Name: "NOERROR", Algo: module.Incremental},
			{ID: "response_per_rcode_FORMERR", Name: "FORMERR", Algo: module.Incremental},
			{ID: "response_per_rcode_SERVFAIL", Name: "SERVFAIL", Algo: module.Incremental},
			{ID: "response_per_rcode_NXDOMAIN", Name: "NXDOMAIN", Algo: module.Incremental},
			{ID: "response_per_rcode_NOTIMP", Name: "NOTIMP", Algo: module.Incremental},
			{ID: "response_per_rcode_REFUSED", Name: "REFUSED", Algo: module.Incremental},
			{ID: "response_per_rcode_YXDOMAIN", Name: "YXDOMAIN", Algo: module.Incremental},
			{ID: "response_per_rcode_YXRRSET", Name: "YXRRSET", Algo: module.Incremental},
			{ID: "response_per_rcode_NXRRSET", Name: "NXRRSET", Algo: module.Incremental},
			{ID: "response_per_rcode_NOTAUTH", Name: "NOTAUTH", Algo: module.Incremental},
			{ID: "response_per_rcode_NOTZONE", Name: "NOTZONE", Algo: module.Incremental},
			{ID: "response_per_rcode_BADSIG", Name: "BADSIG", Algo: module.Incremental},
			{ID: "response_per_rcode_BADKEY", Name: "BADKEY", Algo: module.Incremental},
			{ID: "response_per_rcode_BADTIME", Name: "BADTIME", Algo: module.Incremental},
			{ID: "response_per_rcode_BADMODE", Name: "BADMODE", Algo: module.Incremental},
			{ID: "response_per_rcode_BADNAME", Name: "BADNAME", Algo: module.Incremental},
			{ID: "response_per_rcode_BADALG", Name: "BADALG", Algo: module.Incremental},
			{ID: "response_per_rcode_BADTRUNC", Name: "BADTRUNC", Algo: module.Incremental},
			{ID: "response_per_rcode_BADCOOKIE", Name: "BADCOOKIE", Algo: module.Incremental},
			{ID: "response_per_rcode_other", Name: "other", Algo: module.Incremental},
		},
	},
}

var serverCharts = Charts{
	{
		ID:    "per_%s_%s_dns_request_count_total",
		Title: "Number Of DNS Requests, %s %s",
		Units: "requests/s",
		Fam:   "%s %s",
		Ctx:   "coredns.server_dns_request_count_total",
		Dims: Dims{
			{ID: "%s_request_total", Name: "requests", Algo: module.Incremental},
		},
	},
	{
		ID:    "per_%s_%s_dns_responses_count_total",
		Title: "Number Of DNS Responses, %s %s",
		Units: "responses/s",
		Fam:   "%s %s",
		Ctx:   "coredns.server_dns_responses_count_total",
		Dims: Dims{
			{ID: "%s_response_total", Name: "responses", Algo: module.Incremental},
		},
	},
	{
		ID:    "per_%s_%s_dns_request_count_total_per_status",
		Title: "Number Of Processed And Dropped DNS Requests, %s %s",
		Units: "requests/s",
		Fam:   "%s %s",
		Ctx:   "coredns.server_dns_request_count_total_per_status",
		Type:  module.Stacked,
		Dims: Dims{
			{ID: "%s_request_per_status_processed", Name: "processed", Algo: module.Incremental},
			{ID: "%s_request_per_status_dropped", Name: "dropped", Algo: module.Incremental},
		},
	},
	{
		ID:    "per_%s_%s_dns_requests_count_total_per_proto",
		Title: "Number Of DNS Requests Per Transport Protocol, %s %s",
		Units: "requests/s",
		Fam:   "%s %s",
		Ctx:   "coredns.server_dns_requests_count_total_per_proto",
		Type:  module.Stacked,
		Dims: Dims{
			{ID: "%s_request_per_proto_udp", Name: "udp", Algo: module.Incremental},
			{ID: "%s_request_per_proto_tcp", Name: "tcp", Algo: module.Incremental},
		},
	},
	{
		ID:    "per_%s_%s_dns_requests_count_total_per_ip_family",
		Title: "Number Of DNS Requests Per IP Family, %s %s",
		Units: "requests/s",
		Fam:   "%s %s",
		Ctx:   "coredns.server_dns_requests_count_total_per_ip_family",
		Type:  module.Stacked,
		Dims: Dims{
			{ID: "%s_request_per_ip_family_v4", Name: "v4", Algo: module.Incremental},
			{ID: "%s_request_per_ip_family_v6", Name: "v6", Algo: module.Incremental},
		},
	},
	//{
	//	ID:    "per_%s_%s_dns_requests_duration_seconds",
	//	Title: "Number Of DNS Requests Per Bucket, %s %s",
	//	Units: "requests/s",
	//	Fam:   "%s %s",
	//	Ctx:   "coredns.server_dns_requests_duration_seconds",
	//	Type:  module.Stacked,
	//	Dims: Dims{
	//		{ID: "%s_request_duration_seconds_bucket_0.00025", Name: "0.00025s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_0.0005", Name: "0.0005s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_0.001", Name: "0.001s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_0.002", Name: "0.002s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_0.004", Name: "0.004s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_0.008", Name: "0.008s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_0.016", Name: "0.016s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_0.032", Name: "0.032s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_0.064", Name: "0.064s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_0.128", Name: "0.128s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_0.256", Name: "0.256s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_0.512", Name: "0.512s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_1.024", Name: "1.024s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_2.048", Name: "2.048s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_4.096", Name: "4.096s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_8.192", Name: "8.192s", Algo: module.Incremental},
	//		{ID: "%s_request_duration_seconds_bucket_+Inf", Name: "+Inf", Algo: module.Incremental},
	//	},
	//},
	{
		ID:    "per_%s_%s_dns_requests_count_total_per_type",
		Title: "Number Of DNS Requests Per Type, %s %s",
		Units: "requests/s",
		Fam:   "%s %s",
		Ctx:   "coredns.server_dns_requests_count_total_per_per_type",
		Type:  module.Stacked,
		Dims: Dims{
			{ID: "%s_request_per_type_A", Name: "A", Algo: module.Incremental},
			{ID: "%s_request_per_type_AAAA", Name: "AAAA", Algo: module.Incremental},
			{ID: "%s_request_per_type_MX", Name: "MX", Algo: module.Incremental},
			{ID: "%s_request_per_type_SOA", Name: "SOA", Algo: module.Incremental},
			{ID: "%s_request_per_type_CNAME", Name: "CNAME", Algo: module.Incremental},
			{ID: "%s_request_per_type_PTR", Name: "PTR", Algo: module.Incremental},
			{ID: "%s_request_per_type_TXT", Name: "TXT", Algo: module.Incremental},
			{ID: "%s_request_per_type_NS", Name: "NS", Algo: module.Incremental},
			{ID: "%s_request_per_type_DS", Name: "DS", Algo: module.Incremental},
			{ID: "%s_request_per_type_DNSKEY", Name: "DNSKEY", Algo: module.Incremental},
			{ID: "%s_request_per_type_RRSIG", Name: "RRSIG", Algo: module.Incremental},
			{ID: "%s_request_per_type_NSEC", Name: "NSEC", Algo: module.Incremental},
			{ID: "%s_request_per_type_NSEC3", Name: "NSEC3", Algo: module.Incremental},
			{ID: "%s_request_per_type_IXFR", Name: "IXFR", Algo: module.Incremental},
			{ID: "%s_request_per_type_ANY", Name: "ANY", Algo: module.Incremental},
			{ID: "%s_request_per_type_other", Name: "other", Algo: module.Incremental},
		},
	},
	{
		ID:    "per_%s_%s_dns_responses_count_total_per_rcode",
		Title: "Number Of DNS Responses Per Rcode, %s %s",
		Units: "responses/s",
		Fam:   "%s %s",
		Ctx:   "coredns.server_dns_responses_count_total_per_rcode",
		Type:  module.Stacked,
		Dims: Dims{
			{ID: "%s_response_per_rcode_NOERROR", Name: "NOERROR", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_FORMERR", Name: "FORMERR", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_SERVFAIL", Name: "SERVFAIL", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_NXDOMAIN", Name: "NXDOMAIN", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_NOTIMP", Name: "NOTIMP", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_REFUSED", Name: "REFUSED", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_YXDOMAIN", Name: "YXDOMAIN", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_YXRRSET", Name: "YXRRSET", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_NXRRSET", Name: "NXRRSET", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_NOTAUTH", Name: "NOTAUTH", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_NOTZONE", Name: "NOTZONE", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_BADSIG", Name: "BADSIG", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_BADKEY", Name: "BADKEY", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_BADTIME", Name: "BADTIME", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_BADMODE", Name: "BADMODE", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_BADNAME", Name: "BADNAME", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_BADALG", Name: "BADALG", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_BADTRUNC", Name: "BADTRUNC", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_BADCOOKIE", Name: "BADCOOKIE", Algo: module.Incremental},
			{ID: "%s_response_per_rcode_other", Name: "other", Algo: module.Incremental},
		},
	},
}

var zoneCharts = func() Charts {
	c := serverCharts.Copy()
	_ = c.Remove("per_%s_%s_dns_request_count_total_per_status")
	return *c
}()