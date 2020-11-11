package model

import(
	"time"
)

type Podcast struct {
    Id             string    `json:"id,omitempty"`
    Title          string    `json:"title,omitempty"`
    Author         string    `json:"author,omitempty"`
    Topic          string    `json:"topic,omitempty"`
    Episodes       []Episode `json:"episodes,omitempty"`
    Created        time.Time `json:"created,omitempty"`
    Updated        time.Time `json:"updated,omitempty"`
}

type Episode struct {
    Number         int       `json:"number,omitempty"`
    Title          string    `json:"title,omitempty"`
    Created        time.Time `json:"created,omitempty"`
    Updated        time.Time `json:"updated,omitempty"`
    Location       string    `json:"location,omitempty"`
}