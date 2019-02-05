package main

import (
	"bufio"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const aMain = `package main

func main() {

}
`

// Problem ...
type Problem struct {
	ID      string
	Link    string
	Title   string
	Problem string
	Summary string
}

// Fetch ...
func (p *Problem) Fetch() error {
	p.Link = "https://projecteuler.net/problem=" + p.ID
	client := &http.Client{}
	req, err := http.NewRequest("GET", p.Link, nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}
	p.Title = doc.Find("#content > h2").Text()
	if p.Title == "" {
		return errors.New("#404")
	}
	p.Problem = doc.Find("#problem_info > h3").Children().Remove().End().Text()
	p.Summary = doc.Find("#content > div.problem_content").Text()
	p.Summary = strings.Trim(p.Summary, " \n")
	return nil
}

// Markdown ...
func (p *Problem) Markdown() string {
	md := ""
	md += "# " + p.Title
	md += "\n\n"
	md += p.Problem + "  "
	md += "\n\n"
	for _, line := range strings.Split(p.Summary, "\n") {
		md += line + "  \n"
	}
	md += "\n\n"
	md += p.Link + "\n"
	return md
}

func (p *Problem) String() string {
	str := ""
	str += "ID: " + p.ID
	str += "\n"
	str += "Link: " + p.Link
	str += "\n"
	str += "Title: " + p.Title
	str += "\n"
	str += "Problem: " + p.Problem
	str += "\n"
	str += "Summary: " + p.Summary
	return str
}

// Solution ...
func (p *Problem) Solution() string {
	file, err := os.Open("solutions/solutions.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	toFind := p.ID + ". "
	sol := ""
	for scanner.Scan() {
		line := scanner.Text()
		if len(toFind) > len(line) {
			continue
		}
		if line[:len(toFind)] == toFind {
			sol = line[len(toFind):]
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sol
}

func createProject(main, dir string) {
	if _, err := os.Stat(main); os.IsNotExist(err) {
		p := Problem{ID: os.Args[2]}
		if err := p.Fetch(); err != nil {
			panic(err)
		}
		os.MkdirAll(dir, os.ModePerm)

		if _, err := os.Stat(dir + "README.md"); os.IsNotExist(err) {
			err := ioutil.WriteFile(dir+"README.md", []byte(p.Markdown()), os.ModePerm)
			if err != nil {
				panic(err)
			}
		}
		err := ioutil.WriteFile(dir+"main.go", []byte("// "+p.Link+"\n"+aMain+"// Solution is "+p.Solution()+"\n"), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func executeProject(main string) {
	cmd := exec.Command("go", "run", main)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println("Run failed", err)
	}
}

func openInVSCode(main string) {
	exec.Command("code", main).Run()
}

func main() {
	nArgs := len(os.Args)
	if nArgs != 3 {
		return
	}
	action := os.Args[1]
	dir := "problems/" + os.Args[2] + "/"
	main := dir + "main.go"
	switch action {
	case "init", "install", "i", "g", "get":
		createProject(main, dir)
	case "run", "r":
		executeProject(main)
	case "d", "do": // mother of lazy
		createProject(main, dir)
		openInVSCode(main)
		executeProject(main)
	}
}
