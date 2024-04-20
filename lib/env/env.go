package env

import (
	"os"
	"runtime"
	"time"

	"github.com/spf13/viper"
)

type Configuration interface {
	SetDefault(key string, value interface{})

	GetString(key string) string

	GetBool(key string) bool

	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64

	GetUint(key string) uint
	GetUint32(key string) uint32
	GetUint64(key string) uint64

	GetFloat64(key string) float64

	GetTime(key string) time.Time
	GetDuration(key string) time.Duration

	GetInts(key string) []int
	GetStrings(key string) []string

	GetVersion() *Version
}

type configuration struct {
	Version *Version
}

type Version struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	BuildDate string `json:"build_date"`
	Go        string `json:"go_version"`
}

func NewConfiguration(appName string) Configuration {
	viper.AutomaticEnv()

	configuration := &configuration{}

	version := &Version{
		Name: appName,
		Go:   runtime.Version(),
	}

	configuration.Version = version

	return configuration
}

func (c *configuration) SetDefault(key string, value interface{}) { viper.SetDefault(key, value) }

func (c *configuration) GetString(key string) string { return viper.GetString(key) }

func (c *configuration) GetBool(key string) bool { return viper.GetBool(key) }

func (c *configuration) GetInt(key string) int     { return viper.GetInt(key) }
func (c *configuration) GetInt32(key string) int32 { return viper.GetInt32(key) }
func (c *configuration) GetInt64(key string) int64 { return viper.GetInt64(key) }

func (c *configuration) GetUint(key string) uint     { return viper.GetUint(key) }
func (c *configuration) GetUint32(key string) uint32 { return viper.GetUint32(key) }
func (c *configuration) GetUint64(key string) uint64 { return viper.GetUint64(key) }

func (c *configuration) GetFloat64(key string) float64 { return viper.GetFloat64(key) }

func (c *configuration) GetTime(key string) time.Time         { return viper.GetTime(key) }
func (c *configuration) GetDuration(key string) time.Duration { return viper.GetDuration(key) }

func (c *configuration) GetInts(key string) []int       { return viper.GetIntSlice(key) }
func (c *configuration) GetStrings(key string) []string { return viper.GetStringSlice(key) }

func (c *configuration) GetVersion() *Version { return c.Version }

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
