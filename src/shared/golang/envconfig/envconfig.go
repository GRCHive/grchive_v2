package envconfig

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var int32Type = reflect.TypeOf((int32)(0))
var stringType = reflect.TypeOf((string)(""))

func Load(output interface{}) error {
	interfaceType := reflect.TypeOf(output).Elem()
	interfaceValue := reflect.ValueOf(output).Elem()

	for i := 0; i < interfaceType.NumField(); i++ {
		fieldType := interfaceType.Field(i)
		fieldValue := interfaceValue.Field(i)

		// Deliberate design choice to only handle fields with an explicit env tag.
		tag, ok := fieldType.Tag.Lookup("env")
		if !ok {
			continue
		}

		splitTag := strings.Split(tag, ",")
		fieldName := splitTag[0]
		required := false

		if len(splitTag) > 1 {
			required = splitTag[1] == "required"
		}

		val, ok := os.LookupEnv(fieldName)
		if !ok {
			if required {
				return errors.New("Failed to find required environment variable: " + fieldName)
			} else {
				continue
			}
		}

		var dataValue reflect.Value
		switch fieldType.Type {
		case int32Type:
			cv, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				return err
			}
			dataValue = reflect.ValueOf(int32(cv))
		case stringType:
			dataValue = reflect.ValueOf(val)
		}

		fieldValue.Set(dataValue)
	}
	return nil
}
