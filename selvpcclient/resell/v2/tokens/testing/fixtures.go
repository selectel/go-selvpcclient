package testing

import "github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/tokens"

// TestCreateTokenOptsRaw represents marshalled options for the Create request.
const TestCreateTokenOptsRaw = `
{
    "token": {
        "project_id": "4aa8f79c86744f13aa80c324b8d24158"
    }
}
`

// TestCreateTokenOpts represent options for the Create request.
var TestCreateTokenOpts = tokens.TokenOpts{
	ProjectID: "4aa8f79c86744f13aa80c324b8d24158",
}

// TestCreateTokenResponseRaw represents a raw response from the Create request.
const TestCreateTokenResponseRaw = `
{
    "token": {
        "id": "gAAAAABa3IpNNDDRzihlPrit6xFuCw7V2uMFcP9yGAYahg-xdxmKJ2QMVKtU_lJm0gfJL15GSegslW9IUBajeraM1y2oo0Fds7yzynVJKweCcsqpkArT2OHJfj_Nfqub62Ffv17SeFx7F-3c8-9P-xJ3McG93Cdd18oI_fkwZMEBTq5hzqSmbNQ"
    }
}
`

// TestCreateTokenResponse represents the unmarshalled TestCreateTokenResponseRaw response.
var TestCreateTokenResponse = &tokens.Token{
	ID: "gAAAAABa3IpNNDDRzihlPrit6xFuCw7V2uMFcP9yGAYahg-xdxmKJ2QMVKtU_lJm0gfJL15GSegslW9IUBajeraM1y2oo0Fds7yzynVJKweCcsqpkArT2OHJfj_Nfqub62Ffv17SeFx7F-3c8-9P-xJ3McG93Cdd18oI_fkwZMEBTq5hzqSmbNQ",
}

// TestTokenInvalidResponseRaw represents a raw invalid response with a single token.
const TestTokenInvalidResponseRaw = `
{
    "token": {
        "id": 123
    }
}
`
