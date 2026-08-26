import { describe, expect, it } from 'vitest'
import { prettyBytes, prettyBytesRate, prettyTime, us2ns } from '@/utils/format'

describe('format utilities', () => {
  it('formats nanosecond durations at representative scales', () => {
    expect(prettyTime(4500000000)).toBe('4.5 s')
    expect(prettyTime(2500000)).toBe('3 ms')
    expect(prettyTime(12500)).toBe('13 µs')
    expect(prettyTime(500)).toBe('500 ns')
  })

  it('leaves invalid duration values unchanged', () => {
    expect(prettyTime('1000')).toBe('1000')
    expect(prettyTime(Number.NaN)).toBeNaN()
  })

  it('converts microseconds to nanoseconds', () => {
    expect(us2ns(2.5)).toBe(2500)
    expect(us2ns(0)).toBe(0)
  })

  it('formats byte counts using decimal units', () => {
    expect(prettyBytes(999)).toBe('999 B')
    expect(prettyBytes(1500)).toBe('1.5 kB')
    expect(prettyBytes(2500000)).toBe('2.5 MB')
  })

  it('preserves signs and sub-byte values', () => {
    expect(prettyBytes(-1500)).toBe('-1.5 kB')
    expect(prettyBytes(0.5)).toBe('0.5 B')
  })

  it('formats byte rates and accepts numeric strings', () => {
    expect(prettyBytesRate(0)).toBe('no')
    expect(prettyBytesRate('2500')).toBe('2.5 kB/s')
    expect(prettyBytesRate(-1000)).toBe('-1 kB/s')
  })

  it('reports an invalid byte rate as NaN', () => {
    expect(prettyBytesRate('not-a-rate')).toBeNaN()
  })
})
