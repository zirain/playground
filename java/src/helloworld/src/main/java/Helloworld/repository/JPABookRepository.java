package Helloworld.repository;

import Helloworld.entity.JPABook;
import org.springframework.data.jpa.repository.JpaRepository;

public interface JPABookRepository extends JpaRepository<JPABook,Long> {
}
