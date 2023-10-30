package utils

func StringSliceHas(slice []string, tgt string) bool {
	for _, tst := range slice {
		if tst == tgt {
			return true
		}
	}
	return false
}

// FilterSlice takes a slice and a function to test. The test function is run
// on every entry in the slice, if it returns true the entry is kept. It returns
// the slice trimmed to size.
func FilterSlice[T any](slice []T, test func(in T) bool) []T {
	tgt := slice[:0]
	max := 0
	for _, entry := range slice {
		if test(entry) {
			tgt = append(tgt, entry)
			max++
		}
	}
	return tgt[:max]
}

// CleanErrorSlice removes any indices that are nil and returns the new slice
// to length.
func CleanErrorSlice(slice []error) []error {
	return FilterSlice[error](slice, func(err error) bool {
		return err != nil
	})
}

// CleanStringSlice takes  a slice of strings and removes any empty
// entries, trimming the original slice down.
func CleanStringSlice(slice []string) []string {
	return FilterSlice[string](slice, func(str string) bool {
		return len(str) > 0
	})
}
