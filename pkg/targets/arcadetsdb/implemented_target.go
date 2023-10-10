package arcadetsdb

import (
	"github.com/blagojts/viper"
	"github.com/spf13/pflag"
	"github.com/timescale/tsbs/pkg/data/serialize"
	"github.com/timescale/tsbs/pkg/data/source"
	"github.com/timescale/tsbs/pkg/targets"
	"github.com/timescale/tsbs/pkg/targets/constants"
)

func NewTarget() targets.ImplementedTarget {
	return &arcadetsdbTarget{}
}

type arcadetsdbTarget struct {
}

func (a arcadetsdbTarget) Benchmark(targetDB string, dataSourceConfig *source.DataSourceConfig, v *viper.Viper) (targets.Benchmark, error) {
	return newBenchmark(targetDB, dataSourceConfig, parseSpecificConfig(v))
}

func (a arcadetsdbTarget) Serializer() serialize.PointSerializer {
	return &serializer{}
}

func (a arcadetsdbTarget) TargetSpecificFlags(flagPrefix string, flagSet *pflag.FlagSet) {
	flagSet.String(flagPrefix+"host", "localhost", "ArcadeTSDB host address.")
	flagSet.Int(flagPrefix+"port", 8809, "ArcadeTSDB port.")
}

func (a arcadetsdbTarget) TargetName() string {
	return constants.FormatArcadeTSDB
}
