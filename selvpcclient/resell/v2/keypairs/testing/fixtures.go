package testing

import "github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/keypairs"

// TestListResponseRaw represents a raw response from List requests.
const TestListResponseRaw = `
{
    "keypairs": [
        {
            "name": "key0",
            "public_key": "ssh-rsa AAABBBCCC user0@selectel.com",
            "regions": [
                "ru-1",
                "ru-2",
                "ru-3"
            ],
            "user_id": "82a026cae2104e92b999dbe00cdb9435"
        },
        {
            "name": "key1",
            "public_key": "ssh-rsa BBBAAACCC user1@example.org",
            "regions": [
                "ru-1",
                "ru-2"
            ],
            "user_id": "046ffcab518f430bb6fc50c5edcdd8db"
        },
        {
            "name": "key2",
            "public_key": "ssh-rsa CCCAAABBB user2@selectel.com",
            "regions": [
                "ru-3"
            ],
            "user_id": "6d7eb892ca98413e8621c6366c8416be"
        }
    ]
}
`

// TestListResponseSingleRaw represents a raw response with a single keypair
// from the List requests.
const TestListResponseSingleRaw = `
{
    "keypairs": [
        {
            "name": "key2",
            "public_key": "ssh-rsa CCCAAABBB user2@selectel.com",
            "regions": [
                "ru-3"
            ],
            "user_id": "6d7eb892ca98413e8621c6366c8416be"
        }
    ]
}
`

// TestListResponseSingle represents the unmarshalled TestListResponseSingleRaw
// response.
var TestListResponseSingle = []*keypairs.Keypair{
	{
		Name:      "key2",
		PublicKey: "ssh-rsa CCCAAABBB user2@selectel.com",
		Regions:   []string{"ru-3"},
		UserID:    "6d7eb892ca98413e8621c6366c8416be",
	},
}

// TestManyKeypairsInvalidResponseRaw represents a raw invalid response from the
// List call.
const TestManyKeypairsInvalidResponseRaw = `
{
    "keypairs": [
        {
            "user_id": 123
        }
    ]
}
`
