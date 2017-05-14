package schema

type InboxURLTemplates struct {
	UserWeb             string `json:"user_web"`
	AttachmentWeb       string `json:"attachment_web"`
	AttachmentStream    string `json:"attachment_stream"`
	Attachment          string `json:"attachment"`
	FileLargeIcon       string `json:"file_large_icon"`
	PageWeb             string `json:"page_web"`
	PagePreview         string `json:"page_preview"`
	Group               string `json:"group"`
	Realtime            string `json:"realtime"`
	AttachmentDownload  string `json:"attachment_download"`
	GroupMugshot        string `json:"group_mugshot"`
	MessageWeb          string `json:"message_web"`
	Thread              string `json:"thread"`
	Message             string `json:"message"`
	AttachmentPreview   string `json:"attachment_preview"`
	AttachmentScaled    string `json:"attachment_scaled"`
	TopicWeb            string `json:"topic_web"`
	GroupWeb            string `json:"group_web"`
	Topic               string `json:"topic"`
	AttachmentEdit      string `json:"attachment_edit"`
	UserMugshot         string `json:"user_mugshot"`
	Page                string `json:"page"`
	ThreadWeb           string `json:"thread_web"`
	User                string `json:"user"`
	AttachmentThumbnail string `json:"attachment_thumbnail"`
}

type InboxFeed struct {
	ID                string             `json:"id"`
	ChannelID         string             `json:"channel_id"`
	NetworkID         int                `json:"network_id"`
	HasOlderThreads   bool               `json:"has_older_threads"`
	UnseenThreadCount int                `json:"unseen_thread_count"`
	URLTemplates      *InboxURLTemplates `json:"url_templates"`
}
