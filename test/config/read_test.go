package config

import (
    "app/internal/config"
    "testing"
)

func TestReadConfig(t *testing.T) {
    dbConfig := &config.DbConfigMap{}

    err := config.SetConfigPath("./configs")
    if err != nil {
        t.Errorf("set config path error: %v", err)
        return
    }

    err = config.Unmarshal("database", dbConfig)
    if err != nil {
        t.Errorf("Unmarshal error: %v", err)
        return
    }

    t.Log(dbConfig)

}
