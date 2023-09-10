//go:build db

package storage_test

// TODO: Implement mock DB for testing, end-to-end testing is not ideal
import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"github.com/mjmoshiri/log-lyfe/gol/internal/storage"
	"github.com/mjmoshiri/log-lyfe/gol/internal/utils"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func getTestDBConfig() models.DBConfig {
	return models.DBConfig{
		Cluster:     []string{"localhost"},
		Keyspace:    "testkeyspace",
		Consistency: gocql.LocalOne,
		Username:    "cassandra",
		Password:    "cassandra",
		Table:       "events",
	}
}

func TestNew(t *testing.T) {
	cfg := getTestDBConfig()
	db, err := storage.New(cfg)
	assert.Nil(t, err)
	assert.NotNil(t, db)
	db.Close()
}

func TestInsertOne(t *testing.T) {
	cfg := getTestDBConfig()
	db, _ := storage.New(cfg)
	defer db.Close()
	events := storage.Events
	// 1. Insert New
	t.Run("Insert", func(t *testing.T) {
		err := db.Insert(&events[0])
		assert.Nil(t, err)
	})
}

func TestFind(t *testing.T) {
	cfg := getTestDBConfig()
	db, _ := storage.New(cfg)
	defer db.Close()
	events := storage.Events[1:]
	for _, event := range events {
		err := db.Insert(&event)
		if err != nil {
			t.Fatal(err)
		}
	}
	// 1. Find Using Action Filter
	t.Run("Find Using Action Filter", func(t *testing.T) {
		filters := map[string]interface{}{"action": `'create'`}
		results, err := db.Find(filters, 0)
		fmt.Println(err)
		assert.Nil(t, err)
		assert.Equal(t, 5, len(results))
		for _, result := range results {
			fmt.Println(result)
		}
	})
	// 2. Find Using Actor Filter
	t.Run("Find Using Actor Filter", func(t *testing.T) {
		filters := map[string]interface{}{"actor": `'user1'`}
		results, err := db.Find(filters, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(results))
		for _, result := range results {
			fmt.Println(result)
		}
	})
	// 3. Wrong Filter
	t.Run("Wrong Filter", func(t *testing.T) {
		filters := map[string]interface{}{"wrong": "filter"}
		_, err := db.Find(filters, 0)
		assert.NotNil(t, err)
	})
	// 5. Fetch Size
	t.Run("Fetch Size", func(t *testing.T) {
		filters := map[string]interface{}{"action": `'create'`}
		results, err := db.Find(filters, 3)
		assert.Nil(t, err)
		assert.Equal(t, 3, len(results))
	})
	// 6. Fetch Size > 10000
	t.Run("Fetch Size > 10000", func(t *testing.T) {
		filters := map[string]interface{}{"action": `'create'`}
		results, err := db.Find(filters, 10001)
		assert.NotNil(t, err)
		assert.Equal(t, 0, len(results))
	})
	// 7. Multiple Filters
	t.Run("Multiple Filters", func(t *testing.T) {
		filters := map[string]interface{}{"action": `'create'`, "bucket": strconv.FormatInt(utils.TimeToBucket(events[7].Timestamp), 10)}
		results, err := db.Find(filters, 0)
		assert.Nil(t, err)
		assert.Equal(t, 3, len(results))
	})
}
