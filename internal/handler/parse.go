package handler

import (
	"os"
	"path"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/jh123x/mermaid-cli-go/internal/common"
)

func GetSVG(config *common.Config) (string, error) {
	if config == nil {
		return "", common.ErrConfigNotFound
	}

	mermaidVal, err := getInputData(config.InputPath)
	if err != nil {
		return "", err
	}

	dir := os.TempDir()
	mermaidPath := path.Join(dir, "index.html")
	if err := GenHTML(&common.Template{
		Mermaid: mermaidVal,
		BgColor: config.BgColor,
	}, mermaidPath); err != nil {
		return "", err
	}

	file, err := os.Create(config.OutputPath)
	if err != nil {
		return "", err
	}

	result, err := launchAndGetSVG(mermaidPath, config)
	if err != nil {
		return "", nil
	}

	if _, err := file.Write([]byte(result)); err != nil {
		return "", err
	}

	return config.OutputPath, nil
}

func launchAndGetSVG(mermaidPath string, config *common.Config) (string, error) {
	path, _ := launcher.LookPath()
	launcher := launcher.New().Bin(path).Headless(true).Leakless(false)
	defer launcher.Cleanup()
	controlURL := launcher.MustLaunch()
	page := rod.New().ControlURL(controlURL).Trace(!config.QuietMode).MustConnect().MustPage(mermaidPath)
	defer func() { _ = page.Close() }()

	svg := page.MustWaitStable().MustElement(common.Selector)

	result, err := svg.HTML()
	if err != nil {
		return "", err
	}

	launcher.Kill()

	return result, nil
}
