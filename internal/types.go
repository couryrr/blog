package internal

import (
    "time"
)

type Article struct {
    Id int `json:"id"`
    Title string `json:"title"`
    Slug string `json:"slug"`
    FilePath string `json:"-"`
    DateCreated time.Time `json:"date_created"`
    DateUpdate *time.Time `json:"date_updated"`
}

