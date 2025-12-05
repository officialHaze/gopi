package initializer

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"projinit/helper"
	"projinit/path"
	"strings"
)

func New(path, name, author, goversion string) *Initializer {
	return &Initializer{
		projectpath: path,
		projectname: name,
		authorname:  author,
		goversion:   goversion,
	}
}

type Initializer struct {
	projectpath string
	projectname string
	authorname  string
	goversion   string
	execfns     []func() error
}

func (i *Initializer) Init() error {
	i.Generate()

	for _, fn := range i.execfns {
		if err := fn(); err != nil {
			return err
		}
	}

	return nil
}

func (i *Initializer) Generate() {
	i.execfns = []func() error{
		i.NormalizeAndSetProjectPath, // primary first

		i.CreateAuthorMD,
		i.CreateReadmeMD,
		i.CreateModFile,
		i.CreateMainPkg,
		i.CreateBinDir,
		i.CreateHelperPkg,
		i.CreateModelPkg,
		i.CreatePublicDir,
		i.CreateServerPkg,
		i.CreateRoutePkg,
		i.CopySettingsPkg,
		i.CopyUtilPkg,
		i.CopyEnvLocal,
		i.CopyGitignore,

		i.DownloadModules, // primary last
	}
}

func (i *Initializer) CreateModFile() error {
	templatepath := path.Join("templates", "mod.txt")

	b, err := helper.ReplaceProjectNamePlaceholder(templatepath, i.projectname)
	if err != nil {
		return fmt.Errorf("error replacing project name placeholder: %v", err)
	}

	b = helper.ReplaceGOVersionPlaceholder(b, i.goversion)

	// Write the mod file in project directory
	return os.WriteFile(filepath.Join(i.projectpath, "go.mod"), b, os.ModePerm)
}

func (i *Initializer) CreateMainPkg() error {
	templatepath := path.Join("templates", "main.txt")

	b, err := helper.ReplaceProjectNamePlaceholder(templatepath, i.projectname)
	if err != nil {
		return fmt.Errorf("error replacing project name placeholder: %v", err)
	}

	// Write the main file in project directory
	return os.WriteFile(filepath.Join(i.projectpath, "main.go"), b, os.ModePerm)
}

func (i *Initializer) CreateAuthorMD() error {
	content := fmt.Appendf(nil, "## %s", i.authorname)
	return os.WriteFile(fmt.Sprintf("%s/AUTHOR.md", i.projectpath), content, os.ModePerm)
}

func (i *Initializer) CreateReadmeMD() error {
	content := []byte(fmt.Sprintf("# %s", strings.ToUpper(i.projectname)))
	return os.WriteFile(fmt.Sprintf("%s/README.md", i.projectpath), content, os.ModePerm)
}

func (i *Initializer) CreateBinDir() error {
	return os.MkdirAll(fmt.Sprintf("%s/bin", i.projectpath), os.ModePerm)
}

func (i *Initializer) CreateHelperPkg() error {
	return os.MkdirAll(fmt.Sprintf("%s/helper", i.projectpath), os.ModePerm)
}

func (i *Initializer) CreateModelPkg() error {
	return os.MkdirAll(fmt.Sprintf("%s/model", i.projectpath), os.ModePerm)
}

func (i *Initializer) CreatePublicDir() error {
	return os.MkdirAll(fmt.Sprintf("%s/public", i.projectpath), os.ModePerm)
}

func (i *Initializer) CreateServerPkg() error {
	templatepath := path.Join("templates", "server.txt")

	b, err := helper.ReplaceProjectNamePlaceholder(templatepath, i.projectname)
	if err != nil {
		return fmt.Errorf("error replacing project name placeholder: %v", err)
	}

	// create the subdirs
	dirpath := fmt.Sprintf("%s/api/REST/server", i.projectpath)
	if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
		return err
	}

	// Write the server boilerplate code
	return os.WriteFile(filepath.Join(dirpath, "server.go"), b, os.ModePerm)
}

func (i *Initializer) CreateRoutePkg() error {
	templatepath := path.Join("templates", "api_route.txt")

	// create the subdirs
	dirpath := fmt.Sprintf("%s/api/REST/server/routes", i.projectpath)
	if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
		return err
	}

	// copy the route boilerplate code
	cmd := exec.Command("cp", templatepath, fmt.Sprintf("%s/api.go", dirpath))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (i *Initializer) CopySettingsPkg() error {
	templatepath := path.Join("templates", "settings")
	cmd := exec.Command("cp", "-r", templatepath, i.projectpath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (i *Initializer) CopyUtilPkg() error {
	templatepath := path.Join("templates", "util")
	cmd := exec.Command("cp", "-r", templatepath, i.projectpath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (i *Initializer) CopyEnvLocal() error {
	envfile := path.Join("templates", ".env.local")
	cmd := exec.Command("cp", envfile, i.projectpath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (i *Initializer) CopyGitignore() error {
	gitignore := path.Join("templates", ".gitignore")
	cmd := exec.Command("cp", gitignore, i.projectpath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (i *Initializer) DownloadModules() error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = i.projectpath

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error downloading modules: %v", err)
	}

	return nil
}

// Setters
func (i *Initializer) NormalizeAndSetProjectPath() error {
	// project path cannot contain both ~ and . (root) at the same time
	if strings.Contains(i.projectpath, "~") && strings.Contains(i.projectpath, ".") {
		return fmt.Errorf("invalid projectpath format: %s", i.projectpath)
	}

	// Replace relative projectpath with absolute
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting pwd at - %s: %v", i.projectpath, err)
	}
	i.projectpath = strings.ReplaceAll(i.projectpath, ".", pwd)

	user, err := user.Current()
	if err != nil {
		return fmt.Errorf("error getting username: %v", err)
	}
	i.projectpath = strings.ReplaceAll(i.projectpath, "~", fmt.Sprintf("/home/%s", user.Username))

	log.Printf("Normalized Path: %s", i.projectpath)

	// take care of directory creation at projectpath
	if err := os.MkdirAll(i.projectpath, os.ModePerm); err != nil {
		return fmt.Errorf("error creating dir at projectpath - %s: %v", i.projectpath, err)
	}

	return nil
}

// Getters
func (i *Initializer) Get_ProjectPath() string {
	return i.projectpath
}
