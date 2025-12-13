package com.user.dto;

import lombok.Data;

import java.time.LocalDate;

/**
 * 用户查询条件 DTO
 */
@Data
public class UserQueryDTO {
    private String name;
    private String email;
    private Integer age;
    private String beType;
    private LocalDate startDate;
    private LocalDate endDate;
}


