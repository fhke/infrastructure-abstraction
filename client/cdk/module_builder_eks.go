package cdk

const eksModule = "eks"

type EKSModuleBuilder struct {
	*GenericModuleBuilder
}

func newEKSModuleBuilder(name string) *EKSModuleBuilder {
	ab := &EKSModuleBuilder{
		GenericModuleBuilder: newGenericModuleBuilder(name, eksModule),
	}
	ab.setDefaults()
	return ab
}

func (a *EKSModuleBuilder) InVPC(vpc *VPCModuleBuilder) *EKSModuleBuilder {
	a.WithVPCID(vpc.GetOutputVPCID())
	a.WithSubnetIDs(vpc.GetOutputPrivateSubnetIDs())
	return a
}

func (e *EKSModuleBuilder) WithClusterName(name string) *EKSModuleBuilder {
	e.WithInput("cluster_name", name)
	return e
}

func (e *EKSModuleBuilder) WithClusterEndpointPublicAccess(enabled bool) *EKSModuleBuilder {
	e.WithInput("cluster_endpoint_public_access", enabled)
	return e
}

func (e *EKSModuleBuilder) WithClusterEndpointPrivateAccess(enabled bool) *EKSModuleBuilder {
	e.WithInput("cluster_endpoint_private_access", enabled)
	return e
}

func (e *EKSModuleBuilder) WithVPCID(id string) *EKSModuleBuilder {
	e.WithInput("vpc_id", id)
	return e
}

func (e *EKSModuleBuilder) WithSubnetIDs(subnetIDs any) *EKSModuleBuilder {
	e.WithInput("subnet_ids", subnetIDs)
	return e
}

func (e *EKSModuleBuilder) setDefaults() {
	e.WithInput("cluster_endpoint_public_access", false)
}
