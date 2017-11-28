ext.checkProperties = { Map<String, String> prop1, Map prop2, String entorno ->
    prop1.findAll({ e -> !e.getKey().toString().contains("snapshot") })
            .each { key, value ->
        if (!prop2.containsKey(key)) {
            def msg = sprintf('Error en checkeo de propiedad de entorno. No existe la propiedad %1$s en el entorno %2$s', key, entorno)
            throw new GradleException(msg)
        } else if (prop1.get(key).toString().isEmpty() || prop2.get(key).toString().isEmpty()) {
            if (project.ext.properties.containsKey(key)) {
                if (prop1.get(key).toString().isEmpty()) prop1[key] = project.ext.properties.get(key)
                if (prop2.get(key).toString().isEmpty()) prop2[key] = project.ext.properties.get(key)
            } else {
                def msg = sprintf('Error en checkeo de propiedad de entorno. La propiedad %1$s esta vacia', key)
                throw new GradleException(msg)
            }

        }
    }
}

ext.addProperties = { Map<String, String> prop1, Map prop2, Map prop3 ->
    def addCommitId = "git rev-parse HEAD".execute().text.trim()
    def addCommitLog = "git log -1".execute().text.trim()
    println "Commit Id con valor $addCommitId"
    println "Commit Info con valor $addCommitLog"

    prop1.each { key, value ->

        Properties configProperties = project.ext.properties;
        configProperties.put("application.info.app.ImplementationTitle", project.name)
        configProperties.put("application.info.app.Version", rootProject.version)
        configProperties.put("application.info.app.Group", rootProject.group)
        configProperties.put("application.info.app.BuiltDate", new Date())
        configProperties.put("application.info.app.BuiltJDK", System.getProperty('java.version'))
        configProperties.put("application.info.app.BuiltGradle", gradle.gradleVersion)
        configProperties.put("application.info.app.CommitId", addCommitId)
        configProperties.put("application.info.app.CommitLog", addCommitLog)
//        configProperties.put("application.spring.jmx.default-domain", 'agentservice')
//        configProperties.put("application.endpoints.jmx.domain", 'agentservice')
//        configProperties.put("application.endpoints.jmx.unique-names", true)

        if (configProperties.containsKey(key)) {
            def valueProp = configProperties.get(key)
            println 'sustituido property ' + key
            prop1[key] = valueProp
            prop2[key] = valueProp
            prop3[key] = valueProp
        }

    }
}

ext.propiedadesDeEntorno = {
    ext.env = project.hasProperty('env') ? project.getProperty('env') : 'local'
    println "Cargando configuracion para entorno '$env."
    def configFile = file("$rootDir/gradle/config/buildConfig.groovy")

    HashMap configMap = [
            'local': new ConfigSlurper("local").parse(configFile.toURL()).flatten(),
            'dev'  : new ConfigSlurper("dev").parse(configFile.toURL()).flatten(),
            'prod' : new ConfigSlurper("prod").parse(configFile.toURL()).flatten()
    ]

    def localMap = configMap.get('local')
    def devMap = configMap.get('dev')
    def prodMap = configMap.get('prod')
    addProperties(localMap, devMap, prodMap)

    checkProperties(localMap, devMap, "dev")
    checkProperties(localMap, prodMap, "prod")
    return configMap.get(env)
}