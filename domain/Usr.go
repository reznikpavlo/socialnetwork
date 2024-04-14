package domain

import (
	"strings"
	"time"
)

type Usr struct {
	Id        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Locale    string    `json:"locale"`
	LastVisit time.Time `json:"lastVisit"`
}

func (u Usr) String() string {
	b := strings.Builder{}
	b.WriteString("{ ")
	b.WriteString(u.Name)
	b.WriteString(" ")
	b.WriteString(u.Email)
	b.WriteString(" ")
	b.WriteString(u.Id)
	b.WriteString(" }")
	return b.String()
}
