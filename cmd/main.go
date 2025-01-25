package main

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/etherealblaade/rattus_rex/internal/chain"
	"github.com/etherealblaade/rattus_rex/internal/ui"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	modelChain := chain.NewModelChain()
	if _, err := tea.NewProgram(ui.NewModel(modelChain)).Run(); err != nil {
		log.Fatal(err)
	}
}
