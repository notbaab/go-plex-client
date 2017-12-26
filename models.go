package plex

import (
	"encoding/xml"
	"net/http"
)

// Plex contains fields that are required to make
// an api call to your plex server
type Plex struct {
	URL        string
	Token      string
	HTTPClient http.Client
}

// SearchResults a list of media returned when searching
// for media via your plex server
type SearchResults struct {
	MediaContainer struct {
		Metadata []struct {
			Genre []struct {
				Tag string `json:"tag"`
			} `json:"Genre"`
			Role []struct {
				Tag string `json:"tag"`
			} `json:"Role"`
			AddedAt               int64   `json:"addedAt"`
			AllowSync             bool    `json:"allowSync"`
			Art                   string  `json:"art"`
			Banner                string  `json:"banner"`
			ChildCount            int64   `json:"childCount"`
			ContentRating         string  `json:"contentRating"`
			Duration              int64   `json:"duration"`
			Index                 int64   `json:"index"`
			Key                   string  `json:"key"`
			LastViewedAt          int64   `json:"lastViewedAt"`
			LeafCount             int64   `json:"leafCount"`
			LibrarySectionID      int64   `json:"librarySectionID"`
			LibrarySectionTitle   string  `json:"librarySectionTitle"`
			LibrarySectionUUID    string  `json:"librarySectionUUID"`
			OriginallyAvailableAt string  `json:"originallyAvailableAt"`
			Personal              bool    `json:"personal"`
			Rating                float64 `json:"rating"`
			RatingKey             string  `json:"ratingKey"`
			SourceTitle           string  `json:"sourceTitle"`
			Studio                string  `json:"studio"`
			Summary               string  `json:"summary"`
			Theme                 string  `json:"theme"`
			Thumb                 string  `json:"thumb"`
			Title                 string  `json:"title"`
			Type                  string  `json:"type"`
			UpdatedAt             int64   `json:"updatedAt"`
			ViewCount             int64   `json:"viewCount"`
			ViewedLeafCount       int64   `json:"viewedLeafCount"`
			Year                  int64   `json:"year"`
		} `json:"Metadata"`
		Provider []struct {
			Key   string `json:"key"`
			Title string `json:"title"`
			Type  string `json:"type"`
		} `json:"Provider"`
		Identifier      string `json:"identifier"`
		MediaTagPrefix  string `json:"mediaTagPrefix"`
		MediaTagVersion int64  `json:"mediaTagVersion"`
		Size            int64  `json:"size"`
	} `json:"MediaContainer"`
}

