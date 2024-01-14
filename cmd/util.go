package cmd

import (
	"errors"
	"fmt"
	"github.com/adel-hadadi/gotisan/config"
	"github.com/adel-hadadi/gotisan/constant"
	"os"
)

var (
	templates map[string]string = map[string]string{
		"handler": constant.HandlerSample,
		"model":   constant.ModelSample,
		"dto":     constant.DTOSample,
	}

	frameworkHandlers map[string]string = map[string]string{
		"gin":   constant.GinHandlerSample,
		"echo":  constant.EchoHandlerSample,
		"fiber": constant.FiberHandlerSample,
	}
)

func copyTemplateFiles(cfg *config.Config) error {
	if cfg.Framework != "" {
		templates["handler"] = frameworkHandlers[cfg.Framework]
	}

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

	_, err := os.Stat(".gitignore")
	if !errors.Is(err, os.ErrNotExist) {
		file, err := os.OpenFile(".gitignore", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		defer file.Close()

		_, err = file.WriteString("\n.gotisan/\n")
		if err != nil {
			return err
		}
	}

	return nil
}
