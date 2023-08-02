/*
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

// this file is auto-generated with mmv1, any changes made here will be overwritten

import jetbrains.buildServer.configs.kotlin.BuildType
import jetbrains.buildServer.configs.kotlin.Project
import jetbrains.buildServer.configs.kotlin.AbsoluteId

const val providerName = "google"

// Google returns an instance of Project,
// which has multiple build configurations defined within it.
// See https://teamcity.jetbrains.com/app/dsl-documentation/root/project/index.html
fun Google(environment: String, manualVcsRoot: AbsoluteId, branchRef: String, configuration: ClientConfiguration) : Project {
    return Project{

        // Create build configs for each package defined in packages.kt
        var buildConfigs = buildConfigurationsForPackages(packages, providerName, "google", environment, manualVcsRoot, branchRef, configuration)
        buildConfigs.forEach { buildConfiguration ->
            buildType(buildConfiguration)
        }
    }
}

fun buildConfigurationsForPackages(packages: Map<String, String>, providerName : String, path : String, environment: String, manualVcsRoot: AbsoluteId, branchRef: String, config: ClientConfiguration): List<BuildType> {
    var list = ArrayList<BuildType>()

    packages.forEach { (packageName, displayName) ->
        if (packageName == "services") {
            // `services` is a folder containing packages, not a package itself; call buildConfigurationsForPackages to iterate through directories found within `services`
            var serviceList = buildConfigurationsForPackages(services, providerName, path+"/"+packageName, environment, manualVcsRoot, branchRef, config)
            list.addAll(serviceList)
        } else {
            // other folders assumed to be packages
            var triggerConfig = NightlyTriggerConfiguration(environment, branchRef)

            var pkg = packageDetails(packageName, displayName, providerName, environment)
            var buildConfig = pkg.buildConfiguration(path, manualVcsRoot, defaultParallelism, triggerConfig)

            buildConfig.params.ConfigureGoogleSpecificTestParameters(config)

            list.add(buildConfig)
        }
    }

    return list
}

class NightlyTriggerConfiguration(environment: String, branchRef: String, nightlyTestsEnabled: Boolean = true, startHour: Int = defaultStartHour, daysOfWeek: String = defaultDaysOfWeek, daysOfMonth: String = defaultDaysOfMonth) {

    // Default values are used below unless
    // - alternatives passed in as arguments
    // - logic in `init` changes them based on environment
    var branchRef = branchRef
    var nightlyTestsEnabled = nightlyTestsEnabled
    var startHour = startHour
    var daysOfWeek = daysOfWeek
    var daysOfMonth = daysOfMonth

    init {
        // If the environment parameter is set to the value of MAJOR_RELEASE_TESTING, 
        // change the days of week to the day for v5.0.0 feature branch testing
        if (environment == MAJOR_RELEASE_TESTING) {
            this.daysOfWeek = "4" // Thursday for GA
        }
    }

}
