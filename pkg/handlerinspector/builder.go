package handlerinspector

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

// HandlerBuilder is a builder interface to generating http.handlers for mock servers
type HandlerBuilder struct {
	// rules is a list of HandlerInspector rules
	rules []Rule
	// called records which rules have been called for the Inspector
	called map[string]int
	// failed records if no non-matching rule was called for the Inspector
	failed bool
}

// NewBuilder creates a new HandlerBuilder. Start here.
func NewBuilder() *HandlerBuilder {
	return &HandlerBuilder{}
}

// WithRule appends a Rule to the builder
func (b *HandlerBuilder) WithRule(r Rule) *HandlerBuilder {
	b.rules = append(b.rules, r)
	return b
}

// Build builds a http.Handler from the rules in the builder
func (b *HandlerBuilder) Build() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Debugf("Checking rules for request %v", r)
		foundRule := false
		for _, rule := range b.rules {
			logrus.Debugf("Checking rule %s", rule.name)
			matches := true
			for _, c := range rule.conditions {
				logrus.Debugf("Checking condition %v", c)
				if !c.Matches(r) {
					matches = false
				}
			}
			if matches {
				foundRule = true
				logrus.Debugf("Carrying out matching rule %s", rule.name)
				if b.called == nil {
					b.called = make(map[string]int)
				}
				if v, ok := b.called[rule.name]; ok {
					v++
				} else {
					b.called[rule.name] = 1
				}
				for key, value := range rule.headers {
					w.Header().Add(key, value)
				}
				w.WriteHeader(rule.code)
				if rule.useBodyFunc {
					_, _ = fmt.Fprint(w, rule.bodyFunc(r))
				} else {
					_, _ = fmt.Fprint(w, rule.body)
				}
				return
			}
		}

		if !foundRule {
			logrus.Errorf("Didn't find a rule for request %v with body %v", r, r.Body)
			b.failed = true
		}
	})
}
