package Helloworld.service.impl;

import Helloworld.entity.JPABook;
import Helloworld.repository.JPABookRepository;
import Helloworld.service.JPABookService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class JPABookServiceImpl implements JPABookService {

    private final JPABookRepository jpaBookRepository;

    public JPABookServiceImpl(JPABookRepository jpaBookRepository) {
        this.jpaBookRepository = jpaBookRepository;
    }


    @Override
    public List<JPABook> findAll() {
        return  jpaBookRepository.findAll();
    }

    @Override
    public JPABook findById(Long id) {
        return jpaBookRepository.getOne(id);
    }
}
