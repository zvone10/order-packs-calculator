/**
 * Safely join URL parts by removing double slashes
 * @param base The base URL (can have trailing slash or not)
 * @param path The path to append (can have leading slash or not)
 * @returns The properly joined URL
 */
export const joinUrl = (base: string, path: string): string => {
  const baseTrimmed = base.endsWith('/') ? base.slice(0, -1) : base
  const pathTrimmed = path.startsWith('/') ? path : `/${path}`
  return `${baseTrimmed}${pathTrimmed}`
}
