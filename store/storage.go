package store

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
)

const (
	DB_NAME = "notifications.db"
)

func AddNotification(notification NotificationEvent) {
	db, err := bolt.Open(DB_NAME, 0600, nil)
	if err != nil {
		fmt.Errorf("Error opening database, %s", err)
		return
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(notification.Provider))
		if err != nil {
			return fmt.Errorf("Error opening/creating bucket: %s", err)
		}
		key := notification.Id
		value, err := json.Marshal(notification)

		bucket.Put([]byte(key), value)

		return nil
	})

	if err != nil {
		fmt.Errorf("Error updating database %s", err)
	}
}

func NotificationExist(notification NotificationEvent) bool {
	db, err := bolt.Open(DB_NAME, 0600, nil)
	if err != nil {
		fmt.Errorf("Error opening database, %s", err)
		return false
	}
	defer db.Close()

	tx, err := db.Begin(false)
	defer tx.Rollback()

	if err != nil {
		fmt.Errorf("Error creating read-only transaction", err)
		return false
	}

	bucket := tx.Bucket([]byte(notification.Provider))
	if bucket == nil {
		return false
	}

	return len(bucket.Get([]byte(notification.Id))) > 0
}
