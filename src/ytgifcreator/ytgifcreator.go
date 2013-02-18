package ytgifcreator

import (
    "net/http"
    "html/template"
)

var templates = template.Must(template.ParseFiles("templates/index.html"))

func init() {
    http.HandleFunc("/", indexHandler)
    /*ytVideo := new(ytlib.YTVideo).Init("8k9XnnqACdc")
    ytVideo.GetFormatList()
    title := ytVideo.VideoInformation["title"][0]

    fmt.Println(title, ytVideo.FormatList)

    format, _ := ytVideo.GetBestQuality()
    ytVideo.DownloadVideo(title,
        ytlib.DownloadOptions { Format: format, Begin: 60000 },
    )*/
}

func renderTemplate(w http.ResponseWriter, template string, data interface{}) {
    err := templates.ExecuteTemplate(w, template, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "index.html", struct {}{})
}
