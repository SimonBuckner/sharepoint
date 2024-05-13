package sharepoint

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/koltyakov/gosip/api"
)

type List struct {
	SP   *api.SP
	Api  *api.List
	Info *api.ListInfo
}

type SharePointField struct {
	Id                 string
	EntityPropertyName string
	InternalName       string
	StaticName         string
	Title              string
	// Group              string
	// LookupWebId        string
	// PrimaryFieldId     string
	// ReadOnlyField      bool
	// Required           bool
	// Scope              string
}

func connectToList(sp *api.SP, listName string) (*List, error) {

	listUri := fmt.Sprintf("lists/%s", listName)
	listApi := sp.Web().GetList(listUri)

	out := List{
		SP:  sp,
		Api: listApi,
	}
	info, err := out.Api.Get()
	if err != nil {
		return &out, err
	}
	out.Info = info.Data()
	slog.Info("list connection successful", "title", out.Info.Title)
	return &out, nil
}

func (list *List) GetTitle() string {
	return list.Info.Title
}

func (list *List) Get(columnNames string, out any) error {
	resp, err := list.Api.Items().Select(columnNames).Get()
	if err != nil {
		return err
	}
	return json.Unmarshal(resp.Normalized(), out)
}

func (list *List) GetFields() ([]SharePointField, error) {

	fieldsRes, err := list.Api.Fields().Get()
	if err != nil {
		return nil, err
	}
	slog.Info("got fields from", "list", list.GetTitle())

	fields := make([]SharePointField, 0)
	err = json.Unmarshal(fieldsRes.Normalized(), &fields)

	return fields, err
}
