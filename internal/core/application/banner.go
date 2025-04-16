package application

import (
	"context"
	"errors"
)

type Banner struct {
	ID        int64
	Title     string
	ImgURL    string
	TargetURL string
	Text      string
}

func (a *Application) CreateBanner(ctx context.Context, title, text, imgUrl, targetUrl string) error {
	if title == "" || text == "" || imgUrl == "" || targetUrl == "" {
		return errors.New("invalid parameters")
	}

	a.mux.Lock()
	defer a.mux.Unlock()
	var newID int64 = 1
	if n := len(a.banners); n > 0 {
		newID = a.banners[n-1].ID + 1
	}
	banner := Banner{
		ID:        int64(newID),
		Title:     title,
		Text:      text,
		ImgURL:    imgUrl,
		TargetURL: targetUrl,
	}

	a.banners = append(a.banners, banner)

	return nil
}

func (a *Application) GetBannerByID(ctx context.Context, id int64) (Banner, error) {
	a.mux.Lock()
	defer a.mux.Unlock()

	for _, banner := range a.banners {
		if banner.ID == id {
			return banner, nil
		}
	}

	return Banner{}, ErrNotFound
}

func (a *Application) UpdateBanner(ctx context.Context, id int64, title, text, imgUrl, targetUrl string) error {
	a.mux.Lock()
	defer a.mux.Unlock()

	for i, banner := range a.banners {
		if banner.ID == id {
			a.banners[i] = Banner{
				ID:        id,
				Title:     title,
				Text:      text,
				ImgURL:    imgUrl,
				TargetURL: targetUrl,
			}
			return nil
		}
	}

	return ErrNotFound
}

func (a *Application) DeleteBanner(ctx context.Context, id int64) error {
	a.mux.Lock()
	defer a.mux.Unlock()

	for i, banner := range a.banners {
		if banner.ID == id {
			a.banners = append(a.banners[:i], a.banners[i+1:]...)
			return nil
		}
	}

	return ErrNotFound
}

func (a *Application) ListBanners(ctx context.Context) []Banner {
	a.mux.Lock()
	defer a.mux.Unlock()

	// возвращает копию
	result := make([]Banner, len(a.banners))
	copy(result, a.banners)
	return result
}
