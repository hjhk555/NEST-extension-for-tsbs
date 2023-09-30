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
	//TODO implement me
	return &benchmark{}, nil
}

func (a arcadetsdbTarget) Serializer() serialize.PointSerializer {
	//TODO implement me
	panic("implement me")
}

func (a arcadetsdbTarget) TargetSpecificFlags(flagPrefix string, flagSet *pflag.FlagSet) {
	//TODO implement me
	flagSet.String(flagPrefix+"host", "localhost", "ArcadeTSDB host address.")
	flagSet.Int(flagPrefix+"port", 8809, "ArcadeTSDB port.")
}

func (a arcadetsdbTarget) TargetName() string {
	//TODO implement me
	return constants.FormatArcadeTSDB
}
