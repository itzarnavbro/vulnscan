package elasticsearch

import (
	"vulnscan/pkg/iac/providers/aws/elasticsearch"
	"vulnscan/pkg/iac/scanners/cloudformation/parser"
)

// Adapt adapts an ElasticSearch instance
func Adapt(cfFile parser.FileContext) elasticsearch.Elasticsearch {
	return elasticsearch.Elasticsearch{
		Domains: getDomains(cfFile),
	}
}
