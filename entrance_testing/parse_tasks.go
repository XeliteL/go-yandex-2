package entrance_testing

import (
	"strings"
	"time"
)

func ParseTasks(text string) (Ticket, bool) {
	if !strings.HasPrefix(text, "TICKET-") {
		return Ticket{}, false
	}

	parseTicket := strings.Split(text, "_")
	if len(parseTicket) != 4 {
		return Ticket{}, false
	}

	ticketNumber := parseTicket[0]
	ticketUser := parseTicket[1]
	ticketStatus := parseTicket[2]
	ticketDateStr := parseTicket[3]

	ticketDate, err := time.Parse("2006-01-02", ticketDateStr)
	if err != nil {
		return Ticket{}, false
	}

	return Ticket{
		Ticket: ticketNumber,
		User:   ticketUser,
		Status: ticketStatus,
		Date:   ticketDate,
	}, true
}
