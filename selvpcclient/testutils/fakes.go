package testutils

// FakeTokenID is a fake token for the Selectel VPC API.
const FakeTokenID = "fakeUUID"

const TokenInfo = `
{
  "token": {
    "catalog": [
      {
        "endpoints": [
          {
            "id": "8daa99650cff4502874f30bafa6fc8a8",
            "interface": "public",
            "region_id": "ru-1",
            "url": "{{ .QuotaManagerEndpoint}}",
            "region": "ru-1"
          },
          {
            "id": "5040e0815ace4a7e9eeac3a1706ad899",
            "interface": "public",
            "region_id": "ru-3",
            "url": "https://api.selvpc/quota-manager/v1",
            "region": "ru-3"
          },
          {
            "id": "95a24577049349309bb949a5f8bcc253",
            "interface": "admin",
            "region_id": "ru-1",
            "url": "http://ru-1.openstack:63079",
            "region": "ru-1"
          }
        ],
        "id": "73bbae11073b4e34b7c8fcf7dfc27ef1",
        "type": "quota-manager",
        "name": "bizeff"
      },
      {
        "endpoints": [
          {
            "id":"8a2bd27225434e098068f49af7cd6f79",
            "interface":"public",
            "region_id":"ru-1",
            "url":"{{ .ResellEndpoint}}",
            "region":"ru-1"
          },
          {
            "id":"2934bf5234e54bd3bfef8544893a5e71",
            "interface":"public",
            "region_id":"ru-2",
            "url":"{{ .ResellEndpoint}}",
            "region":"ru-2"
          }
        ],
        "id":"5549d38ff47b4dbfb406d5b459023090",
        "type":"resell",
        "name":"hong"
      }
    ]
  }
}
`

type TokenInfoTemplate struct {
	QuotaManagerEndpoint string
	ResellEndpoint       string
}
