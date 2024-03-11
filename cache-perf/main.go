package main

type Image []byte

type User struct {
	Login   string
	Active  bool
	Icon    Image
	Country string
}

// CountryCount returns a map of countries to the number of active users.
func CountryCount(users []User) map[string]int {
	counts := make(map[string]int) // country -> count
	for _, u := range users {
		if !u.Active {
			continue
		}
		counts[u.Country]++
	}

	return counts
}
