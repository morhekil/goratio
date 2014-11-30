package elastic

import "github.com/morhekil/goratio/analyser"

// Repository backed by ElasticSearch
type Repository struct{}

// Props generates a list of all known events based on the list of
// terms in ElasticSearc
func (r Repository) Props() []analyser.Prop {
	return []analyser.Prop{}
}
