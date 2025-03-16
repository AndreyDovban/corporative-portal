package post

import (
	"gorm.io/datatypes"
)

type PostCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type PostUpdateRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	MailInstruction string `json:"mail_instruction"`
	WebInstruction  string `json:"web_instruction"`
}

type GetPostsRequest struct {
	Limit   int      `json:"limit"`
	Offset  int      `json:"offset"`
	Columns []string `json:"columns"`
}

type PostResponse struct {
	Uid             string         `json:"uid"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	MailInstruction string         `json:"mail_instruction"`
	WebInstruction  string         `json:"web_instruction"`
	CreatedAt       datatypes.Date `json:"created_at"`
	UpdatedAt       datatypes.Date `json:"updated_at"`
}

type GetPostsResponse struct {
	Columns []string       `json:"columns"`
	Data    []PostResponse `json:"data"`
	Count   int64          `json:"count"`
}
