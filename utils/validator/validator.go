package validator

import (
	"fmt"
	"go-forum-api/app/models"
	"regexp"
	"strconv"
	"sync"
)

const slugRegEx = "^(\\d|\\w|-|_)*(\\w|-|_)(\\d|\\w|-|_)*$"

type Validator struct {
	slugRegExCompiled *regexp.Regexp
}

var validatorInstanceLock = &sync.Mutex{}
var validatorInstance *Validator

func GetInstance() (*Validator, error) {
	if validatorInstance == nil {
		validatorInstanceLock.Lock()
		defer validatorInstanceLock.Unlock()
		if validatorInstance == nil {
			var err error
			validatorInstance, err = createValidator()
			if err != nil {
				return nil, err
			}
		}
	}
	return validatorInstance, nil
}

func createValidator() (validator *Validator, err error) {
	validator = &Validator{}
	validator.slugRegExCompiled, err = regexp.Compile(slugRegEx)
	if err != nil {
		return nil, err
	}
	return
}

func (validator *Validator) ValidateSlug(slug string) bool {
	return validator.slugRegExCompiled.MatchString(slug)
}

func (validator *Validator) ValidateForumQuery(query *models.ForumQueryParams) {
	if query.Limit == 0 {
		query.Limit = 100
	}
}

func (validator *Validator) GetSlugOrIdOrErr(slugOrId string) (slug string, id int, err error) {
	if slugOrId == "" {
		err = fmt.Errorf("пустой slug or id")
		return
	}

	id, err = strconv.Atoi(slugOrId)
	if err == nil {
		return
	}

	if validator.ValidateSlug(slugOrId) == false {
		err = fmt.Errorf("неверный slug or id")
		return
	}

	err = nil
	slug = slugOrId
	return
}
