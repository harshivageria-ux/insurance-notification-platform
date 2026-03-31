import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class LoggerService {

  constructor() { }

  info(message: string): void {
    console.log(`[NOTIFICATION_ADMIN_PORTAL] INFO: ${message}`);
  }

  error(message: string): void {
    console.error(`[NOTIFICATION_ADMIN_PORTAL] ERROR: ${message}`);
  }

  debug(message: string): void {
    console.debug(`[NOTIFICATION_ADMIN_PORTAL] DEBUG: ${message}`);
  }

  warn(message: string): void {
    console.warn(`[NOTIFICATION_ADMIN_PORTAL] WARN: ${message}`);
  }
}