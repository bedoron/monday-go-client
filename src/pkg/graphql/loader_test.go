package graphql

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"

	"github.com/bedoron/monday-go-client/src/pkg/entities"
	"github.com/stretchr/testify/assert"
)

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}

type MondayTableSchema struct {
	Followers             *entities.ColumnNumeric `mapstructure:"__followers,omitempty"`
	Comments              *entities.ColumnNumeric `mapstructure:"comments,omitempty"`
	Date                  *entities.ColumnDate    `mapstructure:"date,omitempty"`
	Date1                 *entities.ColumnDate    `mapstructure:"date_1,omitempty"`
	PostPublificationTime *entities.ColumnDate    `mapstructure:"date2,omitempty"`
	Language              *entities.ColumnStatus  `mapstructure:"language,omitempty"`
	Likes                 *entities.ColumnNumeric `mapstructure:"likes,omitempty"`
	Media                 *entities.ColumnString  `mapstructure:"link_to_media,omitempty"`
	Platform              *entities.ColumnStatus  `mapstructure:"platform,omitempty"`
	Category              *entities.ColumnString  `mapstructure:"post_category,omitempty"`
	PostId                *entities.ColumnString  `mapstructure:"post_id,omitempty"`
	ProfileName           *entities.ColumnString  `mapstructure:"profile_name,omitempty"`
	Shares                *entities.ColumnNumeric `mapstructure:"shares,omitempty"`
	Reasp                 *entities.ColumnStatus  `mapstructure:"status_18,omitempty"`
	Intent                *entities.ColumnStatus  `mapstructure:"support___report,omitempty"`
	OriginalComment       *entities.ColumnString  `mapstructure:"text7,omitempty"`
	Source                *entities.ColumnStatus  `mapstructure:"source,omitempty"`
	ReasonToReport        *entities.ColumnString  `mapstructure:"status_1,omitempty"`
	ReasonToReportHeb     *entities.ColumnString  `mapstructure:"status_18,omitempty"`
	CreationLog1          *entities.ColumnDate    `mapstructure:"creation_log_1,omitempty"`
	TargetRegion          *entities.ColumnString  `mapstructure:"target_region,omitempty"`
	PostText              *entities.ColumnString  `mapstructure:"text_of_the_post,omitempty"`
	People                *entities.ColumnString  `mapstructure:"people,omitempty"`

	Name     *entities.ColumnString `mapstructure:"name,omitempty"`
	Priority *entities.ColumnString `mapstructure:"priority,omitempty"`
	Status   *entities.ColumnStatus `mapstructure:"status,omitempty"` // Approved or not

	// Aux Fields
	Text1 *entities.ColumnString `mapstructure:"text_1,omitempty"`
	Text2 *entities.ColumnString `mapstructure:"text_2,omitempty"`
	Text3 *entities.ColumnString `mapstructure:"text_3,omitempty"`
	Text4 *entities.ColumnString `mapstructure:"text_4,omitempty"`
	Text5 *entities.ColumnString `mapstructure:"text_5,omitempty"`
	Text6 *entities.ColumnString `mapstructure:"text_6,omitempty"`
	Text7 *entities.ColumnString `mapstructure:"text_7,omitempty"`
	Text8 *entities.ColumnString `mapstructure:"text_8,omitempty"`
}

type EmptyMondayTableSchema struct {
	JunkField *entities.ColumnNumeric `mapstructure:"doesnt_exists,omitempty"`
}

func Test_NonExistingFields(t *testing.T) {
	b := Must(os.ReadFile("testdata/single_api_response.json"))

	var v map[string]interface{}
	err := json.Unmarshal(b, &v)
	assert.NoError(t, err)

	mtfEmpty := &EmptyMondayTableSchema{}
	gqEmpty := NewLoader[EmptyMondayTableSchema](mtfEmpty)

	gEmptyItem := Must(Parse[GQLItem](v))
	gqEmpty.Load(gEmptyItem)

	assert.Nil(t, mtfEmpty.JunkField, "expected empy object")
}

func Test_NewGraphQLLoader(t *testing.T) {
	b := Must(os.ReadFile("testdata/single_api_response.json"))

	var v map[string]interface{}
	err := json.Unmarshal(b, &v)
	assert.NoError(t, err)

	mtfTarget := &MondayTableSchema{}
	gq := NewLoader[MondayTableSchema](mtfTarget) // This will load the schema we need to parse

	assert.NotNil(t, gq)

	gItem := Must(Parse[GQLItem](v)) // Parse generic structure as a graphql item entry
	gq.Load(gItem)

	// Do some manual tests for problematic fields
	assert.Equal(t, "Approved", *mtfTarget.Status.GetText())
	assert.Equal(t, Must(entities.DecodeTimeBestEffort("2023-10-30 11:27:19 UTC")), mtfTarget.CreationLog1.GetDate())
	assert.Equal(t, float64(118800), *mtfTarget.Followers.Get())
	assert.Equal(t, Must(entities.DecodeTimeBestEffort("2023-10-26 11:30")), mtfTarget.PostPublificationTime.GetDate())

	b = Must(os.ReadFile("testdata/single_api_schema.json"))
	mtfExpected := &MondayTableSchema{}
	if err := json.Unmarshal(b, mtfExpected); err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(mtfExpected, mtfTarget) {
		gotJson, _ := json.MarshalIndent(mtfTarget, "", "  ")
		expectedJson, _ := json.MarshalIndent(mtfExpected, "", "  ")
		t.Errorf("Parse() = %v, want %v", string(gotJson), string(expectedJson))
	}

}
