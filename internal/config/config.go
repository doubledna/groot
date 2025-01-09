package config

import (
	"fmt"
	"github.com/spf13/pflag"
	c "github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	File string
)

func init() {
	pflag.StringVarP(&File, "config.file", "c", "", "specify config file path")
	loadConfigFile()
}

// loadConfigFile load configuration file
func loadConfigFile() {
	// if a configuration file is specified, use the specified
	if File != "" {
		c.SetConfigFile(File)
	} else {
		// if not a configuration file is specified, go to the conf directory to search according to the environment
		// variable
		env := strings.ToLower(os.Getenv("CONF_ENV"))
		if env == "" {
			env = "dev"
		}

		fileName := fmt.Sprintf("%s.yaml", env)
		file, err := findConfigFile(fileName, "./conf", "../conf", filepath.Join(getRootDir(), "./conf"))
		if err != nil {
			panic(err)
		}

		c.SetConfigFile(file)
	}
	err := c.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func getRootDir() string {
	_, f, _, _ := runtime.Caller(0)
	root := path.Join(path.Dir(f), "../..")
	return root
}

func findConfigFile(filename string, paths ...string) (string, error) {
	for _, p := range paths {
		file := filepath.Join(p, filename)
		if ok, _ := exists(file); ok {
			return file, nil
		}
	}
	return "", fmt.Errorf("file %s not found in %v", filename, paths)
}

func exists(filename string) (bool, error) {
	stat, err := os.Stat(filename)
	if err == nil {
		return !stat.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Get can retrieve any value given the key to use.
// Get is case-insensitive for a key.
// Get has the behavior of returning the value associated with the first
// place from where it is set. Viper will check in the following order:
// override, flag, env, config file, key/value store, default
//
// Get returns an interface. For a specific value use one of the Get____ methods.
func Get(key string) interface{} {
	return c.Get(key)
}

// GetString returns the value associated with the key as a string.
func GetString(key string) string {
	return c.GetString(key)
}

// GetBool returns the value associated with the key as a boolean.
func GetBool(key string) bool {
	return c.GetBool(key)
}

// GetInt returns the value associated with the key as an integer.
func GetInt(key string) int {
	return c.GetInt(key)
}

// GetInt32 returns the value associated with the key as an integer.
func GetInt32(key string) int32 {
	return c.GetInt32(key)
}

// GetInt64 returns the value associated with the key as an integer.
func GetInt64(key string) int64 {
	return c.GetInt64(key)
}

// GetUint returns the value associated with the key as an unsigned integer.
func GetUint(key string) uint {
	return c.GetUint(key)
}

// GetUint32 returns the value associated with the key as an unsigned integer.
func GetUint32(key string) uint32 {
	return c.GetUint32(key)
}

// GetUint64 returns the value associated with the key as an unsigned integer.
func GetUint64(key string) uint64 {
	return c.GetUint64(key)
}

// GetFloat64 returns the value associated with the key as a float64.
func GetFloat64(key string) float64 {
	return c.GetFloat64(key)
}

// GetTime returns the value associated with the key as time.
func GetTime(key string) time.Time {
	return c.GetTime(key)
}

// GetDuration returns the value associated with the key as a duration.
func GetDuration(key string) time.Duration {
	return c.GetDuration(key)
}

// GetIntSlice returns the value associated with the key as a slice of int values.
func GetIntSlice(key string) []int {
	return c.GetIntSlice(key)
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func GetStringSlice(key string) []string {
	return c.GetStringSlice(key)
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func GetStringMap(key string) map[string]interface{} {
	return c.GetStringMap(key)
}

// GetStringMapString returns the value associated with the key as a map of strings.
func GetStringMapString(key string) map[string]string {
	return c.GetStringMapString(key)
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func GetStringMapStringSlice(key string) map[string][]string {
	return c.GetStringMapStringSlice(key)
}

// GetSizeInBytes returns the size of the value associated with the given key
// in bytes.
func GetSizeInBytes(key string) uint {
	return c.GetSizeInBytes(key)
}
