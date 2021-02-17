package types

import "strings"

// QueryListWhois - list-whois command
const QueryListWhois = "list-whois"

// QueryGetWhois - get-whois command
const QueryGetWhois = "get-whois"

// QueryResolveName - resolve-name command
const QueryResolveName = "resolve-name"

// QueryResResolve Queries Result Payload for a resolve query
type QueryResResolve struct {
	Value string `json:"value"`
}

// implement fmt.Stringer
func (r QueryResResolve) String() string {
	return r.Value
}

// QueryResNames Queries Result Payload for a names query
type QueryResNames []string

// implement fmt.Stringer
func (n QueryResNames) String() string {
	return strings.Join(n[:], "\n")
}
