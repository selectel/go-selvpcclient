/*
Package quotamanager provides methods and structures to work with the Quotas Manager API.

Example of creating QuotaRegional client

	  resellClient := resell.NewV2ResellClient(APIToken)
	  ctx := context.Background()
	  accountName := "123456"
	  token, _, _ := reselTokens.Create(ctx, resellClient, reselTokens.TokenOpts{
		  AccountName: accountName,
	  })

	  OpenstackClient := resell.NewOpenstackClient(token.ID)
	  identity := quotamanager.NewIdentityManager(resellClient, OpenstackClient, accountName)
	  QuotaRegionalClient := quotamanager.NewQuotaRegionalClient(selvpcclient.NewHTTPClient(), identity)
*/
package quotamanager
