package webhooks

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"gopkg.in/yaml.v2"
)

type testCase struct {
	Action   string
	Payload  string
	WantErr  bool
	Expected string
}

type testFile struct {
	Webhookevents []testCase
}

func testEvents(t *testing.T, testFilePath string) {
	b, err := os.ReadFile(testFilePath)
	if err != nil {
		panic(err)
	}

	tf := &testFile{}
	if err := yaml.Unmarshal(b, tf); err != nil {
		panic(err)
	}

	for _, tt := range tf.Webhookevents {
		t.Run(tt.Action, func(t *testing.T) {
			var rawPayload map[string]interface{}
			if err := json.Unmarshal([]byte(tt.Payload), &rawPayload); err != nil {
				panic(err)
			}

			got, err := Parse[Payload](rawPayload)
			if err == nil && tt.WantErr {
				t.Log("Expected error")
				t.FailNow()
			}

			if (err != nil) != tt.WantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.WantErr)
				return
			}

			expected := &Payload{}
			if err := json.Unmarshal([]byte(tt.Expected), expected); err != nil {
				panic(err)
			}
			if !reflect.DeepEqual(got, expected) {
				gotJson, _ := json.MarshalIndent(got, "", "  ")
				expectedJson, _ := json.MarshalIndent(expected, "", "  ")
				t.Errorf("Parse() = %v, want %v", string(gotJson), string(expectedJson))
			}
		})
	}
}

func TestParseBasic(t *testing.T) {
	testEvents(t, "testdata/events.yaml")
}

func TestParseColumns(t *testing.T) {
	testEvents(t, "testdata/column-events.yaml")
}
