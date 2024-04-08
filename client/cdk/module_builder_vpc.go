package cdk

const vpcModule = "vpc"

type VPCModuleBuilder struct {
	*GenericModuleBuilder
}

func newVPCModuleBuilder(name string) *VPCModuleBuilder {
	ab := &VPCModuleBuilder{
		GenericModuleBuilder: newGenericModuleBuilder(name, vpcModule),
	}
	ab.setDefaults()
	return ab
}

/*
	Inputs
*/

func (e *VPCModuleBuilder) WithName(name string) *VPCModuleBuilder {
	e.WithInput("name", name)
	return e
}

func (e *VPCModuleBuilder) WithAZs(azs ...string) *VPCModuleBuilder {
	e.WithInput("azs", azs)
	return e
}

func (e *VPCModuleBuilder) WithPrivateSubnets(cidrs ...string) *VPCModuleBuilder {
	e.WithInput("private_subnets", cidrs)
	return e
}

func (e *VPCModuleBuilder) WithDatabaseSubnets(cidrs ...string) *VPCModuleBuilder {
	e.WithInput("database_subnets", cidrs)
	return e
}

func (e *VPCModuleBuilder) WithPublicSubnets(cidrs ...string) *VPCModuleBuilder {
	e.WithInput("public_subnets", cidrs)
	return e
}

func (e *VPCModuleBuilder) WithCIDR(cidr string) *VPCModuleBuilder {
	e.WithInput("cidr", cidr)
	return e
}

/*
Outputs
*/
func (e *VPCModuleBuilder) GetOutputVPCID() string {
	return e.GetOutput("vpc_id")
}

func (e *VPCModuleBuilder) GetOutputPrivateSubnetARNs() string {
	return e.GetOutput("private_subnet_arns")
}

func (e *VPCModuleBuilder) GetOutputPrivateSubnetIDs() string {
	return tfRefARNsToIDs(e.GetOutputRef("private_subnet_arns"))
}

func (e *VPCModuleBuilder) GetOutputDBSubnetGroupName() string {
	return e.GetOutput("database_subnet_group_name")
}

/*
Helpers
*/
func (e *VPCModuleBuilder) setDefaults() {
	e.WithInput("create_database_subnet_group", true)
}
