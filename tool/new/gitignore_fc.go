package new

func init() {
	Files = append(Files, gitIgnorefc)
}

var (
	gitIgnorefc = &FileContent{
		FileName: ".gitignore",
		Dir:      ".",
		Content: `/{{.ServerName}}
lastupdate.tmp
*.tar.gz
.com.apple*
.idea
*.svg
*.proto`,
	}
)
