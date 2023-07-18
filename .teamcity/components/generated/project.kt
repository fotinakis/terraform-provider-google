// this file is auto-generated with mmv1, any changes made here will be overwritten

import jetbrains.buildServer.configs.kotlin.BuildType
import jetbrains.buildServer.configs.kotlin.Project

const val providerName = "google"

fun Google(environment: String, branch: String, configuration : ClientConfiguration) : Project {
    return Project{
        var providerRepository = getProviderRepository(branch)
        vcsRoot(providerRepository)

        var buildConfigs = buildConfigurationsForPackages(packages, providerName, "google", environment, configuration)
        buildConfigs.forEach { buildConfiguration ->
            buildType(buildConfiguration)
        }
    }
}

fun buildConfigurationsForPackages(packages: Map<String, String>, providerName : String, path : String, environment: String, config : ClientConfiguration): List<BuildType> {
    var list = ArrayList<BuildType>()

    packages.forEach { (packageName, displayName) ->
        if (packageName == "services") {
            var serviceList = buildConfigurationsForPackages(services, providerName, path+"/"+packageName, environment, config)
            list.addAll(serviceList)
        } else {
            var testConfig = testConfiguration(environment)

            var pkg = packageDetails(packageName, displayName, environment, branch)
            var buildConfig = pkg.buildConfiguration(providerName, path, true, testConfig.startHour, testConfig.parallelism, testConfig.daysOfWeek, testConfig.daysOfMonth)

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
