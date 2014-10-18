package core

import (
	"fmt"
	"sync"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type CoreSuite struct {
	sync.WaitGroup
}

var _ = Suite(&CoreSuite{})

func (s *CoreSuite) TestVCS_IsValid(c *C) {
	vcs := VCS("git@github.com:mcuadros/dockership.git")
	c.Assert(vcs.IsValid(), Equals, true)
}

func (s *CoreSuite) TestVCS_NotIsValid(c *C) {
	vcs := VCS("foo")
	c.Assert(vcs.IsValid(), Equals, false)
}

func (s *CoreSuite) TestVCS_Info(c *C) {
	vcs := VCS("git@github.com:mcuadros/dockership.git")

	c.Assert(vcs.Info().Name, Equals, "dockership")
	c.Assert(vcs.Info().Username, Equals, "mcuadros")
	c.Assert(vcs.Info().Branch, Equals, "master")
}

func (s *CoreSuite) TestVCS_InfoBranch(c *C) {
	vcs := VCS("git@github.com:mcuadros/dockership.git!branch")

	c.Assert(vcs.Info().Name, Equals, "dockership")
	c.Assert(vcs.Info().Username, Equals, "mcuadros")
	c.Assert(vcs.Info().Branch, Equals, "branch")
}

func (s *CoreSuite) TestCommit_GetShort(c *C) {
	commit := Commit("123456789123456789")

	c.Assert(commit.GetShort(), Equals, "123456789123")
}

func (s *CoreSuite) TestRevision_Get(c *C) {
	revision := Revision{"foo": "123456789123456789", "bar:": "123456789123456789"}

	c.Assert(revision.Get(), Equals, "28a247e8ba3ab48ab72dd196f1052f8a")
}

func (s *CoreSuite) TestRevision_String(c *C) {
	revisionAZ := Revision{"foo": "ZZZZZZZZZZZZZZZZZZ", "bar:": "123456789123456789"}
	c.Assert(fmt.Sprintf("%s", revisionAZ), Equals, "e1ba1f05de5f184fe94ec745250b5d9e")

	revisionZA := Revision{"foo": "123456789123456789", "bar:": "ZZZZZZZZZZZZZZZZZZ"}
	c.Assert(fmt.Sprintf("%s", revisionZA), Equals, "e1ba1f05de5f184fe94ec745250b5d9e")
}
