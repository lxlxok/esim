package new

func init() {
	Files = append(Files, modfc)
}

var (
	modfc = &FileContent{
		FileName: "go.mod",
		Dir:      ".",
		Content:  `module {{.ProPath}}{{.ServerName}}`,
	}
)
