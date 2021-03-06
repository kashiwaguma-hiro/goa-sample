package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("admin", func() {
	Title("Admin Service")
	Description("Service for admin")
    Server("admin", func() {
        Host("localhost", func() {
            URI("http://localhost:8000")
        })
    })
})

var _ = Service("admin", func() {
	Description("The admin service performs operations on users.")

	Method("users", func() {
		Payload(func() {
			users_request()
			Required("first-name", "last-name", "email", "tel")
		})

		Result(func(){
			users_response()
		})

		HTTP(func() {
			POST("/users/")
		})

	})

	Files("/openapi.json", "./gen/http/openapi.json")
})

func users_request(){
	Field(1, "first-name", String, "first-name", func() {
		Example("fn")
		MinLength(1)
		MaxLength(128)
		Format(FormatRegexp)
		Pattern("^[0-9a-zA-Z\\.\\-_~]*$") // URIに含まれるので、利用できる記号はRFC3986のみ
	})
	Field(2, "last-name", String, "last-name", func() {
		Example("ln")
		MinLength(1)
		MaxLength(128)
		Format(FormatRegexp)
		Pattern("^[0-9a-zA-Z\\.\\-_~]*$") // URIに含まれるので、利用できる記号はRFC3986のみ
	})
	Field(3, "email", String, "Email", func() {
		Example("tom@example.com")
		MinLength(1)
		MaxLength(256)
		Format(FormatEmail)
	})
	Field(4, "tel", String, "tel", func() {
		Example("000-0000-0000")
		MinLength(1)
		MaxLength(128)
		Format(FormatRegexp)
		Pattern("^[1-9]\\d{3}\\.\\d{4}\\.\\d{4}") // FIXME tekitou
	})
}

func users_response(){
	users_request()
	Field(5, "id", Int, "id", func() {
		Example(101)
	})
}