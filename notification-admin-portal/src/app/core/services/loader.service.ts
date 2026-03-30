import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';

/**
 * Loader Service
 * Global loading state management using RxJS BehaviorSubject
 * 
 * Usage:
 * - Subscribe to isLoading$ to show/hide loader in UI
 * - Call show() when starting async operations
 * - Call hide() when async operations complete
 * - Track multiple operations with counter
 */
@Injectable({
  providedIn: 'root'
})
export class LoaderService {

  private loadingCounter = 0;
  private isLoadingSubject = new BehaviorSubject<boolean>(false);

  /**
   * Observable stream of loading state
   * Subscribe to this in your UI components
   */
  public isLoading$: Observable<boolean> = this.isLoadingSubject.asObservable();

  constructor() { }

  /**
   * Show loader (increment counter)
   * Useful when multiple async operations are running
   */
  show(): void {
    this.loadingCounter++;
    this.updateLoadingState();
  }

  /**
   * Hide loader (decrement counter)
   * Only hide when all async operations are complete
   */
  hide(): void {
    if (this.loadingCounter > 0) {
      this.loadingCounter--;
    }
    this.updateLoadingState();
  }

  /**
   * Forcefully show loader
   */
  forceShow(): void {
    this.loadingCounter = 1;
    this.updateLoadingState();
  }

  /**
   * Forcefully hide loader
   */
  forceHide(): void {
    this.loadingCounter = 0;
    this.updateLoadingState();
  }

  /**
   * Get current loading state synchronously
   * Use only when necessary; prefer isLoading$ observable
   */
  isLoading(): boolean {
    return this.isLoadingSubject.value;
  }

  /**
   * Update loading state based on counter
   */
  private updateLoadingState(): void {
    this.isLoadingSubject.next(this.loadingCounter > 0);
  }

  /**
   * Reset loading state (useful for cleanup)
   */
  reset(): void {
    this.loadingCounter = 0;
    this.updateLoadingState();
  }
}
