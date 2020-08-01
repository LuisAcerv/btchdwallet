package db

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/LuisAcerv/btchdwallet/models"
	"github.com/boltdb/bolt"
)

// Initialize database
func Initialize() (*bolt.DB, error) {
	bucket := []byte("WALLETS")

	// It will be created if it doesn't exist.
	db, err := bolt.Open("repository.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB Setup Done")
	return db, nil
}

// SaveWallet is in charge of saving the encrypted wallet to db
func SaveWallet(name string, key []byte, wallet string) error {
	bucket := []byte("WALLETS")

	db, err := bolt.Open("repository.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	createdAt := time.Now().Format("01-01-2020")
	entry := models.EncryptedWallet{Name: name, Encrypted: wallet, CreatedAt: string(createdAt)}

	entryBytes, _ := json.Marshal(entry)

	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket(bucket).Put(key, entryBytes)
		if err != nil {
			return fmt.Errorf("could not insert entry: %v", err)
		}

		return nil
	})

	fmt.Println()
	fmt.Println("Saved")
	fmt.Println()

	return err
}

// GetWallet by name
func GetWallet(name []byte) *models.EncryptedWallet {
	bucket := []byte("WALLETS")
	db, err := bolt.Open(fmt.Sprintf("repository.db"), 0600, &bolt.Options{Timeout: 10 * time.Second})
	defer db.Close()
	if err != nil {
		panic(err)
	}

	repo := &models.EncryptedWallet{}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		if b == nil {
			fmt.Println("Bucket SECRETS not found!")
			return nil
		}
		r := b.Get(name)

		err := json.Unmarshal(r, repo)
		if err != nil {
			return nil
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Readed secret")
	return repo
}
