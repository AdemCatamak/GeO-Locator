package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"sync"
)

type IConfigManager interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	GetObj(key string, obj any)
}

var (
	configManager IConfigManager
	lock          sync.Mutex
)

func GetConfigManager() IConfigManager {
	if configManager == nil {
		lock.Lock()
		defer lock.Unlock()
		if configManager == nil {
			log.Println("Creating configManager instance now.")
			configManager = newViperConfigManager()
		} else {
			fmt.Println("configManager instance already created.")
		}
	} else {
		fmt.Println("configManager instance already created.")
	}

	return configManager
}

type viperConfigManager struct {
}

func newViperConfigManager() *viperConfigManager {
	log.Print("newViperConfigManager function was triggered")

	loadConfig()

	// Watch for changes in the config file
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
		loadConfig()
	})

	log.Print("loadConfig function was completed")

	manager := viperConfigManager{}
	return &manager
}

func loadConfig() {
	viper.SetConfigFile("configs/config.default.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	viper.AutomaticEnv()
}

func (c viperConfigManager) GetString(key string) string {
	value := viper.GetString(key)
	return value
}

func (c viperConfigManager) GetInt(key string) int {
	value := viper.GetInt(key)
	return value
}

func (c viperConfigManager) GetBool(key string) bool {
	value := viper.GetBool(key)
	return value
}

func (c viperConfigManager) GetObj(key string, obj any) {
	err := viper.UnmarshalKey(key, obj)
	if err != nil {
		panic(err)
	}
}
