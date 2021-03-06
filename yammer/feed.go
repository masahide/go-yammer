package yammer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/masahide/go-yammer/schema"
)

func (c *Client) Realtime() (*schema.RealtimeURL, error) {
	url := fmt.Sprintf("%s/api/v1/realtime.json?_=%d", c.baseURL, time.Now().Unix())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &schema.RealtimeURL{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Add("Yammer-Capabilities", "external_messaging,external_groups,system_request,user_sidebar,parsed_body_only2")

	resp, err := c.connection.Do(req)
	if err != nil {
		log.Println(err)
		return &schema.RealtimeURL{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &schema.RealtimeURL{}, err
	}

	var realtime schema.RealtimeURL
	err = json.Unmarshal(body, &realtime)
	if err != nil {
		return &schema.RealtimeURL{}, fmt.Errorf("json.Unmarshal err:%s, body:%s", err, body)
	}

	return &realtime, nil
}

func (c *Client) GroupFeed(id int) (*schema.MessageFeed, error) {
	url := fmt.Sprintf("%s/api/v1/messages/in_group/%d.json", c.baseURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &schema.MessageFeed{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Add("Yammer-Capabilities", "external_messaging,external_groups,system_request,user_sidebar,parsed_body_only2")

	resp, err := c.connection.Do(req)
	if err != nil {
		log.Println(err)
		return &schema.MessageFeed{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &schema.MessageFeed{}, err
	}

	var feed schema.MessageFeed
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return &schema.MessageFeed{}, err
	}

	return &feed, nil
}

func (c *Client) ThreadFeed(id int) (*schema.MessageFeed, error) {
	url := fmt.Sprintf("%s/api/v1/messages/in_thread/%d.json", c.baseURL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &schema.MessageFeed{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Add("Yammer-Capabilities", "external_messaging,external_groups,system_request,user_sidebar,parsed_body_only2")

	resp, err := c.connection.Do(req)
	if err != nil {
		log.Println(err)
		return &schema.MessageFeed{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &schema.MessageFeed{}, err
	}

	var feed schema.MessageFeed
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return &schema.MessageFeed{}, err
	}

	return &feed, nil
}

type FeedOptions struct {
	IncludeThreadStarter bool `url:"include_thread_starter"`
	ThreadsCount         int  `url:"threads_count"`
	MessagesCount        int  `url:"messages_count"`
	OlderThan            int  `url:"older_than,omitempty"`
	NewerThan            int  `url:"newer_than,omitempty"`
}

var DefaultHydrantOptions = "feed(id,network_id,channel_id,last_seen_message_id,has_older_threads,unseen_thread_count,unviewed_threads,threads(id,network_id,group_ids,thread_starter_id,first_reply_id,latest_reply_at,latest_reply_id,read_only,has_attachments,messages_count,share_count,topics,participants,in_private_conversation,messages(id,network_id,group_id,thread_id,sender_id,replied_to_id,shared_message_id,file_ids,page_ids,client_type_id,created_at,language,message_type,message_subtype,sender_type,body,title,is_bookmarked,is_liked,like_ids,likes_count,in_private_conversation,references,poll_options(option,answer,count,selected),praise(id,network_id,message_id,praised_user_ids,comment,icon),ymodule(id,network_id,ymodule_application_id)),state(unseen_message_count,last_read_message,is_following,priority_reason)),reference_data(client_applications(id,name,url),networks(id,name),groups(id,network_id,mugshot_id,permalink,name,private,moderated,color),users(id,network_id,mugshot_id,activated_at,permalink,full_name,job_title,state,email,is_following),guides(id,network_id,mugshot_id,permalink,full_name),files(id,network_id,group_id,latest_version_id,created_at,last_uploaded_at,name,original_name,description,content_type,size,width,height,transcoded,in_private_conversation,in_private_group,icon,state,is_following,is_editable),opengraphobjects(id,record_id,network_id,url,image,title,description,state,is_following),pages(id,network_id,group_id,latest_version_id,updated_at,title,official,state,is_following),tags(id),topics(id,tag_id,network_id,name,normalized_name),shared_messages(id,network_id,group_id,thread_id,sender_id,replied_to_id,shared_message_id,file_ids,page_ids,client_type_id,created_at,language,message_type,message_subtype,sender_type,body,title,is_bookmarked,is_liked,like_ids,likes_count,in_private_conversation,references,poll_options(option,answer,count,selected),praise(id,network_id,message_id,praised_user_ids,comment,icon),ymodule(id,network_id,ymodule_application_id)),replied_to_messages(id,network_id,group_id,thread_id,sender_id,replied_to_id,shared_message_id,file_ids,page_ids,client_type_id,created_at,language,message_type,message_subtype,sender_type,body,title,is_bookmarked,is_liked,like_ids,likes_count,in_private_conversation,references,poll_options(option,answer,count,selected),praise(id,network_id,message_id,praised_user_ids,comment,icon),ymodule(id,network_id,ymodule_application_id)),thread_summaries(id,network_id,group_ids,thread_starter_id,in_private_conversation,messages_count,share_count)),url_templates(thread,thread_web,message,message_web,attachment,attachment_web,attachment_download,attachment_thumbnail,attachment_scaled,attachment_preview,attachment_stream,attachment_edit,file_large_icon,page,page_web,page_preview,user,user_web,user_mugshot,group,group_web,group_mugshot,topic,topic_web,tag))"

func (c *Client) GroupFeedV2(id int, options FeedOptions) (*schema.HydrantFeed, error) {
	querystring, _ := query.Values(options)
	url := fmt.Sprintf("%s/api/v2/networks/107/feeds/group/%d?%s", c.baseURL, id, querystring.Encode())

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(DefaultHydrantOptions))
	if err != nil {
		return &schema.HydrantFeed{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "text/plain; charset=UTF-8")
	req.Header.Set("Accept", "application/json")

	resp, err := c.connection.Do(req)
	if err != nil {
		return &schema.HydrantFeed{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &schema.HydrantFeed{}, err
	}

	var feed schema.HydrantFeed
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return &schema.HydrantFeed{}, err
	}

	return &feed, nil
}

func (c *Client) ThreadFeedV2(id int) (*schema.HydrantThreadFeed, error) {
	url := fmt.Sprintf("%s/api/v2/threads/%d", c.baseURL, id)

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(DefaultHydrantOptions))
	if err != nil {
		return &schema.HydrantThreadFeed{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "text/plain; charset=UTF-8")
	req.Header.Set("Accept", "application/json")

	resp, err := c.connection.Do(req)
	if err != nil {
		return &schema.HydrantThreadFeed{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &schema.HydrantThreadFeed{}, err
	}

	var feed schema.HydrantThreadFeed
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return &schema.HydrantThreadFeed{}, err
	}

	return &feed, nil
}

var DefaultInboxOptions = `inboxes(id,network_id,channel_id,last_seen_message_id,has_older_threads,unseen_thread_count,threads(id,network_id,group_ids,thread_starter_id,first_reply_id,latest_reply_at,latest_reply_id,read_only,has_attachments,messages_count,share_count,topics,participants,in_private_conversation,messages(id,network_id,group_id,thread_id,sender_id,replied_to_id,shared_message_id,file_ids,page_ids,client_type_id,created_at,language,message_type,message_subtype,sender_type,body,title,in_private_conversation,references,version_num,updated_at,praise(id,network_id,message_id,praised_user_ids,comment,icon),is_bookmarked,is_liked,like_ids,likes_count,is_edited,is_editable,poll_options(option,answer,count,selected),connector_card,previous_versions(id,network_id,group_id,thread_id,sender_id,replied_to_id,shared_message_id,file_ids,page_ids,client_type_id,created_at,language,message_type,message_subtype,sender_type,body,title,in_private_conversation,references,version_num,updated_at,praise(id,network_id,message_id,praised_user_ids,comment,icon))),state(unseen_message_count,last_read_message,is_following,priority_reason)),reference_data(client_applications(id,name,url),networks(id,name),groups(id,network_id,mugshot_id,permalink,name,private,moderated,color),users(id,network_id,mugshot_id,activated_at,permalink,full_name,job_title,state,email,is_following),guides(id,network_id,mugshot_id,permalink,full_name),files(id,network_id,group_id,latest_version_id,created_at,last_uploaded_at,name,original_name,description,content_type,size,width,height,transcoded,in_private_conversation,in_private_group,icon,state,is_following,is_editable),opengraphobjects(id,record_id,network_id,url,image,title,description,state,is_following,video_url,provider_name),pages(id,network_id,group_id,latest_version_id,updated_at,title,official,state,is_following),tags(id),topics(id,tag_id,network_id,name,normalized_name),shared_messages(id,network_id,group_id,thread_id,sender_id,replied_to_id,shared_message_id,file_ids,page_ids,client_type_id,created_at,language,message_type,message_subtype,sender_type,body,title,in_private_conversation,references,version_num,updated_at,praise(id,network_id,message_id,praised_user_ids,comment,icon),is_bookmarked,is_liked,like_ids,likes_count,is_edited,is_editable,poll_options(option,answer,count,selected),connector_card),replied_to_messages(id,network_id,group_id,thread_id,sender_id,replied_to_id,shared_message_id,file_ids,page_ids,client_type_id,created_at,language,message_type,message_subtype,sender_type,body,title,in_private_conversation,references,version_num,updated_at,praise(id,network_id,message_id,praised_user_ids,comment,icon),is_bookmarked,is_liked,like_ids,likes_count,is_edited,is_editable,poll_options(option,answer,count,selected),connector_card),thread_summaries(id,network_id,group_ids,thread_starter_id,in_private_conversation,messages_count,share_count)),url_templates(thread,thread_web,message,message_web,attachment,attachment_web,attachment_download,attachment_thumbnail,attachment_scaled,attachment_preview,attachment_stream,attachment_edit,file_large_icon,page,page_web,page_preview,user,user_web,user_mugshot,group,group_web,group_mugshot,topic,topic_web,tag))`

func (c *Client) InboxFeedV2() (*schema.InboxFeed, error) {
	url := fmt.Sprintf("%s/api/v2/inboxes?include_thread_starter=true&threads_count=10&messages_count=1&followed_threads=true&private_threads=true&for_feed=inboxChat&fetch_type=newerThan_Backfill", c.baseURL)

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(DefaultInboxOptions))
	if err != nil {
		return &schema.InboxFeed{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "text/plain; charset=UTF-8")
	req.Header.Set("Accept", "application/json")

	resp, err := c.connection.Do(req)
	if err != nil {
		return &schema.InboxFeed{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &schema.InboxFeed{}, err
	}

	var feed schema.InboxFeed
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return &schema.InboxFeed{}, err
	}

	return &feed, nil
}

func (c *Client) InboxFeed() (*schema.MessageFeed, error) {
	url := fmt.Sprintf("%s/api/v1/messages/inbox.json", c.baseURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &schema.MessageFeed{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))

	if c.DebugMode {
		debug(httputil.DumpRequestOut(req, true))
	}

	resp, err := c.connection.Do(req)
	if err != nil {
		log.Println(err)
		return &schema.MessageFeed{}, err
	}

	if c.DebugMode {
		debug(httputil.DumpResponse(resp, true))
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &schema.MessageFeed{}, err
	}

	var feed schema.MessageFeed
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return &schema.MessageFeed{}, err
	}

	return &feed, nil
}

type InboxFeedV2Options struct {
	ThreadReadState      string `url:"thread_read_state,omitempty"`
	ThreadsCount         int    `url:"threads_count,omitempty"`
	MessagesCount        int    `url:"messages_count,omitempty"`
	FollowedThreads      bool   `url:"followed_threads"`
	IncludeThreadStarter bool   `url:"include_thread_starter"`
}

var SimpleInboxHydrantOptions = `inboxes(id,threads(id,thread_starter_id,read_only,messages(id,thread_id,sender_id,replied_to_id,created_at,body,title,references,updated_at,like_ids,additional_data),state(unseen_message_count,last_read_message)),reference_data(users(id,full_name,email)))`

func (c *Client) _InboxFeedV2(options InboxFeedV2Options, payload string) (*schema.HydrantFeed, error) {
	if payload == "" {
		payload = SimpleInboxHydrantOptions
	}
	querystring, _ := query.Values(options)
	url := fmt.Sprintf("%s/api/v2/inboxes?%s", c.baseURL, querystring.Encode())

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(payload))
	if err != nil {
		return &schema.HydrantFeed{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("Content-Type", "text/plain; charset=UTF-8")
	req.Header.Set("Accept", "application/json")

	if c.DebugMode {
		debug(httputil.DumpRequestOut(req, true))
	}

	resp, err := c.connection.Do(req)
	if err != nil {
		return &schema.HydrantFeed{}, err
	}
	defer resp.Body.Close()

	if c.DebugMode {
		debug(httputil.DumpResponse(resp, true))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &schema.HydrantFeed{}, err
	}

	var feed schema.HydrantFeed
	err = json.Unmarshal(body, &feed)
	if err != nil {
		return &schema.HydrantFeed{}, err
	}

	return &feed, nil
}
