package schema

type Reference struct {
	ID   int    `json:"id"`
	Type string `json:"type"`

	// Common fields across multiple types
	FullName   string `json:"full_name"` // user, group
	FirstName  string `json:"first_name"`
	MugshotURL string `json:"mugshot_url"` // user, group
	WebURL     string `json:"web_url"`     // user, group

	// Type == "user"
	Email string `json:"email"`
}
