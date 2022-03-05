package main

import (
	"fmt"
	"path/filepath"

	"github.com/brittonhayes/pillager/pkg/hunter"
)

const OldConfigBanner = `██▓███   ██▓ ██▓     ██▓    ▄▄▄        ▄████ ▓█████  ██▀███
▓██░  ██▒▓██▒▓██▒    ▓██▒   ▒████▄     ██▒ ▀█▒▓█   ▀ ▓██ ▒ ██▒
▓██░ ██▓▒▒██▒▒██░    ▒██░   ▒██  ▀█▄  ▒██░▄▄▄░▒███   ▓██ ░▄█ ▒
▒██▄█▓▒ ▒░██░▒██░    ▒██░   ░██▄▄▄▄██ ░▓█  ██▓▒▓█  ▄ ▒██▀▀█▄
▒██▒ ░  ░░██░░██████▒░██████▒▓█   ▓██▒░▒▓███▀▒░▒████▒░██▓ ▒██▒
▒▓▒░ ░  ░░▓  ░ ▒░▓  ░░ ▒░▓  ░▒▒   ▓▒█░ ░▒   ▒ ░░ ▒░ ░░ ▒▓ ░▒▓░
░▒ ░      ▒ ░░ ░ ▒  ░░ ░ ▒  ░ ▒   ▒▒ ░  ░   ░  ░ ░  ░  ░▒ ░ ▒░
░░        ▒ ░  ░ ░     ░ ░    ░   ▒   ░ ░   ░    ░     ░░   ░`

const ConfigBanner = `
░█▀█░▀█▀░█░░░█░░░█▀█░█▀▀░█▀▀░█▀▄
░█▀▀░░█░░█░░░█░░░█▀█░█░█░█▀▀░█▀▄
░▀░░░▀▀▀░▀▀▀░▀▀▀░▀░▀░▀▀▀░▀▀▀░▀░▀
Pillage filesystems for loot.
`

func HunterConfigToMap(h *hunter.Hunter) map[string]string {
	absScanPath, err := filepath.Abs(h.ScanPath)
	if err != nil {
		absScanPath = "./"
	}

	absRulesPath, err := filepath.Abs(h.Gitleaks.Path)
	if err != nil {
		absRulesPath = "./"
	}

	templateStr := h.Template
	if templateStr == "" {
		templateStr = "default"
	}

	return map[string]string{
		AboutView:     ConfigBanner,
		ScanPathView:  absScanPath,
		RulesPathView: absRulesPath,
		FormatView:    fmt.Sprintf("%T", h.Reporter)[7:],
		TemplateView:  templateStr,
		WorkersView:   fmt.Sprintf("%d", h.Workers),
		VerboseView:   fmt.Sprintf("%v", h.Verbose),
		RedactView:    fmt.Sprintf("%v", h.Redact),
	}
}
