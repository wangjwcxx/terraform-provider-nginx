package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/wangjwcxx/terraform-provider-nginx/pkg"
)

var (
	version string = "dev"
)

func main() {
	opts := providerserver.ServeOpts{Address: "registry.terraform.io/wangjwcxx/terraform-provider-nginx"}

	err := providerserver.Serve(context.Background(), pkg.New(version), opts)

	if err != nil {
		fmt.Println(err)
	}
}
