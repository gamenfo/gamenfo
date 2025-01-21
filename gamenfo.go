package gamenfo

import (
	"encoding/xml"
	"fmt"
	"time"
)

type Game struct {
	XMLName     xml.Name     `xml:"game" json:"-"`
	Title       []LangText   `xml:"title" json:"title"`
	Release     []Platform   `xml:"release" json:"release"`
	Genre       []string     `xml:"genre" json:"genre"`
	Developers  []string     `xml:"developer" json:"developer"`
	Publishers  []string     `xml:"publisher" json:"publisher"`
	Plot        []LangText   `xml:"plot" json:"plot"`
	Rating      float64      `xml:"rating" json:"rating"`
	AgeRating   []LangText   `xml:"agerating" json:"agerating"`
	Icon        string       `xml:"icon" json:"icon"`
	Thumbnail   string       `xml:"thumb" json:"thumb"`
	CreatorInfo *CreatorInfo `xml:"creatorinfo" json:"creatorinfo"`
	DLCs        []DLC        `xml:"dlcs" json:"dlcs"`

	CreatedAt time.Time `xml:"created" json:"created"`
	UpdatedAt time.Time `xml:"updated" json:"updated"`
}

type Platform struct {
	Platform string        `xml:"platform" json:"platform"`
	Releases []CountryDate `xml:"releases" json:"releases"`
}

type CountryDate struct {
	Country string    `xml:"country,attr" json:"country"`
	Date    time.Time `xml:",chardata" json:"date"`
}

type LangText struct {
	Lang  string `xml:"lang,attr,omitempty" json:"lang,omitempty"`
	Value string `xml:",chardata" json:"value"`
}

type CreatorInfo struct {
	IsStreamingAllowed     bool   `xml:"stream,omitempty" json:"stream,omitempty"`
	IsVideoOnDemandAllowed bool   `xml:"vod,omitempty" json:"vod,omitempty"`
	MusicUsageAllowed      bool   `xml:"musicusage,omitempty" json:"musicusage,omitempty"`
	MonetisationAllowed    bool   `xml:"monetisation,omitempty" json:"monetisation,omitempty"`
	AdditionalInfo         string `xml:"additionalinfo,omitempty" json:"additionalinfo,omitempty"`
}

type DLC struct {
	Title       []LangText  `xml:"title" json:"title"`
	ReleaseDate CountryDate `xml:"releasedate" json:"releasedate"`
	Developers  []string    `xml:"developer" json:"developer"`
	Publishers  []string    `xml:"publisher" json:"publisher"`
	Plot        []LangText  `xml:"plot" json:"plot"`
	AgeRating   []LangText  `xml:"agerating" json:"agerating"`
}

func (g *Game) ToNFO() (string, error) {
	xmlData, err := xml.MarshalIndent(g, "", "\t")
	if err != nil {
		return "", err
	}

	nfo := fmt.Sprintf("<?xml version=\"1.0\" encoding=\"UTF-8\" standalone=\"yes\"?>\n%s", xmlData)
	return nfo, nil
}
