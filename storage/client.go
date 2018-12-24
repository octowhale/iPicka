package storage

import (
	"errors"

	"github.com/octowhale/iPicka/storage/aliyunoss"
)

type StorageClient interface {
	Put(object, filepath string) error
}

func New(config Config) (StorageClient, error) {

	// var client StorageClient
	switch config.Storage {
	case "aliyunoss":
		client, err := aliyunoss.NewAliyunOSSClient(config.AccKey, config.AccSec, config.Region, config.Bucket, config.Internal)
		// client = aliyunoss.NewAliyunOSSClient()
		return client, err
	}

	return nil, errors.New("Invalid backend")
}
