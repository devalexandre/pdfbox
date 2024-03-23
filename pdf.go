package pdfbox

import (
	"embed"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

//go:embed util/pdfbox-app-3.0.2.jar
var pdfBoxJar embed.FS

func ExtractTextFromPdf(pdfFilePath string) (string, error) {
	// Cria um arquivo temporário para a saída de texto
	tmpOutputFile, err := ioutil.TempFile("", "*.txt")
	if err != nil {
		return "", fmt.Errorf("erro ao criar arquivo temporário de texto: %w", err)
	}
	defer os.Remove(tmpOutputFile.Name()) // Limpa o arquivo temporário ao final

	// Extrai o arquivo JAR incorporado para o sistema de arquivos
	jarData, err := pdfBoxJar.ReadFile("util/pdfbox-app-3.0.2.jar")
	if err != nil {
		return "", fmt.Errorf("erro ao ler o arquivo JAR incorporado: %w", err)
	}

	// Cria um arquivo temporário para o JAR
	tmpJarFile, err := ioutil.TempFile("", "pdfbox-*.jar")
	if err != nil {
		return "", fmt.Errorf("erro ao criar arquivo temporário para JAR: %w", err)
	}
	defer os.Remove(tmpJarFile.Name()) // Limpa o arquivo temporário do JAR ao final

	// Escreve o conteúdo do JAR incorporado no arquivo temporário
	if _, err := tmpJarFile.Write(jarData); err != nil {
		return "", fmt.Errorf("erro ao escrever no arquivo temporário do JAR: %w", err)
	}
	if err := tmpJarFile.Close(); err != nil {
		return "", fmt.Errorf("erro ao fechar o arquivo temporário do JAR: %w", err)
	}

	// Usa o arquivo JAR em um comando Java para extrair o texto
	cmd := exec.Command("java", "-jar", tmpJarFile.Name(), "export:text", "-i", pdfFilePath, "-o", tmpOutputFile.Name())

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("falha na execução do comando: %v", err)
	}

	// Lê o conteúdo do arquivo de texto temporário
	textContent, err := ioutil.ReadFile(tmpOutputFile.Name())
	if err != nil {
		return "", fmt.Errorf("erro ao ler o arquivo de texto gerado: %w", err)
	}

	return string(textContent), nil
}
