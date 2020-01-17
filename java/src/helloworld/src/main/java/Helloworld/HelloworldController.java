package Helloworld.controller;


import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

@Controller
public class HelloworldController {

    @RequestMapping("/")
    public @ResponseBody String say(){
        return "Hello Spring Boot!";
    }

    @RequestMapping("/hello")
    public @ResponseBody String hello(){
        return "Hello Spring Boot2!";
    }
}
