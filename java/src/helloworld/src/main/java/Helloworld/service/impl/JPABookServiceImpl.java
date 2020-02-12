package Helloworld.service.impl;

import Helloworld.dao.JPABookMapper;
import Helloworld.entity.JPABook;
import Helloworld.service.JPABookService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class JPABookServiceImpl implements JPABookService {

    private final JPABookMapper jpaBookMapper;

    public JPABookServiceImpl(JPABookMapper jpaBookMapper) {
        this.jpaBookMapper = jpaBookMapper;
    }


    @Override
    public List<JPABook> findAll() {
        return  jpaBookMapper.getAllBooks();
    }

    @Override
    public JPABook findById(Long id) {
        return jpaBookMapper.getById(id);
    }
}
