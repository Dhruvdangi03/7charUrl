package service

import (
	"7charUrl/internal/store"
	"7charUrl/internal/util"
)

type URLService struct {
}

func NewURLService() *URLService {
	return &URLService{}
}

func (u *URLService) Shorten(original string) string {
	short, err := util.GenerateShortCode()
	if err != nil {
		return "There was an error while Generating the ShortCode" + err.Error()
	}

	store.CreateURL(short, original)
	return short
}

func (u *URLService) Custom(original string, custom string) (string, bool) {
	url, err := store.GetURL(custom)
	if err != nil {
		return "There is some error" + err.Error(), false
	}
	if url != nil {
		return "Already Exist, Choose Another Url", false
	}
	store.CreateURL(custom, original)
	return custom, true
}

func (u *URLService) Resolve(short string) (string, bool) {
	url, err := store.GetURL(short)
	if err != nil {
		return "There is some error" + err.Error(), false
	}
	return url.LongURL, true
}
