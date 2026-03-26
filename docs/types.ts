export interface TerraformResource {
  type: string;
  name: string;
  properties: Record<string, any>;
}

export interface TerraformModule {
  source: string;
  version?: string;
  inputs?: Record<string, any>;
}

export interface TerraformPlan {
  resources: TerraformResource[];
  modules: TerraformModule[];
}

export interface TerraformOutput {
  name: string;
  value: any;
  sensitive?: boolean;
}

export interface TerraformState {
  resources: TerraformResource[];
  outputs: TerraformOutput[];
}

export interface TerraformConfig {
  backend: {
    type: string;
    config: Record<string, any>;
  };
  provider: {
    name: string;
    version: string;
    config: Record<string, any>;
  };
  variables: Record<string, any>;
}

export type TerraformAction = 'apply' | 'plan' | 'destroy';

export interface TerraformExecutionOptions {
  action: TerraformAction;
  config: TerraformConfig;
  variables?: Record<string, any>;
}