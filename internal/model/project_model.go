package model

import "time"

type Project struct {
	ID          int       `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	GithubLink  string    `json:"github_link"`
	Image       string    `json:"image"`
	Created     time.Time `json:"created_at"`
	Updated     time.Time `json:"updated"`
}
