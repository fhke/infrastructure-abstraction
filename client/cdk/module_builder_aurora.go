package cdk

const auoraModule = "aurora"

type AuroraModuleBuilder struct {
	*GenericModuleBuilder
}

func newAuroraModuleBuilder(name string) *AuroraModuleBuilder {
	ab := &AuroraModuleBuilder{
		GenericModuleBuilder: newGenericModuleBuilder(name, auoraModule),
	}
	ab.setDefaults()
	return ab
}

func (a *AuroraModuleBuilder) InVPC(vpc *VPCModuleBuilder) *AuroraModuleBuilder {
	a.WithVPCID(vpc.GetOutputVPCID())
	a.WithDBSubnetGroupName(vpc.GetOutputDBSubnetGroupName())
	return a
}

func (a *AuroraModuleBuilder) WithName(name string) *AuroraModuleBuilder {
	a.WithInput("name", name)
	return a
}

func (a *AuroraModuleBuilder) WithEngine(engine string) *AuroraModuleBuilder {
	a.WithInput("engine", engine)
	return a
}

func (a *AuroraModuleBuilder) WithMasterUsername(username string) *AuroraModuleBuilder {
	a.WithInput("master_username", username)
	return a
}

func (a *AuroraModuleBuilder) WithEngineVersion(engineVersion string) *AuroraModuleBuilder {
	a.WithInput("engine_version", engineVersion)
	return a
}

func (a *AuroraModuleBuilder) WithVPCID(vpcID string) *AuroraModuleBuilder {
	a.WithInput("vpc_id", vpcID)
	return a
}

func (a *AuroraModuleBuilder) WithDBSubnetGroupName(name string) *AuroraModuleBuilder {
	a.WithInput("db_subnet_group_name", name)
	return a
}

func (a *AuroraModuleBuilder) setDefaults() {
	a.WithInput("engine", "aurora-postgresql")
	a.WithInput("engine_version", "14.5")
}
