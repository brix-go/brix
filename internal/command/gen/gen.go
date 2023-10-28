package gen

import (
	"fmt"
	tpl "github.com/brix-go/brix/internal/template"
	"github.com/brix-go/brix/pkg/helper"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

type Create struct {
	ProjectName        string
	CreateType         string
	FilePath           string
	FileName           string
	FileNameTitleLower string
	FileNameFirstChar  string
	IsFull             bool
}

func NewCreate() *Create {
	return &Create{}
}

var CmdCreate = &cobra.Command{
	Use:     "gen [type] [controller-name]",
	Short:   "gen a new controller/service/repository/model",
	Example: "brix gen controller user",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var (
	tplPath string
)

func init() {
	CmdCreatecontroller.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdCreateService.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdCreateRepository.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdCreateDomain.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdCreateAll.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")

}

var CmdCreatecontroller = &cobra.Command{
	Use:     "controller",
	Short:   "gen a new controller",
	Example: "brix gen controller user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CmdCreateService = &cobra.Command{
	Use:     "service",
	Short:   "gen a new service",
	Example: "brix gen service user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CmdCreateRepository = &cobra.Command{
	Use:     "repository",
	Short:   "gen a new repository",
	Example: "brix gen repository user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CmdCreateDomain = &cobra.Command{
	Use:     "model",
	Short:   "gen a new domain",
	Example: "brix gen model user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CmdCreateAll = &cobra.Command{
	Use:     "all",
	Short:   "gen a new controller & service & repository & domain",
	Example: "brix gen all user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}

func runCreate(cmd *cobra.Command, args []string) {
	c := NewCreate()
	c.ProjectName = helper.GetProjectName(".")
	c.CreateType = cmd.Use
	c.FilePath, c.FileName = filepath.Split(args[0])
	c.FileName = strings.ReplaceAll(strings.ToUpper(string(c.FileName[0]))+c.FileName[1:], ".go", "")
	c.FileNameTitleLower = strings.ToLower(string(c.FileName[0])) + c.FileName[1:]
	c.FileNameFirstChar = string(c.FileNameTitleLower[0])

	switch c.CreateType {
	case "controller", "service", "repository", "domain":
		c.genFile()
	case "all":
		c.CreateType = "controller"
		c.genFile()

		c.CreateType = "service"
		c.genFile()

		c.CreateType = "repository"
		c.genFile()

		c.CreateType = "domain"
		c.genFile()
	default:
		log.Fatalf("Invalid generate type: %s", c.CreateType)
	}

}

func (c *Create) genFile() {
	var f *os.File
	filePath := c.FilePath
	if filePath == "" {
		filePath = fmt.Sprintf("internal/domain/%s/", c.FileName)
	}
	if c.CreateType != "domain" {
		f = createFile(fmt.Sprintf("%s/%s", filePath, c.CreateType), strings.ToLower(c.FileName)+".go")
	} else {
		f = createFile(filePath, strings.ToLower(c.FileName)+".go")

	}
	if f == nil {
		log.Printf("warn: file %s%s %s", filePath, strings.ToLower(c.FileName)+".go", "already exists.")
		return
	}
	defer f.Close()
	var t *template.Template
	var err error
	if tplPath == "" {
		t, err = template.ParseFS(tpl.CreateTemplateFS, fmt.Sprintf("create/%s.tpl", c.CreateType))
	} else {
		t, err = template.ParseFiles(path.Join(tplPath, fmt.Sprintf("%s.tpl", c.CreateType)))
	}
	if err != nil {
		log.Fatalf("generate %s error: %s", c.CreateType, err.Error())
	}
	err = t.Execute(f, c)
	if err != nil {
		log.Fatalf("generate %s error: %s", c.CreateType, err.Error())
	}
	log.Printf("Created new %s: %s", c.CreateType, filePath+strings.ToLower(c.FileName)+".go")

}
func createFile(dirPath string, filename string) *os.File {
	filePath := filepath.Join(dirPath, filename)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to generate dir %s: %v", dirPath, err)
	}
	stat, _ := os.Stat(filePath)
	if stat != nil {
		return nil
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to generate file %s: %v", filePath, err)
	}

	return file
}
