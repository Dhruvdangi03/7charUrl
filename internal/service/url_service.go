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
	short := util.GenerateShortCode(7) // 7-char short URL
	u.store.Save(short, original)
	return short
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
