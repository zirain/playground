package Helloworld.service;

import Helloworld.entity.JPABook;

import java.util.List;

public interface JPABookService {

    List<JPABook> findAll();

    JPABook findById(Long id);
}
