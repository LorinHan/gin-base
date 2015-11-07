package serivces

import (
	"errors"
	"fmt"
	"gin-base/models"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

type ts struct {
	Name string
}

func trans(val interface{}) (*ts, error) {
	switch val.(type) {
	case ts:
		if t, ok := val.(ts); ok {
			return &t, nil
		}
	}
	return nil, errors.New("trans error")
}

func TestCurrentFile(t *testing.T) {
	t.Run("te", func(t *testing.T) {
		curPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		fmt.Println(curPath)
		_, file, _, ok := runtime.Caller(1)
		if !ok {
			panic(errors.New("Can not get current file info"))
		}
		fmt.Println(file)
	})
}

func TestTrans(t *testing.T) {
	t.Run("TestTrans", func(t *testing.T) {
		entity := ts{"Lorin"}
		if entityCopy, err := trans(entity); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(reflect.TypeOf(entityCopy))
		}
	})
}

func Test_userService_FindUserByPwd(t *testing.T) {
	t.Run("FindUserByPwd", func(t *testing.T) {
		if users, err := FindUserByPwd("123456"); err != nil {
			log.Println(err)
		} else {
			for i := range users {
				log.Println(users[i])
			}
		}
	})
}

func TestRowsToMaps(t *testing.T) {
	t.Run("test rows to maps", func(t *testing.T) {
		result := models.TestRows2Maps(0, 10)
		fmt.Println(result)
		for i, val := range *result {
			fmt.Println(i, *val)
		}
	})
}
