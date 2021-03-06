package govalidator

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// validateFiles validate file size, mimes, extension etc
func validateFiles(r *http.Request, field, rule, msg string, errsBag url.Values) {
	fExist, fn, ext, mime, size, fErr := getFileInfo(r, field)
	if fExist {
		r.Form.Add(field, "$^__file__exist__&#")
	}
	// check size
	if strings.HasPrefix(rule, "size:") {
		l, err := strconv.ParseInt(strings.TrimPrefix(rule, "size:"), 10, 64)
		if err != nil {
			panic(errStringToInt)
		}
		if size > l {
			if msg != "" {
				errsBag.Add(field, msg)
			} else {
				errsBag.Add(field, fmt.Sprintf("The %s field size is can not be greater than %d bytes", field, l))
			}
		}
		if fErr != nil {
			errsBag.Add(field, fmt.Sprintf("The %s field file %s failed to read", field, fn))
		}
	}

	// check extension
	if strings.HasPrefix(rule, "ext:") {
		exts := strings.Split(strings.TrimPrefix(rule, "ext:"), ",")
		f := false
		for _, e := range exts {
			if e == ext {
				f = true
			}
		}
		if !f {
			if msg != "" {
				errsBag.Add(field, msg)
			} else {
				errsBag.Add(field, fmt.Sprintf("The %s field file extension %s is invalid", field, ext))
			}
		}
		if fErr != nil {
			errsBag.Add(field, fmt.Sprintf("The %s field file %s failed to read", field, fn))
		}
	}

	// check mimes
	if strings.HasPrefix(rule, "mime:") {
		mimes := strings.Split(strings.TrimPrefix(rule, "mime:"), ",")
		f := false
		for _, m := range mimes {
			if m == mime {
				f = true
			}
		}
		if !f {
			if msg != "" {
				errsBag.Add(field, msg)
			} else {
				errsBag.Add(field, fmt.Sprintf("The %s field file mime %s is invalid", field, mime))
			}
		}
		if fErr != nil {
			errsBag.Add(field, fmt.Sprintf("The %s field file %s failed to read", field, fn))
		}
	}
}
