package repo

import (
	"github.com/glasskube/glasskube/internal/repo/client"
	"github.com/glasskube/glasskube/internal/repo/types"
)

type (
	PackageIndex         = types.PackageIndex
	PackageIndexItem     = types.PackageIndexItem
	PackageRepoIndex     = types.PackageRepoIndex
	PackageRepoIndexItem = types.PackageRepoIndexItem
)

var NewClientset = client.NewClientset
