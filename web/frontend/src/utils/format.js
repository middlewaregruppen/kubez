/**
 * Formatting utilities — replaces the Vue 2 filters that were
 * used in CGroup.vue and ApiEndpoint.vue.
 */

/**
 * Format a nanosecond duration into a human-readable string.
 */
export function prettyTime(num) {
  if (typeof num !== 'number' || isNaN(num)) return num
  if (num > Math.pow(10, 9) * 3) return (num / Math.pow(10, 9)).toFixed(1) + ' s'
  if (num > Math.pow(10, 6)) return (num / Math.pow(10, 6)).toFixed(0) + ' ms'
  if (num > Math.pow(10, 3)) return (num / Math.pow(10, 3)).toFixed(0) + ' µs'
  return num + ' ns'
}

/**
 * Convert microseconds to nanoseconds.
 */
export function us2ns(num) {
  return num * 1000
}

/**
 * Format a byte count into a human-readable size string (no /s suffix).
 * Used for memory sizes in CGroup.vue.
 */
export function prettyBytes(num) {
  if (typeof num !== 'number' || isNaN(num)) return num
  const units = ['B', 'kB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']
  const neg = num < 0
  if (neg) num = -num
  if (num < 1) return (neg ? '-' : '') + num + ' B'
  const exponent = Math.min(Math.floor(Math.log(num) / Math.log(1000)), units.length - 1)
  const formatted = (num / Math.pow(1000, exponent)).toFixed(2) * 1
  return (neg ? '-' : '') + formatted + ' ' + units[exponent]
}

/**
 * Format a byte-per-second rate into a human-readable string (with /s suffix).
 * Returns 'no' for zero. Used for request/response rates in ApiEndpoint.vue.
 */
export function prettyBytesRate(num) {
  num = parseInt(num)
  if (typeof num !== 'number' || isNaN(num)) return num
  if (num === 0) return 'no'
  const units = ['B', 'kB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']
  const neg = num < 0
  if (neg) num = -num
  if (num < 1) return (neg ? '-' : '') + num + ' B/s'
  const exponent = Math.min(Math.floor(Math.log(num) / Math.log(1000)), units.length - 1)
  const formatted = (num / Math.pow(1000, exponent)).toFixed(2) * 1
  return (neg ? '-' : '') + formatted + ' ' + units[exponent] + '/s'
}
