package testing

// TestGetCapabilitiesRaw represents a raw response from the Get request.
const TestGetCapabilitiesRaw = `
{
    "capabilities": {
        "licenses": [
            {
                "availability": [
                    "ru-1",
                    "ru-2",
                    "ru-3"
                ],
                "type": "license_windows_2012_standard"
            },
            {
                "availability": [
                    "ru-1",
                    "ru-2",
                    "ru-3"
                ],
                "type": "license_windows_2016_standard"
            }
        ],
    		"logo": {
    		    "max_size_bytes": 65536
    		},
    		"regions": [
    				{
    				    "description": "Saint Petersburg 2",
    				    "is_default": false,
    				    "name": "ru-3",
    				    "zones": [
    				        {
    				            "description": "Tsvetochnaya-1 (ru-3a)",
    				            "enabled": true,
    				            "is_default": true,
    				            "name": "ru-3a"
    				        }
    				    ]
    				},
    				{
    				    "description": "Moscow",
    				    "is_default": false,
    				    "name": "ru-2",
    				    "zones": [
    				        {
    				            "description": "Berzarina-1 (ru-2a)",
    				            "enabled": true,
    				            "is_default": true,
    				            "name": "ru-2a"
    				        }
    				    ]
    				},
    				{
    				    "description": "Saint Petersburg",
    				    "is_default": true,
    				    "name": "ru-1",
    				    "zones": [
    				        {
    				            "description": "Dubrovka-1 (ru-1a)",
    				            "enabled": true,
    				            "is_default": false,
    				            "name": "ru-1a"
    				        },
    				        {
    				            "description": "Dubrovka-2 (ru-1b)",
    				            "enabled": true,
    				            "is_default": true,
    				            "name": "ru-1b"
    				        }
    				    ]
    				}
    		],
    		"resources": [
    		    {
    		        "name": "compute_cores",
    		        "preordered": false,
    		        "quota_scope": "zone",
    		        "quotable": true,
    		        "unbillable": true
    		    },
    		    {
    		        "name": "compute_ram",
    		        "preordered": false,
    		        "quota_scope": "zone",
    		        "quotable": true,
    		        "unbillable": true
    		    },
    		    {
    		        "name": "volume_gigabytes_fast",
    		        "preordered": false,
    		        "quota_scope": "zone",
    		        "quotable": true,
    		        "unbillable": false
    		    },
    		    {
    		        "name": "network_floatingips",
    		        "preordered": false,
    		        "quota_scope": null,
    		        "quotable": false,
    		        "unbillable": false
    		    },
    		    {
    		        "name": "volume_gigabytes_basic",
    		        "preordered": false,
    		        "quota_scope": "zone",
    		        "quotable": true,
    		        "unbillable": false
    		    },
    		    {
    		        "name": "image_gigabytes",
    		        "preordered": false,
    		        "quota_scope": "region",
    		        "quotable": true,
    		        "unbillable": false
    		    },
    		    {
    		        "name": "network_subnets_29",
    		        "preordered": false,
    		        "quota_scope": null,
    		        "quotable": false,
    		        "unbillable": false
    		    },
    		    {
    		        "name": "license_windows_2012_standard",
    		        "preordered": false,
    		        "quota_scope": null,
    		        "quotable": false,
    		        "unbillable": true
    		    },
    		    {
    		        "name": "volume_gigabytes_universal",
    		        "preordered": false,
    		        "quota_scope": "zone",
    		        "quotable": true,
    		        "unbillable": false
    		    },
    		    {
    		        "name": "network_subnets_25",
    		        "preordered": true,
    		        "quota_scope": null,
    		        "quotable": false,
    		        "unbillable": false
    		    },
    		    {
    		        "name": "network_subnets_28",
    		        "preordered": true,
    		        "quota_scope": null,
    		        "quotable": false,
    		        "unbillable": false
    		    },
    		    {
    		        "name": "network_subnets_27",
    		        "preordered": true,
    		        "quota_scope": null,
    		        "quotable": false,
    		        "unbillable": false
    		    },
    		    {
    		        "name": "network_subnets_26",
    		        "preordered": true,
    		        "quota_scope": null,
    		        "quotable": false,
    		        "unbillable": false
    		    },
    		    {
    		        "name": "network_subnets_24",
    		        "preordered": true,
    		        "quota_scope": null,
    		        "quotable": false,
    		        "unbillable": false
    		    },
    		    {
    		        "name": "network_subnets_29_vrrp",
    		        "preordered": true,
    		        "quota_scope": null,
    		        "quotable": false,
    		        "unbillable": false
    		    },
    		    {
    		        "name": "license_windows_2016_standard",
    		        "preordered": false,
    		        "quota_scope": null,
    		        "quotable": false,
    		        "unbillable": true
    		    }
    		],
        "subnets": [
            {
                "availability": [
                    "ru-1",
                    "ru-2",
                    "ru-3"
                ],
                "prefix_length": "29",
                "type": "ipv4"
            }
        ],
        "traffic": {
            "granularities": [
                {
                    "granularity": 3600,
                    "timespan": 96
                },
                {
                    "granularity": 1,
                    "timespan": 32
                },
                {
                    "granularity": 86400,
                    "timespan": 1825
                }
            ]
        }
    }
}
`

// TestGetCapabilitiesInvalidRaw represents a raw response from the Get request.
const TestGetCapabilitiesInvalidRaw = `
{
    "capabilities": {
        123: {}
    }
}
`
