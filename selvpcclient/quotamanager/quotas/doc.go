/*
Package quotas provides the ability to retrieve and update quotas through the
Quota Manager API.

Example of getting quota limits for a single project

	limits, _, err := quotas.GetLimits(ctx, QuotaRegionalClient)
	if err != nil {
	  log.Fatal(err)
	}
	for _, limit := range limits {
	  fmt.Println(limit)
	}

Example of getting quotas for a single project in specific region

	singleProjectQuotas, _, err := quotas.GetProjectQuotas(ctx, ResellClient, QuotaRegionalClient, projectID, regionName)
	if err != nil {
	  log.Fatal(err)
	}
	for _, singleProjectQuota := range singleProjectQuotas {
	  fmt.Println(singleProjectQuota)
	}

Example of updating quotas for a single project in specific region

	projectQuotaUpdateOpts := quotas.UpdateProjectQuotasOpts{
	  QuotasOpts: []*quota-manager.QuotaOpts{
	    {
	      Name: "image_gigabytes",
	      ResourceQuotasOpts: []quotas.ResourceQuotaOpts{
	        {
	          Value:  10,
	        },
	        {
	          Value:  20,
	        },
	      },
	    },
	  },
	}
	updatedProjectQuotas, _, err := quotas.UpdateProjectQuotas(context, ResellClient, QuotaRegionalClient, projectID, regionName, projectQuotaUpdateOpts)
	if err != nil {
	  log.Fatal(err)
	}
	for _, updatedProjectQuota := range updatedProjectQuotas {
	  fmt.Println(updatedProjectQuota)
	}
*/
package quotas
