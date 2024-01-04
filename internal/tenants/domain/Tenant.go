package domain

type TenantRequestDto struct {
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type TenantResponseDto struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

type Tenant struct {
	ID      string
	Name    string
	Active  bool
	Created string
}
