package main

import (
	"testing"

	"oh-my-posh/runtime"

	"github.com/stretchr/testify/assert"
)

func TestOSInfo(t *testing.T) {
	cases := []struct {
		Case              string
		ExpectedString    string
		GOOS              string
		WSLDistro         string
		Platform          string
		DisplayDistroName bool
	}{
		{
			Case:           "WSL debian - icon",
			ExpectedString: "WSL at \uf306",
			GOOS:           "linux",
			WSLDistro:      "debian",
			Platform:       "debian",
		},
		{
			Case:              "WSL debian - name",
			ExpectedString:    "WSL at burps",
			GOOS:              "linux",
			WSLDistro:         "burps",
			Platform:          "debian",
			DisplayDistroName: true,
		},
		{
			Case:           "plain linux - icon",
			ExpectedString: "\uf306",
			GOOS:           "linux",
			Platform:       "debian",
		},
		{
			Case:              "plain linux - name",
			ExpectedString:    "debian",
			GOOS:              "linux",
			Platform:          "debian",
			DisplayDistroName: true,
		},
		{
			Case:           "windows",
			ExpectedString: "windows",
			GOOS:           "windows",
		},
		{
			Case:           "darwin",
			ExpectedString: "darwin",
			GOOS:           "darwin",
		},
		{
			Case:           "unknown",
			ExpectedString: "unknown",
			GOOS:           "unknown",
		},
	}
	for _, tc := range cases {
		env := new(runtime.MockedEnvironment)
		env.On("GetRuntimeGOOS", nil).Return(tc.GOOS)
		env.On("Getenv", "WSL_DISTRO_NAME").Return(tc.WSLDistro)
		env.On("GetPlatform", nil).Return(tc.Platform)
		props := &properties{
			values: map[Property]interface{}{
				WSL:               "WSL",
				WSLSeparator:      " at ",
				DisplayDistroName: tc.DisplayDistroName,
				Windows:           "windows",
				MacOS:             "darwin",
			},
		}
		osInfo := &osInfo{
			env:   env,
			props: props,
		}
		assert.Equal(t, tc.ExpectedString, osInfo.string(), tc.Case)
		if tc.WSLDistro != "" {
			assert.Equal(t, tc.WSLDistro, osInfo.OS, tc.Case)
		} else if tc.Platform != "" {
			assert.Equal(t, tc.Platform, osInfo.OS, tc.Case)
		} else {
			assert.Equal(t, tc.GOOS, osInfo.OS, tc.Case)
		}
	}
}
