rootProject.name = "Shadowrocket"

include(":app")
include(":core")
include(":service")
include(":design")
include(":common")
include(":hideapi")

pluginManagement {
    repositories {
        mavenLocal()
        mavenCentral()
        gradlePluginPortal()
    }
}
