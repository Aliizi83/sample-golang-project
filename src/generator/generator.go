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
	EntityName           string
	CamelCasedEntityName string
	SnakeCasedEntityName string
	PluralName           string
	CamelCasedPluralName string
}

func main() {
	entityFlag := flag.String("entity", "", "Entity name")
	camelEntityFlag := flag.String("camel-entity", "", "Camel cased entity name")
	snakeCaseEntityFlag := flag.String("snake-entity", "", "Snake cased entity name")
	pluralFlag := flag.String("plural", "", "Plural name")
	camelPluralFlag := flag.String("camel-plural", "", "Camel cased plural name")

	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	data := generatorData{
		EntityName:           resolveField(reader, "EntityName", *entityFlag, "UserProfile"),
		CamelCasedEntityName: resolveField(reader, "camelCasedEntityName", *camelEntityFlag, "userProfile"),
		SnakeCasedEntityName: resolveField(reader, "snake_cased_entity_name", *snakeCaseEntityFlag, "user_profile"),
		PluralName:           resolveField(reader, "PluralName", *pluralFlag, "UserProfiles"),
		CamelCasedPluralName: resolveField(reader, "camelCasedPluralName", *camelPluralFlag, "userProfiles"),
	}

	// Domain Layer
	buildGoFilesFromTemplate("./templates/model.tmpl", "../domain/models", data.SnakeCasedEntityName, model, data)
	appendCodeToGoFileFromTemplate("./templates/repository.tmpl", "../domain/repositories/repositories.go", data)

	// Dependency Injection
	appendCodeToGoFileFromTemplate("./templates/dependency.tmpl", "../dependencies/dependency.go", data)

	// Service Layer
	buildGoFilesFromTemplate("./templates/service_dto.tmpl", "../services/dto", data.SnakeCasedEntityName, dto, data)
	buildGoFilesFromTemplate("./templates/service.tmpl", "../services", data.SnakeCasedEntityName, service, data)

	// Api Layer
	buildGoFilesFromTemplate("./templates/api_dto.tmpl", "../api/dto", data.SnakeCasedEntityName, dto, data)
	buildGoFilesFromTemplate("./templates/handler.tmpl", "../api/handlers", data.SnakeCasedEntityName, handler, data)
	buildGoFilesFromTemplate("./templates/router.tmpl", "../api/routers", data.SnakeCasedEntityName, router, data)

}

func resolveField(reader *bufio.Reader, fieldName, flagValue, example string) string {
	if flagValue != "" {
		err := validateInput(flagValue)
		if err != nil {
			fmt.Printf("Using flag value for %s: %s\n", fieldName, flagValue)
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

	matched, _ := regexp.MatchString("^[a-zA-Z]+$", input)
	if !matched {
		return errors.New("must contain only English letters (no numbers, spaces, or special characters)")
	}

	return nil
}

func buildGoFilesFromTemplate(templatePath, dstFolder, fileName string, codeType CodeType, data generatorData) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(io.Discard, data); err != nil {
		return err
	}

	formattedFileName := fmt.Sprintf(string(codeType), data.SnakeCasedEntityName)
	fullPathFile := fmt.Sprintf("%s/%s.go", dstFolder, formattedFileName)

	outputFile, err := os.Create(fullPathFile)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		return err
	}

	fmt.Printf("Successfully loaded %s file of type ( %s )", fullPathFile, codeType)
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
