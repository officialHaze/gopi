package main

import (
	"log"
	"os"
	"projinit/initializer"
	"projinit/path"
	"projinit/settings"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	// Generate my settings
	if err := settings.Generate(); err != nil {
		log.Fatalln(err)
	}

	// load the env file
	p := path.Join(settings.MySettings.ENV_FILE_NAME)
	if err := godotenv.Load(p); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	log.Printf(" ****** Project Initializer Running In %s Environment ******", strings.ToUpper(os.Getenv("ENV")))

	projectpath, projectname, author, goversion := getArgs()

	i := initializer.New(projectpath, projectname, author, goversion)

	// Initialize
	if err := i.Init(); err != nil {
		log.Fatalln(err)
	}
}

func getArgs() (string, string, string, string) {
	args := os.Args

	execname := normalizeArg(args[0])
	log.Printf("Executable Name: %s", execname)

	if len(args) <= 1 {
		log.Fatalln("No args provided")
	}

	projectpath := normalizeArg(args[1])
	projectname := normalizeArg(args[2])
	author := normalizeArg(args[3])
	goversion := normalizeArg(args[4])

	return projectpath, projectname, author, goversion
}

func normalizeArg(arg string) string {
	return strings.TrimSpace(strings.ToLower(arg))
}
