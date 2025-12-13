package com.user.common.exception;

import lombok.Getter;

/**
 * 自定义应用异常
 */
@Getter
public class AppException extends RuntimeException {
    private final Integer code;

    public AppException(Integer code, String message) {
        super(message);
        this.code = code;
    }
}


