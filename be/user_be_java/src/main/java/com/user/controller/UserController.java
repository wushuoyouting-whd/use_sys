package com.user.controller;

import com.user.common.response.ApiResponse;
import com.user.domain.User;
import com.user.dto.PageResponse;
import com.user.dto.UserCreateDTO;
import com.user.dto.UserQueryDTO;
import com.user.dto.UserUpdateDTO;
import com.user.service.UserService;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

/**
 * 用户控制器
 */
@RestController
@RequestMapping("/api/users")
@RequiredArgsConstructor
@Tag(name = "Users", description = "用户管理接口")
public class UserController {

    private final UserService userService;

    /**
     * 获取用户列表（分页）
     */
    @GetMapping
    @Operation(summary = "获取用户列表", description = "支持分页和条件查询")
    public ResponseEntity<ApiResponse<PageResponse<User>>> list(
            @RequestParam(defaultValue = "1") Integer page,
            @RequestParam(defaultValue = "10") Integer limit,
            @RequestParam(required = false) String name,
            @RequestParam(required = false) String email,
            @RequestParam(required = false) Integer age,
            @RequestParam(required = false) String beType,
            @RequestParam(required = false) String startDate,
            @RequestParam(required = false) String endDate
    ) {

        UserQueryDTO queryDTO = new UserQueryDTO();
        queryDTO.setName(name);
        queryDTO.setEmail(email);
        queryDTO.setAge(age);
        queryDTO.setBeType(beType);
        if (startDate != null) {
            queryDTO.setStartDate(java.time.LocalDate.parse(startDate));
        }
        if (endDate != null) {
            queryDTO.setEndDate(java.time.LocalDate.parse(endDate));
        }

        PageResponse<User> result = userService.getUserWithPageAndCount(page, limit, queryDTO);
        System.out.println("进来了=="+result);
        return ResponseEntity.ok(ApiResponse.success(result));
    }

    /**
     * 获取用户详情
     */
    @GetMapping("/{id}")
    @Operation(summary = "获取用户详情")
    public ResponseEntity<ApiResponse<User>> getById(@PathVariable Integer id) {
        User user = userService.getUserById(id);
        return ResponseEntity.ok(ApiResponse.success(user));
    }

    /**
     * 创建用户
     */
    @PostMapping
    @Operation(summary = "创建用户")
    public ResponseEntity<ApiResponse<User>> create(@Valid @RequestBody UserCreateDTO dto) {
        User user = userService.createUser(dto);
        return ResponseEntity.ok(ApiResponse.success(user, "user created success"));
    }

    /**
     * 更新用户
     */
    @PutMapping("/{id}")
    @Operation(summary = "更新用户")
    public ResponseEntity<ApiResponse<User>> update(
            @PathVariable Integer id,
            @Valid @RequestBody UserUpdateDTO dto
    ) {
        User user = userService.updateUser(id, dto);
        return ResponseEntity.ok(ApiResponse.success(user, "user updated success"));
    }

    /**
     * 删除用户
     */
    @DeleteMapping("/{id}")
    @Operation(summary = "删除用户")
    public ResponseEntity<ApiResponse<Object>> delete(@PathVariable Integer id) {
        userService.deleteUser(id);
        return ResponseEntity.ok(ApiResponse.success(null, "user deleted success"));
    }
}

