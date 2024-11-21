package traffic

import (
	"encoding/json"
	"time"

	"github.com/selectel/go-selvpcclient/v4/selvpcclient"
)

// Traffic contains information about used and paid traffic.
type Traffic struct {
	// Type is a human-readable name of the type of traffic.
	Type string `json:"-"`

	// TrafficData contains information about traffic.
	TrafficData Data `json:"-"`
}

// Data represents information about traffic in the specified period.
type Data struct {
	// Start contains the start timestamp.
	Start time.Time `json:"-"`

	// Stop contains the stop timestamp.
	Stop time.Time `json:"-"`

	// Unit represents a unit that is used to represent traffic data.
	Unit string `json:"unit"`

	// Value contains traffic value for the specified period.
	Value int `json:"value"`
}

// UnmarshalJSON helps to unmarshal Data timestamp fields into the needed values.
func (r *Data) UnmarshalJSON(b []byte) error {
	type tmp Data
	var s struct {
		tmp
		Start selvpcclient.JSONRFC3339NoZTimezone `json:"start"`
		Stop  selvpcclient.JSONRFC3339NoZTimezone `json:"stop"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Data(s.tmp)

	// Collect traffic timestamps.
	r.Start = time.Time(s.Start)
	r.Stop = time.Time(s.Stop)

	return err
}

// DomainTraffic represents domain traffic information.
type DomainTraffic struct {
	// DomainData contains data about domain traffic.
	DomainData []*Traffic `json:"domain"`
}

/*
UnmarshalJSON implements custom unmarshalling method for the DomainTraffic type.

We need it to work with a JSON structure that the Resell v2 API responses with:

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
	}
*/
func (result *DomainTraffic) UnmarshalJSON(b []byte) error {
	// Populate temporary structure with resource quotas represented as maps.
	var s struct {
		DomainTraffic map[string]Data `json:"domain"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	// Populate the result with an empty slice in case of empty traffic list.
	*result = DomainTraffic{
		DomainData: []*Traffic{},
	}

	if len(s.DomainTraffic) != 0 {
		// Convert domain traffic maps to the slice of Traffic types.
		// Here we're allocating memory in advance because we already know the length
		// of a result slice from the JSON bytearray.
		domainTrafficSlice := make([]*Traffic, len(s.DomainTraffic))
		i := 0
		for trafficType, trafficData := range s.DomainTraffic {
			domainTrafficSlice[i] = &Traffic{
				Type:        trafficType,
				TrafficData: trafficData,
			}
			i++
		}

		// Add the unmarshalled traffic slice to the result.
		result.DomainData = domainTrafficSlice
	}

	return nil
}
