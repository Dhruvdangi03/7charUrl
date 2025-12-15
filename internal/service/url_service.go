package service

import (
	"7charUrl/internal/store"
	"7charUrl/internal/util"
)

type URLService struct {
	store *store.MemoryStore
}

func NewURLService(s *store.MemoryStore) *URLService {
	return &URLService{store: s}
}

func (u *URLService) Shorten(original string) string {
	shortCode := ""
	for true {
		short, err := util.GenerateShortCode()
		if err != nil {
			return "There was an error while Generating the ShortCode" + err.Error()
		}
		_, ok := u.store.Get(short)
		if !ok {
			shortCode = short
			break
		}
	}

	u.store.Save(shortCode, original)
	return shortCode
}

func (u *URLService) Custom(original string, custom string) (string, bool) {
	_, ok := u.store.Get(custom)
	if ok {
		return "Already Exist, Choose Another Url", false
	}
	u.store.Save(custom, original)
	return custom, true
}

func (u *URLService) Resolve(short string) (string, bool) {
	return u.store.Get(short)
}
