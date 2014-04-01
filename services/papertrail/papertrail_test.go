package papertrail

import(
	"reflect"
	"testing"
	"github.com/remind101/slack-hooks/slack"
)

var(
	papertrail = Service{}
	parseTests = []struct {
		input string
		expected *Payload
	}{
		{
			`{"events":[{"hostname":"api","message":"measure.time=1"}]}`,
			&Payload{ Events: []Event{ { Hostname: "api", Message: "measure.time=1" } } },
		},
	}
	handleTests = []struct {
		input  *Payload
		expected *slack.Message
	}{
		{
			&Payload{ Events: []Event{ { Hostname: "api", Message: "message.time=1" } } },
			&slack.Message{ Username: "Papertrail", Text: "Search" },
		},
	}

)

func TestParse(t *testing.T) {
	for _, test := range parseTests {
		payload := papertrail.Parse(test.input)
		equal   := reflect.DeepEqual(payload, test.expected)

		if !equal {
			t.Log(payload)
			t.Log(test.expected)
			t.FailNow()
		}
	}
}

func TestHandle(t *testing.T) {
	for _, test := range handleTests {
		message := papertrail.Handle(test.input)
		equal   := reflect.DeepEqual(message, test.expected)

		if !equal {
			t.Log(message)
			t.Log(test.expected)
			t.FailNow()
		}
	}
}
