package spotify

import "github.com/zmb3/spotify"

type Range struct {
	Start int
	End   int
}

func batches(length, batch int) []Range {
	if length <= 0 {
		return nil
	}

	var ranges []Range
	i := 0
	for ; i < (length - batch); i += batch {
		ranges = append(ranges, Range{i, i + batch})
	}
	return append(ranges, Range{i, length})
}

func spotifyIdsInterface(s []interface{}) []spotify.ID {
	output := make([]spotify.ID, len(s))
	for i, v := range s {
		output[i] = spotify.ID(v.(string))
	}
	return output
}

func difference(slice1 []interface{}, slice2 []interface{}) []interface{} {
	var diff []interface{}

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}
