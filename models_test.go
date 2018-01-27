// +build integration
package plex

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"os"
	"testing"
)

var (
	u                       string
	token                   string
	movie_meta_data_id      int
	tv_top_meta_data_id     int
	tv_season_meta_data_id  int
	tv_episode_meta_data_id int
	testConn                *Plex
)

func init() {
	u = os.Getenv("PLEX_URL")
	token = os.Getenv("PLEX_TOKEN")

	movie_meta_data_id = 1
	tv_top_meta_data_id = 9
	tv_season_meta_data_id = 10
	tv_episode_meta_data_id = 11

	if u != "" {
		var err error
		if testConn, err = New(u, token); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func sharedTestData(plexConn *Plex, query string, config *mapstructure.DecoderConfig, t *testing.T) {

	fullUrl := fmt.Sprintf("%s/%s", testConn.URL, query)

	resp, err := testConn.get(fullUrl, defaultHeaders())

	if err != nil {
		t.Error(err)
	}

	// Unauthorized
	if resp.StatusCode == 401 {
		t.Error("Unauthorized")
	}

	defer resp.Body.Close()

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		panic(err)
	}

	var x interface{}

	if err := json.NewDecoder(resp.Body).Decode(&x); err != nil {
		t.Error(err)
	}

	if err := decoder.Decode(x); err != nil {
		panic(err)
	}

	for _, key := range config.Metadata.Unused {
		t.Errorf("Unused key: %#v\n", key)
	}

	// for _, key := range config.Metadata.Keys {
	// 	t.Logf("Used keys: %#v\n", key)
	// }
}

func sectionTest(plexConn *Plex) func(t *testing.T) {
	query := "library/sections/"

	var results Response
	var md mapstructure.Metadata
	config := &mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &results,
	}

	return func(t *testing.T) {
		sharedTestData(plexConn, query, config, t)
	}
}

func sectionAllTest(plexConn *Plex) func(t *testing.T) {
	query := "library/sections/all"

	var results Response
	var md mapstructure.Metadata
	config := &mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &results,
	}

	return func(t *testing.T) {
		sharedTestData(plexConn, query, config, t)
	}
}

func metadataTop(plexConn *Plex, metadata_id int) func(t *testing.T) {
	query := fmt.Sprintf("library/metadata/%d", metadata_id)

	var results Response
	var md mapstructure.Metadata
	config := &mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &results,
	}

	return func(t *testing.T) {
		sharedTestData(plexConn, query, config, t)
	}
}

func metadataChildren(plexConn *Plex, metadata_id int) func(t *testing.T) {
	query := fmt.Sprintf("library/metadata/%d/children", metadata_id)

	var results Response
	var md mapstructure.Metadata
	config := &mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &results,
	}

	return func(t *testing.T) {
		sharedTestData(plexConn, query, config, t)
	}
}

func serverSessions(plexConn *Plex) func(t *testing.T) {
	query := "status/sessions"

	var results Response
	var md mapstructure.Metadata
	config := &mapstructure.DecoderConfig{
		Metadata: &md,
		Result:   &results,
	}

	return func(t *testing.T) {
		sharedTestData(plexConn, query, config, t)
	}
}

func TestCompareJson(t *testing.T) {
	t.Run("Section", sectionTest(testConn))
	t.Run("SectionAll", sectionAllTest(testConn))
	t.Run("Metadata Top", metadataTop(testConn, movie_meta_data_id))
	t.Run("Metadata Show", metadataTop(testConn, tv_top_meta_data_id))
	t.Run("Metadata Season", metadataTop(testConn, tv_season_meta_data_id))
	t.Run("Metadata Episode", metadataTop(testConn, tv_episode_meta_data_id))

	t.Run("Metadata Children show", metadataChildren(testConn, tv_top_meta_data_id))
	t.Run("Metadata Children season", metadataChildren(testConn, tv_season_meta_data_id))
	t.Run("Sessions", serverSessions(testConn))
}