// MediaMetadata ...
type MediaMetadata struct {
	MediaContainer struct {
		Metadata []struct {
			Director []struct {
				Filter string `json:"filter"`
				ID     int64  `json:"id"`
				Tag    string `json:"tag"`
			} `json:"Director"`
			Media []struct {
				Part []struct {
					Stream []struct {
						AudioChannelLayout string  `json:"audioChannelLayout"`
						BitDepth           int64   `json:"bitDepth"`
						Bitrate            int64   `json:"bitrate"`
						BitrateMode        string  `json:"bitrateMode"`
						Cabac              bool    `json:"cabac"`
						Channels           int64   `json:"channels"`
						ChromaSubsampling  string  `json:"chromaSubsampling"`
						Codec              string  `json:"codec"`
						CodecID            string  `json:"codecID"`
						ColorRange         string  `json:"colorRange"`
						ColorSpace         string  `json:"colorSpace"`
						Default            bool    `json:"default"`
						DialogNorm         int64   `json:"dialogNorm"`
						Duration           int64   `json:"duration"`
						FrameRate          float64 `json:"frameRate"`
						FrameRateMode      string  `json:"frameRateMode"`
						HasScalingMatrix   bool    `json:"hasScalingMatrix"`
						Height             int64   `json:"height"`
						ID                 int64   `json:"id"`
						Index              int64   `json:"index"`
						Language           string  `json:"language"`
						LanguageCode       string  `json:"languageCode"`
						Level              int64   `json:"level"`
						PixelFormat        string  `json:"pixelFormat"`
						Profile            string  `json:"profile"`
						RefFrames          int64   `json:"refFrames"`
						SamplingRate       int64   `json:"samplingRate"`
						ScanType           string  `json:"scanType"`
						Selected           bool    `json:"selected"`
						StreamType         int64   `json:"streamType"`
						Width              int64   `json:"width"`
					} `json:"Stream"`
					Container    string `json:"container"`
					Duration     int64  `json:"duration"`
					File         string `json:"file"`
					ID           int64  `json:"id"`
					Key          string `json:"key"`
					Size         int64  `json:"size"`
					VideoProfile string `json:"videoProfile"`
				} `json:"Part"`
				AspectRatio     float64 `json:"aspectRatio"`
				AudioChannels   int64   `json:"audioChannels"`
				AudioCodec      string  `json:"audioCodec"`
				Bitrate         int64   `json:"bitrate"`
				Container       string  `json:"container"`
				Duration        int64   `json:"duration"`
				Height          int64   `json:"height"`
				ID              int64   `json:"id"`
				VideoCodec      string  `json:"videoCodec"`
				VideoFrameRate  string  `json:"videoFrameRate"`
				VideoProfile    string  `json:"videoProfile"`
				VideoResolution string  `json:"videoResolution"`
				Width           int64   `json:"width"`
			} `json:"Media"`
			Writer []struct {
				Filter string `json:"filter"`
				ID     int64  `json:"id"`
				Tag    string `json:"tag"`
			} `json:"Writer"`
			AddedAt               int64   `json:"addedAt"`
			Art                   string  `json:"art"`
			ContentRating         string  `json:"contentRating"`
			Duration              int64   `json:"duration"`
			GrandparentArt        string  `json:"grandparentArt"`
			GrandparentKey        string  `json:"grandparentKey"`
			GrandparentRatingKey  string  `json:"grandparentRatingKey"`
			GrandparentTheme      string  `json:"grandparentTheme"`
			GrandparentThumb      string  `json:"grandparentThumb"`
			GrandparentTitle      string  `json:"grandparentTitle"`
			GUID                  string  `json:"guid"`
			Index                 int64   `json:"index"`
			Key                   string  `json:"key"`
			LastViewedAt          int64   `json:"lastViewedAt"`
			LibrarySectionID      int64   `json:"librarySectionID"`
			LibrarySectionKey     string  `json:"librarySectionKey"`
			OriginallyAvailableAt string  `json:"originallyAvailableAt"`
			ParentIndex           int64   `json:"parentIndex"`
			ParentKey             string  `json:"parentKey"`
			ParentRatingKey       string  `json:"parentRatingKey"`
			ParentThumb           string  `json:"parentThumb"`
			ParentTitle           string  `json:"parentTitle"`
			Rating                float64 `json:"rating"`
			RatingKey             string  `json:"ratingKey"`
			Summary               string  `json:"summary"`
			Thumb                 string  `json:"thumb"`
			Title                 string  `json:"title"`
			Type                  string  `json:"type"`
			UpdatedAt             int64   `json:"updatedAt"`
			ViewCount             int64   `json:"viewCount"`
			ViewOffset            int64   `json:"viewOffset"`
			Year                  int64   `json:"year"`
		} `json:"Metadata"`
		AllowSync           bool   `json:"allowSync"`
		Identifier          string `json:"identifier"`
		LibrarySectionID    int64  `json:"librarySectionID"`
		LibrarySectionTitle string `json:"librarySectionTitle"`
		LibrarySectionUUID  string `json:"librarySectionUUID"`
		MediaTagPrefix      string `json:"mediaTagPrefix"`
		MediaTagVersion     int64  `json:"mediaTagVersion"`
		Size                int64  `json:"size"`
	} `json:"MediaContainer"`

	Size                string `json:"size" xml:"size,attr"`
	AllowSync           string `json:"allowSync" xml:"allowSync,attr"`
	Identifier          string `json:"identifier" xml:"identifier,attr"`
	LibrarySectionID    string `json:"librarySectionID" xml:"librarySectionID,attr"`
	LibrarySectionTitle string `json:"librarySectionTitle" xml:"librarySectionTitle,attr"`
	LibrarySectionUUID  string `json:"librarySectionUUID" xml:"librarySectionUUID,attr"`
	MediaTagPrefix      string `json:"mediaTagPrefix" xml:"mediaTagPrefix,attr"`
	MediaTagVersion     string `json:"mediaTagVersion" xml:"mediaTagVersion,attr"`
	Video               struct {
		RatingKey            string `json:"ratingKey" xml:"ratingKey,attr"`
		Key                  string `json:"key" xml:"key,attr"`
		GrandparentTitle     string `json:"grandparentTitle" xml:"grandparentTitle,attr"`
		GrandparentRatingKey string `json:"grandparentRatingKey" xml:"grandparentRatingKey,attr"`
		ParentRatingKey      string `json:"parentRatingKey" xml:"parentRatingKey,attr"`
		ParentYear           string `json:"parentYear" xml:"parentYear,attr"`
		ParentTitle          string `json:"parentTitle" xml:"parentTitle,attr"`
		GUID                 string `json:"guid" xml:"guid,attr"`
		LibrarySectionID     string `json:"librarySectionID" xml:"librarySectionID,attr"`
		Type                 string `json:"type" xml:"type,attr"`
		Title                string `json:"title" xml:"title,attr"`
		Summary              string `json:"summary" xml:"summary,attr"`
		ViewCount            string `json:"viewCount" xml:"viewCount,attr"`
		LastViewedAt         string `json:"lastViewedAt" xml:"lastViewedAt,attr"`
		Year                 string `json:"year" xml:"year,attr"`
		Thumb                string `json:"thumb" xml:"thumb,attr"`
		Art                  string `json:"art" xml:"art,attr"`
		Duration             string `json:"duration" xml:"duration,attr"`
		AddedAt              string `json:"addedAt" xml:"addedAt,attr"`
		UpdatedAt            string `json:"updatedAt" xml:"updatedAt,attr"`
		ChapterSource        string `json:"chapterSource" xml:"chapterSource,attr"`
		Media                struct {
			VideoResolution string `json:"videoResolution" xml:"videoResolution,attr"`
			ID              string `json:"id" xml:"id,attr"`
			Duration        string `json:"duration" xml:"duration,attr"`
			Bitrate         string `json:"bitrate" xml:"bitrate,attr"`
			Width           string `json:"width" xml:"width,attr"`
			Height          string `json:"height" xml:"height,attr"`
			AspectRatio     string `json:"aspectRatio" xml:"aspectRatio,attr"`
			AudioChannels   string `json:"audioChannels" xml:"audioChannels,attr"`
			AudioCodec      string `json:"audioCodec" xml:"audioCodec,attr"`
			VideoCodec      string `json:"videoCodec" xml:"videoCodec,attr"`
			Container       string `json:"container" xml:"container,attr"`
			VideoFrameRate  string `json:"videoFrameRate" xml:"videoFrameRate,attr"`
			VideoProfile    string `json:"videoProfile" xml:"videoProfile,attr"`
			Part            struct {
				ID           string `json:"id" xml:"id,attr"`
				Key          string `json:"key" xml:"key,attr"`
				Duration     string `json:"duration" xml:"duration,attr"`
				File         string `json:"file" xml:"file,attr"`
				Size         string `json:"size" xml:"size,attr"`
				Container    string `json:"container" xml:"container,attr"`
				VideoProfile string `json:"videoProfile" xml:"videoProfile,attr"`
				Stream       []struct {
					ID                 string `json:"id" xml:"id,attr"`
					StreamType         string `json:"streamType" xml:"streamType,attr"`
					Default            string `json:"default" xml:"default,attr"`
					Codec              string `json:"codec" xml:"codec,attr"`
					Index              string `json:"index" xml:"index,attr"`
					Bitrate            string `json:"bitrate" xml:"bitrate,attr"`
					BitDepth           string `json:"bitDepth" xml:"bitDepth,attr"`
					Cabac              string `json:"cabac" xml:"cabac,attr"`
					ChromaSubsampling  string `json:"chromaSubsampling" xml:"chromaSubsampling,attr"`
					CodecID            string `json:"codecID" xml:"codecID,attr"`
					ColorRange         string `json:"colorRange" xml:"colorRange,attr"`
					ColorSpace         string `json:"colorSpace" xml:"colorSpace,attr"`
					Duration           string `json:"duration" xml:"duration,attr"`
					FrameRate          string `json:"frameRate" xml:"frameRate,attr"`
					FrameRateMode      string `json:"frameRateMode" xml:"frameRateMode,attr"`
					HasScalingMatrix   string `json:"hasScalingMatrix" xml:"hasScalingMatrix,attr"`
					HeaderStripping    string `json:"headerStripping" xml:"headerStripping,attr"`
					Height             string `json:"height" xml:"height,attr"`
					Level              string `json:"level" xml:"level,attr"`
					PixelFormat        string `json:"pixelFormat" xml:"pixelFormat,attr"`
					Profile            string `json:"profile" xml:"profile,attr"`
					RefFrames          string `json:"refFrames" xml:"refFrames,attr"`
					ScanType           string `json:"scanType" xml:"scanType,attr"`
					Width              string `json:"width" xml:"width,attr"`
					Selected           string `json:"selected" xml:"selected,attr"`
					Channels           string `json:"channels" xml:"channels,attr"`
					AudioChannelLayout string `json:"audioChannelLayout" xml:"audioChannelLayout,attr"`
					BitrateMode        string `json:"bitrateMode" xml:"bitrateMode,attr"`
					DialogNorm         string `json:"dialogNorm" xml:"dialogNorm,attr"`
					SamplingRate       string `json:"samplingRate" xml:"samplingRate,attr"`
				} `json:"stream" xml:"Stream"`
			} `json:"part" xml:"Part"`
		} `json:"media" xml:"Media"`
		Label struct {
			ID  string `json:"id" xml:"id,attr"`
			Tag string `json:"tag" xml:"tag,attr"`
		} `json:"label" xml:"Label"`
		Field []struct {
			Name   string `json:"name" xml:"name,attr"`
			Locked string `json:"locked" xml:"locked,attr"`
		} `json:"field" xml:"Field"`
	} `json:"video" xml:"Video"`
	Directory struct {
		RatingKey             string `json:"ratingKey" xml:"ratingKey,attr"`
		Key                   string `json:"key" xml:"key,attr"`
		GUID                  string `json:"guid" xml:"guid,attr"`
		LibrarySectionID      string `json:"librarySectionID" xml:"librarySectionID,attr"`
		ParentTitle           string `json:"parentTitle" xml:"parentTitle,attr"`
		ParentYear            string `json:"parentYear" xml:"parentYear,attr"`
		Studio                string `json:"studio" xml:"studio,attr"`
		Type                  string `json:"type" xml:"type,attr"`
		Title                 string `json:"title" xml:"title,attr"`
		TitleSort             string `json:"titleSort" xml:"titleSort,attr"`
		ContentRating         string `json:"contentRating" xml:"contentRating,attr"`
		Summary               string `json:"summary" xml:"summary,attr"`
		Index                 string `json:"index" xml:"index,attr"`
		Rating                string `json:"rating" xml:"rating,attr"`
		ViewCount             string `json:"viewCount" xml:"viewCount,attr"`
		LastViewedAt          string `json:"lastViewedAt" xml:"lastViewedAt,attr"`
		Year                  string `json:"year" xml:"year,attr"`
		Thumb                 string `json:"thumb" xml:"thumb,attr"`
		Art                   string `json:"art" xml:"art,attr"`
		Banner                string `json:"banner" xml:"banner,attr"`
		Theme                 string `json:"theme" xml:"theme,attr"`
		Duration              string `json:"duration" xml:"duration,attr"`
		OriginallyAvailableAt string `json:"originallyAvailableAt" xml:"originallyAvailableAt,attr"`
		LeafCount             string `json:"leafCount" xml:"leafCount,attr"`
		ViewedLeafCount       string `json:"viewedLeafCount" xml:"viewedLeafCount,attr"`
		ChildCount            string `json:"childCount" xml:"childCount,attr"`
		AddedAt               string `json:"addedAt" xml:"addedAt,attr"`
		UpdatedAt             string `json:"updatedAt" xml:"updatedAt,attr"`
		Genre                 []struct {
			ID  string `json:"id" xml:"id,attr"`
			Tag string `json:"tag" xml:"tag,attr"`
		} `json:"genre" xml:"Genre"`
		Role []struct {
			ID    string `json:"id" xml:"id,attr"`
			Tag   string `json:"tag" xml:"tag,attr"`
			Role  string `json:"role" xml:"role,attr"`
			Thumb string `json:"thumb" xml:"thumb,attr"`
		} `json:"role" xml:"Role"`
		Field struct {
			Name   string `json:"name" xml:"name,attr"`
			Locked string `json:"locked" xml:"locked,attr"`
		} `json:"field" xml:"Field"`
		Location string `json:"location" xml:"Location"`
	} `json:"directory" xml:"Directory"`
}

