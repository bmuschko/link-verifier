package com.bmuschko.linkverifier

import org.gradle.api.DefaultTask
import org.gradle.api.Project
import org.gradle.api.tasks.TaskAction
import org.gradle.api.tasks.InputDirectory
import org.gradle.api.tasks.OutputDirectory

class Distributions extends DefaultTask {
    @InputDirectory
    File inputDir = project.file('.gogradle')
    
    @OutputDirectory
    File outputDir = project.file("$project.buildDir/distributions")
    
    Distributions() {
        group = 'GoGradle'
        description = 'Builds packaged distributions for all prebuilt binaries.'
    }
    
    @TaskAction
    void create() {
        inputDir.listFiles(new CrossCompiledFileFilter(project)).each { preBuiltLib ->
            if (preBuiltLib.name.contains('windows')) {
                ant.zip(destfile: "${outputDir}/${preBuiltLib.name}.zip") {
                    fileset(dir: inputDir) {
                        include(name: preBuiltLib.name)
                    }
                }
            } else {
                def tarFile = "${outputDir}/${preBuiltLib.name}.tar"
                ant.tar(destfile: tarFile) {
                    tarfileset(dir: inputDir) {
                        include(name: preBuiltLib.name)
                    }
                }
                ant.gzip(destfile: "${tarFile}.gz", src: tarFile)
                project.delete(tarFile)
            }
        }
    }
    
    private static class CrossCompiledFileFilter implements FilenameFilter {
        private final Project project
        
        CrossCompiledFileFilter(Project project) {
            this.project = project
        }
        
        boolean accept(File f, String filename) {
            filename.endsWith("_${project.name}")
        }  
    } 
}