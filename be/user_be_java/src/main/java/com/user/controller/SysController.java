package com.user.controller;

import com.user.common.response.ApiResponse;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;
import java.util.Map;

/**
 * 系统控制器
 */
@RestController
@RequestMapping("/api/sys")
@RequiredArgsConstructor
@Tag(name = "Sys", description = "系统接口")
public class SysController {

    @Value("${app.be-type}")
    private String beType;

    /**
     * 获取后端类型
     */
    @GetMapping("/be_type")
    @Operation(summary = "获取后端类型")
    public ResponseEntity<ApiResponse<Map<String, String>>> getBeType() {
        Map<String, String> result = new HashMap<>();
        result.put("beType", beType);
        return ResponseEntity.ok(ApiResponse.success(result));
    }
}

