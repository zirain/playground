package Helloworld.entity;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

import javax.persistence.*;


@Entity
@Table(name = "JPABook")
@JsonIgnoreProperties(value = { "hibernateLazyInitializer"})
public class JPABook {
    /**
     * id
     */
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    /**
     * 书名
     */
    @Column(nullable = false)
    private String name;

    /**
     * 作者
     */
    @Column(nullable = false)
    private String writer;

    /**
     * 简介
     */
    @Column(nullable = false)
    private String introduction;

    /**
     * 空参构造
     */
    public JPABook() {
    }

    /**
     * 有参构造
     */
    public JPABook(String name, String writer, String introduction) {
        this.name = name;
        this.writer = writer;
        this.introduction = introduction;
    }


    public String getWriter() {
        return writer;
    }

    public void setWriter(String writer) {
        this.writer = writer;
    }

    public String getIntroduction() {
        return introduction;
    }

    public void setIntroduction(String introduction) {
        this.introduction = introduction;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }
}