// Location is the path of a plex server directory
type Location struct {
	ID   int    `json:"id"`
	Path string `json:"path"`
}

// Directory shows plex directory metadata
type Directory struct {
	Location   []Location `json:"Location"`
	Agent      string     `json:"agent"`
	AllowSync  bool       `json:"allowSync"`
	Art        string     `json:"art"`
	Composite  string     `json:"composite"`
	CreatedAt  int        `json:"createdAt"`
	Filter     bool       `json:"filters"`
	Key        string     `json:"key"`
	Language   string     `json:"language"`
	Refreshing bool       `json:"refreshing"`
	Scanner    string     `json:"scanner"`
	Thumb      string     `json:"thumb"`
	Title      string     `json:"title"`
	Type       string     `json:"type"`
	UpdatedAt  int        `json:"updatedAt"`
	UUID       string     `json:"uuid"`
}

// LibrarySections metadata of your library contents
type LibrarySections struct {
	MediaContainer struct {
		Directory []Directory `json:"Directory"`
	} `json:"MediaContainer"`
}

// TaggedData ...
type TaggedData struct {
	Tag string `json:"tag"`
}

// Media ...
type Media struct {
	Part            []Part  `json:"Part"`
	AspectRatio     float64 `json:"aspectRatio"`
	AudioChannels   int8    `json:"audioChannels"`
	AudioCodec      string  `json:"audioCodec"`
	Bitrate         int64   `json:"bitrate"`
	Container       string  `json:"container"`
	Duration        int64   `json:"duration"`
	Height          int64   `json:"height"`
	ID              int64   `json:"id"`
	Selected        bool    `json:"selected"`
	VideoCodec      string  `json:"videoCodec"`
	VideoFrameRate  string  `json:"videoFrameRate"`
	VideoProfile    string  `json:"videoProfile"`
	VideoResolution string  `json:"videoResolution"`
	Width           int64   `json:"width"`
}

type SectionType string

const (
	SECTION_TYPES_MOVIE   SectionType = "movie"
	SECTION_TYPES_TV_SHOW SectionType = "show"
)

type SectionMetadata struct {
	Country               []TaggedData `json: "Country"`
	Director              []TaggedData `json: "Director"`
	Genre                 []TaggedData `json: "Genre"`
	Role                  []TaggedData `json: "Role"`
	Writer                []TaggedData `json: "Writer"`
	Media                 []Media      `json: "Media"`
	AddedAt               int          `json: "addedAt"`
	Art                   string       `json: "art"`
	AudienceRating        float64      `json: "audienceRating"`
	AudienceRatingImage   string       `json: "audienceRatingImage"`
	ChapterSource         string       `json: "chapterSource"`
	ContentRating         string       `json: "contentRating"`
	Duration              int          `json: "duration"`
	Key                   string       `json: "key"`
	OriginallyAvailableAt string       `json: "originallyAvailableAt"`
	PrimaryExtraKey       string       `json: "primaryExtraKey"`
	Rating                float64      `json: "rating"`
	RatingImage           string       `json: "ratingImage"`
	RatingKey             string       `json: "ratingKey"`
	Studio                string       `json: "studio"`
	Summary               string       `json: "summary"`
	Tagline               string       `json: "tagline"`
	Thumb                 string       `json: "thumb"`
	Title                 string       `json: "title"`
	Type                  SectionType  `json: "type"`
	UpdatedAt             int          `json: "updatedAt"`
	Year                  int          `json: "year"`
}

type Section struct {
	MediaContainer struct {
		Metadata []SectionMetadata `json: "Metadata"`
	} `json: "MedaiContainer"`
}

// MetadataChildren returns metadata about a piece of media (tv show, movie, music, etc)
type MetadataChildren struct {
	MediaContainer struct {
		Directory []struct {
			Key             string `json:"key"`
			LeafCount       int64  `json:"leafCount"`
			Thumb           string `json:"thumb"`
			Title           string `json:"title"`
			ViewedLeafCount int64  `json:"viewedLeafCount"`
		} `json:"Directory"`
		Metadata []struct {
			AddedAt         int64  `json:"addedAt"`
			Art             string `json:"art"`
			Index           int64  `json:"index"`
			Key             string `json:"key"`
			LastViewedAt    int64  `json:"lastViewedAt"`
			LeafCount       int64  `json:"leafCount"`
			ParentIndex     int64  `json:"parentIndex"`
			ParentKey       string `json:"parentKey"`
			ParentRatingKey string `json:"parentRatingKey"`
			ParentTheme     string `json:"parentTheme"`
			ParentThumb     string `json:"parentThumb"`
			ParentTitle     string `json:"parentTitle"`
			RatingKey       string `json:"ratingKey"`
			Summary         string `json:"summary"`
			Thumb           string `json:"thumb"`
			Title           string `json:"title"`
			Type            string `json:"type"`
			UpdatedAt       int64  `json:"updatedAt"`
			ViewCount       int64  `json:"viewCount"`
			ViewedLeafCount int64  `json:"viewedLeafCount"`
		} `json:"Metadata"`
		AllowSync           bool   `json:"allowSync"`
		Art                 string `json:"art"`
		Banner              string `json:"banner"`
		Identifier          string `json:"identifier"`
		Key                 string `json:"key"`
		LibrarySectionID    int64  `json:"librarySectionID"`
		LibrarySectionTitle string `json:"librarySectionTitle"`
		LibrarySectionUUID  string `json:"librarySectionUUID"`
		MediaTagPrefix      string `json:"mediaTagPrefix"`
		MediaTagVersion     int64  `json:"mediaTagVersion"`
		Nocache             bool   `json:"nocache"`
		ParentIndex         int64  `json:"parentIndex"`
		ParentTitle         string `json:"parentTitle"`
		ParentYear          int64  `json:"parentYear"`
		Size                int64  `json:"size"`
		Summary             string `json:"summary"`
		Theme               string `json:"theme"`
		Thumb               string `json:"thumb"`
		Title1              string `json:"title1"`
		Title2              string `json:"title2"`
		ViewGroup           string `json:"viewGroup"`
		ViewMode            int64  `json:"viewMode"`
	} `json:"MediaContainer"`
}

