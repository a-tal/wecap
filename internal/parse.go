package internal

import "strings"

func parseQuery(uri string) map[string]string {
	// this uses non-standard query strings...
	// they don't add a ?, they start with &'s

	parts := strings.Split(uri, "&")

	query := map[string]string{}
	for i := 1; i < len(parts); i++ {
		part := parts[i]
		if strings.Contains(part, "=") {
			sections := strings.SplitN(part, "=", 2)
			if len(sections) == 2 {
				key, value := sections[0], sections[1]
				if previous, ok := query[key]; ok {
					ll.Printf(
						"duplicate key found: %q. previous value of %q will be overwritten with %q",
						key,
						previous,
						value,
					)
				}
				query[key] = value
			} else {
				ll.Printf("failed to parse query part: %s", part)
			}
		}
	}

	return query
}
