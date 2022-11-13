package service

import "errors"

var (
	ErrorBanner    = errors.New("something is wrong with the banner")
	ErrorCharacter = errors.New("undefined characters")
	ErrorNoBanner  = errors.New("can't get the banner")
)