// SearchResultsEpisode contains metadata about an episode
type SearchResultsEpisode struct {
	Children []struct {
		Children []struct {
			Children []struct {
				ElementType           string `json:"_elementType"`
				Container             string `json:"container"`
				Duration              int    `json:"duration"`
				File                  string `json:"file"`
				Has64bitOffsets       bool   `json:"has64bitOffsets"`
				HasThumbnail          string `json:"hasThumbnail"`
				ID                    string `json:"id"`
				Key                   string `json:"key"`
				OptimizedForStreaming bool   `json:"optimizedForStreaming"`
				Size                  int    `json:"size"`
				VideoProfile          string `json:"videoProfile"`
			} `json:"_children"`
			ElementType           string `json:"_elementType"`
			AspectRatio           string `json:"aspectRatio"`
			AudioChannels         int    `json:"audioChannels"`
			AudioCodec            string `json:"audioCodec"`
			Bitrate               int    `json:"bitrate"`
			Container             string `json:"container"`
			Duration              int    `json:"duration"`
			Has64bitOffsets       bool   `json:"has64bitOffsets"`
			Height                int    `json:"height"`
			ID                    int    `json:"id"`
			OptimizedForStreaming int    `json:"optimizedForStreaming"`
			VideoCodec            string `json:"videoCodec"`
			VideoFrameRate        string `json:"videoFrameRate"`
			VideoProfile          string `json:"videoProfile"`
			VideoResolution       string `json:"videoResolution"`
			Width                 int    `json:"width"`
		} `json:"_children"`
		ElementType           string `json:"_elementType"`
		AddedAt               int    `json:"addedAt"`
		ChapterSource         string `json:"chapterSource"`
		Duration              int    `json:"duration"`
		Index                 int    `json:"index"`
		Key                   string `json:"key"`
		LastViewedAt          int    `json:"lastViewedAt"`
		OriginallyAvailableAt string `json:"originallyAvailableAt"`
		ParentKey             string `json:"parentKey"`
		ParentRatingKey       int    `json:"parentRatingKey"`
		Rating                string `json:"rating"`
		RatingKey             int    `json:"ratingKey"`
		Summary               string `json:"summary"`
		Thumb                 string `json:"thumb"`
		Title                 string `json:"title"`
		Type                  string `json:"type"`
		UpdatedAt             int    `json:"updatedAt"`
		ViewCount             int    `json:"viewCount"`
		Year                  int    `json:"year"`
	} `json:"_children"`
	ElementType              string `json:"_elementType"`
	AllowSync                string `json:"allowSync"`
	Art                      string `json:"art"`
	Banner                   string `json:"banner"`
	GrandparentContentRating string `json:"grandparentContentRating"`
	GrandparentRatingKey     string `json:"grandparentRatingKey"`
	GrandparentStudio        string `json:"grandparentStudio"`
	GrandparentTheme         string `json:"grandparentTheme"`
	GrandparentThumb         string `json:"grandparentThumb"`
	GrandparentTitle         string `json:"grandparentTitle"`
	Identifier               string `json:"identifier"`
	Key                      string `json:"key"`
	LibrarySectionID         string `json:"librarySectionID"`
	LibrarySectionTitle      string `json:"librarySectionTitle"`
	LibrarySectionUUID       string `json:"librarySectionUUID"`
	MediaTagPrefix           string `json:"mediaTagPrefix"`
	MediaTagVersion          string `json:"mediaTagVersion"`
	Nocache                  string `json:"nocache"`
	ParentIndex              string `json:"parentIndex"`
	ParentTitle              string `json:"parentTitle"`
	ParentYear               string `json:"parentYear"`
	SortAsc                  string `json:"sortAsc"`
	Theme                    string `json:"theme"`
	Thumb                    string `json:"thumb"`
	Title1                   string `json:"title1"`
	Title2                   string `json:"title2"`
	ViewGroup                string `json:"viewGroup"`
	ViewMode                 string `json:"viewMode"`
}

type plexResponse struct {
	Children []struct {
		ElementType string `json:"_elementType"`
		Count       string `json:"count"`
		Key         string `json:"key"`
		Title       string `json:"title"`
	} `json:"_children"`
	ElementType                   string `json:"_elementType"`
	AllowCameraUpload             string `json:"allowCameraUpload"`
	AllowChannelAccess            string `json:"allowChannelAccess"`
	AllowSync                     string `json:"allowSync"`
	BackgroundProcessing          string `json:"backgroundProcessing"`
	Certificate                   string `json:"certificate"`
	CompanionProxy                string `json:"companionProxy"`
	FriendlyName                  string `json:"friendlyName"`
	HubSearch                     string `json:"hubSearch"`
	MachineIdentifier             string `json:"machineIdentifier"`
	Multiuser                     string `json:"multiuser"`
	MyPlex                        string `json:"myPlex"`
	MyPlexMappingState            string `json:"myPlexMappingState"`
	MyPlexSigninState             string `json:"myPlexSigninState"`
	MyPlexSubscription            string `json:"myPlexSubscription"`
	MyPlexUsername                string `json:"myPlexUsername"`
	Platform                      string `json:"platform"`
	PlatformVersion               string `json:"platformVersion"`
	RequestParametersInCookie     string `json:"requestParametersInCookie"`
	Sync                          string `json:"sync"`
	TranscoderActiveVideoSessions string `json:"transcoderActiveVideoSessions"`
	TranscoderAudio               string `json:"transcoderAudio"`
	TranscoderLyrics              string `json:"transcoderLyrics"`
	TranscoderPhoto               string `json:"transcoderPhoto"`
	TranscoderSubtitles           string `json:"transcoderSubtitles"`
	TranscoderVideo               string `json:"transcoderVideo"`
	TranscoderVideoBitrates       string `json:"transcoderVideoBitrates"`
	TranscoderVideoQualities      string `json:"transcoderVideoQualities"`
	TranscoderVideoResolutions    string `json:"transcoderVideoResolutions"`
	UpdatedAt                     string `json:"updatedAt"`
	Version                       string `json:"version"`
}

type killTranscodeResponse struct {
	Children []struct {
		ElementType   string  `json:"_elementType"`
		AudioChannels int     `json:"audioChannels"`
		AudioCodec    string  `json:"audioCodec"`
		AudioDecision string  `json:"audioDecision"`
		Container     string  `json:"container"`
		Context       string  `json:"context"`
		Duration      int     `json:"duration"`
		Height        int     `json:"height"`
		Key           string  `json:"key"`
		Progress      float64 `json:"progress"`
		Protocol      string  `json:"protocol"`
		Remaining     int     `json:"remaining"`
		Speed         float64 `json:"speed"`
		Throttled     bool    `json:"throttled"`
		VideoCodec    string  `json:"videoCodec"`
		VideoDecision string  `json:"videoDecision"`
		Width         int     `json:"width"`
	} `json:"_children"`
	ElementType string `json:"_elementType"`
}

// CreateLibraryParams params required to create a library
type CreateLibraryParams struct {
	Name        string
	Location    string
	LibraryType string
	Agent       string
	Scanner     string
	Language    string
}

// DevicesResponse  metadata of a device that has connected to your server
type DevicesResponse struct {
	ID         int    `json:"id"`
	LastSeenAt string `json:"lastSeenAt"`
	Name       string `json:"name"`
	Product    string `json:"product"`
	Version    string `json:"version"`
}

