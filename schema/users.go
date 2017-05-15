package schema

type User struct {
	Type               string `json:"type"`
	ID                 int    `json:"id"`
	Email              string `json:"email"`
	FullName           string `json:"full_name"`
	NetworkID          int    `json:"network_id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	Privacy            string `json:"privacy"`
	URL                string `json:"url"`
	WebURL             string `json:"web_url"`
	MugshotURL         string `json:"mugshot_url"`
	MugshotURLTemplate string `json:"mugshot_url_template"`
	MugshotID          string `json:"mugshot_id"`
	ShowInDirectory    string `json:"show_in_directory"`
	CreatedAt          string `json:"created_at"`
	Color              string `json:"color"`
	External           bool   `json:"external"`
	Moderated          bool   `json:"moderated"`
	CreatorType        string `json:"creator_type"`
	CreatorID          int    `json:"creator_id"`
	State              string `json:"state"`
	Stats              struct {
		Members       int    `json:"members"`
		Updates       int    `json:"updates"`
		LastMessageID int    `json:"last_message_id"`
		LastMessageAt string `json:"last_message_at"`
	} `json:"stats"`
	Admin       string `json:"admin"`
	NetworkName string `json:"network_name"`
}
