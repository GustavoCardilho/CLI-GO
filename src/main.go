package main

import (
	"bufio"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "gustavo-cli-teste"}

	var nome, email, senha string

	var cmd = &cobra.Command{
		Use:   "Create",
		Short: "Crie um novo usuario",
		Run: func(cmd *cobra.Command, args []string) {

			println("Nome: ", nome)
			println("Email: ", email)
			println("Senha: ", senha)
		},
	}

	cmd.Flags().StringVarP(&nome, "nome", "n", "", "Nome do usuario")
	cmd.Flags().StringVarP(&email, "email", "e", "", "Email do usuario")
	cmd.Flags().StringVarP(&senha, "senha", "s", "", "Senha do usuario")

	rootCmd.AddCommand(cmd)
	rootCmd.Execute()

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
}
