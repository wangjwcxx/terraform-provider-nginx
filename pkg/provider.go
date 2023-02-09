package pkg

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"golang.org/x/crypto/ssh"
)

type nginxPorvider struct {
	version string
}

var _ provider.Provider = &nginxPorvider{}

type nginxProviderModel struct {
	Ip       types.String `tfsdk:"ip"`
	Port     types.String `tfsdk:"port"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (p *nginxPorvider) Metadata(ctx context.Context, req provider.MetadataRequest, rep *provider.MetadataResponse) {
	rep.TypeName = "nginx"
	rep.Version = p.version
}

// provider 的属性
func (p *nginxPorvider) Schema(ctx context.Context, req provider.SchemaRequest, rep *provider.SchemaResponse) {
	rep.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ip": schema.StringAttribute{
				Optional: true,
			},
			"username": schema.StringAttribute{
				Optional: true,
			},
			"password": schema.StringAttribute{
				Optional: true,
			},
			"port": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *nginxPorvider) Configure(ctx context.Context, req provider.ConfigureRequest, rep *provider.ConfigureResponse) {

	var data nginxProviderModel

	rep.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if rep.Diagnostics.HasError() {
		return
	}
}

func (p *nginxPorvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewNginxDataSource,
	}
}

func (p *nginxPorvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewNginxResource,
	}
}

func (p *nginxPorvider) ValidateConfig(ctx context.Context, req provider.ValidateConfigRequest, rep *provider.ValidateConfigResponse) {
	var data nginxProviderModel

	rep.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if rep.Diagnostics.HasError() {
		return
	}

}

func (p *nginxPorvider) Create(ctx context.Context, req resource.CreateRequest, rep *resource.CreateResponse) {
	var data *nginxProviderModel

	rep.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if rep.Diagnostics.HasError() {
		return
	}

	// ssh 登录服务器安装nginx

	client, err := ssh.Dial("tcp", data.Ip.String()+":"+data.Port.String(), &ssh.ClientConfig{
		User:            data.Username.String(),
		Auth:            []ssh.AuthMethod{ssh.Password(data.Password.String())},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})

	if err != nil {

	}
	session, err := client.NewSession()

	if err != nil {

	}
	defer session.Close()

	if err := session.Run("yum install nginx -y"); err != nil {

	}

}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &nginxPorvider{
			version: version,
		}
	}
}
