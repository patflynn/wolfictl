package vex

import (
	"path/filepath"
	"testing"
	"time"

	"chainguard.dev/melange/pkg/build"
	"chainguard.dev/vex/pkg/vex"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestFromPackageConfiguration(t *testing.T) {
	buildCfg, err := build.ParseConfiguration(filepath.Join("testdata", "git.yaml"))
	if err != nil {
		return
	}
	vexCfg := Config{
		Distro: "wolfi",
	}

	doc, err := FromPackageConfiguration(vexCfg, buildCfg)
	require.NoError(t, err)

	// zero out non-deterministic fields
	doc.Timestamp = nil
	timePointer := func(t time.Time) *time.Time { return &t }
	tz := time.FixedZone("-0500", -5*3600)
	expected := &vex.VEX{
		Metadata: vex.Metadata{
			ID:     "vex-60d2bb8952362a0a7bf52b2ac2619a7846cd2394a5cdb3dfe83a66f5f9838e7d",
			Format: "text/vex",
		},
		Statements: []vex.Statement{
			{
				Vulnerability: "CVE-1234-5678",
				Products: []string{
					"pkg:apk/wolfi/git@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-daemon@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-email@2.39.0-r0?distro=wolfi",
				},
				Status: "not_affected",
			},
			{
				Vulnerability: "CVE-2022-1111",
				Timestamp:     timePointer(time.Date(2022, 12, 23, 1, 28, 16, 0, tz)),
				Products: []string{
					"pkg:apk/wolfi/git@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-daemon@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-email@2.39.0-r0?distro=wolfi",
				},
				Status: "under_investigation",
			},
			{
				Vulnerability: "CVE-2022-1111",
				Timestamp:     timePointer(time.Date(2022, 12, 23, 2, 11, 57, 0, tz)),
				Products: []string{
					"pkg:apk/wolfi/git@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-daemon@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-email@2.39.0-r0?distro=wolfi",
				},
				Status:        "not_affected",
				Justification: "component_not_present",
			},
			{
				Vulnerability: "CVE-2022-2222",
				Timestamp:     timePointer(time.Date(2022, 12, 24, 1, 28, 16, 0, tz)),
				Products: []string{
					"pkg:apk/wolfi/git@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-daemon@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-email@2.39.0-r0?distro=wolfi",
				},
				Status: "under_investigation",
			},
			{
				Vulnerability: "CVE-2022-2222",
				Timestamp:     timePointer(time.Date(2022, 12, 24, 2, 12, 49, 0, tz)),
				Products: []string{
					"pkg:apk/wolfi/git@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-daemon@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-email@2.39.0-r0?distro=wolfi",
				},
				Status:          "affected",
				ActionStatement: "action statement",
			},
			{
				Vulnerability: "CVE-2022-2222",
				Timestamp:     timePointer(time.Date(2022, 12, 24, 2, 50, 18, 0, tz)),
				Products: []string{
					"pkg:apk/wolfi/git@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-daemon@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-email@2.39.0-r0?distro=wolfi",
				},
				Status: "fixed",
			},
			{
				Vulnerability: "CVE-2022-39253",
				Products: []string{
					"pkg:apk/wolfi/git@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-daemon@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-email@2.39.0-r0?distro=wolfi",
				},
				Status: "fixed",
			},
			{
				Vulnerability: "CVE-2022-39260",
				Products: []string{
					"pkg:apk/wolfi/git@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-daemon@2.39.0-r0?distro=wolfi",
					"pkg:apk/wolfi/git-email@2.39.0-r0?distro=wolfi",
				},
				Status: "fixed",
			},
		},
	}

	if diff := cmp.Diff(expected, doc); diff != "" {
		t.Errorf("Unexpected result from FromPackageConfiguration (-want, +got):\n%s", diff)
	}
}
