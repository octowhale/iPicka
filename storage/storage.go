package storage

import (
	"errors"

	"github.com/octowhale/iPicka/storage/aliyunoss"
)

type StorageClient interface {
	Put(object, file string) error
}

func New(config *Config) (StorageClient, error) {

	switch config.Driver {

	case "aliyunoss":
		return aliyunoss.NewOSSClient(
			config.Acckey,
			config.Accsec,
			config.Bucketname,
			config.Region,
			config.Internal,
		), nil

	case "qcloudcos":
		return nil, errors.New("Storage Qcloud COS Invalid")

	}
	return nil, errors.New("Invalid Storage Client")
}
