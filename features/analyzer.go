// analyzer/repo_analyzer.go
package analyzer

import (
	"fmt"
	"github.com/Dmaddu/devpilot/client"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"
)

type FileSummary struct {
	FilePath    string
	PackageName string
	Functions   []string
	Structs     []string
}

func AnalyzeRepo(rootDir string) ([]FileSummary, error) {
	var summaries []FileSummary

	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
		if err != nil {
			return nil // Skip files that can't be parsed
		}

		fileSum := FileSummary{
			FilePath:    path,
			PackageName: node.Name.Name,
		}

		for _, decl := range node.Decls {
			switch d := decl.(type) {
			case *ast.FuncDecl:
				fileSum.Functions = append(fileSum.Functions, d.Name.Name)
			case *ast.GenDecl:
				for _, spec := range d.Specs {
					if ts, ok := spec.(*ast.TypeSpec); ok {
						if _, ok := ts.Type.(*ast.StructType); ok {
							fileSum.Structs = append(fileSum.Structs, ts.Name.Name)
						}
					}
				}
			}
		}

		summaries = append(summaries, fileSum)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return summaries, nil
}

func FormatForGPT(summaries []FileSummary) string {
	builder := strings.Builder{}
	builder.WriteString("Here is a summary of a Go repository's structure:\n\n")

	for _, file := range summaries {
		builder.WriteString(fmt.Sprintf("File: %s\n", file.FilePath))
		builder.WriteString(fmt.Sprintf("  Package: %s\n", file.PackageName))
		if len(file.Structs) > 0 {
			builder.WriteString("  Structs: " + strings.Join(file.Structs, ", ") + "\n")
		}
		if len(file.Functions) > 0 {
			builder.WriteString("  Functions: " + strings.Join(file.Functions, ", ") + "\n")
		}
		builder.WriteString("\n")
	}

	return builder.String()
}

func GetArchitectureSummary(repoRoot string) (string, error) {
	summaries, err := AnalyzeRepo(repoRoot)
	if err != nil {
		return "", err
	}

	formatted := FormatForGPT(summaries)

	prompt := fmt.Sprintf(`
You are a senior software architect. Given this Go repository summary, explain the overall architecture:
- What are the main components/modules?
- What does each package seem to do?
- How is the code organized?

%s
`, formatted)
	client := client.NewAzureOpenAIClient()
	return client.SendPrompt(prompt)
}
