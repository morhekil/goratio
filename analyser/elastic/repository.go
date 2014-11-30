package elastic

import "github.com/morhekil/goratio/analyser/core"

// Repository backed by ElasticSearch
type Repository struct{}

// Props generates a list of all known events based on the list of
// terms in ElasticSearc
func (r Repository) Props() []core.Prop {
	return []core.Prop{}
}
