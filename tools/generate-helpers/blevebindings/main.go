package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode"

	. "github.com/dave/jennifer/jen"
	"github.com/spf13/cobra"
	v1 "github.com/stackrox/rox/generated/api/v1"
	"github.com/stackrox/rox/pkg/utils"
	"github.com/stackrox/rox/tools/generate-helpers/blevebindings/operations"
	"github.com/stackrox/rox/tools/generate-helpers/common/packagenames"
)

const (
	headerComment = "Code generated by blevebindings generator. DO NOT EDIT."
)

func newFile() *File {
	f := NewFile("index")
	f.HeaderComment(headerComment)
	return f
}

func makeTag(object string) string {
	var split []string
	i := 0
	for s := object; s != ""; s = s[i:] {
		i = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if i <= 0 {
			i = len(s)
		}
		split = append(split, strings.ToLower(s[:i]))
	}
	return strings.Join(split, "_")
}

func generateOptionsFile(props operations.GeneratorProperties) error {
	if !props.WriteOptions {
		return nil
	}
	tagString := makeTag(props.Object)
	f := NewFile("mappings")
	f.HeaderComment(headerComment)
	f.Line()
	f.Var().Id("OptionsMap").Op("=").Qual(packagenames.RoxSearch, "Walk").Call(
		Qual(packagenames.V1, props.SearchCategory),
		Lit(tagString),
		Parens(Op("*").Qual(props.Pkg, props.Object)).Parens(Nil()),
	)
	goSubPackage := operations.GenerateMappingGoSubPackageWithinCentral(props)

	// Hack to figure out the real directory corresponding to the central package.
	// It's fine for this to be hacky since bleve is going away soon.
	// First, get the working directory. We know this will be inside central
	// since all go generate commands for bleve indexes as inside central.
	// Then, strip out paths from the end until we're in the central directory.
	workingDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("couldn't determine working directory: %w", err)
	}
	remainingWorkingDir := workingDir
	var centralFilePath string
	for {
		if remainingWorkingDir == string(os.PathSeparator) {
			return fmt.Errorf("couldn't find central path in working directory %q", workingDir)
		}
		lastComponent := filepath.Base(remainingWorkingDir)
		if lastComponent == "central" {
			centralFilePath = remainingWorkingDir
			break
		}
		remainingWorkingDir = filepath.Dir(remainingWorkingDir)
	}
	return f.Save(filepath.Join(centralFilePath, goSubPackage, "options.go"))
}

func generateIndexInterfaceFile(interfaceMethods []Code) error {
	f := newFile()
	f.Type().Id("Indexer").Interface(interfaceMethods...)
	return f.Save("indexer.go")
}

func generateMocks(props operations.GeneratorProperties) error {
	if !props.GenerateMockIndexer {
		return nil
	}
	if props.MockgenWrapperExecutablePath == "" {
		return errors.New("mockgen-wrapper path not specified")
	}
	cmd := exec.Command(props.MockgenWrapperExecutablePath)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "GOFILE=indexer.go")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("couldn't exec mockgen-wrapper: %w", err)
	}
	return nil
}

func generate(props operations.GeneratorProperties) error {
	interfaceMethods := operations.GenerateInterfaceAndImplementation(props)

	if err := generateIndexInterfaceFile(interfaceMethods); err != nil {
		return err
	}
	if err := generateOptionsFile(props); err != nil {
		return err
	}

	if err := generateMocks(props); err != nil {
		return err
	}
	return nil
}

func renderSearchCategories() string {
	allCategories := make([]string, 0, len(v1.SearchCategory_value))

	for category := range v1.SearchCategory_value {
		allCategories = append(allCategories, category)
	}
	return strings.Join(allCategories, ",")
}

func checkSupported(searchCategory string) error {
	if _, ok := v1.SearchCategory_value[searchCategory]; !ok {
		return fmt.Errorf("search category %s is unsupported", searchCategory)
	}
	return nil
}

func main() {
	c := &cobra.Command{
		Use: "generate store implementations",
	}

	props := operations.GeneratorProperties{}
	c.Flags().StringVar(&props.Pkg, "package", "github.com/stackrox/rox/generated/storage", "the package of the object being indexed")

	c.Flags().StringVar(&props.Object, "object", "", "the (Go) name of the object being indexed")
	utils.Must(c.MarkFlagRequired("object"))

	c.Flags().StringVar(&props.Singular, "singular", "", "the singular name of the object")
	utils.Must(c.MarkFlagRequired("singular"))

	c.Flags().StringVar(&props.Plural, "plural", "", "the plural name of the object (optional; appends 's' to singular by default")

	c.Flags().StringVar(&props.IDFunc, "id-func", "GetId", "the method to invoke on the proto object to get an id out")

	c.Flags().StringVar(&props.SearchCategory, "search-category", "", fmt.Sprintf("the search category to index under (supported - %s)", renderSearchCategories()))
	utils.Must(c.MarkFlagRequired("search-category"))

	c.Flags().BoolVar(&props.WriteOptions, "write-options", true, "enable writing out the options map")
	c.Flags().StringVar(&props.OptionsPath, "options-path", "/index/mappings", "path to write out the options to")
	c.Flags().StringVar(&props.ObjectPathName, "object-path-name", "", "overwrite the object path underneath Central")
	c.Flags().StringVar(&props.Tag, "tag", "", "use the specified json tag")

	c.Flags().BoolVar(&props.GenerateMockIndexer, "generate-mock-indexer", false, "whether to generate a mock for the indexer")
	c.Flags().StringVar(&props.MockgenWrapperExecutablePath, "mockgen-executable-path", "", "path to mockgen-wrapper executable")

	c.RunE = func(*cobra.Command, []string) error {
		if props.Plural == "" {
			props.Plural = fmt.Sprintf("%ss", props.Singular)
		}
		if err := checkSupported(props.SearchCategory); err != nil {
			return err
		}
		props.SearchCategory = fmt.Sprintf("SearchCategory_%s", props.SearchCategory)
		return generate(props)
	}

	if err := c.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
