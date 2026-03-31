export const environment = {
  production: true,
  // Go backend runs on :9000 (see notification-service/cmd/server/main.go).
  // If the portal is started in production build mode, it must still call :9000.
  apiBaseUrl: 'http://localhost:9000'
};
