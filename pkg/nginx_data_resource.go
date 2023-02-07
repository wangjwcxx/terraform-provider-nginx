package pkg

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NewNginxDataSource() datasource.DataSource {
	return &nginxDataSource{}
}

type nginxDataSource struct {
}

type nginxDataSourceModel struct {
	Listen types.String `tfsdk:"listen"`
	Port   types.String `tfsdk:"port"`
}

func (d *nginxDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nginx"
}

func (d *nginxDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "nginx http data source",
		Attributes: map[string]schema.Attribute{
			"listen": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
			"port": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "",
			},
		},
	}
}

func (d *nginxDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

func (d *nginxDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data nginxDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	data.Listen = types.StringValue(data.Listen.String())
	data.Port = types.StringValue(data.Port.String())

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
