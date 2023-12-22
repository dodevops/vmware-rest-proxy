package handlerinspector

// The Inspector inspects the given HandlerBuilder instance and answers questions about its usage
type Inspector struct {
	hb *HandlerBuilder
}

// NewInspector creates a new Inspector
func NewInspector(hb *HandlerBuilder) *Inspector {
	return &Inspector{hb: hb}
}

// Called returns how often rule ruleName was called
func (i *Inspector) Called(ruleName string) int {
	if v, ok := i.hb.called[ruleName]; ok {
		return v
	} else {
		return 0
	}
}

// Failed returns if the handler fails at least once
func (i *Inspector) Failed() bool {
	return i.hb.failed
}

// AllWereCalled returns whether all rules were called at least once
func (i *Inspector) AllWereCalled() bool {
	allWereCalled := true
	for _, rule := range i.hb.rules {
		if i.Called(rule.name) == 0 {
			allWereCalled = false
		}
	}
	return allWereCalled
}
