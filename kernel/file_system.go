package kernel

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

type Content struct {
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Extension string `json:"extension,omitempty"`
	IsDir     bool   `json:"is_dir"`
	ModTime   string `json:"mod_time"`
	Mode      string `json:"mode"`
}

func (k *Kernel) GetDirContent(p string) ([]byte, error) {
	files, _ := ioutil.ReadDir(p)

	var contents []Content

	for _, f := range files {
		content := Content{
			Name:    f.Name(),
			Size:    f.Size(),
			IsDir:   f.IsDir(),
			ModTime: f.ModTime().String(),
			Mode:    f.Mode().String(),
		}
		if !f.IsDir() {
			content.Extension = path.Ext(p + f.Name())
		}
		contents = append(contents, content)
	}
	b, err := json.Marshal(contents)
	if err != nil {
		return nil, err
	}
	return b, nil
}
