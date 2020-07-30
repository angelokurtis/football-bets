package br.dev.kurtis

import io.micronaut.runtime.Micronaut.*

fun main(args: Array<String>) {
	build()
	    .args(*args)
		.packages("br.dev.kurtis")
		.start()
}

