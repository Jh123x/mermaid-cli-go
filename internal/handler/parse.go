package handler

import (
	"io"
	"os"
	"path"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
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
	template := config.ToTemplate()
	template.Mermaid = mermaidVal

	if err := GenHTML(template, mermaidPath); err != nil {
		return "", err
	}

	file, err := os.Create(config.OutputPath)
	if err != nil {
		return "", err
	}

	result, err := launchAndGetImg(mermaidPath, config, file)
	if err != nil {
		return "", nil
	}

	if _, err := file.Write(result); err != nil {
		return "", err
	}

	return config.OutputPath, nil
}

func launchAndGetImg(mermaidPath string, config *common.Config, file *os.File) ([]byte, error) {
	path, _ := launcher.LookPath()
	launcher := launcher.New().Bin(path).Headless(true).Leakless(false)
	defer launcher.Cleanup()
	defer launcher.Kill()

	controlURL := launcher.MustLaunch()
	page := rod.New().
		ControlURL(controlURL).
		Trace(!config.QuietMode).
		MustConnect().
		MustPage(mermaidPath).
		MustWaitDOMStable()
	defer func() { _ = page.Close() }()

	switch config.OutputFormat {
	case common.FORMAT_SVG:
		return getSVG(page)
	case common.FORMAT_PNG, common.FORMAT_JPEG, common.FORMAT_WEBP:
		return getImg(page, config.OutputFormat)
	case common.FORMAT_PDF:
		return getPDF(page, config)
	case common.FORMAT_MD:
		return nil, common.ErrNotSupported
	default:
		return nil, common.ErrInvalidOutputFormat
	}

}

func getSVG(page *rod.Page) ([]byte, error) {
	svg := page.MustElement(common.Selector)
	res, err := svg.HTML()
	if err != nil {
		return nil, err
	}

	return []byte(res), nil
}

func getImg(page *rod.Page, format string) ([]byte, error) {
	captureConfig := &proto.PageCaptureScreenshot{}
	switch format {
	case common.FORMAT_PNG:
		captureConfig.Format = proto.PageCaptureScreenshotFormatPng
	case common.FORMAT_WEBP:
		captureConfig.Format = proto.PageCaptureScreenshotFormatWebp
	case common.FORMAT_JPEG:
		captureConfig.Format = proto.PageCaptureScreenshotFormatJpeg
	default:
		return nil, common.ErrNotSupported
	}
	return page.Screenshot(true, captureConfig)
}

func getPDF(page *rod.Page, config *common.Config) ([]byte, error) {
	reader, err := page.PDF(&proto.PagePrintToPDF{
		DisplayHeaderFooter: false,
		Scale:               common.GetPtrOf(float64(config.Scale)),
	})
	if err != nil {
		return nil, err
	}

	return io.ReadAll(reader)
}
