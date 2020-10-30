/*
@Time : 2020/10/29 4:59 下午
@Author : mmy83
@File : envconfig
@Software: GoLand
*/

package envconfig

import (
	"github.com/joho/godotenv"
	"github.com/wonderivan/logger"
	"os"
	"reflect"
	"strconv"
)

func GetConfig(ptr interface{},filenames ...string){
	godotenv.Load(filenames...)
	v:= reflect.ValueOf(ptr).Elem()
	if v.Kind().String() == "struct" {

		for i := 0; i < v.NumField(); i++ {

			tagName := v.Type().Field(i).Tag.Get("envconfig")
			logger.Debug("tag :",tagName)
			if tagName!="" {

				tagVal := os.Getenv(tagName)
				logger.Debug("env :",tagVal)
				attrType := v.Field(i).Kind().String()
				logger.Debug(attrType)
				switch attrType {
				case "int32":
					num,err:=strconv.ParseInt(tagVal, 10, 64)
					if err!=nil{
						logger.Fatal("string conv int64 error!")
					}
					v.Field(i).SetInt(num)
				case "string":
					v.Field(i).SetString(tagVal)
				case "bool":
					boolVal, err := strconv.ParseBool(tagVal)
					if err!=nil{
						logger.Fatal("string conv to bool error!")
					}
					v.Field(i).SetBool(boolVal)
				case "float64":
					boolVal, err := strconv.ParseFloat(tagVal,64)
					if err!=nil{
						logger.Fatal("string conv to bool error!")
					}
					v.Field(i).SetFloat(boolVal)
				default:
					logger.Warn("Unsupported type！")
				}

			}
		}
	}
	logger.Debug(ptr)

}