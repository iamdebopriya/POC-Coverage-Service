import type { Coverage, RegisteredService } from './types'

const BASE = '/api'

export async function fetchRegisteredServices(): Promise<RegisteredService[]> {
  const res = await fetch(`${BASE}/registered-services`)
  if (!res.ok) throw new Error('Failed to fetch registered services')
  return res.json()
}

export async function fetchCoverages(params?: { service?: string; from?: string; to?: string }): Promise<Coverage[]> {
  const q = new URLSearchParams()
  if (params?.service) q.set('service', params.service)
  if (params?.from) q.set('from', params.from)
  if (params?.to) q.set('to', params.to)
  
  q.set('offset', String(new Date().getTimezoneOffset()))
  
  const res = await fetch(`${BASE}/coverage${q.toString() ? '?' + q : ''}`)
  if (!res.ok) throw new Error('Failed to fetch coverages')
  return res.json()
}

export async function fetchCoverageServices(): Promise<string[]> {
  const res = await fetch(`${BASE}/coverage/services`)
  if (!res.ok) throw new Error('Failed to fetch services')
  return res.json()
}

export function downloadCoverageUrl(params?: { service?: string; from?: string; to?: string }) {
  const q = new URLSearchParams()
  if (params?.service) q.set('service', params.service)
  if (params?.from) q.set('from', params.from)
  if (params?.to) q.set('to', params.to)
  return `${BASE}/coverage/download${q.toString() ? '?' + q : ''}`
}

// Opens an EventSource SSE connection for streaming test output.
// onLine is called for each line, onDone when test run is saved, onError on failure.
export function runServiceTests(
  serviceName: string,
  onLine: (line: string) => void,
  onDone: () => void,
  onError: (msg: string) => void,
): EventSource {
  const es = new EventSource(`${BASE}/run/${encodeURIComponent(serviceName)}`)

  es.onmessage = (e) => {
    onLine(e.data)
  }

  es.addEventListener('done', () => {
    es.close()
    onDone()
  })

  es.addEventListener('error', (e: Event) => {
    const me = e as MessageEvent
    es.close()
    onError(me.data ?? 'Test run failed')
  })

  es.onerror = () => {
    es.close()
    onDone() // connection closed after stream ends = done
  }

  return es
}
