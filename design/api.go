package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("stress", func() {
	Title("Stress test service for kagome")
	Description("Service for kagome stress tests")
	Server("stress", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

func RequiredAttribute(name string, args ...interface{}) {
	Attribute(name, args...)
	Required(name)
}

var TokenResult = ResultType("application/vnd.token+json", func() {
	RequiredAttribute("surface", String)
	RequiredAttribute("pos", String)
	RequiredAttribute("start", Int)
	RequiredAttribute("end", Int)
	RequiredAttribute("type", String)
})

var _ = Service("stress", func() {
	Description("Stress test performs to tokenize sentences.")

	Method("start", func() {
		HTTP(func() {
			POST("/start")
		})
	})
	Method("stop", func() {
		HTTP(func() {
			POST("/stop")
		})
	})
	Method("tokenize", func() {
		Payload(func() {
			Attribute("sentence", String)
			Required("sentence")
		})
		Result(CollectionOf(TokenResult))
		HTTP(func() {
			POST("/tokenize")
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
