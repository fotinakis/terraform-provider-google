// this file is auto-generated with mmv1, any changes made here will be overwritten

import jetbrains.buildServer.configs.kotlin.vcs.GitVcsRoot
import jetbrains.buildServer.configs.kotlin.vcs.GitVcsRoot.AgentCleanPolicy.ON_BRANCH_CHANGE
import jetbrains.buildServer.configs.kotlin.vcs.GitVcsRoot.AgentCleanFilesPolicy.ALL_UNTRACKED


fun getProviderRepository(branchRef: String) : GitVcsRoot{
    val repoName = "terraform-provider-google"
    val vcsId = repoName + branchRef
    return GitVcsRoot {
        id(vcsId)
        name = repoName
        url = "https://github.com/hashicorp/terraform-provider-google.git"
        agentCleanPolicy = ON_BRANCH_CHANGE
        agentCleanFilesPolicy = ALL_UNTRACKED
        branchSpec = "+:*"
        branch = branchRef
        authMethod = anonymous()
    }
}
