package Helloworld.rest;

import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class UserReditRestController {
    @RequestMapping(value = "/usercredit/{id}")
    public Integer getCreditLevel(@PathVariable String id)
    {
        return 3;
    }
}
