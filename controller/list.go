package controller
import (
	"net/http"
	"path/filepath"
	"os"
	"github.com/phpor/go_simple_fileserver/config"
)

func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	w.Write([]byte(`已上传文件（点击下载）<br>`))
	w.Write([]byte(`<ul>`))
	filepath.Walk(config.Get("data_dir"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		w.Write([]byte("<li><a href='/download?name="+path+"'>"+path+"</a></li>"))
		return nil
	})
	w.Write([]byte(`</ul>`))

	html := `
			<form method="post" action="/upload">
			文件名：<input type="text" name="filename" value=""/><br/>
			<input type="file" name="file" /> <br/>
			<input type="submit" value="提交"/>
			</form>
		`


	w.Write([]byte(html))

}
