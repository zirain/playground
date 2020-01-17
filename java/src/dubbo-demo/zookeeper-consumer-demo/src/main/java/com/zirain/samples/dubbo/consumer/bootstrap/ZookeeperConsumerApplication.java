package com.zirain.samples.dubbo.consumer.bootstrap;

import org.apache.dubbo.config.annotation.Reference;
import org.apache.dubbo.config.spring.context.annotation.EnableDubbo;
import org.springframework.boot.ApplicationRunner;
import org.springframework.boot.WebApplicationType;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.context.annotation.Bean;

import com.zirain.samples.dubbo.provider.service.DemoService;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@SpringBootApplication
@EnableDubbo
public class ZookeeperConsumerApplication {

	private final Logger logger = LoggerFactory.getLogger(getClass());

	@Reference(version = "1.0.0")
	private DemoService demoService;

	public static void main(String[] args) {
		new SpringApplicationBuilder(ZookeeperConsumerApplication.class).web(WebApplicationType.NONE).run(args);
	}

	@Bean
	public ApplicationRunner runner() {
		return args -> logger.info(demoService.sayHello());
	}
}
