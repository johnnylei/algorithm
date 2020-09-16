package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"icode.baidu.com/baidu/gdp/gdp"
	"icode.baidu.com/baidu/inputmethod/servergoapi/library/ecode"
	"icode.baidu.com/baidu/inputmethod/servergoapi/models/review"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync"
)

var (
	defaultValidate    *ValidateEngine
	TextReviewValidate *validator
	lock               sync.RWMutex
)

func RegisterValidate(name string, v *validator) error {
	validators := defaultValidate.validators
	lock.RLock()
	if _, okay := validators[name]; okay {
		lock.RUnlock()
		return fmt.Errorf("%s repeated", name)
	}
	lock.RUnlock()

	lock.Lock()
	defer lock.Unlock()
	validators[name] = v
	return nil
}

type validatorErrHandler func(interface{}, interface{}, error) error

type validator struct {
	handler    func(interface{}, interface{}) error
	errHandler map[string]validatorErrHandler
	lock       sync.RWMutex
}

func (v *validator) RegisterErrHandler(key string, handler validatorErrHandler) error {
	v.lock.RLock()
	if _, okay := v.errHandler[key]; okay {
		v.lock.RUnlock()
		return fmt.Errorf("%s has been registered", key)
	}
	v.lock.RUnlock()

	v.lock.Lock()
	defer v.lock.Unlock()
	if v.errHandler == nil {
		v.errHandler = make(map[string]validatorErrHandler)
	}
	v.errHandler[key] = handler
	return nil
}

type ValidateEngine struct {
	Tag        string
	validators map[string]*validator
}

func (v *ValidateEngine) ValidateStruct(object interface{}) error {
	value := reflect.ValueOf(object)

	switch value.Kind() {
	case reflect.Ptr:
		value = value.Elem()
		return v.ValidateStruct(value.Interface())

	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			tag := value.Type().Field(i).Tag
			tagValue, okay := tag.Lookup(v.Tag)
			if !okay {
				continue
			}

			field := value.Field(i)
			switch field.Kind() {
			case reflect.Ptr:
				if field.Elem().Kind() == reflect.Struct {
					if err := v.ValidateStruct(field.Elem().Interface()); err != nil {
						return err
					}
				} else {
					validatorKeys := strings.Split(tagValue, ",")
					for i := 0; i < len(validatorKeys); i++ {
						validator, okay := v.validators[validatorKeys[i]]
						if !okay {
							continue
						}

						if err := validator.handler(object, field.Elem().Interface()); err != nil {
							return err
						}
					}
				}
			case reflect.Struct:
				if err := v.ValidateStruct(field.Interface()); err != nil {
					return err
				}
			case reflect.Slice:
				for i := 0; i < field.Len(); i++ {
					if err := v.ValidateStruct(field.Index(i).Interface()); err != nil {
						return err
					}
				}
			default:
				validatorKeys := strings.Split(tagValue, ",")
				for i := 0; i < len(validatorKeys); i++ {
					validator, okay := v.validators[validatorKeys[i]]
					if !okay {
						continue
					}

					if err := validator.handler(object, field.Interface()); err != nil {
						return err
					}
				}
			}
		}

	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			if err := v.ValidateStruct(value.Index(i).Interface()); err != nil {
				return err
			}
		}
		return nil
	}

	return nil
}

func (v *ValidateEngine) Engine() interface{} {
	return v
}

func getContext() *gdp.WebContext {
	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	ginContext.Request, _ = http.NewRequest("GET", "/get", nil)
	return gdp.NewWebContext(ginContext)
}

func init() {
	TextReviewValidate = &validator{
		handler: func(object interface{}, field interface{}) error {
			text, okay := field.(string)
			if !okay {
				return errors.New("invalid type")
			}

			if len(text) == 0 {
				return nil
			}

			if okay, err := review.IsValidText(getContext(), []string{
				text,
			}); !okay {
				objectType := reflect.TypeOf(object)
				objectTypeName := objectType.Name()
				if fn, okay := TextReviewValidate.errHandler[objectTypeName]; okay {
					return fn(object, field, err)
				}

				//return fmt.Errorf("\"%s\" text review denined", text)
				return ecode.TextCheckError
			}

			return nil
		},
	}
	defaultValidate = &ValidateEngine{
		validators: map[string]*validator{
			"text_review": TextReviewValidate,
		},
		Tag: "binding",
	}
	binding.Validator = defaultValidate
}
