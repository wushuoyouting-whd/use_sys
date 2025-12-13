package com.user.dto;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

/**
 * 分页响应 DTO
 * DTO 的全称是 Data Transfer Object（数据传输对象）
 */
@Data
@NoArgsConstructor
@AllArgsConstructor
public class PageResponse<T> {
    private List<T> data;
    private Long total;
    private Integer page;
    private Integer limit;
    private Integer totalPages;
}


