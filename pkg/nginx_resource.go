package pkg

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type NginxResource struct {
}

func NewNginxResource() resource.Resource {
	return &NginxResource{}
}

func (r *NginxResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {

}

func (r *NginxResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

}

func (r *NginxResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

}

func (r *NginxResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

}

func (r *NginxResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

}

func (r *NginxResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

func (r *NginxResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}
