// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import "github.com/sah4ez/dapr-example/interfaces/types"

type requestUserGetNameByID struct {
	Id types.ID `json:"id,omitempty"`
}

type responseUserGetNameByID struct {
	User types.User `json:"user,omitempty"`
}
