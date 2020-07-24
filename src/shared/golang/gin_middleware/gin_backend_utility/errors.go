package gin_backend_utility

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/lib/pq"
	"gitlab.com/grchive/grchive-v2/shared/backend/roles"
	"strings"
)

type GrchiveErrorCodes int

const (
	GECNoError    GrchiveErrorCodes = 0
	GECBadRequest                   = 1
	// Internal: something happened due to a messed up SQL query or something because we messed up.
	// External: The user put in some value that triggered a constraint (e.g. a duplicate).
	GECDbInternal   = 2
	GECDbExternal   = 3
	GECUnauthorized = 10000
	GECGeneric      = 99999
)

type GrchiveErrorMessage string

const (
	GEMBadRequest         GrchiveErrorMessage = "Your browser sent an improperly formed request. Please try again."
	GEMBadRequestTree                         = "Double check your resource identififers and ensure you're using the correct org and engagement ID."
	GEMDbGeneric                              = "An unknown error occurred on our servers. Please try again/contact us."
	GEMDbUniqueRequest                        = "An object with the given name/ID already exists. Choose another name/ID."
	GEMValueDbGeneric                         = "You provided invalid values. Please try again."
	GEMUnauthorizedLogin                      = "You don't have permission to view this page. Ensure that you are logged in."
	GEMUnauthorizedVerify                     = "Please verify your email address."
	GEMUnauthorizedRoles                      = "You do not have an assigned role."
	GEMUnknown                                = "Unknown"
)

func CreateErrorMessageForLackingPermissions(allPermissions ...roles.Permission) GrchiveErrorMessage {
	builder := strings.Builder{}
	builder.WriteString("You are missing the permissions: ")
	for idx, p := range allPermissions {
		builder.WriteString("\"")
		builder.WriteString(p)
		builder.WriteString("\"")

		if idx != len(allPermissions)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString(".")
	return GrchiveErrorMessage(builder.String())
}

// "Message" here should give the user some clue on how to fix their problem.
type ApiErrorResponse struct {
	GrchiveCode    GrchiveErrorCodes
	GrchiveMessage GrchiveErrorMessage
}

type WebappError struct {
	Err     error
	Context string
	Code    GrchiveErrorCodes
	Message GrchiveErrorMessage
}

func (w *WebappError) Error() string {
	if w.Err != nil {
		return fmt.Sprintf("ERR: %s - [%s]", w.Context, w.Err.Error())
	} else {
		return fmt.Sprintf("ERR: %s", w.Context)
	}
}

// This is mainly here to handle some database errors that would be
// tedious to handle elsewhere. Other times the code and message is probably
// better crafted at the place of the error with a best guess.
func CreateGrchiveCodeMessageFromError(err error) (GrchiveErrorCodes, GrchiveErrorMessage) {
	if err == nil {
		return GECGeneric, GEMUnknown
	}

	switch terr := err.(type) {
	case *pq.Error:
		switch terr.Code {
		case "23505":
			return GECDbExternal, GEMDbUniqueRequest
		case "P0001":
			return GECDbExternal, GEMValueDbGeneric
		default:
			return GECDbInternal, GEMDbGeneric
		}
	}

	return GECGeneric, GEMUnknown
}

func ApiMiddlewareErrorHandling(c *gin.Context) {
	c.Next()

	if len(c.Errors) == 0 {
		return
	}

	// Only handle the first error that has the expected WebappError type.
loop:
	for _, err := range c.Errors {
		switch terr := err.Err.(type) {
		case *WebappError:
			resp := ApiErrorResponse{}
			resp.GrchiveCode, resp.GrchiveMessage = CreateGrchiveCodeMessageFromError(terr.Err)

			if terr.Code != GECNoError {
				resp.GrchiveCode = terr.Code
			}

			if terr.Message != "" {
				resp.GrchiveMessage = terr.Message
			}

			render.WriteJSON(c.Writer, resp)
			break loop
		default:
			continue
		}
	}
}
