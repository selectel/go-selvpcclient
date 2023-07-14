package testing

import (
	"time"

	"github.com/selectel/go-selvpcclient/v3/selvpcclient"
	"github.com/selectel/go-selvpcclient/v3/selvpcclient/resell/v2/traffic"
)

// TestGetTrafficRaw represents a raw response from the Get request.
const TestGetTrafficRaw = `
{
    "traffic": {
        "domain": {
            "paid": {
                "start": "2018-04-01T00:00:00",
                "stop": "2018-04-30T23:59:59",
                "unit": "B",
                "value": 0
            },
            "prepaid": {
                "start": "2018-04-01T00:00:00",
                "stop": "2018-04-30T23:59:59",
                "unit": "B",
                "value": 3000000000000
            },
            "used": {
                "start": "2018-04-01T00:00:00",
                "stop": "2018-04-30T23:59:59",
                "unit": "B",
                "value": 658003816
            }
        },
        "projects": {}
    }
}
`

// TestGetTrafficUsedRaw represents a raw response from the Get request with
// only used traffic.
const TestGetTrafficUsedRaw = `
{
    "traffic": {
        "domain": {
            "used": {
                "start": "2018-04-01T00:00:00",
                "stop": "2018-04-30T23:59:59",
                "unit": "B",
                "value": 658003816
            }
        },
        "projects": {}
    }
}
`

var (
	trafficStartTimeStamp, _ = time.Parse(selvpcclient.RFC3339NoZ, "2018-04-01T00:00:00")
	trafficStopTimeStamp, _  = time.Parse(selvpcclient.RFC3339NoZ, "2018-04-30T23:59:59")
)

// TestGetTrafficUsed represents the unmarshalled TestGetTrafficUsedRaw response.
var TestGetTrafficUsed = &traffic.DomainTraffic{
	DomainData: []*traffic.Traffic{
		{
			Type: "used",
			TrafficData: traffic.Data{
				Start: trafficStartTimeStamp,
				Stop:  trafficStopTimeStamp,
				Unit:  "B",
				Value: 658003816,
			},
		},
	},
}

// TestGetTrafficInvalidTimestampsRaw represents a raw response from the Get request
// with invalid timestamps.
const TestGetTrafficInvalidTimestampsRaw = `
{
    "traffic": {
        "domain": {
            "paid": {
                "start": "2006-01-02T15:04:05.999999+00:00",
                "stop": "2006-01-20T15:04:05.999999+00:00",
                "unit": "B",
                "value": 0
            },
            "prepaid": {
                "start": "2006-01-02T15:04:05.999999+00:00",
                "stop": "2006-01-20T15:04:05.999999+00:00",
                "unit": "B",
                "value": 3000000000000
            },
            "used": {
                "start": "2006-01-02T15:04:05.999999+00:00",
                "stop": "2006-01-20T15:04:05.999999+00:00",
                "unit": "B",
                "value": 658003816
            }
        },
        "projects": {}
    }
}
`

// TestGetTrafficInvalidDataResponseRaw represents a raw response from the Get request
// with invalid traffic data.
const TestGetTrafficInvalidDataResponseRaw = `
{
    "traffic": {
        "domain": {
            123: {}
        }
    }
}
`
