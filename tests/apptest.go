package tests

import (
	"github.com/revel/revel/testing"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatHealthCheckWorks() {
	t.Get("/lb-status")
	t.AssertOk()
}

func (t *AppTest) After() {
	println("Tear down")
}
