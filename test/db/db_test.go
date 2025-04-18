package db

import (
    "encoding/json"
    "gourd/internal/config"
    "gourd/internal/global"
    "gourd/internal/initialize"
    "gourd/internal/orm/query"
    "testing"
)

func TestDB(t *testing.T) {

    err := config.SetConfigPath("../configs")
    if err != nil {
        t.Error(err)
        return
    }

    err = initialize.InitDatabase()
    if err != nil {
        t.Error(err)
        return
    }

    takeUser, err := query.Q.User.Take()
    if err != nil {
        t.Error(err)
        return
    }

    jsonStr, err := json.Marshal(takeUser)
    if err != nil {
        t.Error(err)
        return
    }

    t.Logf("User %s", jsonStr)

    t.Log(global.GetDb("mysql"))

}
