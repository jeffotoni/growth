val ktor_version: String by project
val kotlin_version: String by project
val logback_version: String by project
val ktorm_version: String by project
val h2_version: String by project

plugins {
    application
    kotlin("jvm") version "1.5.21"
}

group = "com.growth"
version = "0.0.1"
application {
    mainClass.set("com.growth.ApplicationKt")
}

repositories {
    mavenCentral()
}

dependencies {
    implementation("io.ktor:ktor-server-core:$ktor_version")
    implementation("io.ktor:ktor-locations:$ktor_version")
    implementation("io.ktor:ktor-jackson:$ktor_version")
    implementation("io.ktor:ktor-server-netty:$ktor_version")
    implementation("org.ktorm:ktorm-core:${ktorm_version}")
    implementation("com.h2database:h2:${h2_version}")
    implementation("ch.qos.logback:logback-classic:$logback_version")
    testImplementation("io.ktor:ktor-server-tests:$ktor_version")
    testImplementation("org.jetbrains.kotlin:kotlin-test:$kotlin_version")
}