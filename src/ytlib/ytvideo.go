package ytlib

import (
    "net/http"
    "io/ioutil"
    "regexp"
    "strconv"
    "os"
    "strings"
    "fmt"
)

// YouTube Video url
const YT_URL = "http://www.youtube.com/watch?v="

// Represents a YouTube video
type YTVideo struct {
    VideoId string
    FormatList map[int]string
    qualityOrder []int
    filetypes map[int]string
}

// Get worst possible format id
func (this *YTVideo) getWorstQuality() (int, error) {
    for i := len(this.qualityOrder) - 1; i > -1; i-- {
        format := this.qualityOrder[i]
        if _, ok := this.FormatList[format]; ok {
            return format, nil
        }
    }

    return 0, fmt.Errorf("Could not find a format!")
}

// Get best possible format id
func (this *YTVideo) getBestQuality() (int, error) {
    for _, format := range this.qualityOrder {
        if _, ok := this.FormatList[format]; ok {
            return format, nil
        }
    }

    return 0, fmt.Errorf("Could not find a format!")
}

// Download worst possible quality
func (this *YTVideo) DownloadWorstQuality(filename string) {
    worstFormat, _ := this.getWorstQuality()
    this.DownloadVideo(worstFormat, filename)
}

// Download best possible quality
func (this *YTVideo) DownloadBestQuality(filename string) {
    format, _ := this.getBestQuality()
    this.DownloadVideo(format, filename)
}

// Download video for given format and save it into given name
func (this *YTVideo) DownloadVideo(format int, name string) error {
    if url, ok := this.FormatList[format]; ok {
        filename := name + this.filetypes[format]

        return this.download(url, filename)
    }

    return fmt.Errorf("Format not found!")
}

// Create a new YTVideo struct
func (this *YTVideo) Init(videoId string) *YTVideo {
    return &YTVideo {
        videoId,
        make(map[int]string, 0),
        []int { 37, 46, 22, 45, 35, 44, 18, 34, 43, 36, 5, 17 },
        map[int]string {
            37: ".mpg", 46: ".flv", // 1080p
            22: ".mpg", 45: ".flv", // 720p
            35: ".mpg", 44: ".flv", // 480p
            18: ".mpg", 34: ".mpg", 43: ".flv", // 360p
            36: ".mpg", 5: ".flv", // 240p
            17: ".mpg", // 114p
        },
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
    this.FormatList = make(map[int]string, 0)
    reg, _ := regexp.Compile("\"url_encoded_fmt_stream_map\": \"([^\"]*)")
    regItag, _ := regexp.Compile("itag=([0-9]+)")

    url_encoded := reg.FindSubmatch(body)

    resultString := string(url_encoded[1])
    resultString = replace(resultString, []string { "%25", "\\u0026", "\\" }, []string { "%", "&", "" })

    // do some replace magic on the url for each video format
    for _, v := range strings.Split(resultString, ",") {

        t := strings.SplitN(v, "url=http", 2)
        v = "url=http" + t[1] + "&" + t[0]
        v = strings.Replace(v, "url=http%3A%2F%2F", "http://", 1)
        v = replace(v, []string { "%3F", "%2F", "%3D", "%26", "%252C", "\\u0026", "sig=" },
                       []string { "?",   "/",   "=",   "&",   "%2C",   "&",       "signature=" })

        itag, _ := strconv.Atoi(regItag.FindStringSubmatch(v)[1])

        if strings.Count(v, "itag=") > 1 {
            v = strings.Replace(v, fmt.Sprintf("&itag=%d", itag), "", 1)
        }

        // Add video url to result
        this.FormatList[itag] = v
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


// Replace each string in search with corresponding string in replace
func replace(value string, search []string, replace []string) string {
    for i := range replace {
        value = strings.Replace(value, search[i], replace[i], -1)
    }

    return value
}
