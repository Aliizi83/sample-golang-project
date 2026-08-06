package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	"github.com/jinzhu/inflection"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CodeType string

const (
	dto     CodeType = "%s.go"
	handler CodeType = "%s.go"
	model   CodeType = "%s.go"
	router  CodeType = "%s.go"
	service CodeType = "%s_service.go"
)

type generatorData struct {
	Entity            string
	CamelEntity       string
	SnakeEntity       string
	EntityPlural      string
	CamelEntityPlural string
}

func main() {
	entityFlag := flag.String("entity", "", "Entity name")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	entityName := resolveField(reader, "EntityName", *entityFlag, "UserProfile")

	data := generatorData{
		Entity:            toPascalCase(entityName),
		CamelEntity:       toCamelCase(entityName),
		SnakeEntity:       toSnakeCase(entityName),
		EntityPlural:      toPlural(entityName),
		CamelEntityPlural: toCamelCase(toPlural(entityName)),
	}

	// Domain Layer
	if err := buildGoFilesFromTemplate("./templates/model.tmpl", "../domain/models", data.SnakeEntity, model, data); err != nil {
		fmt.Println(err)
	}
	if err := appendCodeToGoFileFromTemplate("./templates/repository.tmpl", "../domain/repositories/base_repository.go", data); err != nil {
		fmt.Println(err)
	}

	// Dependency Injection
	if err := appendCodeToGoFileFromTemplate("./templates/dependency.tmpl", "../dependencies/dependency.go", data); err != nil {
		fmt.Println(err)
	}

	// Service Layer
	if err := buildGoFilesFromTemplate("./templates/service_dto.tmpl", "../services/dto", data.SnakeEntity, dto, data); err != nil {
		fmt.Println(err)
	}
	if err := buildGoFilesFromTemplate("./templates/service.tmpl", "../services", data.SnakeEntity, service, data); err != nil {
		fmt.Println(err)
	}

	// Api Layer
	if err := buildGoFilesFromTemplate("./templates/api_dto.tmpl", "../api/dto", data.SnakeEntity, dto, data); err != nil {
		fmt.Println(err)
	}
	if err := buildGoFilesFromTemplate("./templates/handler.tmpl", "../api/handlers", data.SnakeEntity, handler, data); err != nil {
		fmt.Println(err)
	}
	if err := buildGoFilesFromTemplate("./templates/router.tmpl", "../api/routers", data.SnakeEntity, router, data); err != nil {
		fmt.Println(err)
	}
}

func resolveField(reader *bufio.Reader, fieldName, flagValue, example string) string {
	if flagValue != "" {
		err := validateInput(flagValue)
		if err == nil {
			fmt.Printf("✔ Using flag value for %s: %s\n", fieldName, flagValue)
			return flagValue
		}
		fmt.Printf("Invalid flag passed for %s (%s). Falling back to prompt...\n", fieldName, err)
	}

	for {
		fmt.Printf("Enter %s (e.g., %s): ", fieldName, example)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if err := validateInput(input); err != nil {
			fmt.Printf("Error: %s. Please try again.\n\n", err)
			continue
		}
		return input
	}
}

func validateInput(input string) error {
	if input == "" {
		return errors.New("input cannot be empty")
	}

	matched, _ := regexp.MatchString("^[a-zA-Z_]+$", input)
	if !matched {
		return errors.New("must contain only English letters and underscores")
	}

	return nil
}

func toCamelCase(input string) string {
	words := splitWords(input)
	if len(words) == 0 {
		return ""
	}

	parts := make([]string, 0, len(words))
	for i, word := range words {
		lowerWord := strings.ToLower(word)
		if i == 0 {
			parts = append(parts, lowerWord)
			continue
		}
		parts = append(parts, cases.Title(language.Und).String(lowerWord))
	}

	return strings.Join(parts, "")
}

func toSnakeCase(input string) string {
	words := splitWords(input)
	if len(words) == 0 {
		return ""
	}

	parts := make([]string, 0, len(words))
	for _, word := range words {
		parts = append(parts, strings.ToLower(word))
	}

	return strings.Join(parts, "_")
}

func toPlural(input string) string {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return ""
	}

	words := splitWords(trimmed)
	if len(words) == 0 {
		return ""
	}

	lastWord := strings.ToLower(words[len(words)-1])
	pluralLastWord := inflection.Plural(lastWord)

	if len(words) == 1 {
		return toPascalCase(pluralLastWord)
	}

	prefixWords := words[:len(words)-1]
	parts := make([]string, 0, len(prefixWords)+1)
	for _, word := range prefixWords {
		parts = append(parts, word)
	}
	parts = append(parts, pluralLastWord)

	return toPascalCase(strings.Join(parts, "_"))
}

func toPascalCase(input string) string {
	words := splitWords(input)
	if len(words) == 0 {
		return ""
	}

	parts := make([]string, 0, len(words))
	for _, word := range words {
		parts = append(parts, cases.Title(language.Und).String(strings.ToLower(word)))
	}

	return strings.Join(parts, "")
}

func splitWords(input string) []string {
	cleaned := strings.ReplaceAll(strings.ReplaceAll(input, "_", " "), "-", " ")
	parts := regexp.MustCompile(`[A-Z]+(?=[A-Z][a-z]|[0-9]|$)|[A-Z]?[a-z]+|[A-Z]+|[0-9]+`).FindAllString(cleaned, -1)
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func isVowel(r rune) bool {
	switch unicode.ToLower(r) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}

func buildGoFilesFromTemplate(templatePath, dstFolder, fileName string, codeType CodeType, data generatorData) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(io.Discard, data); err != nil {
		return err
	}

	formattedFileName := fmt.Sprintf(string(codeType), fileName)
	fullPathFile := fmt.Sprintf("%s/%s", dstFolder, formattedFileName)

	outputFile, err := os.Create(fullPathFile)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, data)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully loaded %s file of type ( %s )\n", fullPathFile, codeType)
	return nil
}

func appendCodeToGoFileFromTemplate(templatePath, dstFile string, data generatorData) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(io.Discard, data); err != nil {
		return err
	}

	file, err := os.OpenFile(dstFile, os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		return err
	}

	return nil
}
