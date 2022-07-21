package utils_test

import (
	"go-app/pkg/utils"
	"testing"
)

type ExpectedResult struct {
	title string
	slug  string
}

var expectedResults = []ExpectedResult{
	{
		title: "Fix low wifi speed on Linux (Ubuntu) with chip Atheros AR9285",
		slug:  "fix-low-wifi-speed-on-linux-ubuntu-with-chip-atheros-ar9285",
	},
	{
		title: "Python and Scala smoke the peace pipe",
		slug:  "python-and-scala-smoke-the-peace-pipe",
	},
	{
		title: "Graphite, Carbon and Diamond",
		slug:  "graphite-carbon-and-diamond",
	},
	{
		title: "Here I go PyGrunn'13!",
		slug:  "here-i-go-pygrunn13",
	},
	{
		title: "How-to install GNOME 3 instead of Unity on Ubuntu 11.04",
		slug:  "how-to-install-gnome-3-instead-of-unity-on-ubuntu-1104",
	},
}

func TestSlugify(t *testing.T) {
	t.Parallel()
	for testNumber, testExpected := range expectedResults {
		title := testExpected.title
		expectedSlug := testExpected.slug

		if result := utils.Slugify(title); result != expectedSlug {
			t.Errorf("#%d (%s)\n+++ %s\n--- %s", testNumber, title, result, expectedSlug)
		}
	}
}
