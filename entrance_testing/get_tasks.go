package entrance_testing

import (
	"bufio"
	"strings"
)

func GetTasks(text string, user *string, status *string) []Ticket {
	var result []Ticket
	scanner := bufio.NewScanner(strings.NewReader(text))

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		ticket, err := ParseTasks(line)
		if err {
			continue
		}

		if user != nil && ticket.User != *user {
			continue
		}

		if status != nil && ticket.Status != *status {
			continue
		}

		result = append(result, ticket)
	}

	return result
}
