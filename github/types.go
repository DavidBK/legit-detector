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
	Team   struct {
		Name string `json:"name"`
	} `json:"team"`
}

type RepositoryPayload struct {
	Action     string `json:"action"`
	Repository struct {
		Name      string    `json:"name"`
		Id        int       `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"repository"`
}
