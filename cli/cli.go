package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/flavio/flang/runtime"
)

const version = "0.2.0"

const banner = `
  ███████╗██╗      █████╗ ███╗   ██╗ ██████╗
  ██╔════╝██║     ██╔══██╗████╗  ██║██╔════╝
  █████╗  ██║     ███████║██╔██╗ ██║██║  ███╗
  ██╔══╝  ██║     ██╔══██╗██║╚██╗██║██║   ██║
  ██║     ███████╗██║  ██║██║ ╚████║╚██████╔╝
  ╚═╝     ╚══════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝
  v%s - Tudo roda direto do .fg
`

// Run executes the CLI.
func Run(args []string) {
	if len(args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch args[1] {
	case "run":
		if len(args) < 3 {
			fmt.Println("Uso: flang run <arquivo.fg>")
			os.Exit(1)
		}
		porta := "8080"
		if len(args) >= 4 {
			porta = args[3]
		}
		if err := runtime.Executar(args[2], porta); err != nil {
			fmt.Printf("[flang] ERRO: %s\n", err)
			os.Exit(1)
		}

	case "check":
		if len(args) < 3 {
			fmt.Println("Uso: flang check <arquivo.fg>")
			os.Exit(1)
		}
		if err := runtime.Verificar(args[2]); err != nil {
			fmt.Printf("[flang] ERRO: %s\n", err)
			os.Exit(1)
		}

	case "new":
		if len(args) < 3 {
			fmt.Println("Uso: flang new <nome>")
			os.Exit(1)
		}
		cmdNew(args[2])

	case "version":
		fmt.Printf(banner, version)

	case "help":
		printUsage()

	default:
		// If arg ends in .fg, treat it as "run"
		if strings.HasSuffix(args[1], ".fg") {
			porta := "8080"
			if len(args) >= 3 {
				porta = args[2]
			}
			if err := runtime.Executar(args[1], porta); err != nil {
				fmt.Printf("[flang] ERRO: %s\n", err)
				os.Exit(1)
			}
		} else {
			fmt.Printf("Comando desconhecido: %s\n", args[1])
			printUsage()
			os.Exit(1)
		}
	}
}

func printUsage() {
	fmt.Printf(banner, version)
	fmt.Println(`
Uso: flang <comando> [argumentos]

Comandos:
  run <arquivo.fg> [porta]  Executa o arquivo .fg (porta padrão: 8080)
  check <arquivo.fg>        Verifica sintaxe sem executar
  new <nome>                Cria um novo projeto .fg
  version                   Mostra a versão
  help                      Mostra esta ajuda

Atalho:
  flang inicio.fg           Mesmo que "flang run inicio.fg"

Exemplo:
  flang run inicio.fg
  flang inicio.fg 3000
`)
}

func cmdNew(name string) {
	dir := name
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("Erro: %s\n", err)
		os.Exit(1)
	}

	fg := `sistema ` + name + `

dados

  item
    nome: texto
    descricao: texto
    preco: numero

telas

  tela itens

    titulo "` + strings.ToUpper(name[:1]) + name[1:] + `"

    lista itens

      mostrar nome
      mostrar descricao
      mostrar preco

    botao azul
      texto "Novo"

eventos

  quando clicar "Novo"
    criar item
`

	fgPath := filepath.Join(dir, "inicio.fg")
	if err := os.WriteFile(fgPath, []byte(fg), 0644); err != nil {
		fmt.Printf("Erro: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("[flang] Projeto '%s' criado!\n", name)
	fmt.Printf("[flang] Execute: flang %s\n", fgPath)
}
