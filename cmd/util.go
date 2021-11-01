package cmd

import (
	"fmt"
	"regexp"
	"strings"
)

func parseTagsFlag(tagsInput []string) (map[string]string, error) {
	result := map[string]string{}

	for _, t := range tagsInput {
		match, err := regexp.MatchString("^([^=])+=([^=])+$", t)
		if err != nil {
			return nil, err
		}
		if !match {
			return nil, fmt.Errorf("Invalid tag format: %s", t)
		}

		tag := strings.Split(t, "=")
		result[tag[0]] = tag[1]
	}
	return result, nil
}
