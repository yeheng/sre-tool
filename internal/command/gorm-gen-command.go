package command

import (
	"fmt"
	"sre-tool/internal/utils"
	"strings"

	"github.com/spf13/cobra"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type GormGenCommand struct {
	generator    *gen.Generator
	db           *gorm.DB
	config       gen.Config
	tables       []string
	ignoreFileds []string
	methods      map[string][]any
	dataMap      map[string]func(detailType gorm.ColumnType) (dataType string)
}

type OptionFunc func(*GormGenCommand)

type OptionInterface interface {
	apply(*GormGenCommand)
}

func (f OptionFunc) apply(g *GormGenCommand) {
	f(g)
}

func NewGenCommand(opts ...OptionInterface) Interface {
	g := &GormGenCommand{}
	for _, opt := range opts {
		opt.apply(g)
	}
	gormGen := gen.NewGenerator(g.config)
	gormGen.UseDB(g.db)
	g.generator = gormGen
	return g
}

func (g *GormGenCommand) genModel() {
	g.genModelMethod()
	if g.dataMap != nil {
		g.generator.WithDataTypeMap(g.dataMap)
	}
	jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
		if len(g.ignoreFileds) > 0 {
			for _, v := range g.ignoreFileds {
				if strings.Contains(v, columnName) {
					return "-"
				}
			}
		}
		return columnName
	})
	fieldOpts := []gen.ModelOpt{jsonField}
	models := make([]any, len(g.tables))
	if len(g.tables) > 0 {
		for _, table := range g.tables {
			model := g.generator.GenerateModelAs(table, utils.CaseToCamel(table), fieldOpts...)
			models = append(models, model)
		}
	} else {
		models = g.generator.GenerateAllTable(fieldOpts...)
	}
	g.generator.ApplyBasic(models...)
	g.generator.Execute()
}

func (g *GormGenCommand) genModelMethod() {
	for table, methods := range g.methods {
		if len(methods) == 0 || table == "" {
			continue
		}
		for _, method := range methods {
			g.generator.ApplyInterface(method, g.generator.GenerateModel(table))
		}
	}
}

func (g *GormGenCommand) Command() *cobra.Command {
	return &cobra.Command{
		Use:   "gen:model",
		Short: "创建model或model对应方法",
		Long:  `基于gorm的gen的代码生成器，生成数据表model，并生成model对应的方法。`,
		Run: func(cmd *cobra.Command, args []string) {
			g.genModel()
			fmt.Printf("\033[1;32;42m%s\033[0m\n", "generated successfully.")
		},
	}
}

func (g *GormGenCommand) Flags(root *cobra.Command) {}

func WithDB(db *gorm.DB) OptionFunc {
	return func(ggc *GormGenCommand) {
		ggc.db = db
	}
}

func WithTables(tables []string) OptionFunc {
	return func(ggc *GormGenCommand) {
		ggc.tables = tables
	}
}

func WithIgnoreFileds(ignoreFileds []string) OptionFunc {
	return func(ggc *GormGenCommand) {
		ggc.ignoreFileds = ignoreFileds
	}
}

func WithMethods(methods map[string][]any) OptionFunc {
	return func(ggc *GormGenCommand) {
		ggc.methods = methods
	}
}

func WithDataMap(dataMap map[string]func(detailType gorm.ColumnType) (dataType string)) OptionFunc {
	return func(ggc *GormGenCommand) {
		ggc.dataMap = dataMap
	}
}

func WithConfig(conf gen.Config) OptionFunc {
	return func(ggc *GormGenCommand) {
		ggc.config = conf
	}
}