// Friends are the plex accounts that have access to your server
type Friends struct {
	ID                        int    `xml:"id,attr"`
	Title                     string `xml:"title,attr"`
	Thumb                     string `xml:"thumb,attr"`
	Protected                 string `xml:"protected,attr"`
	Home                      string `xml:"home,attr"`
	AllowSync                 string `xml:"allowSync,attr"`
	AllowCameraUpload         string `xml:"allowCameraUpload,attr"`
	AllowChannels             string `xml:"allowChannels,attr"`
	FilterAll                 string `xml:"filterAll,attr"`
	FilterMovies              string `xml:"filterMovies,attr"`
	FilterMusic               string `xml:"filterMusic,attr"`
	FilterPhotos              string `xml:"filterPhotos,attr"`
	FilterTelevision          string `xml:"filterTelevision,attr"`
	Restricted                string `xml:"restricted,attr"`
	Username                  string `xml:"username,attr"`
	Email                     string `xml:"email,attr"`
	RecommendationsPlaylistID string `xml:"recommendationsPlaylistId,attr"`
	Server                    struct {
		ID                string `xml:"id,attr"`
		ServerID          string `xml:"serverId,attr"`
		MachineIdentifier string `xml:"machineIdentifier,attr"`
		Name              string `xml:"name,attr"`
		LastSeenAt        string `xml:"lastSeenAt,attr"`
		NumLibraries      string `xml:"numLibraries,attr"`
		AllLibraries      string `xml:"allLibraries,attr"`
		Owned             string `xml:"owned,attr"`
		Pending           string `xml:"pending,attr"`
	} `xml:"Server"`
}

type friendsResponse struct {
	XMLName           xml.Name  `xml:"MediaContainer"`
	FriendlyName      string    `xml:"friendlyName,attr"`
	Identifier        string    `xml:"identifier,attr"`
	MachineIdentifier string    `xml:"machineIdentifier,attr"`
	TotalSize         string    `xml:"totalSize,attr"`
	Size              int       `xml:"size,attr"`
	User              []Friends `xml:"User"`
}

type resultResponse struct {
	XMLName  xml.Name `xml:"Response"`
	Response struct {
		Code   int    `xml:"code,attr"`
		Status string `xml:"status,attr"`
	} `xml:"Response"`
}

type inviteFriendResponse struct {
	XMLName           xml.Name `xml:"MediaContainer"`
	FriendlyName      string   `xml:"friendlyName,attr"`
	Identifier        string   `xml:"identifier,attr"`
	MachineIdentifier string   `xml:"machineIdentifier,attr"`
	Size              string   `xml:"size,attr"`
	SharedServer      struct {
		ID                string `xml:"id,attr"`
		Username          string `xml:"username,attr"`
		Email             string `xml:"email,attr"`
		UserID            int    `xml:"userID,attr"`
		AccessToken       string `xml:"accessToken,attr"`
		Name              string `xml:"name,attr"`
		AcceptedAt        string `xml:"acceptedAt,attr"`
		InvitedAt         string `xml:"invitedAt,attr"`
		AllowSync         string `xml:"allowSync,attr"`
		AllowCameraUpload string `xml:"allowCameraUpload,attr"`
		AllowChannels     string `xml:"allowChannels,attr"`
		Owned             string `xml:"owned,attr"`
		Section           struct {
			ID     string `xml:"id,attr"`
			Key    string `xml:"key,attr"`
			Title  string `xml:"title,attr"`
			Type   string `xml:"type,attr"`
			Shared string `xml:"shared,attr"`
		} `xml:"Section"`
	} `xml:"SharedServer"`
}

// InviteFriendParams are the params to invite a friend
type InviteFriendParams struct {
	UsernameOrEmail string
	MachineID       string
	Label           string
	LibraryIDs      []int
}

// UpdateFriendParams optional parameters to update your friends access to your server
type UpdateFriendParams struct {
	AllowSync         string
	AllowCameraUpload string
	AllowChannels     string
	FilterMovies      string
	FilterTelevision  string
	FilterMusic       string
	FilterPhotos      string
}
type inviteFriendBody struct {
	ServerID        string                      `json:"server_id"`
	SharedServer    inviteFriendSharedServer    `json:"shared_server"`
	SharingSettings inviteFriendSharingSettings `json:"sharing_settings"`
}

type inviteFriendSharedServer struct {
	InvitedEmail      string `json:"invited_email"`
	LibrarySectionIDs []int  `json:"library_section_ids"`
}
type inviteFriendSharingSettings struct {
	FilterMovies     string `json:"filterMovies"`
	FilterTelevision string `json:"filterTelevision"`
}

type resourcesResponse struct {
	XMLName xml.Name     `xml:"MediaContainer"`
	Size    int          `xml:"size,attr"`
	Device  []PMSDevices `xml:"Device"`
}

type terminateSessionResponse struct {
	XMLName xml.Name `xml:"MediaContainer"`
	Size    int      `xml:"size,attr"`
}

// PMSDevices is the result of the https://plex.tv/pms/resources endpoint
type PMSDevices struct {
	Name                 string       `json:"name" xml:"name,attr"`
	Product              string       `json:"product" xml:"product,attr"`
	ProductVersion       string       `json:"productVersion" xml:"productVersion,attr"`
	Platform             string       `json:"platform" xml:"platform,attr"`
	PlatformVersion      string       `json:"platformVersion" xml:"platformVersion,attr"`
	Device               string       `json:"device" xml:"device,attr"`
	ClientIdentifier     string       `json:"clientIdentifier" xml:"clientIdentifier,attr"`
	CreatedAt            string       `json:"createdAt" xml:"createdAt,attr"`
	LastSeenAt           string       `json:"lastSeenAt" xml:"lastSeenAt,attr"`
	Provides             string       `json:"provides" xml:"provides,attr"`
	Owned                string       `json:"owned" xml:"owned,attr"`
	AccessToken          string       `json:"accessToken" xml:"accessToken,attr"`
	HTTPSRequired        int          `json:"httpsRequired" xml:"httpsRequired,attr"`
	Synced               string       `json:"synced" xml:"synced,attr"`
	Relay                int          `json:"relay" xml:"relay,attr"`
	PublicAddressMatches string       `json:"publicAddressMatches" xml:"publicAddressMatches,attr"`
	PublicAddress        string       `json:"publicAddress" xml:"publicAddress,attr"`
	Presence             string       `json:"presence" xml:"presence,attr"`
	Connection           []Connection `json:"connection" xml:"Connection"`
}

// Connection lists options to connect to a device
type Connection struct {
	Protocol string `json:"protocol" xml:"protocol,attr"`
	Address  string `json:"address" xml:"address,attr"`
	Port     string `json:"port" xml:"port,attr"`
	URI      string `json:"uri" xml:"uri,attr"`
	Local    int    `json:"local" xml:"local,attr"`
}

