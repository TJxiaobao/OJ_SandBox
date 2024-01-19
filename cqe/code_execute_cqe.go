package cqe

import (
	"OJ_sandbox/utils/errno"
)

type CodeRequestCmd struct {
	Language string `json:"language" form:"language"`
	Code     string `json:"code" form:"code"`
	Input    string `json:"input" form:"input"`
}

func (c *CodeRequestCmd) Validate() error {
	if c.Language == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "language")
	}
	if c.Input == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "input")
	}
	if c.Code == "" {
		return errno.NewSimpleBizError(errno.ErrMissingParameter, nil, "code")
	}
	return nil
}
