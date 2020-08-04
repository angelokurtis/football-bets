package br.dev.kurtis

import io.micronaut.context.annotation.ConfigurationProperties

@ConfigurationProperties("championships")
data class ChampionshipConfiguration(var directory: String? = null)