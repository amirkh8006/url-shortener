package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService *StoreService

func init()  {
	testStoreService = initStoreService()
}

func TestInsertaitionAndRetrieve(t *testing.T){
	initialLink := "https://google.com"
	userUUID := "2e5793df-b86b-4e6a-bda9-90f364547a5b"
	shortUrl := "BPiZViQX24DSqTDLv7X"

	SaveUrlMapping(shortUrl, initialLink , userUUID)

	retreievedUrl := RetrieveInitialUrl(shortUrl)

	assert.Equal(t, initialLink, retreievedUrl)
}