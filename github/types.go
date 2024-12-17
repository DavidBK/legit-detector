// TODO: generate this file
package github

import "time"

type PushPayload struct {
	Repository struct {
		PushedAt int64  `json:"pushed_at"`
		Name     string `json:"name"`
		Id       int    `json:"id"`
	} `json:"repository"`
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
