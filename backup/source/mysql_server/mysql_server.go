package mysql_server

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type MysqlServerSource struct {
	Host               string   `yaml:"Host"`
	Port               string   `yaml:"Port"`
	User               string   `yaml:"User"`
	Password           string   `yaml:"Password"`
	MysqldumpExtraArgs []string `yaml:"MysqldumpExtraArgs"`
}

func (s MysqlServerSource) Validate() error {
	if s.Host == "" {
		return fmt.Errorf("MysqlServerSource need to have a Host")
	}
	if s.Port == "" {
		return fmt.Errorf("MysqlServerSource need to have a Port")
	}
	if s.User == "" {
		return fmt.Errorf("MysqlServerSource need to have a User")
	}
	if s.Password == "" {
		return fmt.Errorf("MysqlServerSource need to have a Password")
	}
	return nil
}

func (s MysqlServerSource) Backup() (io.ReadSeeker, error) {
	var err error

	outputFile, err := os.CreateTemp("", "tergum-dump-mysql-")
	if err != nil {
		return nil, err
	}
	defer os.Remove(outputFile.Name())

	args := []string{
		"-h", s.Host,
		"-P", s.Port,
		"-u", s.User,
		"-p" + s.Password,
		"--all-databases",
	}
	cmd := exec.Command(
		"mysqldump",
		append(s.MysqldumpExtraArgs, args...)...,
	)
	cmd.Stdout = outputFile

	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	outputFile.Seek(0, 0)
	return outputFile, nil
}
