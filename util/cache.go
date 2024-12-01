package util

import (
	"fmt"
	"os"
	"path/filepath"
)

const CacheDir = "cache"

func ensureCacheDir() error {
	if _, err := os.Stat(CacheDir); os.IsNotExist(err) {
		return os.MkdirAll(CacheDir, 0755)
	}
	return nil
}

func GetCachedInput(day int) (string, error) {
	if err := ensureCacheDir(); err != nil {
		return "", fmt.Errorf("failed to ensure cache directory: %w", err)
	}

	cacheFile := filepath.Join(CacheDir, fmt.Sprintf("day%d.txt", day))
	data, err := os.ReadFile(cacheFile)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", fmt.Errorf("failed to read cache file: %w", err)
	}

	return string(data), nil
}

func SaveInputToCache(day int, input string) error {
	if err := ensureCacheDir(); err != nil {
		return fmt.Errorf("failed to ensure cache directory: %w", err)
	}

	cacheFile := filepath.Join(CacheDir, fmt.Sprintf("day%d.txt", day))
	err := os.WriteFile(cacheFile, []byte(input), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to cache file: %w", err)
	}

	return nil
}
