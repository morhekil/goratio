package analyser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/morhekil/goratio/analyser"
	"github.com/morhekil/goratio/analyser/timeframe"
)

type mockPropsRepo struct {
	trace []string
}

type mockProp struct {
	repo *mockPropsRepo
	name string
}

func (m mockProp) Analyse(t timeframe.Moment) {
	m.repo.trace = append(m.repo.trace, m.name)
}

func (r *mockPropsRepo) Props() []analyser.Prop {
	return []analyser.Prop{mockProp{r, "Gipsy"}, mockProp{r, "Cherno"}, mockProp{r, "Crimson"}}
}

type mockTimeframe struct {
}

func TestProcessesAllPropsAtTimeframe(t *testing.T) {
	f := mockTimeframe{}
	r := &mockPropsRepo{}
	a := analyser.New(r)
	a.Process(f)

	assert.Equal(t, []string{"Gipsy", "Cherno", "Crimson"}, r.trace)
}
