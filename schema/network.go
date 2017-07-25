package schema

// Network is /api/v1/networks/current.json Returns a list of networks
type Network struct {
	ID                   int    `json:"id"`
	Email                string `json:"email"`
	Name                 string `json:"name"`
	Community            bool   `json:"community"`
	Permalink            string `json:"permalink"`
	WebURL               string `json:"web_url"`
	IsGroupEnabled       bool   `json:"is_group_enabled"`
	IsChatEnabled        bool   `json:"is_chat_enabled"`
	IsTranslationEnabled bool   `json:"is_translation_enabled"`
	CreatedAt            string `json:"created_at"`
	State                string `json:"state"`
}
