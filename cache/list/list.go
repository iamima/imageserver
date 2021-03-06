// Package list provides a list of Image Cache
package list

import (
	"github.com/pierrre/imageserver"
)

// ListCache represents a list of Image Cache
type ListCache []imageserver.Cache

// Get gets an Image from caches in sequential order
//
// If an Image is found, previous caches are filled
func (cache ListCache) Get(key string, parameters imageserver.Parameters) (*imageserver.Image, error) {
	for i, c := range cache {
		image, err := c.Get(key, parameters)
		if err == nil {
			if i > 0 {
				err = cache.set(key, image, parameters, i)
				if err != nil {
					return nil, err
				}
			}
			return image, nil
		}
	}

	return nil, imageserver.NewCacheMissError(key, cache, nil)
}

func (cache ListCache) set(key string, image *imageserver.Image, parameters imageserver.Parameters, indexLimit int) error {
	for i := 0; i < indexLimit; i++ {
		err := cache[i].Set(key, image, parameters)
		if err != nil {
			return err
		}
	}
	return nil
}

// Set sets the image to all caches
func (cache ListCache) Set(key string, image *imageserver.Image, parameters imageserver.Parameters) error {
	return cache.set(key, image, parameters, len(cache))
}
