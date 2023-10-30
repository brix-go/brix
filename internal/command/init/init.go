package new

import (
	"bytes"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/brix-go/brix/config"
	"github.com/brix-go/brix/pkg/helper"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var text = "A tool for generating golang project"
var coloredText = "\x1B[38;2;66;211;146m" + text + "\x1B[39m"

type Bar struct {
	percent int64  // progress percentage
	cur     int64  // current progress
	total   int64  // total value for progress
	rate    string // the actual progress bar to be printed
	graph   string // the fill value for progress bar
}

type Project struct {
	ProjectName string `survey:"name"`
}

var CmdNew = &cobra.Command{
	Use:     "new",
	Example: "brix new demo-api",
	Short:   "create a new project.",
	Long:    `create a new project with brix layout.`,
	Run:     run,
}
var (
	repoURL string
)

func init() {
	CmdNew.Flags().StringVarP(&repoURL, "repo-url", "r", repoURL, "layout repo")

}
func NewProject() *Project {
	return &Project{}
}

func run(cmd *cobra.Command, args []string) {
	p := NewProject()
	if len(args) == 0 {
		err := survey.AskOne(&survey.Input{
			Message: "What is your project name?",
			Help:    "project name.",
			Suggest: nil,
		}, &p.ProjectName, survey.WithValidator(survey.Required))
		if err != nil {
			return
		}
	} else {
		p.ProjectName = args[0]
	}

	// clone repo

	yes, err := p.cloneTemplate()
	if err != nil || !yes {
		return
	}

	err = p.replacePackageName()
	if err != nil || !yes {
		return
	}

	err = p.replacePackageName()
	if err != nil || !yes {
		return
	}
	err = p.modTidy()
	if err != nil || !yes {
		return
	}
	p.rmGit()
	//p.installWire()
	fmt.Sprintf("\n                               \n,-----.         ,--.           \n|  |) /_ ,--.--.`--',--.  ,--. \n|  .-.  \\|  .--',--. \\  `'  /  \n|  '--' /|  |   |  | /  /.  \\  \n`------' `--'   `--''--'  '--' \n                               \n%s", coloredText)
	fmt.Printf("🎉 Project \u001B[36m%s\u001B[0m created successfully!\n\n", p.ProjectName)
	fmt.Printf("Done. Now run:\n\n")
	fmt.Printf("› \033[36mcd %s \033[0m\n", p.ProjectName)
	fmt.Printf("› \033[36mbrix run \033[0m\n\n")
}

func (p *Project) cloneTemplate() (bool, error) {

	stat, _ := os.Stat(p.ProjectName)
	if stat != nil {
		var overwrite = false

		prompt := &survey.Confirm{
			Message: fmt.Sprintf("Folder %s already exists, do you want to overwrite it?", p.ProjectName),
			Help:    "Remove old project and create new project.",
		}
		err := survey.AskOne(prompt, &overwrite)
		if err != nil {
			return false, err
		}
		if !overwrite {
			return false, nil
		}
		err = os.RemoveAll(p.ProjectName)
		if err != nil {
			fmt.Println("remove old project error: ", err)
			return false, err
		}
	}
	repo := config.FiberRepo

	if repoURL == "" {
		layout := ""
		prompt := &survey.Select{
			Message: "Please select a framework:",
			Options: []string{
				"Fiber",
				"Gin Gonic",
				"Mux",
			},
			Description: func(value string, index int) string {
				if index == 1 {
					return " - A project structure with gin framework"
				}
				if index == 2 {
					return "- A project structure with mux framework"
				}
				return " - It has rich functions such as db, jwt, redis, migration, test, etc"
			},
		}
		err := survey.AskOne(prompt, &layout)
		if err != nil {
			return false, err
		}
		if layout == "Fiber" {
			repo = config.FiberRepo
		}
		if layout == "Gin Gonic" {
			fmt.Println("Sorry boilerplate for this framework not yet available: ")
			return false, err
		}
		if layout == "Mux" {
			fmt.Println("Sorry boilerplate for this framework not yet available: ")
			return false, err
		}
		err = os.RemoveAll(p.ProjectName)
		if err != nil {
			fmt.Println("remove old project error: ", err)
			return false, err
		}
	} else {
		repo = repoURL
	}

	cmd := exec.Command("git", "clone", repo, p.ProjectName)

	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("git clone %s error: %s\n", repo, err)
		return false, err
	}
	fmt.Println("Generating project ", p.ProjectName)
	var bar Bar
	bar.NewOption(0, 100)
	for i := 0; i <= 100; i++ {
		time.Sleep(500 * time.Millisecond)
		bar.Play(int64(i))
	}
	bar.Finish()
	return true, nil
}

func (p *Project) replacePackageName() error {
	packageName := helper.GetProjectName(p.ProjectName)

	err := p.replaceFiles(packageName)
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "edit", "-module", p.ProjectName)
	cmd.Dir = p.ProjectName
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println("go mod edit error: ", err)
		return err
	}
	return nil
}
func (p *Project) modTidy() error {
	fmt.Println("go mod tidy")
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = p.ProjectName
	if err := cmd.Run(); err != nil {
		fmt.Println("go mod tidy error: ", err)
		return err
	}
	return nil
}
func (p *Project) rmGit() {
	os.RemoveAll(p.ProjectName + "/.git")
}
func (p *Project) installWire() {
	fmt.Printf("go install %s\n", config.WireCmd)
	cmd := exec.Command("go", "install", config.WireCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("go install %s error\n", err)
	}
}

func (p *Project) replaceFiles(packageName string) error {
	err := filepath.Walk(p.ProjectName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		newData := bytes.ReplaceAll(data, []byte(packageName), []byte(p.ProjectName))
		if err := os.WriteFile(path, newData, 0644); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Println("walk file error: ", err)
		return err
	}
	return nil
}

func (bar *Bar) NewOption(start, total int64) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		bar.graph = "█"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		bar.rate += bar.graph // initial progress position
	}
}

func (bar *Bar) getPercent() int64 {
	return int64((float32(bar.cur) / float32(bar.total)) * 100)
}

func (bar *Bar) Play(cur int64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", bar.rate, bar.percent, bar.cur, bar.total)
}

func (bar *Bar) Finish() {
	fmt.Println()
}
