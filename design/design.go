package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("PANOS API", func() {
	Title("PANOS user-id API")
	Host("10.10.1.213:8080")
	Scheme("https")
	Produces("application/xml")
	BasePath("/api")
})

var _ = Resource("keygen", func() {
	BasePath("/keygen")
	Action("get", func() {
		Routing(
			GET(""),
		)

		Params(func() {
			Param("user", String, "Username")
			Param("password", String, "Password")
		})

		Response(OK, "application/xml")
	})
})

var _ = Resource("user-id", func() {
	BasePath("/user-id")
	Action("get", func() {
		Routing(
			GET(""),
		)

		Params(func() {
			Param("key", String)
			Param("action", String)
			Param("cmd", String)
		})

		Response(OK, "application/xml")
	})
})
