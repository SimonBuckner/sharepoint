package sharepoint

import (
	"log/slog"

	"github.com/koltyakov/gosip"
	"github.com/koltyakov/gosip/api"
)

type Site struct {
	SP   *api.SP
	Api  *api.Site
	Info *api.SiteInfo
}

func connectToSite(spClient *gosip.SPClient) (*Site, error) {
	sp := api.NewSP(spClient)
	out := Site{
		SP:  sp,
		Api: sp.Site(),
	}

	title, err := out.GetTitle()
	if err != nil {
		return &out, err
	}
	slog.Info("site connection successful", "title", title)

	info, err := out.Api.Get()
	if err != nil {
		return &out, err
	}
	out.Info = info.Data()
	slog.Info("site information obtained", "title", title)

	return &out, nil
}

func (site *Site) GetTitle() (string, error) {
	titleRes, err := site.SP.Web().Select("Title").Get()
	if err != nil {
		return "", err
	}
	return titleRes.Data().Title, nil
}

func (site *Site) ConnectToList(listName string) (*List, error) {
	return connectToList(site.SP, listName)
}
