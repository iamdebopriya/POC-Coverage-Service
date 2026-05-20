export interface Coverage {
  id: number
  service_name: string
  backend_coverage: number
  frontend_coverage: number
  total_tests: number
  passed_tests: number
  failed_tests: number
  flaky_tests: number
  avg_execution_time: number
  timestamp: string
}

export interface RegisteredService {
  name: string
  display_name: string
  backend_path: string
  frontend_path: string
  backend_type: string
  frontend_type: string
}
