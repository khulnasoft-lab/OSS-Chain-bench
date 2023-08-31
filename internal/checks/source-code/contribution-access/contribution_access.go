package contributionaccess

import (
	_ "embed"
	"encoding/json"

	"github.com/khulnasoft-lab/oss-chain-bench/internal/checks/common"
	"github.com/khulnasoft-lab/oss-chain-bench/internal/models/checkmodels"
)

var (
	//go:embed rules.metadata.json
	metadataString []byte

	//go:embed rules.rego
	regoQuery string

	checksMetadata checkmodels.CheckMetadataMap

	checks = []*checkmodels.Check{}
)

func init() {
	if err := json.Unmarshal(metadataString, &checksMetadata); err != nil {
		panic(err)
	}

	if err := common.AppendCheck(&checks,
		checkmodels.Check{
			Action:           common.GetRegoRunAction(regoQuery, checksMetadata),
			CheckMetadataMap: checksMetadata,
		},
	); err != nil {
		panic(err)
	}
}

func GetChecks() []*checkmodels.Check {
	return checks
}
