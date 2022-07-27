package _29_unique_email_addresses

import "strings"

func NumUniqueEmails(emails []string) int {
	unique, local := make(map[string]struct{}), strings.Builder{}
	for _, email := range emails {
		local.Reset()
		i := 0
		for ; email[i] != '@'; i++ {
			if email[i] == '.' {
				continue
			}
			if email[i] == '+' {
				j := i + 1
				for ; email[j] != '@'; j++ {
				}
				i = j
				break
			}
			local.WriteByte(email[i])
		}
		unique[local.String()+email[i:]] = struct{}{}
	}
	return len(unique)
}
