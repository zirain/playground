package Helloworld.restController;

import Helloworld.entity.JPABook;
import Helloworld.service.JPABookService;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
public class BookRestController {

    private final JPABookService jpaBookService;

    public BookRestController(JPABookService jpaBookService) {
        this.jpaBookService = jpaBookService;
    }

    @GetMapping(value = "/book")
    public List<JPABook> getAllBooks()
    {
        return jpaBookService.findAll();
    }

    @GetMapping(value = "/book/{id}")
    public JPABook getBookById(@PathVariable Long id)
    {
        return jpaBookService.findById(id);
    }
}
