package handler

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/jh123x/mermaid-cli-go/internal/common"
)

func GetMarkdown(config *common.Config) (string, error) {
	if config == nil {
		return "", common.ErrConfigNotFound
	}

	if !strings.HasSuffix(config.InputPath, "."+common.FORMAT_MD) {
		if !config.QuietMode {
			fmt.Println("only markdown file can be converted to markdown format")
		}

		return "", common.ErrNotSupported
	}

	fileReader, err := os.Open(config.InputPath)
	if err != nil {
		return "", err
	}

	contents, err := io.ReadAll(fileReader)
	if err != nil {
		return "", err
	}

	outputDir := path.Dir(config.OutputPath)
	if !config.QuietMode {
		fmt.Println("Output dir of images is", outputDir)
	}

	counter := 0
	c2 := config.Clone()
	c2.OutputFormat = common.FORMAT_SVG

	results := common.MD_REGEX.ReplaceAllStringFunc(
		string(contents),
		func(data string) string {
			data = data[10 : len(data)-6]
			result, err := getDataFromMermaidInput(c2, data)
			if err != nil {
				return err.Error()
			}

			imgName := fmt.Sprintf("img_%d.svg", counter)
			filePath := path.Join(outputDir, imgName)
			if !config.QuietMode {
				fmt.Println("Saving MD Image to ", filePath)
			}
			counter += 1

			file, err := os.Create(filePath)

			if err != nil {
				return err.Error()
			}

			defer func() { _ = file.Close() }()
			file.Write(result)

			return fmt.Sprintf("![Diagram %d](./%s)", counter, imgName)
		},
	)

	file, err := os.Create(config.OutputPath)
	if err != nil {
		return "", err
	}

	if _, err := file.WriteString(results); err != nil {
		return "", err
	}

	return config.OutputPath, nil
}

func GetDiagram(config *common.Config) (string, error) {
	if config == nil {
		return "", common.ErrConfigNotFound
	}

	mermaidVal, err := getInputData(config.InputPath)
	if err != nil {
		return "", err
	}

	result, err := getDataFromMermaidInput(config, mermaidVal)
	if err != nil {
		return "", err
	}

	file, err := os.Create(config.OutputPath)
	if err != nil {
		return "", err
	}

	if _, err := file.Write(result); err != nil {
		return "", err
	}

	return config.OutputPath, nil
}

func getDataFromMermaidInput(config *common.Config, mermaidVal string) ([]byte, error) {
	dir := os.TempDir()
	mermaidPath := path.Join(dir, "index.html")
	template := config.ToTemplate()
	template.Mermaid = mermaidVal

	if err := GenHTML(template, mermaidPath); err != nil {
		return nil, err
	}

	defer func() { _ = os.Remove(mermaidPath) }()

	return launchAndGetImg(mermaidPath, config)
}

func launchAndGetImg(mermaidPath string, config *common.Config) ([]byte, error) {
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
