package com.zirain.samples.dubbo.provider.bootstrap;

import org.apache.dubbo.config.spring.context.annotation.EnableDubbo;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.WebApplicationType;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;

@SpringBootApplication
@EnableDubbo
public class ZookeeperProviderApplication {

	public static void main(String[] args) {
		//SpringApplication.run(ZookeeperProviderApplication.class, args);
		new SpringApplicationBuilder(ZookeeperProviderApplication.class).web(WebApplicationType.NONE).run(args);
	}
}
