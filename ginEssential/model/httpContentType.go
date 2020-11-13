package model

//map for HTTP Content-Type HTTP文件类型对应的Content-Type
var HttpContentType = map[string]string{
	".avi":"video/avi",
	".mp3":"audio/mp3",
	".mp4":"video/mp4",
	".wmv":"video/x-ms-wmv",
	".asf":"video/x-ms-wmv",
	".rm":   "application/vnd.rn-realmedia",
	".rmvb": "application/vnd.rn-realmedia-vbr",
	".mov":  "video/quicktime",
	".m4v":  "video/mp4",
	".flv":  "video/x-flv",
	".jpg":  "image/jpeg",
	".png":  "image/png",
}
