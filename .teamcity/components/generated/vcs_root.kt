// this file is auto-generated with mmv1, any changes made here will be overwritten

import jetbrains.buildServer.configs.kotlin.vcs.GitVcsRoot


fun getProviderRepository(branch: String) {

    return object providerRepository : GitVcsRoot({
        name = "terraform-provider-google"
        url = "https://github.com/hashicorp/terraform-provider-google.git"
        agentCleanPolicy = AgentCleanPolicy.ON_BRANCH_CHANGE
        agentCleanFilesPolicy = AgentCleanFilesPolicy.ALL_UNTRACKED
        branchSpec = "+:*"
        branch = branch
        authMethod = anonymous()
    })
}