// BaseAPIResponse info about the Plex Media Server
type BaseAPIResponse struct {
	MediaContainer struct {
		Directory []struct {
			Count int64  `json:"count"`
			Key   string `json:"key"`
			Title string `json:"title"`
		} `json:"Directory"`
		AllowCameraUpload             bool   `json:"allowCameraUpload"`
		AllowChannelAccess            bool   `json:"allowChannelAccess"`
		AllowSharing                  bool   `json:"allowSharing"`
		AllowSync                     bool   `json:"allowSync"`
		BackgroundProcessing          bool   `json:"backgroundProcessing"`
		Certificate                   bool   `json:"certificate"`
		CompanionProxy                bool   `json:"companionProxy"`
		CountryCode                   string `json:"countryCode"`
		Diagnostics                   string `json:"diagnostics"`
		EventStream                   bool   `json:"eventStream"`
		FriendlyName                  string `json:"friendlyName"`
		HubSearch                     bool   `json:"hubSearch"`
		ItemClusters                  bool   `json:"itemClusters"`
		Livetv                        int64  `json:"livetv"`
		MachineIdentifier             string `json:"machineIdentifier"`
		MediaProviders                bool   `json:"mediaProviders"`
		Multiuser                     bool   `json:"multiuser"`
		MyPlex                        bool   `json:"myPlex"`
		MyPlexMappingState            string `json:"myPlexMappingState"`
		MyPlexSigninState             string `json:"myPlexSigninState"`
		MyPlexSubscription            bool   `json:"myPlexSubscription"`
		MyPlexUsername                string `json:"myPlexUsername"`
		OwnerFeatures                 string `json:"ownerFeatures"`
		PhotoAutoTag                  bool   `json:"photoAutoTag"`
		Platform                      string `json:"platform"`
		PlatformVersion               string `json:"platformVersion"`
		PluginHost                    bool   `json:"pluginHost"`
		ReadOnlyLibraries             bool   `json:"readOnlyLibraries"`
		RequestParametersInCookie     bool   `json:"requestParametersInCookie"`
		Size                          int64  `json:"size"`
		StreamingBrainABRVersion      int64  `json:"streamingBrainABRVersion"`
		StreamingBrainVersion         int64  `json:"streamingBrainVersion"`
		Sync                          bool   `json:"sync"`
		TranscoderActiveVideoSessions int64  `json:"transcoderActiveVideoSessions"`
		TranscoderAudio               bool   `json:"transcoderAudio"`
		TranscoderLyrics              bool   `json:"transcoderLyrics"`
		TranscoderPhoto               bool   `json:"transcoderPhoto"`
		TranscoderSubtitles           bool   `json:"transcoderSubtitles"`
		TranscoderVideo               bool   `json:"transcoderVideo"`
		TranscoderVideoBitrates       string `json:"transcoderVideoBitrates"`
		TranscoderVideoQualities      string `json:"transcoderVideoQualities"`
		TranscoderVideoResolutions    string `json:"transcoderVideoResolutions"`
		UpdatedAt                     int64  `json:"updatedAt"`
		Updater                       bool   `json:"updater"`
		Version                       string `json:"version"`
		VoiceSearch                   bool   `json:"voiceSearch"`
	} `json:"MediaContainer"`
}

// User ...
type User struct {
	ID           string `json:"id"`
	UUID         string `json:"uuid"`
	Email        string `json:"email"`
	JoinedAt     string `json:"joined_at"`
	Username     string `json:"username"`
	Thumb        string `json:"thumb"`
	AuthToken    string `json:"authToken"`
	Subscription struct {
		Active   bool     `json:"active"`
		Status   string   `json:"Active"`
		Plan     string   `json:"lifetime"`
		Features []string `json:"features"`
	} `json:"subscription"`
	Roles struct {
		Roles []string `json:"roles"`
	} `json:"roles"`
	Entitlements []string `json:"entitlements"`
	ConfirmedAt  string   `json:"confirmedAt"`
	ForumID      string   `json:"forumId"`
	RememberMe   bool     `json:"rememberMe"`
	Title        string   `json:"title"`
}

// SignInResponse ...
type SignInResponse struct {
	User User `json:"user"`
}

// ServerInfo is the result of the https://plex.tv/api/servers endpoint
type ServerInfo struct {
	XMLName           xml.Name `xml:"MediaContainer"`
	FriendlyName      string   `xml:"friendlyName,attr"`
	Identifier        string   `xml:"identifier,attr"`
	MachineIdentifier string   `xml:"machineIdentifier,attr"`
	Size              int      `xml:"size,attr"`
	Server            []struct {
		AccessToken       string `xml:"accessToken,attr"`
		Name              string `xml:"name,attr"`
		Address           string `xml:"address,attr"`
		Port              string `xml:"port,attr"`
		Version           string `xml:"version,attr"`
		Scheme            string `xml:"scheme,attr"`
		Host              string `xml:"host,attr"`
		LocalAddresses    string `xml:"localAddresses,attr"`
		MachineIdentifier string `xml:"machineIdentifier,attr"`
		CreatedAt         string `xml:"createdAt,attr"`
		UpdatedAt         string `xml:"updatedAt,attr"`
		Owned             string `xml:"owned,attr"`
		Synced            string `xml:"synced,attr"`
	} `xml:"Server"`
}

// SectionIDResponse the section id (or library id) of your server
// useful when inviting a user to the server
type SectionIDResponse struct {
	XMLName           xml.Name `xml:"MediaContainer"`
	FriendlyName      string   `xml:"friendlyName,attr"`
	Identifier        string   `xml:"identifier,attr"`
	MachineIdentifier string   `xml:"machineIdentifier,attr"`
	Size              int      `xml:"size,attr"`
	Server            []struct {
		Name              string           `xml:"name,attr"`
		Address           string           `xml:"address,attr"`
		Port              string           `xml:"port,attr"`
		Version           string           `xml:"version,attr"`
		Scheme            string           `xml:"scheme,attr"`
		Host              string           `xml:"host,attr"`
		LocalAddresses    string           `xml:"localAddresses,attr"`
		MachineIdentifier string           `xml:"machineIdentifier,attr"`
		CreatedAt         int              `xml:"createdAt,attr"`
		UpdatedAt         int              `xml:"updatedAt,attr"`
		Owned             int              `xml:"owned,attr"`
		Synced            string           `xml:"synced,attr"`
		Section           []ServerSections `xml:"Section"`
	} `xml:"Server"`
}

// ServerSections contains information of your library sections
type ServerSections struct {
	ID    int    `xml:"id,attr"`
	Key   string `xml:"key,attr"`
	Type  string `xml:"type,attr"`
	Title string `xml:"title,attr"`
}

// LibraryLabels are the existing labels set on your server
type LibraryLabels struct {
	ElementType     string `json:"_elementType"`
	AllowSync       string `json:"allowSync"`
	Art             string `json:"art"`
	Content         string `json:"content"`
	Identifier      string `json:"identifier"`
	MediaTagPrefix  string `json:"mediaTagPrefix"`
	MediaTagVersion string `json:"mediaTagVersion"`
	Thumb           string `json:"thumb"`
	Title1          string `json:"title1"`
	Title2          string `json:"title2"`
	ViewGroup       string `json:"viewGroup"`
	ViewMode        string `json:"viewMode"`
	Children        []struct {
		ElementType string `json:"_elementType"`
		FastKey     string `json:"fastKey"`
		Key         string `json:"key"`
		Title       string `json:"title"`
	} `json:"_children"`
}

type headers struct {
	Platform         string
	PlatformVersion  string
	Provides         string
	ClientIdentifier string
	Product          string
	Version          string
	Device           string
	ContainerSize    string
	ContainerStart   string
	Token            string
	Accept           string
	ContentType      string
}

type request struct {
	headers
}

// Sessions

// TranscodeSessionsResponse is the result for transcode session endpoint /transcode/sessions
type TranscodeSessionsResponse struct {
	Children []struct {
		ElementType   string  `json:"_elementType"`
		AudioChannels int     `json:"audioChannels"`
		AudioCodec    string  `json:"audioCodec"`
		AudioDecision string  `json:"audioDecision"`
		Container     string  `json:"container"`
		Context       string  `json:"context"`
		Duration      int     `json:"duration"`
		Height        int     `json:"height"`
		Key           string  `json:"key"`
		Progress      float64 `json:"progress"`
		Protocol      string  `json:"protocol"`
		Remaining     int     `json:"remaining"`
		Speed         float64 `json:"speed"`
		Throttled     bool    `json:"throttled"`
		VideoCodec    string  `json:"videoCodec"`
		VideoDecision string  `json:"videoDecision"`
		Width         int     `json:"width"`
	} `json:"_children"`
	ElementType string `json:"_elementType"`
}

