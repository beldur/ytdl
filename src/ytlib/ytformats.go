package ytlib

type FormatInfo struct {
    Container string
    Resolution string
    Encoding string
    profile string
    bitrate string
    audioEncoding string
    audioBitrate int
}

var YouTube_Formats = map[int]FormatInfo {
    5:  { "flv", "240p", "Sorenson H.283", "", "0.25", "mp3", 64 },
    6:  { "flv", "270p", "Sorenson H.263", "", "0.8", "mp3", 64 },
    13: { "3gp", "",     "MPEG-4 Visual",  "", "0.5", "aac", -1 },
    17: { "3gp", "144p", "MPEG-4 Visual", "simple", "0.05", "aac", 24 },
    18: { "mp4", "360p", "H.264", "baseline", "0.5", "aac", 96 },
    22: { "mp4", "720p", "H.264", "high", "2-2.9", "aac", 152 },
    34: { "flv", "360p", "H.264", "main", "0.5", "aac", 128 },
    35: { "flv", "280p", "H.264", "main", "0.8-1", "aac", 128 },
    36: { "3gp", "240p", "MPEG-4 Visual", "simple", "0.17", "aac", 38 },
    37: { "mp4", "1080p", "H.264", "high", "3-4.3", "aac", 152 },
    38: { "mp4", "3072p", "H.264", "high", "3.5-5", "aac", 152 },
    43: { "webm", "360p", "VP8", "", "0.5", "vorbis", 128 },
    44: { "webm", "480p", "VP8", "", "1", "vorbis", 128 },
    45: { "webm", "720p", "VP8", "", "2", "vorbis", 192 },
    46: { "webm", "1080p", "VP8", "", "", "vorbis", 192 },
    82: { "mp4", "360p", "H.264", "3d", "0.5", "aac", 96 },
    83: { "mp4", "240p", "H.264", "3d", "0.5", "aac", 96 },
    84: { "mp4", "720p", "H.264", "3d", "2-2.9", "aac", 152 },
    85: { "mp4", "520p", "H.264", "3d", "2-2.9", "aac", 152 },
    100: { "webm", "360p", "VP8", "3d", "", "vorbis", 128 },
    101: { "webm", "360p", "VP8", "3d", "", "vorbis", 192 },
    102: { "webm", "720p", "VP8", "3d", "", "vorbis", 192 },
}
