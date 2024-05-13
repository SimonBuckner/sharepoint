package sharepoint

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/koltyakov/gosip"
	strategy "github.com/koltyakov/gosip/auth/azurecert"
)

type SharePoint struct {
	TenantId       string `json:"tenant_id"`
	ClientId       string `json:"client_id"`
	CertPath       string `json:"cert_path"`
	CertPassphrase string `json:"cert_passphrase"`
}

func ConnectSharePoint(configPath string) (*SharePoint, error) {
	out := SharePoint{}

	content, err := os.ReadFile(configPath)
	if err != nil {
		return &out, err
	}

	if err := json.Unmarshal(content, &out); err != nil {
		return &out, err
	}
	slog.Info("config file loaded successfully")

	return &out, nil
}

func (site *SharePoint) ConnectToSite(siteUrl string) (*Site, error) {

	authCfg := &strategy.AuthCnfg{
		SiteURL:  siteUrl,
		TenantID: site.TenantId,
		ClientID: site.ClientId,
		CertPath: site.CertPath,
		CertPass: site.CertPassphrase,
	}

	client := &gosip.SPClient{AuthCnfg: authCfg}
	return connectToSite(client)
}
