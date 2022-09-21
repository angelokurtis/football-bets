package br.dev.kurtis.bets;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import java.security.SecureRandom;
import java.util.random.RandomGenerator;

@Configuration
public class RandomGeneratorConfig {
    @Bean
    public RandomGenerator newSecureRandom() {
        return new SecureRandom();
    }
}
