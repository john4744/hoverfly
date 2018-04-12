package matching

import (
	"strings"

	glob "github.com/ryanuber/go-glob"
)

func HeaderMatching(requestMatcherHeaders, toMatch map[string][]string) *FieldMatch {

	matched := true
	var matchScore int

	for matcherHeaderKey, matcherHeaderValues := range requestMatcherHeaders {

		// Make everything lowercase, as headers are case insensitive
		for requestHeaderKey, requestHeaderValues := range toMatch {
			delete(toMatch, requestHeaderKey)
			toMatch[strings.ToLower(requestHeaderKey)] = requestHeaderValues
		}

		toMatchHeaderValues, found := toMatch[strings.ToLower(matcherHeaderKey)]
		if !found {
			matched = false
		}

		for _, matcherHeaderValue := range matcherHeaderValues {
			matcherHeaderValueMatched := false
			for _, toMatchHeaderValue := range toMatchHeaderValues {
				if glob.Glob(strings.ToLower(matcherHeaderValue), strings.ToLower(toMatchHeaderValue)) {
					matcherHeaderValueMatched = true
					matchScore++
				}
			}

			if !matcherHeaderValueMatched {
				matched = false
			}
		}
	}
	return &FieldMatch{
		Matched:    matched,
		MatchScore: matchScore,
	}
}
