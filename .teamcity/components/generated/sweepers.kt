import jetbrains.buildServer.configs.kotlin.BuildType
import jetbrains.buildServer.configs.kotlin.AbsoluteId

val preSweeperBuildConfigId = "PRE_SWEEPER_TESTS"
val postSweeperBuildConfigId = "POST_SWEEPER_TESTS"

class sweeperBuildConfigs(packageName: String, providerName: String, environment: String, branchRef: String, clientConfig: ClientConfiguration) {
    val packageName = packageName
    val providerName = providerName
    val environment = environment
    val branchRef = branchRef
    val sweeperRegions = "us-central1"

    fun preSweeperBuildConfig(path: String, manualVcsRoot: AbsoluteId, parallelism: Int, triggerConfig: NightlyTriggerConfiguration) : BuildType {
        val testPrefix = "TestAcc"
        val timeout = "12"
        val sweeperRegions = "us-central1"
        val sweeperRun = "" // Empty string means all sweepers run

        val configName = "Pre-Sweeper"
        val sweeperStepName = "Pre-Sweeper"

        return createBuildConfig(manualVcsRoot, preSweeperBuildConfigId, configName, sweeperStepName, parallelism, testPrefix, timeout, sweeperRegions, sweeperRun, path, packageName, triggerConfig)
   }

    fun postSweeperBuildConfig(providerName: String, path: String, manualVcsRoot: AbsoluteId, parallelism: Int, triggerConfig: NightlyTriggerConfiguration) : BuildType {
        val testPrefix = "TestAcc"
        val timeout = "12"
        val sweeperRegions = "us-central1"
        val sweeperRun = "" // Empty string means all sweepers run

        val configName = "Post-Sweeper"
        val sweeperStepName = "Post-Sweeper"

        return createBuildConfig(manualVcsRoot, preSweeperBuildConfigId, configName, sweeperStepName, parallelism, testPrefix, timeout, sweeperRegions, sweeperRun, path, packageName, triggerConfig)
    }

    fun createBuildConfig(
        manualVcsRoot: AbsoluteId,
        configId: String,
        configName: String,
        sweeperStepName: String,
        parallelism: Int,
        testPrefix: String,
        timeout: String,
        sweeperRegions: String,
        sweeperRun: String,
        path: String,
        packageName: String,
        triggerConfig: NightlyTriggerConfiguration
        ) : BuildType {
        return BuildType {

            id(configId)

            name = configName

            vcs {
                root(rootId = manualVcsRoot)
                cleanCheckout = true
            }

            steps {
                SetGitCommitBuildId()
                ConfigureGoEnv()
                DownloadTerraformBinary()
                // RunSweepers(sweeperStepName)
            }

            failureConditions {
                errorMessage = true
            }

            features {
                Golang()
            }

            params {
                TerraformAcceptanceTestParameters(parallelism, testPrefix, timeout, sweeperRegions, sweeperRun)
                TerraformAcceptanceTestsFlag()
                TerraformCoreBinaryTesting()
                TerraformShouldPanicForSchemaErrors()
                ReadOnlySettings()
                WorkingDirectory(path, packageName)
            }

            triggers {
                RunNightly(triggerConfig)
            }
        }
    }
}
