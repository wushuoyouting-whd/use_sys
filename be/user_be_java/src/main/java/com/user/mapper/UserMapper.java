package com.user.mapper;

import com.user.domain.User;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;

import java.time.LocalDate;
import java.util.List;

/**
 * 用户 Mapper 接口
 */
@Mapper
public interface UserMapper {
    
    /**
     * 根据ID查询用户
     */
    User selectById(@Param("id") Integer id);
    
    /**
     * 根据邮箱查询用户
     */
    User selectByEmail(@Param("email") String email);
    
    /**
     * 检查邮箱是否存在
     */
    boolean existsByEmail(@Param("email") String email);
    
    /**
     * 分页查询用户列表（带条件）
     */
    List<User> selectUsersWithConditions(
            @Param("name") String name,
            @Param("email") String email,
            @Param("age") Integer age,
            @Param("beType") String beType,
            @Param("startDate") LocalDate startDate,
            @Param("endDate") LocalDate endDate,
            @Param("offset") Integer offset,
            @Param("limit") Integer limit
    );
    
    /**
     * 统计符合条件的用户总数
     */
    long countUsersWithConditions(
            @Param("name") String name,
            @Param("email") String email,
            @Param("age") Integer age,
            @Param("beType") String beType,
            @Param("startDate") LocalDate startDate,
            @Param("endDate") LocalDate endDate
    );
    
    /**
     * 插入用户
     */
    int insert(User user);
    
    /**
     * 更新用户
     */
    int updateById(User user);
    
    /**
     * 根据ID删除用户
     */
    int deleteById(@Param("id") Integer id);
}

