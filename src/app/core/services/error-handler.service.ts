import { Injectable } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';

/**
 * Custom Error Interface
 */
export interface AppError {
  code: string;
  message: string;
  details?: any;
  timestamp: Date;
}

/**
 * Error Handler Service
 * Centralized error handling and logging for the application
 * 
 * Responsibilities:
 * - Handle HTTP errors
 * - Log errors for debugging
 * - Return user-friendly error messages
 * - Track error metrics
 */
@Injectable({
  providedIn: 'root'
})
export class ErrorHandlerService {

  constructor() { }

  /**
   * Handle HTTP errors and return user-friendly message
   * @param error - HttpErrorResponse from failed HTTP request
   * @returns AppError with user-friendly message
   */
  handleHttpError(error: HttpErrorResponse): AppError {
    let appError: AppError;

    if (error.error instanceof ErrorEvent) {
      // Client-side error
      appError = {
        code: 'CLIENT_ERROR',
        message: 'A client-side error occurred. Please check your connection.',
        details: error.error.message,
        timestamp: new Date()
      };
    } else {
      // Server-side error
      const status = error.status;
      const serverMessage = error.error?.error || error.error?.message || error.statusText;

      appError = {
        code: `ERROR_${status}`,
        message: this.getErrorMessage(status, serverMessage),
        details: error.error,
        timestamp: new Date()
      };
    }

    // Log error to console in development mode
    this.logError(appError);

    return appError;
  }

  /**
   * Handle generic errors
   * @param error - Any error object
   * @returns AppError
   */
  handleError(error: any): AppError {
    const appError: AppError = {
      code: 'UNKNOWN_ERROR',
      message: error?.message || 'An unexpected error occurred. Please try again.',
      details: error,
      timestamp: new Date()
    };

    this.logError(appError);
    return appError;
  }

  /**
   * Get user-friendly error message based on HTTP status
   * @param status - HTTP status code
   * @param defaultMessage - Default error message from server
   * @returns User-friendly error message
   */
  private getErrorMessage(status: number, defaultMessage: string): string {
    const errorMessages: Record<number, string> = {
      400: 'Invalid request. Please check your input.',
      401: 'Unauthorized. Please login again.',
      403: 'You do not have permission to perform this action.',
      404: 'Resource not found.',
      409: 'Conflict. The resource already exists or has been modified.',
      422: 'Invalid data provided. Please check your input.',
      500: 'Server error. Please try again later.',
      502: 'Bad gateway. Please try again later.',
      503: 'Service unavailable. Please try again later.',
      504: 'Gateway timeout. Please try again later.'
    };

    return errorMessages[status] || defaultMessage || 'An error occurred. Please try again.';
  }

  /**
   * Log error for debugging and monitoring
   * @param error - AppError to log
   */
  private logError(error: AppError): void {
    // In production, you would send this to a logging service (e.g., Sentry, LogRocket)
    console.error('[APP ERROR]', {
      code: error.code,
      message: error.message,
      details: error.details,
      timestamp: error.timestamp
    });
  }

  /**
   * Validate if error is a specific type
   * @param error - AppError
   * @param code - Error code to check
   * @returns true if error matches code
   */
  isErrorCode(error: AppError, code: string): boolean {
    return error.code === code;
  }

  /**
   * Check if error is a network/connectivity error
   * @param error - AppError
   * @returns true if network error
   */
  isNetworkError(error: AppError): boolean {
    return error.code === 'CLIENT_ERROR' || 
           error.code === 'ERROR_0' || 
           error.code === 'ERROR_503' ||
           error.code === 'ERROR_504';
  }
}
