package resultBoard

import (
	"fmt"
	"github.com/fatih/color"
)

func (banner_ *ResultBoard) Banner() {
	color.Cyan("   ______       _      __     __      __ \n")
	color.Cyan("  / __/ /____ _| | /| / /__ _/ /_____/ / \n")
	color.Cyan(" _\\ \\/  '_/ // / |/ |/ / _ `/ __/ __/ _ \\\n")
	color.Cyan("/___/_/\\_\\\\_, /|__/|__/\\_,_/\\__/\\__/_//_/\n")
	color.Cyan("         /___/                           \n")
	fmt.Println()
	c := color.New(color.FgCyan)
	d := color.New(color.FgCyan, color.Bold)

	c.Println("=============================================")
	d.Printf("Author: TuF3i\n")
	d.Printf("Github: https://github.com/TuF3i/SkyWatch\n")
	d.Printf("Version: 1.0.0-dev\n")
	c.Println("=============================================")
}
