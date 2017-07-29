package com.bmuschko.linkverifier

import org.gradle.api.DefaultTask
import org.gradle.api.tasks.TaskAction
import org.gradle.api.tasks.InputDirectory
import org.gradle.api.tasks.OutputFile

class AggregateCoverage extends DefaultTask {
    @InputDirectory
    File inputDir = project.file('.gogradle/reports/coverage/profiles')
    
    @OutputFile
    File outputFile = project.file("$project.buildDir/reports/coverage-aggregate/coverage.txt")
    
    AggregateCoverage() {
        group = 'GoGradle'
        description = 'Aggregates coverage profile data.'
    }
    
    @TaskAction
    void aggregate() {
        StringBuilder aggregatedCoverage = new StringBuilder()
        
        inputDir.listFiles().each {
            aggregatedCoverage << it.text
        }
        
        outputFile.text = aggregatedCoverage.toString()
    }
}