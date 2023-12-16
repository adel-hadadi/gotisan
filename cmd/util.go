package cmd

import (
	"fmt"
	"github.com/adel-hadadi/gotisan/constant"
	"os"
)

var templates map[string]string = map[string]string{
	"handler": constant.HandlerSample,
	"model":   constant.ModelSample,
}

func copyTemplateFiles() error {
	for k, v := range templates {
		file, err := os.Create(fmt.Sprintf(".gotisan/templates/%s.tmpl", k))
		if err != nil {
			return err
		}

		_, err = file.WriteString(v)
		if err != nil {
			return err
		}
	}

	return nil
}
