package main_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var inPath string

var useInstanceProfile = os.Getenv("SEMVER_TESTING_USE_INSTANCE_PROFILE")
var accessKeyID = os.Getenv("SEMVER_TESTING_ACCESS_KEY_ID")
var secretAccessKey = os.Getenv("SEMVER_TESTING_SECRET_ACCESS_KEY")
var bucketName = os.Getenv("SEMVER_TESTING_BUCKET")
var regionName = os.Getenv("SEMVER_TESTING_REGION")

var _ = BeforeSuite(func() {
	var err error

	if useInstanceProfile == "" {
		Expect(accessKeyID).ShouldNot(BeEmpty(), "must specify $SEMVER_TESTING_ACCESS_KEY_ID or SEMVER_TESTING_USE_INSTANCE_PROFILE=true")
		Expect(secretAccessKey).ShouldNot(BeEmpty(), "must specify $SEMVER_TESTING_SECRET_ACCESS_KEY or SEMVER_TESTING_USE_INSTANCE_PROFILE=true")
	}
	Expect(bucketName).ShouldNot(BeEmpty(), "must specify $SEMVER_TESTING_BUCKET")
	Expect(regionName).ShouldNot(BeEmpty(), "must specify $SEMVER_TESTING_REGION")

	inPath, err = gexec.Build("github.com/concourse/semver-resource/in")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func TestIn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "In Suite")
}
