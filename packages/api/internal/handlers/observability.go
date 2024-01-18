package handlers

import (
	"github.com/e2b-dev/infra/packages/api/internal/api"
	"github.com/e2b-dev/infra/packages/api/internal/constants"
	"github.com/e2b-dev/infra/packages/shared/pkg/models"
	"github.com/e2b-dev/infra/packages/shared/pkg/telemetry"
	"github.com/gin-gonic/gin"
)

func (a *APIStore) GetObservabilityInstances(c *gin.Context) {
	ctx := c.Request.Context()

	team := c.Value(constants.TeamContextKey).(models.Team)

	telemetry.ReportEvent(ctx, "get observability instances")

	instanceInfo := a.instanceCache.GetInstances(&team.ID)

	IdentifyAnalyticsTeam(a.posthog, team.ID.String(), team.Name)
	properties := a.GetPackageToPosthogProperties(&c.Request.Header)
	CreateAnalyticsTeamEvent(a.posthog, team.ID.String(), "got instances observability", properties)

	var instances []api.RunningInstance

	for _, info := range instanceInfo {
		if *info.TeamID != team.ID {
			continue
		}

		instances = append(instances, api.RunningInstance{
			ClientID:   info.Instance.ClientID,
			EnvID:      info.Instance.EnvID,
			InstanceID: info.Instance.InstanceID,
			StartedAt:  *info.StartTime,
		})
	}

	c.JSON(200, instances)
}
