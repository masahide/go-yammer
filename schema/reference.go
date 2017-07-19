package schema

type Reference struct {
	// Common fields across multiple types
	ID   int    `json:"id"`
	Type string `json:"type"`

	// user, group
	FullName   string `json:"full_name"` // user, group
	FirstName  string `json:"first_name"`
	MugshotURL string `json:"mugshot_url"` // user, group

	// Type == "user"
	Email string `json:"email"`

	// message
	Body           *MessageBody `json:"body"`
	GroupId        int          `json:"group_id"`
	CreatedAt      *Time        `json:"created_at"`
	ContentExcerpt string       `json:"content_excerpt"`
	DirectMessage  bool         `json:"direct_message"`
	Language       string       `json:"language"`
	MessageType    string       `json:"message_type"`
	NetworkId      int          `json:"network_id"`
	Privacy        string       `json:"privacy"`
	RepliedToId    int          `json:"replied_to_id"`
	SenderId       int          `json:"sender_id"`
	SenderType     string       `json:"sender_type"`
	SystemMessage  bool         `json:"system_message"`
	ThreadId       int          `json:"thread_id"`
	URL            string       `json:"url"`
	WebURL         string       `json:"web_url"`
}

/*
"type": "message",
"url": "https://www.yammer.com/api/v1/messages/761971884",
"web_url": "https://www.yammer.com/cygames.co.jp/messages/761971884",
"id": 761971884,
"sender_id": 1536041975,
"replied_to_id": null,
"created_at": "2016/09/07 08:45:23 +0000",
"network_id": 224696,
"group_id": 322024,
"sender_type": "user",
"thread_id": 761971884,
"message_type": "update",
"system_message": false,
"content_excerpt": "【インフラスレ】",
"language": "ja",
"privacy": "private",
"direct_message": false,
"body": {
"plain": "【インフラスレ】"
}
*/
