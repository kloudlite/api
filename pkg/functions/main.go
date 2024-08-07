package functions

import (
	"encoding/base64"
	"encoding/json"
	"regexp"
	"strings"

	"github.com/kloudlite/api/pkg/errors"
	nanoid "github.com/matoous/go-nanoid/v2"
)

func ToBase64StringFromJson(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", errors.NewE(err)
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func ToBase64UrlFromJson(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", errors.NewE(err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(errors.NewEf(err, "panicking as Must() check failed"))
	}
	return value
}

var re = regexp.MustCompile(`(\W|_)+`)

func CleanerNanoid(n int) (string, error) {
	id, e := nanoid.New(n)
	if e != nil {
		return "", errors.NewEf(e, "could not get nanoid()")
	}
	res := re.ReplaceAllString(id, "-")
	if strings.HasPrefix(res, "-") {
		res = "k" + res
	}
	if strings.HasSuffix(res, "-") {
		res = res + "k"
	}
	return res, nil
}

func CleanerNanoidOrDie(n int) string {
	id, err := CleanerNanoid(n)
	if err != nil {
		panic(err)
	}
	return id
}

// UUID returns a UUID string of the given size, if specified, otherwise a default size of 16 is used.
func UUID(size ...int) string {
	if len(size) > 0 {
		return nanoid.Must(size[0])
	}
	return nanoid.Must(16)
}

func JsonConversion(from any, to any) error {
	if from == nil {
		return nil
	}

	if to == nil {
		return errors.Newf("receiver (to) is nil")
	}

	b, err := json.Marshal(from)
	if err != nil {
		return nil
	}
	if err := json.Unmarshal(b, &to); err != nil {
		return errors.NewE(err)
	}
	return nil
}

func JsonConvertP[T any](from any) (*T, error) {
	var to T
	if from == nil {
		return nil, nil
	}

	b, err := json.Marshal(from)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &to); err != nil {
		return nil, errors.NewE(err)
	}
	return &to, nil
}

func JsonConvert[T any](from any) (T, error) {
	var to T
	if from == nil {
		return to, nil
	}

	b, err := json.Marshal(from)
	if err != nil {
		return to, err
	}

	if err := json.Unmarshal(b, &to); err != nil {
		return to, errors.NewE(err)
	}
	return to, nil
}
