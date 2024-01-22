package uninstaller

import (
	"fmt"
	"go/build"
	"os"
	"strings"
)

const (
	binPath = "/bin/mosint"
)

type UninstallType struct {
	All, Binary bool
}

type Uninstaller struct {
	GoPath string
}

func Server(GoPath string) *Uninstaller {
	return &Uninstaller{
		GoPath: GoPath,
	}
}

func GetGoPath() string {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		goPath = build.Default.GOPATH
	}
	strings.TrimSuffix(goPath, "/")
	return goPath
}

func (u *Uninstaller) uninstallBinary() error {
	err := os.RemoveAll(u.GoPath + binPath + "/")
	if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
		return err
	}
	fmt.Println("Binary Removed")
	return nil
}

func (u *Uninstaller) Uninstall(arg *UninstallType) error {
	if arg.Binary {
		return u.uninstallBinary()
	}
	return u.defaultUninstall()
}

func (u *Uninstaller) defaultUninstall() error {
	err := u.uninstallBinary()
	if err != nil {
		return err
	}
	fmt.Println("Uninstall Completed")
	return nil
}
