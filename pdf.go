package pdfbox

import (
	"embed"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

//go:embed util/pdfbox-app-3.0.2.jar
var pdfBoxJar embed.FS

func ExtractTextFromPdf(pdfFilePath string) (*string, error) {
	baseFileName := strings.TrimSuffix(pdfFilePath, ".pdf")
	txtOutputPath := baseFileName + ".txt"

	// Extrai o arquivo JAR incorporado para o sistema de arquivos
	jarData, err := pdfBoxJar.ReadFile("util/pdfbox-app-3.0.2.jar")
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo JAR incorporado: %w", err)
	}

	// Cria um arquivo temporário para o JAR
	tmpJarFile, err := ioutil.TempFile("", "pdfbox-*.jar")
	if err != nil {
		return nil, fmt.Errorf("erro ao criar arquivo temporário: %w", err)
	}
	defer os.Remove(tmpJarFile.Name()) // Limpa o arquivo temporário ao final

	// Escreve o conteúdo do JAR incorporado no arquivo temporário
	if _, err := tmpJarFile.Write(jarData); err != nil {
		return nil, fmt.Errorf("erro ao escrever no arquivo temporário: %w", err)
	}
	if err := tmpJarFile.Close(); err != nil {
		return nil, fmt.Errorf("erro ao fechar o arquivo temporário: %w", err)
	}

	// Usa o arquivo JAR em um comando Java
	cmd := exec.Command("java", "-jar", tmpJarFile.Name(), "export:text", "-i", pdfFilePath)

	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("falha na execução do comando: %v", err)
	}

	if _, err := os.Stat(txtOutputPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("o arquivo de texto gerado não foi encontrado: %s", txtOutputPath)
	}

	textContent, err := ioutil.ReadFile(txtOutputPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo de texto gerado: %w", err)
	}

	// remove o arquivo de texto gerado
	if err := os.Remove(txtOutputPath); err != nil {
		return nil, fmt.Errorf("erro ao remover o arquivo de texto gerado: %w", err)
	}

	textAsString := string(textContent)
	return &textAsString, nil
}
