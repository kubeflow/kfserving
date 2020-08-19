package agent

import (
	"fmt"
	pro "github.com/kubeflow/kfserving/pkg/agent/protocols"
	"hash/fnv"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Downloader struct {
	// TODO: Add back-off retries
	ModelDir   string
	ProtocolManagers map[pro.Protocol]pro.ProtocolManager
}

var SupportedProtocols = []pro.Protocol{pro.S3}

func (d *Downloader) DownloadModel(event EventWrapper) error {
	modelSpec := event.ModelSpec
	if modelSpec != nil {
		modelUri := modelSpec.StorageURI
		hashString := hash(modelUri)
		log.Println("Processing:", modelUri, "=", hashString)
		successFile := filepath.Join(d.ModelDir, fmt.Sprintf("SUCCESS.%d", hashString))
		if !pro.FileExists(successFile) {
			downloadErr := d.download(modelUri)
			if downloadErr != nil {
				return fmt.Errorf("download error: %v", downloadErr)
			} else {
				file, createErr := os.Create(successFile)
				if createErr != nil {
					return fmt.Errorf("create file error: %v", createErr)
				}
				defer file.Close()
			}
		} else {
			log.Println("Model", modelSpec.StorageURI, "exists already")
		}
	}
	return nil
}

func (d* Downloader) download(storageUri string) error {
	log.Println("Downloading: ", storageUri)
	protocol, err := validateStorageURI(storageUri)
	if err != nil {
		return fmt.Errorf("unsupported protocol: %v", err)
	}
	manager, ok := d.ProtocolManagers[protocol]
	if !ok {
		return fmt.Errorf("protocol manager for %s is not initialized", protocol)
	}
	if err = manager.Download(d.ModelDir, storageUri); err != nil {
		return fmt.Errorf("failure on download: %v", err)
	}

	return nil
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func validateStorageURI(storageURI string) (pro.Protocol, error) {
	if storageURI == "" {
		return "", fmt.Errorf("there is no storageUri supplied")
	}

	if !regexp.MustCompile("\\w+?://").MatchString(storageURI) {
		return "", fmt.Errorf("there is no protocol specificed for the storageUri")
	}

	for _, prefix := range SupportedProtocols {
		if strings.HasPrefix(storageURI, string(prefix)) {
			return prefix, nil
		}
	}
	return "", fmt.Errorf("protocol not supported for storageUri")
}