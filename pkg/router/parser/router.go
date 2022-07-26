package parser

import (
	. "github.com/Gebes/there/v2"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func IntParameter(param string, request HttpRequest) (int, error) {
	p, ok := request.Params.Get(param)
	if !ok {
		return 0, errors.New("parameter " + param + " is required")
	}
	num, err := strconv.Atoi(p)
	if err != nil {
		return 0, errors.Wrap(err, "could not parse parameter "+param)
	}
	return num, nil
}

func IntRouteParameter(param string, request HttpRequest) (int, error) {
	p, ok := request.RouteParams.Get(param)
	if !ok {
		return 0, errors.New("parameter " + param + " is required")
	}
	num, err := strconv.Atoi(p)
	if err != nil {
		return 0, errors.Wrap(err, "could not parse parameter "+param)
	}
	return num, nil
}

func IntParameters(param string, request HttpRequest) ([]int, error) {
	p, ok := request.Params.Get(param)
	if !ok {
		return nil, errors.New("parameter " + param + " is required")
	}

	split := strings.Split(p, ",")
	nums := make([]int, len(split))

	for i, v := range split {
		parsed, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.Wrap(err, "could not parse number \""+v+"\"")
		}
		nums[i] = parsed
	}

	return nums, nil
}
