package model

import(
	"time"
)

type Podcast struct {
    Title          string    `json:"title,omitempty"`
    Author         string    `json:"author,omitempty"`
    Topic          string    `json:"topic,omitempty"`
    Episodes       []Episode `json:"episodes,omitempty"`
    Created        time.Time `json:"created,omitempty"`
    Updated        time.Time `json:"updated,omitempty"`
}

type Episode struct {
    Title          string    `json:"title,omitempty"`
    Created        time.Time `json:"created,omitempty"`
    Updated        time.Time `json:"updated,omitempty"`
    Location       string    `json:"location,omitempty"`
}