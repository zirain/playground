package Helloworld.dao;

import Helloworld.entity.JPABook;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;


import java.util.List;

@Mapper
public interface JPABookMapper {
    @Select("select * from JPABook")
    List<JPABook> getAllBooks();

    @Select("select * from JPABook where id = #{id}")
    JPABook getById(long id);
}
