package deployment

import (
	"os"
	"path"
	"strings"

	"github.com/termie/go-shutil"
)

func deployDir(tomeePath, packageName string) (deployPath, pkgName string) {
	_, pkgName = path.Split(packageName)
	deployPath = path.Join(tomeePath, "webapps")
	if strings.HasSuffix(pkgName, ".ear") {
		deployPath = path.Join(tomeePath, "apps")
		os.Mkdir(deployPath, 0744)
	}
	return
}

func Undeploy(tomeePath, packageForUndeploy string) error {
	deployPath, pkgName := deployDir(tomeePath, packageForUndeploy)

	err := os.RemoveAll(path.Join(deployPath, pkgName))
	if err != nil {
		return err
	}

	err = os.RemoveAll(path.Join(deployPath, strings.TrimSuffix(pkgName, path.Ext(pkgName))))
	if err != nil {
		return err
	}

	return nil
}

func Deploy(tomeePath, packageForDeploy string) error {
	deployPath, pkgName := deployDir(tomeePath, packageForDeploy)

	err := shutil.CopyFile(packageForDeploy, path.Join(deployPath, pkgName), true)
	if err != nil {
		return err
	}

	return nil
}
