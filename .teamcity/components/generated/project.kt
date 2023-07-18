// this file is auto-generated with mmv1, any changes made here will be overwritten

import jetbrains.buildServer.configs.kotlin.BuildType
import jetbrains.buildServer.configs.kotlin.Project

const val providerName = "google"

fun Google(environment: String, configuration : ClientConfiguration) : Project {
    return Project{
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
            var testConfig = testConfiguration()

            var pkg = packageDetails(packageName, displayName, environment)
            var buildConfig = pkg.buildConfiguration(providerName, path, true, testConfig.startHour, testConfig.parallelism, testConfig.daysOfWeek, testConfig.daysOfMonth)

            buildConfig.params.ConfigureGoogleSpecificTestParameters(config)

            list.add(buildConfig)
        }
    }

    return list
}

class testConfiguration(environment: String, parallelism: Int = defaultParallelism, startHour: Int = defaultStartHour, daysOfWeek: String = defaultDaysOfWeek, daysOfMonth: String = defaultDaysOfMonth) {

    // If the TeamCity project has an environment parameter set to "default",
    // or no environment parameter, use these default values
    if (environment == "default") {
        var parallelism = parallelism
        var startHour = startHour
        var daysOfWeek = daysOfWeek
        var daysOfMonth = daysOfMonth
    }

    // environment parameter set to "major-release-5.0.0" changes the day of week
    if (environment == "major-release-5.0.0") {
        var parallelism = parallelism
        var startHour = startHour
        var daysOfWeek = "4" // Thursday for GA
        var daysOfMonth = daysOfMonth
    }
}
