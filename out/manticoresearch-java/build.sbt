lazy val root = (project in file(".")).
  settings(
    organization := "com.manticoresearch",
    name := "manticoresearch",
    version := "9.0.0",
    scalaVersion := "2.11.4",
    scalacOptions ++= Seq("-feature"),
    Compile / javacOptions ++= Seq("-Xlint:deprecation"),
    Compile / packageDoc / publishArtifact := false,
    resolvers += Resolver.mavenLocal,
    libraryDependencies ++= Seq(
      "com.google.code.findbugs" % "jsr305" % "3.0.0",
      "io.swagger" % "swagger-annotations" % "1.6.5",
      "org.glassfish.jersey.core" % "jersey-client" % "3.0.4",
      "org.glassfish.jersey.inject" % "jersey-hk2" % "3.0.4",
      "org.glassfish.jersey.media" % "jersey-media-multipart" % "3.0.4",
      "org.glassfish.jersey.media" % "jersey-media-json-jackson" % "3.0.4",
      "org.glassfish.jersey.connectors" % "jersey-apache-connector" % "3.0.4",
      "com.fasterxml.jackson.core" % "jackson-core" % "2.17.1" % "compile",
      "com.fasterxml.jackson.core" % "jackson-annotations" % "2.17.1" % "compile",
      "com.fasterxml.jackson.core" % "jackson-databind" % "2.17.1" % "compile",
      "com.fasterxml.jackson.datatype" % "jackson-datatype-jsr310" % "2.17.1" % "compile",
      "org.openapitools" % "jackson-databind-nullable" % "0.2.6" % "compile",
      "jakarta.annotation" % "jakarta.annotation-api" % "2.1.0" % "compile",
      "org.junit.jupiter" % "junit-jupiter-api" % "5.8.2" % "test"
    )
  )
