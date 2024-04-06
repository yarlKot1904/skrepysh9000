package server

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"skrepysh-agent/pkg"

	"go.uber.org/zap"
)

type configureRequest struct {
	WriteFiles []writeFile `json:"write-files"`
	AddUsers   []addUser   `json:"add-users"`
}

func (c *configureRequest) exec(log *zap.Logger) error {
	var errs error
	for _, wf := range c.WriteFiles {
		log.Info(fmt.Sprintf("writing file %s", wf.Filepath))
		if err := wf.exec(); err != nil {
			log.Error(fmt.Sprintf("could write file %s: %s", wf.Filepath, err.Error()))
			errs = errors.Join(errs, err)
		}
	}
	for _, au := range c.AddUsers {
		log.Info(fmt.Sprintf("creating user %s", au.Username))
		if err := au.exec(); err != nil {
			log.Error(fmt.Sprintf("could create user %s: %s", au.Username, err.Error()))
			errs = errors.Join(errs, err)
		}
	}

	return errs
}

type writeFile struct {
	Filepath string `json:"filepath"`
	Content  string `json:"content"`
}

func (wf *writeFile) exec() error {
	dir := filepath.Dir(wf.Filepath)
	if err := os.MkdirAll(dir, 0644); err != nil {
		return err
	}
	return os.WriteFile(wf.Filepath, []byte(wf.Content), 0644)
}

type addUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (au *addUser) exec() error {
	cmd := "adduser"
	args := []string{"--disabled-password", "--gecos", "''", au.Username}
	if _, _, err := pkg.RunCmd(cmd, args...); err != nil {
		return err
	}
	cmd = "echo"
	args = []string{fmt.Sprintf("%s:%s", au.Username, au.Password), "|", "chpasswd"}
	_, _, err := pkg.RunCmd(cmd, args...)
	return err
}

type errorResponse struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}
