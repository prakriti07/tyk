package gateway

import (
	"github.com/prakriti07/tyk/apidef"
	"github.com/prakriti07/tyk/user"
)

// parameterBodies
// swagger:response parameterBodies
type swaggerParameterBodies struct {
	// in: body
	APIStatusMessage apiStatusMessage
	// in: body
	APIModifyKeySuccess apiModifyKeySuccess
	// in: body
	NewClientRequest NewClientRequest
	// in: body
	APIDefinition apidef.APIDefinition
	// in: body
	SessionState user.SessionState
	// in:body
	APIAllKeys apiAllKeys
	// in: body
	OAuthClientToken OAuthClientToken
}
