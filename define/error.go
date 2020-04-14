package define

import "errors"

var (
	AuthError       = errors.New("AuthError")
	DiscoveryIsNull = errors.New("Discovery Null")
)
