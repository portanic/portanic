package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/portanic/portanic/internal/components/integrations"
	"github.com/portanic/portanic/internal/plugins"
)

type IntegrationsHandler struct{}

func (h IntegrationsHandler) GetPlugins(c echo.Context) error {
	return c.JSON(http.StatusOK, plugins.Plugins)
}

func (h IntegrationsHandler) DeployPlugin(c echo.Context) error {
	var plugin plugins.Plugin
	if err := c.Bind(&plugin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Implement your Kubernetes deployment logic here
	log.Printf("Deploying plugin: %s\n", plugin.Name)

	return c.NoContent(http.StatusOK)
}

func (h IntegrationsHandler) HandleShowIntegrations(c echo.Context) error {
	pluginList := []Integrations.Plugin{}

	for _, plugin := range plugins.Plugins {
		tempPlugin := Integrations.Plugin{
			Name:        plugin.Name,
			Description: plugin.Description,
			Icon:        plugin.Icon,
		}
		pluginList = append(pluginList, tempPlugin)
	}
	return render(c, Integrations.Show(pluginList))
}
