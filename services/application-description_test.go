// Copyright 2018 gopcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package services

import (
	"testing"

	"github.com/wmnsk/gopcua/utils/codectest"
)

func TestApplicationDescription(t *testing.T) {
	cases := []codectest.Case{
		{
			Name: "Normal",
			Struct: NewApplicationDescription(
				"app-uri",
				"prod-uri",
				"app-name",
				AppTypeServer,
				"gw-uri",
				"prof-uri",
				[]string{"discov-uri-1", "discov-uri-2"},
			),
			Bytes: []byte{ // Single
				// ApplicationURI
				0x07, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d, 0x75, 0x72, 0x69,
				// ProductURI
				0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x64, 0x2d, 0x75, 0x72, 0x69,
				// ApplicationName
				0x02, 0x08, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d,
				0x6e, 0x61, 0x6d, 0x65,
				// ApplicationType
				0x00, 0x00, 0x00, 0x00,
				// GatewayServerURI
				0x06, 0x00, 0x00, 0x00, 0x67, 0x77, 0x2d, 0x75, 0x72, 0x69,
				// DiscoveryProfileURI
				0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x66, 0x2d, 0x75, 0x72, 0x69,
				// DiscoveryURIs
				0x02, 0x00, 0x00, 0x00,
				0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x31,
				0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x32,
			},
		},
	}
	codectest.Run(t, cases, func(b []byte) (codectest.S, error) {
		return DecodeApplicationDescription(b)
	})
}

func TestApplicationDescriptionArray(t *testing.T) {
	cases := []codectest.Case{
		{
			Name: "Normal",
			Struct: NewApplicationDescriptionArray(
				[]*ApplicationDescription{
					NewApplicationDescription(
						"app-uri",
						"prod-uri",
						"app-name",
						AppTypeServer,
						"gw-uri",
						"prof-uri",
						[]string{"discov-uri-1", "discov-uri-2"},
					),
					NewApplicationDescription(
						"app-uri",
						"prod-uri",
						"app-name",
						AppTypeServer,
						"gw-uri",
						"prof-uri",
						[]string{"discov-uri-1", "discov-uri-2"},
					),
				},
			),
			Bytes: []byte{ // Array
				// ArraySize
				0x02, 0x00, 0x00, 0x00,
				// ApplicationURI
				0x07, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d, 0x75, 0x72, 0x69,
				// ProductURI
				0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x64, 0x2d, 0x75, 0x72, 0x69,
				// ApplicationName
				0x02, 0x08, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d,
				0x6e, 0x61, 0x6d, 0x65,
				// ApplicationType
				0x00, 0x00, 0x00, 0x00,
				// GatewayServerURI
				0x06, 0x00, 0x00, 0x00, 0x67, 0x77, 0x2d, 0x75, 0x72, 0x69,
				// DiscoveryProfileURI
				0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x66, 0x2d, 0x75, 0x72, 0x69,
				// DiscoveryURIs
				0x02, 0x00, 0x00, 0x00,
				0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x31,
				0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x32,
				// ApplicationURI
				0x07, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d, 0x75, 0x72, 0x69,
				// ProductURI
				0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x64, 0x2d, 0x75, 0x72, 0x69,
				// ApplicationName
				0x02, 0x08, 0x00, 0x00, 0x00, 0x61, 0x70, 0x70, 0x2d,
				0x6e, 0x61, 0x6d, 0x65,
				// ApplicationType
				0x00, 0x00, 0x00, 0x00,
				// GatewayServerURI
				0x06, 0x00, 0x00, 0x00, 0x67, 0x77, 0x2d, 0x75, 0x72, 0x69,
				// DiscoveryProfileURI
				0x08, 0x00, 0x00, 0x00, 0x70, 0x72, 0x6f, 0x66, 0x2d, 0x75, 0x72, 0x69,
				// DiscoveryURIs
				0x02, 0x00, 0x00, 0x00,
				0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x31,
				0x0c, 0x00, 0x00, 0x00, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69, 0x2d, 0x32,
			},
		},
	}
	codectest.Run(t, cases, func(b []byte) (codectest.S, error) {
		return DecodeApplicationDescriptionArray(b)
	})
}
