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
	  identity := quota-manager.NewIdentityManager(resellClient, OpenstackClient, accountName)
	  quotaMgr := quota-manager.NewQuotaRegionalClient(selvpcclient.NewHTTPClient(), identity)

Example of getting quota limits for a single project

	limits, _, err := quota-manager.GetLimits(ctx, QuotaRegionalClient)
	if err != nil {
	  log.Fatal(err)
	}
	for _, limit := range limits {
	  fmt.Println(limit)
	}

Example of getting quotas for a single project in specific region

	singleProjectQuotas, _, err := quota-manager.GetProjectQuotas(ctx, ResellClient, QuotaRegionalClient, projectID, regionName)
	if err != nil {
	  log.Fatal(err)
	}
	for _, singleProjectQuota := range singleProjectQuotas {
	  fmt.Println(singleProjectQuota)
	}

Example of updating quotas for a single project in specific region

	projectQuotaUpdateOpts := quota-manager.UpdateProjectQuotasOpts{
	  QuotasOpts: []*quota-manager.QuotaOpts{
	    {
	      Name: "image_gigabytes",
	      ResourceQuotasOpts: []quota-manager.ResourceQuotaOpts{
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
	updatedProjectQuotas, _, err := quota-manager.UpdateProjectQuotas(context, ResellClient, QuotaRegionalClient, projectID, regionName, projectQuotaUpdateOpts)
	if err != nil {
	  log.Fatal(err)
	}
	for _, updatedProjectQuota := range updatedProjectQuotas {
	  fmt.Println(updatedProjectQuota)
	}
*/
package quotas
