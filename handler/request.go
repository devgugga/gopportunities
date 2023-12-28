package handler

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (createRequest *CreateOpeningRequest) Validate() error {

	if createRequest.Role == "" && createRequest.Company == "" && createRequest.Location == "" && createRequest.Link == "" && createRequest.Remote == nil && createRequest.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if createRequest.Role == "" {
		return errParamIsRequired("role", "string")
	}

	if createRequest.Company == "" {
		return errParamIsRequired("company", "string")
	}

	if createRequest.Location == "" {
		return errParamIsRequired("location", "string")
	}

	if createRequest.Link == "" {
		return errParamIsRequired("link", "string")
	}

	if createRequest.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}

	if createRequest.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}
