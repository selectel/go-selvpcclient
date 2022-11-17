/*
Package quotas provides the ability to retrieve and update quotas through the
Resell v2 API.

Example of creating QuotaRegional client

	  resellClient := resell.NewV2ResellClient(APIToken)
	  ctx := context.Background()
	  accountName := "123456"
	  token, _, _ := reselTokens.Create(ctx, resellClient, reselTokens.TokenOpts{
		  AccountName: accountName,
	  })

	  OpenstackClient := resell.NewOpenstackClient(token.ID)
	  identity := quotas.NewIdentityManager(resellClient, OpenstackClient, accountName)
	  quotaMgr := quotas.NewQuotaRegionalClient(selvpcclient.NewHTTPClient(), identity)

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
	  QuotasOpts: []*quotas.QuotaOpts{
	    {
	      Name: "image_gigabytes",
	      ResourceQuotasOpts: []quotas.ResourceQuotaOpts{
	        {
	          Region: "ru-1",
	          Value:  10,
	        },
	        {
	          Region: "ru-2",
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
