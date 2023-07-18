// this file is auto-generated with mmv1, any changes made here will be overwritten

import jetbrains.buildServer.configs.kotlin.BuildType
import jetbrains.buildServer.configs.kotlin.Project
import jetbrains.buildServer.configs.kotlin.vcs.GitVcsRoot

const val providerName = "google"

fun Google(environment: String, branchRef: String, configuration : ClientConfiguration) : Project {
    val gitVcsRoot = getProviderRepository(branchRef)
    return Project{
        vcsRoot(gitVcsRoot)

        var buildConfigs = buildConfigurationsForPackages(packages, providerName, "google", environment, branchRef, root, configuration)
        buildConfigs.forEach { buildConfiguration ->
            buildType(buildConfiguration)
        }
    }
}

fun buildConfigurationsForPackages(packages: Map<String, String>, providerName : String, path : String, environment: String, branchRef: String, gitVcsRoot GitVcsRoot, config : ClientConfiguration): List<BuildType> {
    var list = ArrayList<BuildType>()

    packages.forEach { (packageName, displayName) ->
        if (packageName == "services") {
            var serviceList = buildConfigurationsForPackages(services, providerName, path+"/"+packageName, environment, branchRef, gitVcsRoot, config)
            list.addAll(serviceList)
        } else {
            var testConfig = testConfiguration(environment)

            var pkg = packageDetails(packageName, displayName, environment, branchRef)
            var buildConfig = pkg.buildConfiguration(providerName, path, true, testConfig.startHour, testConfig.parallelism, testConfig.daysOfWeek, testConfig.daysOfMonth, gitVcsRoot)

            buildConfig.params.ConfigureGoogleSpecificTestParameters(config)

            list.add(buildConfig)
        }
    }

    return list
}

class testConfiguration(environment: String, parallelism: Int = defaultParallelism, startHour: Int = defaultStartHour, daysOfWeek: String = defaultDaysOfWeek, daysOfMonth: String = defaultDaysOfMonth) {

    // Default values are present if init doesn't change them
    var parallelism = parallelism
    var startHour = startHour
    var daysOfWeek = daysOfWeek
    var daysOfMonth = daysOfMonth

    init {
        // environment parameter set to "majorRelease5.0.0" changes the day of week
        if (environment == "majorRelease500") {
            this.parallelism = parallelism
            this.startHour = startHour
            this.daysOfWeek = "4" // Thursday for GA
            this.daysOfMonth = daysOfMonth
        }
    }

}
