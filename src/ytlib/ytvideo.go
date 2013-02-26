package ytlib

import (
    "net/http"
    "io/ioutil"
    "os"
    "strings"
    "fmt"
    "net/url"
    "strconv"
)

// Represents a YouTube video
type YTVideo struct {
    VideoId string
    FormatList map[int]string
    qualityOrder []int
    VideoInformation map[string][]string
}

// Get worst possible format id
func (this *YTVideo) GetWorstQuality() (int, error) {
    for i := len(this.qualityOrder) - 1; i > -1; i-- {
        format := this.qualityOrder[i]
        if _, ok := this.FormatList[format]; ok {
            return format, nil
        }
    }

    return 0, fmt.Errorf("Could not find a format!")
}

// Get best possible format id
func (this *YTVideo) GetBestQuality() (int, error) {
    for _, format := range this.qualityOrder {
        if _, ok := this.FormatList[format]; ok {
            return format, nil
        }
    }

    return 0, fmt.Errorf("Could not find a format!")
}

// Download worst possible quality
func (this *YTVideo) DownloadWorstQuality(filename string) {
    format, _ := this.GetWorstQuality()
    this.DownloadVideo(filename, DownloadOptions { Format: format })
}

// Download best possible quality
func (this *YTVideo) DownloadBestQuality(filename string) {
    format, _ := this.GetBestQuality()
    this.DownloadVideo(filename, DownloadOptions{ Format: format })
}

// Download video for given format and save it into given name
func (this *YTVideo) DownloadVideo(name string, downloadOptions DownloadOptions) error {
    if url, ok := this.FormatList[downloadOptions.Format]; ok {
        filename := name + "." + YouTube_Formats[downloadOptions.Format].Container

        if downloadOptions.Start > 0 {
            url = url + "&begin=" + strconv.Itoa(downloadOptions.Start)
        }

        fmt.Println(url)
        return this.download(url, filename)
    }

    return fmt.Errorf("Format not found!")
}

// Create a new YTVideo struct
func (this *YTVideo) Init(videoId string) *YTVideo {
    return &YTVideo {
        videoId,
        map[int]string {},
        []int { 37, 46, 22, 45, 35, 44, 18, 34, 43, 36, 5, 17 },
        map[string][]string {},
    }
}

// Get Formatlist for given videoId
func (this *YTVideo) GetFormatList() (map[int]string, error) {
    resp, err := http.Get(fmt.Sprintf("%s%s", YT_URL, this.VideoId))
    defer resp.Body.Close()

    if err != nil {
        return nil, err
    }

    body, err := ioutil.ReadAll(resp.Body)

    return this.parseBody(body)
}

// Parse Page source for formatlist
func (this *YTVideo) parseBody(body []byte) (map[int]string, error) {

    videoInfo, err := url.ParseQuery(string(body))
    if err != nil {
        return nil, fmt.Errorf("Could not parse video info")
    }

    this.VideoInformation = videoInfo
    this.FormatList = make(map[int]string, 0)

    // Split format list
    for _, v := range strings.Split(videoInfo["url_encoded_fmt_stream_map"][0], ",") {
        formatValues, err := url.ParseQuery(v)
        if err != nil {
            continue
        }

        itag, err := strconv.Atoi(formatValues["itag"][0])
        if err != nil {
            continue
        }

        url := formatValues["url"][0]
        if sig, ok := formatValues["sig"]; ok {
            url += "&signature=" + sig[0]
        }

        // Add video url to result
        this.FormatList[itag] = url
    }

    return this.FormatList, nil
}

// Download Video from given url
func (this *YTVideo) download(url string, filename string) error {
    buffer := make([]byte, 1024)

    fmt.Println("Create file", filename)

    file, err := os.Create(filename)
    if err != nil { return err }
    defer file.Close()

    resp, err := http.Get(url)
    if err != nil { return err }
    defer resp.Body.Close()

    fmt.Printf("Reading %d kB", resp.ContentLength / 1000)

    writeCounter := 0
    for {
        n, err := resp.Body.Read(buffer);

        if n == 0 || err != nil {
            break
        }

        n2, _ := file.Write(buffer[:n])
        writeCounter += n2

        // Print point each mb
        if writeCounter > 1000000 {
            writeCounter = 0
            fmt.Print(".");
        }
    }

    fmt.Println("done")

    return nil
}

func (this *YTVideo) HasFormat(format int) bool {
    _, ok := this.FormatList[format]

    return ok
}

// Replace each string in search with corresponding string in replace
func replace(value string, search []string, replace []string) string {
    for i := range replace {
        value = strings.Replace(value, search[i], replace[i], -1)
    }

    return value
}
