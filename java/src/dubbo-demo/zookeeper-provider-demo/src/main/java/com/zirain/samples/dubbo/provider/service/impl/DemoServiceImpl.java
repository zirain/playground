package com.zirain.samples.dubbo.provider.service.impl;

import org.apache.dubbo.config.annotation.Service;

import com.zirain.samples.dubbo.provider.service.DemoService;

@Service(version = "${dubbo.version}")
public class DemoServiceImpl implements DemoService {

	@Override
	public String sayHello() {
		// TODO Auto-generated method stub
		return "Hello Dubbo!";
	}

}
