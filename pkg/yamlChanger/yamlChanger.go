package yamlChanger

import (
	"errors"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// ErrAlreadyUpdated is a sentinel error returned when no change was made
// because the value of the node already matches the given new value.
var ErrAlreadyUpdated = errors.New("node value already up to date")

type PathError struct{}

func (p *PathError) Error() string {
	return fmt.Sprintf("valuePath or part of it is empty")
}

func GetPathSplits(path string) (res []string, err error) {
	splits := strings.Split(path, ".")
	if len(splits) == 0 {
		return nil, &PathError{}
	}
	for _, value := range splits {
		if value == "" {
			return nil, &PathError{}
		}
		res = append(res, value)
	}

	return res, nil
}

func findNodeValue(node *yaml.Node, path []string) *yaml.Node {
	found := false
	for _, n := range node.Content {
		if found == true {
			if len(path) == 1 {
				return n
			}
			return findNodeValue(n, path[1:])
		}
		if n.Value == path[0] {
			found = true
		}
	}

	return nil
}

func ChangeYaml(body *yaml.Node, newValue string, path []string) error {
	node := findNodeValue(body.Content[0], path)

	if node == nil {
		return errors.New("path not found")
	}

	node.Style = yaml.DoubleQuotedStyle
	node.Tag = ""
	if node.Value == newValue {
		return ErrAlreadyUpdated
	}
	node.Value = newValue

	return nil
}
