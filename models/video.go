package models

type VideoSearchResult struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:""`
	TotalResults int32   `json:"total_results"`
	NextPage     string  `json:"next_page,omitempty"`
	PrevPage     string  `json:"prev_page,omitempty"`
	Videos       []Video `json:"videos"`
}

type Video struct {
	Id            int32          `json:"id"`
	Width         int32          `json:"width"`
	Height        int32          `json:"height"`
	Url           string         `json:"url"`
	Image         string         `json:"image"`
	FullRes       interface{}    `json:"full_res"`
	Duration      float64        `json:"duration"`
	VideoFiles    []VideoFile    `json:"video_files"`
	VideoPictures []VideoPicture `json:"video_pictures"`
}

type PopularVideos struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	Url          string  `json:"url"`
	Videos       []Video `json:"videos"`
	NextPage     string  `json:"next_page,omitempty"`
	PrevPage     string  `json:"prev_page,omitempty"`
}

type VideoFile struct {
	Id       int32   `json:"id"`
	Quality  string  `json:"quality"`
	FileType string  `json:"file_type"`
	Width    int32   `json:"width"`
	Height   int32   `json:"height"`
	Fps      float64 `json:"fps"`
	Link     string  `json:"link"`
}

type VideoPicture struct {
	Id      int32  `json:"id"`
	Picture string `json:"picture"`
	Nr      int32  `json:"nr"`
}