// CurrentSessionsVideo are current video sessions
type CurrentSessionsVideo struct {
	AddedAt               string `xml:"addedAt,attr"`
	Art                   string `xml:"art,attr"`
	ChapterSource         string `xml:"chapterSource,attr"`
	ContentRating         string `xml:"contentRating,attr"`
	Duration              int    `xml:"duration,attr"`
	GUID                  string `xml:"guid,attr"`
	Key                   string `xml:"key,attr"`
	LibrarySectionID      string `xml:"librarySectionID,attr"`
	OriginallyAvailableAt string `xml:"originallyAvailableAt,attr"`
	PrimaryExtraKey       string `xml:"primaryExtraKey,attr"`
	Rating                string `xml:"rating,attr"`
	RatingKey             string `xml:"ratingKey,attr"`
	SessionKey            string `xml:"sessionKey,attr"`
	Studio                string `xml:"studio,attr"`
	Summary               string `xml:"summary,attr"`
	Tagline               string `xml:"tagline,attr"`
	Thumb                 string `xml:"thumb,attr"`
	Title                 string `xml:"title,attr"`
	TitleSort             string `xml:"titleSort,attr"`
	Type                  string `xml:"type,attr"`
	UpdatedAt             string `xml:"updatedAt,attr"`
	ViewOffset            int    `xml:"viewOffset,attr"`
	Year                  string `xml:"year,attr"`
	Media                 struct {
		AspectRatio           string `xml:"aspectRatio,attr"`
		AudioChannels         string `xml:"audioChannels,attr"`
		AudioCodec            string `xml:"audioCodec,attr"`
		AudioProfile          string `xml:"audioProfile,attr"`
		Bitrate               string `xml:"bitrate,attr"`
		Container             string `xml:"container,attr"`
		Duration              string `xml:"duration,attr"`
		Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
		Height                string `xml:"height,attr"`
		ID                    string `xml:"id,attr"`
		OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
		VideoCodec            string `xml:"videoCodec,attr"`
		VideoFrameRate        string `xml:"videoFrameRate,attr"`
		VideoProfile          string `xml:"videoProfile,attr"`
		VideoResolution       string `xml:"videoResolution,attr"`
		Width                 string `xml:"width,attr"`
		Part                  struct {
			AudioProfile          string `xml:"audioProfile,attr"`
			Container             string `xml:"container,attr"`
			Duration              string `xml:"duration,attr"`
			File                  string `xml:"file,attr"`
			Has64bitOffsets       string `xml:"has64bitOffsets,attr"`
			ID                    string `xml:"id,attr"`
			Indexes               string `xml:"indexes,attr"`
			Key                   string `xml:"key,attr"`
			OptimizedForStreaming string `xml:"optimizedForStreaming,attr"`
			Size                  string `xml:"size,attr"`
			VideoProfile          string `xml:"videoProfile,attr"`
			Stream                []struct {
				BitDepth           string `xml:"bitDepth,attr"`
				Bitrate            string `xml:"bitrate,attr"`
				Cabac              string `xml:"cabac,attr"`
				ChromaSubsampling  string `xml:"chromaSubsampling,attr"`
				Codec              string `xml:"codec,attr"`
				CodecID            string `xml:"codecID,attr"`
				ColorRange         string `xml:"colorRange,attr"`
				ColorSpace         string `xml:"colorSpace,attr"`
				Default            string `xml:"default,attr"`
				Duration           string `xml:"duration,attr"`
				FrameRate          string `xml:"frameRate,attr"`
				FrameRateMode      string `xml:"frameRateMode,attr"`
				HasScalingMatrix   string `xml:"hasScalingMatrix,attr"`
				Height             string `xml:"height,attr"`
				ID                 string `xml:"id,attr"`
				Index              string `xml:"index,attr"`
				Level              string `xml:"level,attr"`
				PixelFormat        string `xml:"pixelFormat,attr"`
				Profile            string `xml:"profile,attr"`
				RefFrames          string `xml:"refFrames,attr"`
				ScanType           string `xml:"scanType,attr"`
				StreamIdentifier   string `xml:"streamIdentifier,attr"`
				StreamType         string `xml:"streamType,attr"`
				Width              string `xml:"width,attr"`
				AudioChannelLayout string `xml:"audioChannelLayout,attr"`
				BitrateMode        string `xml:"bitrateMode,attr"`
				Channels           string `xml:"channels,attr"`
				Language           string `xml:"language,attr"`
				LanguageCode       string `xml:"languageCode,attr"`
				SamplingRate       string `xml:"samplingRate,attr"`
				Selected           string `xml:"selected,attr"`
				Format             string `xml:"format,attr"`
				Key                string `xml:"key,attr"`
			} `xml:"Stream"`
		} `xml:"Part"`
	} `xml:"Media"`
	Genre []struct {
		Count string `xml:"count,attr"`
		ID    string `xml:"id,attr"`
		Tag   string `xml:"tag,attr"`
	} `xml:"Genre"`
	Writer []struct {
		ID    string `xml:"id,attr"`
		Tag   string `xml:"tag,attr"`
		Count string `xml:"count,attr"`
	} `xml:"Writer"`
	Director struct {
		Count string `xml:"count,attr"`
		ID    string `xml:"id,attr"`
		Tag   string `xml:"tag,attr"`
	} `xml:"Director"`
	Producer []struct {
		Count string `xml:"count,attr"`
		ID    string `xml:"id,attr"`
		Tag   string `xml:"tag,attr"`
	} `xml:"Producer"`
	Country struct {
		Count string `xml:"count,attr"`
		ID    string `xml:"id,attr"`
		Tag   string `xml:"tag,attr"`
	} `xml:"Country"`
	Role []struct {
		Count string `xml:"count,attr"`
		ID    string `xml:"id,attr"`
		Role  string `xml:"role,attr"`
		Tag   string `xml:"tag,attr"`
	} `xml:"Role"`
	Collection struct {
		Count string `xml:"count,attr"`
		ID    string `xml:"id,attr"`
		Tag   string `xml:"tag,attr"`
	} `xml:"Collection"`
	Label struct {
		ID  string `xml:"id,attr"`
		Tag string `xml:"tag,attr"`
	} `xml:"Label"`
	Field struct {
		Locked string `xml:"locked,attr"`
		Name   string `xml:"name,attr"`
	} `xml:"Field"`
	User struct {
		ID    int    `xml:"id,attr"`
		Title string `xml:"title,attr"`
		Thumb string `xml:"thumb,attr"`
	} `xml:"User"`
	Player struct {
		Address           string `xml:"address,attr"`
		Device            string `xml:"device,attr"`
		MachineIdentifier string `xml:"machineIdentifier,attr"`
		Model             string `xml:"model,attr"`
		Platform          string `xml:"platform,attr"`
		PlatformVersion   string `xml:"platformVersion,attr"`
		Product           string `xml:"product,attr"`
		Profile           string `xml:"profile,attr"`
		State             string `xml:"state,attr"`
		Title             string `xml:"title,attr"`
		Vendor            string `xml:"vendor,attr"`
		Version           string `xml:"version,attr"`
	} `xml:"Player"`
	GrandparentArt       string `xml:"grandparentArt,attr"`
	GrandparentKey       string `xml:"grandparentKey,attr"`
	GrandparentRatingKey string `xml:"grandparentRatingKey,attr"`
	GrandparentTheme     string `xml:"grandparentTheme,attr"`
	GrandparentThumb     string `xml:"grandparentThumb,attr"`
	GrandparentTitle     string `xml:"grandparentTitle,attr"`
	Index                string `xml:"index,attr"`
	LastViewedAt         string `xml:"lastViewedAt,attr"`
	ParentIndex          string `xml:"parentIndex,attr"`
	ParentKey            string `xml:"parentKey,attr"`
	ParentRatingKey      string `xml:"parentRatingKey,attr"`
	ParentThumb          string `xml:"parentThumb,attr"`
	ViewCount            string `xml:"viewCount,attr"`
	Session              struct {
		ID        string `xml:"id,attr"`
		Bandwidth string `xml:"bandwidth,attr"`
		Location  string `xml:"location,attr"`
	}
	TranscodeSession struct {
		Key           string `xml:"key,attr"`
		Throttled     string `xml:"throttled,attr"`
		Progress      string `xml:"progress,attr"`
		Speed         string `xml:"speed,attr"`
		Duration      string `xml:"duration,attr"`
		Remaining     string `xml:"remaining,attr"`
		Context       string `xml:"context,attr"`
		VideoDecision string `xml:"videoDecision,attr"`
		AudioDecision string `xml:"audioDecision,attr"`
		Protocol      string `xml:"protocol,attr"`
		Container     string `xml:"container,attr"`
		VideoCodec    string `xml:"videoCodec,attr"`
		AudioCodec    string `xml:"audioCodec,attr"`
		AudioChannels string `xml:"audioChannels,attr"`
		Width         string `xml:"width,attr"`
		Height        string `xml:"height,attr"`
	} `xml:"TranscodeSession"`
}

