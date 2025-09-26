package entrance_testing

import (
	"testing"
	"time"
)

func TestParseTasks(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		want   Ticket
		wantOk bool
	}{
		{
			name:   "корректная строка",
			input:  "TICKET-12345_Паша Попов_Готово_2024-01-01",
			wantOk: true,
			want: Ticket{
				Ticket: "TICKET-12345",
				User:   "Паша Попов",
				Status: "Готово",
				Date:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name:   "статус с числом",
			input:  "TICKET-12347_Анна_Статус123_2024-01-03",
			wantOk: true,
			want: Ticket{
				Ticket: "TICKET-12347",
				User:   "Анна",
				Status: "Статус123",
				Date:   time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name:   "длинный статус",
			input:  "TICKET-12348_User_Очень длинный статус с пробелами_2024-01-04",
			wantOk: true,
			want: Ticket{
				Ticket: "TICKET-12348",
				User:   "User",
				Status: "Очень длинный статус с пробелами",
				Date:   time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name:   "неправильный префикс",
			input:  "TICKET123_Паша_Готово_2024-01-01",
			wantOk: false,
		},
		{
			name:   "неправильное количество частей",
			input:  "TICKET-12345_Паша_Готово",
			wantOk: false,
		},
		{
			name:   "неправильный формат даты",
			input:  "TICKET-12345_Паша_Готово_2024/01/01",
			wantOk: false,
		},
		{
			name:   "пустая строка",
			input:  "",
			wantOk: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := ParseTasks(tt.input)

			if ok != tt.wantOk {
				t.Errorf("ParseTasks() ok = %v, wantOk %v", ok, tt.wantOk)
				return
			}

			if ok && tt.wantOk {
				if got.Ticket != tt.want.Ticket {
					t.Errorf("Ticket = %v, want %v", got.Ticket, tt.want.Ticket)
				}
				if got.User != tt.want.User {
					t.Errorf("User = %v, want %v", got.User, tt.want.User)
				}
				if got.Status != tt.want.Status {
					t.Errorf("Status = %v, want %v", got.Status, tt.want.Status)
				}
				if !got.Date.Equal(tt.want.Date) {
					t.Errorf("Date = %v, want %v", got.Date, tt.want.Date)
				}
			}
		})
	}
}
