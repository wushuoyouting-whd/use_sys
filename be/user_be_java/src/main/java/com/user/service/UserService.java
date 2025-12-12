package com.user.service;

import com.user.domain.User;
import com.user.dto.PageResponse;
import com.user.dto.UserCreateDTO;
import com.user.dto.UserQueryDTO;
import com.user.dto.UserUpdateDTO;
import com.user.common.exception.AppException;
import com.user.mapper.UserMapper;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.time.LocalDateTime;
import java.util.List;

/**
 * 用户服务类
 */
@Service
@RequiredArgsConstructor
public class UserService {

    @Autowired
    private final UserMapper userMapper;

    @Value("${app.be-type}")
    private String beType;

    /**
     * 分页查询用户列表
     */
    public PageResponse<User> getUserWithPageAndCount(Integer page, Integer limit, UserQueryDTO queryDTO) {
        // 计算偏移量
        int offset = (page - 1) * limit;
        
        // 查询数据
        List<User> users = userMapper.selectUsersWithConditions(
                queryDTO.getName(),
                queryDTO.getEmail(),
                queryDTO.getAge(),
                queryDTO.getBeType(),
                queryDTO.getStartDate(),
                queryDTO.getEndDate(),
                offset,
                limit
        );
        
        // 统计总数
        long total = userMapper.countUsersWithConditions(
                queryDTO.getName(),
                queryDTO.getEmail(),
                queryDTO.getAge(),
                queryDTO.getBeType(),
                queryDTO.getStartDate(),
                queryDTO.getEndDate()
        );
        
        // 计算总页数
        int totalPages = (int) Math.ceil((double) total / limit);
        
        return new PageResponse<>(users, total, page, limit, totalPages);
    }

    /**
     * 根据 ID 获取用户
     */
    public User getUserById(Integer id) {
        User user = userMapper.selectById(id);
        if (user == null) {
            throw new AppException(404, "User not found");
        }
        return user;
    }

    /**
     * 创建用户
     */
    @Transactional
    public User createUser(UserCreateDTO dto) {
        // 检查邮箱是否已存在
        if (userMapper.existsByEmail(dto.getEmail())) {
            throw new AppException(409, "Email already exists");
        }

        User user = new User();
        BeanUtils.copyProperties(dto, user);
        user.setAge(dto.getAge() != null ? dto.getAge() : 0);
        user.setBeType(beType); // 设置后端类型
        user.setCreatedAt(LocalDateTime.now());
        user.setUpdatedAt(LocalDateTime.now());

        userMapper.insert(user);
        return user;
    }

    /**
     * 更新用户
     */
    @Transactional
    public User updateUser(Integer id, UserUpdateDTO dto) {
        User user = userMapper.selectById(id);
        if (user == null) {
            throw new AppException(404, "User not found");
        }

        // 如果邮箱改变，检查新邮箱是否已存在
        if (dto.getEmail() != null && !dto.getEmail().equals(user.getEmail())) {
            if (userMapper.existsByEmail(dto.getEmail())) {
                throw new AppException(409, "email already exists");
            }
        }

        // 更新用户信息
        user.setName(dto.getName());
        user.setEmail(dto.getEmail());
        user.setAge(dto.getAge());
        user.setBirthDate(dto.getBirthDate());
        user.setUpdatedAt(LocalDateTime.now());

        userMapper.updateById(user);
        return user;
    }

    /**
     * 删除用户
     */
    @Transactional
    public void deleteUser(Integer id) {
        User user = userMapper.selectById(id);
        if (user == null) {
            throw new AppException(404, "User not found");
        }
        userMapper.deleteById(id);
    }
}
