package docs

import (
	"github.com/godemo/model"
)

// swagger:operation GET /users/{id} users getuserbyid
// ---
// summary: Get user by id.
// description: Returns user with specified id.
// parameters:
// - name: id
//   in: path
//   description: User ID
//   type: string
//   required: true
// responses:
//   200:
//      $ref: "#/responses/userResponse"

// User Response Payload.
// swagger:response userResponse
type userResponseWrapper struct {
	// in:body
	Body model.User
}

// // swagger:parameters idOfFoobarEndpoint
// type foobarParamsWrapper struct {
// 	// This text will appear as description of your request body.
// 	// in:body
// 	Body api.FooBarRequest
// }