// Stream ...
type Stream struct {
	Anamorphic         string `json:"anamorphic"`
	AudioChannelLayout string `json:"audioChannelLayout"`
	BitDepth           string `json:"bitDepth"`
	Bitrate            string `json:"bitrate"`
	BitrateMode        string `json:"bitrateMode"`
	Channels           string `json:"channels"`
	ChromaSubsampling  string `json:"chromaSubsampling"`
	Codec              string `json:"codec"`
	Default            string `json:"default"`
	Duration           string `json:"duration"`
	FrameRate          string `json:"frameRate"`
	ID                 string `json:"id"`
	HasScalingMatrix   string `json:"hasScalingMatrix"`
	Height             string `json:"height"`
	Index              string `json:"index"`
	Language           string `json:"language"`
	LanguageCode       string `json:"languageCode"`
	Level              string `json:"level"`
	Location           string `json:"location"`
	PixelAspectRatio   string `json:"pixelAspectRatio"`
	Profile            string `json:"profile"`
	RefFrames          string `json:"refFrames"`
	SamplingRate       string `json:"samplingRate"`
	ScanType           string `json:"scanType"`
	Selected           string `json:"selected"`
	StreamType         string `json:"streamType"`
	Width              string `json:"width"`
}

// Part ...
type Part struct {
	Stream                []Stream `json:"Stream"`
	Container             string   `json:"container"`
	Decision              string   `json:"decision"`
	Duration              int      `json:"duration"`
	File                  string   `json:"file"`
	HasThumbnail          string   `json:"hasThumbnail"`
	ID                    int      `json:"id"`
	Key                   string   `json:"key"`
	Selected              bool     `json:"selected"`
	Size                  int      `json:"size"`
	VideoProfile          string   `json:"videoProfile"`
	AudioProfile          string   `json:"audioProfile"`
	Has64bitOffsets       bool     `json:"has64bitOffsets"`
	OptimizedForStreaming bool     `json:"optimizedForStreaming"`
}

// Player ...
type Player struct {
	Address             string `json:"address"`
	Device              string `json:"device"`
	Local               bool   `json:"local"`
	MachineIdentifier   string `json:"machineIdentifier"`
	Model               string `json:"model"`
	Platform            string `json:"platform"`
	PlatformVersion     string `json:"platformVersion"`
	Product             string `json:"product"`
	Profile             string `json:"profile"`
	RemotePublicAddress string `json:"remotePublicAddress"`
	State               string `json:"state"`
	Title               string `json:"title"`
	UserID              int    `json:"userID"`
	Vendor              string `json:"vendor"`
	Version             string `json:"version"`
}

// Session ...
type Session struct {
	Bandwidth int    `json:"bandwidth"`
	ID        string `json:"id"`
	Location  string `json:"location"`
}

// CurrentSessions metadata of users consuming media
type CurrentSessions struct {
	MediaContainer struct {
		Track []struct {
			Media []Media `json:"Media"`
			Mood  []struct {
				Filter string `json:"filter"`
				ID     string `json:"id"`
				Tag    string `json:"tag"`
			} `json:"Mood"`
			Player               Player  `json:"Player"`
			Session              Session `json:"Session"`
			User                 User    `json:"User"`
			AddedAt              string  `json:"addedAt"`
			Art                  string  `json:"art"`
			Duration             string  `json:"duration"`
			GrandparentArt       string  `json:"grandparentArt"`
			GrandparentKey       string  `json:"grandparentKey"`
			GrandparentRatingKey string  `json:"grandparentRatingKey"`
			GrandparentThumb     string  `json:"grandparentThumb"`
			GrandparentTitle     string  `json:"grandparentTitle"`
			GUID                 string  `json:"guid"`
			Index                string  `json:"index"`
			Key                  string  `json:"key"`
			LastViewedAt         string  `json:"lastViewedAt"`
			LibrarySectionID     string  `json:"librarySectionID"`
			LibrarySectionKey    string  `json:"librarySectionKey"`
			ParentIndex          string  `json:"parentIndex"`
			ParentKey            string  `json:"parentKey"`
			ParentRatingKey      string  `json:"parentRatingKey"`
			ParentThumb          string  `json:"parentThumb"`
			ParentTitle          string  `json:"parentTitle"`
			RatingKey            string  `json:"ratingKey"`
			SessionKey           string  `json:"sessionKey"`
			Summary              string  `json:"summary"`
			Thumb                string  `json:"thumb"`
			Title                string  `json:"title"`
			Type                 string  `json:"type"`
			UpdatedAt            string  `json:"updatedAt"`
			ViewCount            string  `json:"viewCount"`
			ViewOffset           string  `json:"viewOffset"`
			Year                 string  `json:"year"`
		} `json:"Track"`
		Video []struct {
			Media                 []Media `json:"Media"`
			Player                Player  `json:"Player"`
			Session               Session `json:"Session"`
			User                  User    `json:"User"`
			AddedAt               string  `json:"addedAt"`
			Art                   string  `json:"art"`
			ContentRating         string  `json:"contentRating"`
			Duration              string  `json:"duration"`
			GrandparentArt        string  `json:"grandparentArt"`
			GrandparentKey        string  `json:"grandparentKey"`
			GrandparentRatingKey  string  `json:"grandparentRatingKey"`
			GrandparentTheme      string  `json:"grandparentTheme"`
			GrandparentThumb      string  `json:"grandparentThumb"`
			GrandparentTitle      string  `json:"grandparentTitle"`
			GUID                  string  `json:"guid"`
			Index                 string  `json:"index"`
			Key                   string  `json:"key"`
			LibrarySectionID      string  `json:"librarySectionID"`
			LibrarySectionKey     string  `json:"librarySectionKey"`
			OriginallyAvailableAt string  `json:"originallyAvailableAt"`
			ParentIndex           string  `json:"parentIndex"`
			ParentKey             string  `json:"parentKey"`
			ParentRatingKey       string  `json:"parentRatingKey"`
			ParentThumb           string  `json:"parentThumb"`
			ParentTitle           string  `json:"parentTitle"`
			RatingKey             string  `json:"ratingKey"`
			SessionKey            string  `json:"sessionKey"`
			Summary               string  `json:"summary"`
			Thumb                 string  `json:"thumb"`
			Title                 string  `json:"title"`
			Type                  string  `json:"type"`
			UpdatedAt             string  `json:"updatedAt"`
			ViewOffset            string  `json:"viewOffset"`
			Year                  string  `json:"year"`
		} `json:"Video"`
		Size int `json:"size"`
	} `json:"MediaContainer"`
}
