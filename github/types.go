// TODO: generate this file
package github

import "time"

type PushPayload struct {
	Ref        string `json:"ref"`
	Before     string `json:"before"`
	After      string `json:"after"`
	Repository struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		Private bool   `json:"private"`
		Owner   struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"owner"`
		PushedAt int64  `json:"pushed_at"`
		Language string `json:"language"`
	} `json:"repository"`
	Pusher struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"pusher"`
	Organization struct {
		Login string `json:"login"`
		ID    int    `json:"id"`
	} `json:"organization"`
}

type TeamPayload struct {
	Action string `json:"action"`
	Scope  string `json:"scope"`
	Member struct {
		Login string `json:"login"`
		ID    int64  `json:"id"`
	} `json:"member"`
	Team struct {
		Name string `json:"name"`
		ID   int64  `json:"id"`
		Slug string `json:"slug"`
	} `json:"team"`
	Organization struct {
		Login string `json:"login"`
		ID    int64  `json:"id"`
	} `json:"organization"`
	Sender struct {
		Login string `json:"login"`
		ID    int64  `json:"id"`
	} `json:"sender"`
}

type RepositoryPayload struct {
	Action     string `json:"action"`
	Repository struct {
		ID         int       `json:"id"`
		Name       string    `json:"name"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
		PushedAt   time.Time `json:"pushed_at"`
		Visibility string    `json:"visibility"`
		Owner      struct {
			Login string `json:"login"`
			ID    int    `json:"id"`
		} `json:"owner"`
	} `json:"repository"`
	Organization struct {
		Login string `json:"login"`
		ID    int    `json:"id"`
	} `json:"organization"`
	Sender struct {
		Login string `json:"login"`
		ID    int    `json:"id"`
	} `json:"sender"`
}
