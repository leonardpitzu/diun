package matcher

import (
	"regexp"
	"sync"
)

var (
	cache   = make(map[string]*regexp.Regexp)
	cacheMu sync.RWMutex
)

// compile returns a cached compiled regexp or compiles and caches it.
func compile(exp string) (*regexp.Regexp, error) {
	cacheMu.RLock()
	re, ok := cache[exp]
	cacheMu.RUnlock()
	if ok {
		return re, nil
	}

	re, err := regexp.Compile(exp)
	if err != nil {
		return nil, err
	}

	cacheMu.Lock()
	cache[exp] = re
	cacheMu.Unlock()
	return re, nil
}

// MatchString reports whether s contains any match of exp.
func MatchString(exp string, s string) bool {
	re, err := compile(exp)
	if err != nil {
		return false
	}
	return re.MatchString(s)
}

// IsIncluded checks if s matches any include pattern.
// If includes is empty, assume true.
func IsIncluded(s string, includes []string) bool {
	if len(includes) == 0 {
		return true
	}
	for _, include := range includes {
		if MatchString(include, s) {
			return true
		}
	}
	return false
}

// IsExcluded checks if s matches any exclude pattern.
// If excludes is empty, assume false.
func IsExcluded(s string, excludes []string) bool {
	if len(excludes) == 0 {
		return false
	}
	for _, exclude := range excludes {
		if MatchString(exclude, s) {
			return true
		}
	}
	return false
}
