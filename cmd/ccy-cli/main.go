package main

import (
	"os"
	"path"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	useLog bool
)

func main() {
	// Try to load environment variables
	godotenv.Load()
	godotenv.Load(path.Join(os.Getenv("HOME"), ".ccyrc"))

	root := &cobra.Command{
		Use:   "experimental",
		Short: "Experimental commands",
	}
	maxCmd := &cobra.Command{
		Use:   "max",
		Short: "Max commands",
		Run:   Max,
	}
	balanceCmd := &cobra.Command{
		Use:   "balance",
		Short: "Get balance",
		RunE:  Balance,
	}
	root.PersistentFlags().BoolVarP(&useLog, "log", "l", false, "Use log")
	root.AddCommand(maxCmd)
	root.AddCommand(balanceCmd)
	root.Execute()
}
